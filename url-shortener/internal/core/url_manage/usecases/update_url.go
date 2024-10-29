package usecases

import (
	"context"
	"errors"
	coreDomain "github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
	"log"
)

// UpdateUrl updates a URL mapping in the repository after validating input.
func (u *urlManageUseCase) UpdateUrl(ctx context.Context, urlMapping coreDomain.UrlMapping) (string, error) {
	if urlMapping.ShortUrlId == "" {
		return "", errors.New("short URL ID cannot be empty")
	}
	if urlMapping.LongUrl == "" {
		return "", errors.New("long URL cannot be empty")
	}

	if err := u.UrlMappingRepository.SaveUrlMapping(ctx, urlMapping); err != nil {
		log.Printf("Error updating URL mapping: %v", err)
		return "", err
	}

	return urlMapping.LongUrl, nil
}
