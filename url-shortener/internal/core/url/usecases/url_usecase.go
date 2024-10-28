package usecases

import (
	"context"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/repository"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url/domain"
)

type urlUseCase struct {
	urlMappingRepository repository.UrlMappingRepository
	urlIdsRepository     repository.UrlIdsRepository
}

type UrlUseCase interface {
	ShortenUrl(ctx context.Context, shortenUrlRequest domain.ShortenUrlRequest) (res domain.ShortenUrlResponse, err error)
	GetLongUrl(ctx context.Context, shortUrlId string) (string, error)
}

func NewUrlUseCase(mappingRepository repository.UrlMappingRepository, idsRepository repository.UrlIdsRepository) UrlUseCase {
	return &urlUseCase{
		urlMappingRepository: mappingRepository,
		urlIdsRepository:     idsRepository,
	}
}
