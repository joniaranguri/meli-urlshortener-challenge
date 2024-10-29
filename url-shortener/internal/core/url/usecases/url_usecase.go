package usecases

import (
	"context"
	coreDomain "github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/repository"
)

type urlUseCase struct {
	urlMappingRepository repository.UrlMappingRepository
}

type UrlUseCase interface {
	ShortenUrl(ctx context.Context, urlMapping coreDomain.UrlMapping) (string, error)
	GetLongUrl(ctx context.Context, shortUrlId string) (string, error)
}

func NewUrlUseCase(mappingRepository repository.UrlMappingRepository) UrlUseCase {
	return &urlUseCase{
		urlMappingRepository: mappingRepository,
	}
}
