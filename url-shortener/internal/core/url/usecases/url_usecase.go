package usecases

import (
	"context"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url/repository"
)

type urlUseCase struct {
	urlRepository repository.UrlRepository
}

type UrlUseCase interface {
	ShortenUrl(ctx context.Context, shortenUrlRequest domain.ShortenUrlRequest) (res domain.ShortenUrlResponse, err error)
	GetLongUrl(ctx context.Context, shortUrlId string) (string, error)
}

func NewUrlUseCase(repo repository.UrlRepository) UrlUseCase {
	return &urlUseCase{
		urlRepository: repo,
	}
}
