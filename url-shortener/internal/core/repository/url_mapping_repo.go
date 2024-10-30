package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
	"gorm.io/gorm"
)

type urlMappingRepository struct {
	db           *gorm.DB
	cache        *redis.Client
	statisticsDb *redis.Client
}

type UrlMappingRepository interface {
	GetLongUrl(ctx context.Context, shortUrl string) (string, error)
	SaveUrlMapping(ctx context.Context, urlMapping domain.UrlMapping) error
	UpdateLongUrl(ctx context.Context, urlMapping domain.UrlMapping) error
	UpdateStatus(ctx context.Context, urlMapping domain.UrlMapping) error
	GetNewUniqueId(ctx context.Context) (string, error)
	SaveClickCountMetrics(ctx context.Context, shortUrlId string) error
}

// GetLongUrl retrieves the long URL associated with the given short URL.
func (ur *urlMappingRepository) GetLongUrl(ctx context.Context, shortUrl string) (string, error) {
	// Try to get the mapping from cache first
	cacheMapping, err := ur.getMappingFromCache(ctx, shortUrl)
	if err != nil {
		log.Printf("Error retrieving from cache: %v", err)
	}
	if cacheMapping.Active {
		return cacheMapping.LongUrl, nil
	}

	// If not found in cache, retrieve from the database
	var urlMapping domain.UrlMapping
	if err := ur.db.WithContext(ctx).Where("short_url = ?", shortUrl).First(&urlMapping).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("short URL not found")
		}
		return "", fmt.Errorf("database error: %v", err)
	}
	if !urlMapping.Active {
		return "", errors.New("the URL is disabled")
	}

	// Update the cache with the new mapping after a successful DB fetch
	if err := ur.sendMappingToCache(ctx, urlMapping); err != nil {
		log.Printf("Error updating cache after fetching from DB: %v", err)
	}

	return urlMapping.LongUrl, nil
}

// UpdateLongUrl updates the long url in the database and updates the cache.
func (ur *urlMappingRepository) UpdateLongUrl(ctx context.Context, urlMapping domain.UrlMapping) error {
	return ur.updateMappingProperty(ctx, urlMapping, "long_url", urlMapping.LongUrl)
}

// UpdateStatus updates the url status in the database and updates the cache.
func (ur *urlMappingRepository) UpdateStatus(ctx context.Context, urlMapping domain.UrlMapping) error {
	return ur.updateMappingProperty(ctx, urlMapping, "active", urlMapping.Active)
}

// SaveUrlMapping saves the URL mapping to the database and updates the cache.
func (ur *urlMappingRepository) SaveUrlMapping(ctx context.Context, urlMapping domain.UrlMapping) error {
	if err := ur.db.WithContext(ctx).Save(&urlMapping).Error; err != nil {
		return fmt.Errorf("error saving URL mapping: %v", err)
	}
	// Update the cache with the new mapping
	if err := ur.sendMappingToCache(ctx, urlMapping); err != nil {
		log.Printf("Error updating cache after saving: %v", err)
	}
	return nil
}

// updateMappingProperty updates a single property of url mapping the database and updates the cache.
func (ur *urlMappingRepository) updateMappingProperty(ctx context.Context, urlMapping domain.UrlMapping, propertyName string, propertyValue any) error {
	var existingMapping domain.UrlMapping

	// Check if the URL mapping already exists
	if err := ur.db.WithContext(ctx).First(&existingMapping, "short_url = ?", urlMapping.ShortUrlId).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("error checking existing URL mapping: %v", err)
		}
	} else {
		// If there is an existing mapping, check the userId
		if existingMapping.UserId != "" && existingMapping.UserId != urlMapping.UserId {
			return fmt.Errorf("unauthorized: user cannot modify this mapping")
		}
	}

	if err := ur.db.WithContext(ctx).Model(&urlMapping).Update(propertyName, propertyValue).Error; err != nil {
		return fmt.Errorf("error updating URL mapping: %v", err)
	}

	switch propertyName {
	case "active":
		if activeValue, ok := propertyValue.(bool); ok {
			existingMapping.Active = activeValue
		} else {
			return fmt.Errorf("propertyValue for 'active' must be of type bool")
		}
	case "long_url":
		if longURLValue, ok := propertyValue.(string); ok {
			existingMapping.LongUrl = longURLValue
		} else {
			return fmt.Errorf("propertyValue for 'long_url' must be of type string")
		}
	default:
		return fmt.Errorf("unsupported property name: %s", propertyName)
	}

	// Update the cache with the new mapping
	if err := ur.sendMappingToCache(ctx, existingMapping); err != nil {
		log.Printf("Error updating cache after saving: %v", err)
	}
	return nil
}

// GetNewUniqueId retrieves a new unique ID for a URL mapping.
func (ur *urlMappingRepository) GetNewUniqueId(ctx context.Context) (string, error) {
	var urlMapping domain.UrlMapping

	// Use a transaction to ensure atomicity
	err := ur.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Try to retrieve an available unique ID
		if err := tx.Model(&urlMapping).Where("long_url IS NULL").Select("short_url").First(&urlMapping).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("no available unique ID found")
			}
			return fmt.Errorf("database error: %v", err)
		}

		// Mark it as used
		if err := tx.Model(&urlMapping).Update("long_url", "Assigned").Error; err != nil {
			return fmt.Errorf("error updating long URL: %v", err)
		}
		return nil
	})

	if err != nil {
		return "", err
	}
	return urlMapping.ShortUrlId, nil
}

// SaveClickCountMetrics increments the click count metric for a short URL ID.
func (ur *urlMappingRepository) SaveClickCountMetrics(ctx context.Context, shortUrlId string) error {
	return ur.statisticsDb.Incr(ctx, shortUrlId).Err()
}

// getMappingFromCache retrieves the URL mapping from Redis cache.
func (ur *urlMappingRepository) getMappingFromCache(ctx context.Context, shortUrlId string) (domain.UrlMapping, error) {
	value, err := ur.cache.Get(ctx, shortUrlId).Result()
	if err == redis.Nil {
		return domain.UrlMapping{}, nil // Cache miss
	} else if err != nil {
		return domain.UrlMapping{}, fmt.Errorf("cache error: %v", err)
	}

	var mapping domain.UrlMapping
	if err := json.Unmarshal([]byte(value), &mapping); err != nil {
		return domain.UrlMapping{}, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return mapping, nil
}

// sendMappingToCache stores the URL mapping in Redis cache.
func (ur *urlMappingRepository) sendMappingToCache(ctx context.Context, mapping domain.UrlMapping) error {
	jsonData, err := json.Marshal(mapping)
	if err != nil {
		return fmt.Errorf("failed to marshal UrlMapping to JSON: %v", err)
	}

	err = ur.cache.Set(ctx, mapping.ShortUrlId, jsonData, 0).Err()
	return err
}

// NewUrlMappingRepository initializes a new UrlMappingRepository.
func NewUrlMappingRepository(db *gorm.DB, cache *redis.Client, statisticsDb *redis.Client) UrlMappingRepository {
	return &urlMappingRepository{
		db:           db,
		cache:        cache,
		statisticsDb: statisticsDb,
	}
}
