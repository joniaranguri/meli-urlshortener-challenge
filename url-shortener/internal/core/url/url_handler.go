package url

import (
	"context"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url/usecases"
)

type urlHandler struct {
	userCases usecases.UrlUseCase
}

type Handler interface {
	ShortenUrl(ctx context.Context, request domain.ShortenUrlRequest) (domain.ShortenUrlResponse, error)
	GetLongUrl(ctx context.Context, shortUrlId string) (string, error)
}

func NewUrlHandler(us usecases.UrlUseCase) Handler {
	return &urlHandler{us}
}

func (handler *urlHandler) ShortenUrl(ctx context.Context, request domain.ShortenUrlRequest) (domain.ShortenUrlResponse, error) {
	return handler.userCases.ShortenUrl(ctx, request)
}

func (handler *urlHandler) GetLongUrl(ctx context.Context, shortUrlId string) (string, error) {
	return handler.userCases.GetLongUrl(ctx, shortUrlId)
}
