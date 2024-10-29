package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"

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
	GetNewUniqueId(ctx context.Context) (string, error)
	SaveClickCountMetrics(ctx context.Context, shortUrlId string) error
}

// GetLongUrl implements repository.UrlMappingRepository
func (ur *urlMappingRepository) GetLongUrl(ctx context.Context, shortUrl string) (string, error) {
	var urlMapping domain.UrlMapping

	// Try to get the mapping from cache first
	cacheMapping, err := ur.getMappingFromCache(ctx, shortUrl)
	if err == nil && cacheMapping.ShortUrlId != "" {
		return cacheMapping.LongUrl, nil
	}

	// If not found in cache, retrieve from the database
	if err := ur.db.WithContext(ctx).Where("short_url = ?", shortUrl).First(&urlMapping).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", fmt.Errorf("short URL not found")
		}
		return "", err
	}

	// Update the cache with the new mapping after a successful DB fetch
	if err := ur.sendMappingToCache(ctx, urlMapping); err != nil {
		log.Printf("Error setting data to cache: %v", err)
	}

	return urlMapping.LongUrl, nil
}

// SaveUrlMapping implements repository.UrlMappingRepository
func (ur *urlMappingRepository) SaveUrlMapping(ctx context.Context, urlMapping domain.UrlMapping) error {
	// Save to database
	if err := ur.db.WithContext(ctx).Model(&urlMapping).Updates(urlMapping).Error; err != nil {
		return err
	}

	// Update the cache with the new mapping
	if err := ur.sendMappingToCache(ctx, urlMapping); err != nil {
		log.Printf("Error updating cache after saving: %v", err)
	}
	return nil
}

// GetNewUniqueId implements repository.UrlMappingRepository
func (ur *urlMappingRepository) GetNewUniqueId(ctx context.Context) (string, error) {
	var urlMapping domain.UrlMapping

	tx := ur.db.WithContext(ctx)
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&urlMapping).Select("short_url").Where("long_url IS NULL").First(&urlMapping.ShortUrlId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("no available unique ID found")
			}
			return err
		}
		urlMapping.LongUrl = "Assigned"
		if err := tx.Model(&urlMapping).Updates(map[string]interface{}{
			"long_url": "Assigned",
		}).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return "", err
	}
	return urlMapping.ShortUrlId, nil
}

func (ur *urlMappingRepository) SaveClickCountMetrics(ctx context.Context, shortUrlId string) error {
	return ur.statisticsDb.Incr(ctx, shortUrlId).Err()
}

func (ur *urlMappingRepository) getMappingFromCache(ctx context.Context, shortUrlId string) (mapping domain.UrlMapping, err error) {
	value, err := ur.cache.Get(ctx, shortUrlId).Result()
	if err == redis.Nil {
		return mapping, nil // Cache miss
	} else if err != nil {
		return mapping, err // Other error
	}

	// Deserialize the JSON string into the UrlMapping struct
	err = json.Unmarshal([]byte(value), &mapping)
	if err != nil {
		log.Printf("Failed to unmarshal JSON: %v", err)
		return mapping, err
	}

	// Cache hit
	return mapping, nil
}

func (ur *urlMappingRepository) sendMappingToCache(ctx context.Context, mapping domain.UrlMapping) error {
	// Serialize the UrlMapping struct to JSON
	jsonData, err := json.Marshal(mapping)
	if err != nil {
		log.Printf("Failed to marshal UrlMapping to JSON: %v", err)
		return err
	}

	// Store the JSON string in Redis with LRU eviction policy
	err = ur.cache.Set(ctx, mapping.ShortUrlId, jsonData, 0).Err() // You may specify an expiration time if needed
	return err
}

// NewUrlMappingRepository initializes a new UrlMappingRepository
func NewUrlMappingRepository(db *gorm.DB, cache *redis.Client, statisticsDb *redis.Client) UrlMappingRepository {
	return &urlMappingRepository{
		db:           db,
		cache:        cache,
		statisticsDb: statisticsDb,
	}
}
