package url

import (
	"context"
	coreDomain "github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
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
	shortUrl, err := handler.userCases.ShortenUrl(ctx, coreDomain.UrlMapping{
		LongUrl: request.LongUrl,
		UserId:  request.UserId,
		Active:  true,
	})
	return domain.ShortenUrlResponse{
		ShortUrl: shortUrl,
	}, err
}

func (handler *urlHandler) GetLongUrl(ctx context.Context, shortUrlId string) (string, error) {
	return handler.userCases.GetLongUrl(ctx, shortUrlId)
}
