package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
	"gorm.io/gorm"
)

type urlMappingRepository struct {
	db *gorm.DB
}

type UrlMappingRepository interface {
	GetLongUrl(ctx context.Context, shortUrl string) (string, error)
	SaveUrlMapping(ctx context.Context, urlMapping domain.UrlMapping) error
	GetNewUniqueId(ctx context.Context) (string, error)
}

// GetLongUrl implements repository.UrlMappingRepository
func (ur *urlMappingRepository) GetLongUrl(ctx context.Context, shortUrl string) (string, error) {
	var urlMapping domain.UrlMapping
	if err := ur.db.WithContext(ctx).Where("short_url = ?", shortUrl).First(&urlMapping).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", fmt.Errorf("short URL not found")
		}
		return "", err
	}
	return urlMapping.LongUrl, nil
}

// SaveUrlMapping implements repository.UrlMappingRepository
func (ur *urlMappingRepository) SaveUrlMapping(ctx context.Context, urlMapping domain.UrlMapping) error {
	return ur.db.WithContext(ctx).
		Model(&urlMapping).
		Updates(urlMapping).Error
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

// NewUrlMappingRepository initializes a new UrlMappingRepository
func NewUrlMappingRepository(db *gorm.DB) UrlMappingRepository {
	return &urlMappingRepository{
		db: db,
	}
}
