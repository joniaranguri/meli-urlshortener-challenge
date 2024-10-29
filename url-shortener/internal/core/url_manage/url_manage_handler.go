package url_manage

import (
	"context"
	coreDomain "github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url_manage/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url_manage/usecases"
)

type urlManageHandler struct {
	useCases usecases.UrlManageUseCase
}

// Handler interface defines the methods for managing URLs.
type Handler interface {
	UpdateUrl(ctx context.Context, request domain.PatchUrlRequest) (domain.PatchUrlResponse, error)
	EnableUrl(ctx context.Context, shortUrlId string, userId string) error
	DisableUrl(ctx context.Context, shortUrlId string, userId string) error
}

// NewUrlManageHandler initializes a new urlManageHandler.
func NewUrlManageHandler(us usecases.UrlManageUseCase) Handler {
	return &urlManageHandler{useCases: us}
}

// UpdateUrl handles the request to update a URL.
func (h *urlManageHandler) UpdateUrl(ctx context.Context, request domain.PatchUrlRequest) (domain.PatchUrlResponse, error) {
	updatedLongUrl, err := h.useCases.UpdateUrl(ctx, coreDomain.UrlMapping{
		ShortUrlId: request.ShortUrlId,
		LongUrl:    request.LongUrl,
		UserId:     request.UserId,
	})
	if err != nil {
		return domain.PatchUrlResponse{}, err
	}
	return domain.PatchUrlResponse{LongUrl: updatedLongUrl}, nil
}

// EnableUrl handles the request to enable a URL.
func (h *urlManageHandler) EnableUrl(ctx context.Context, shortUrlId string, userId string) error {
	return h.useCases.ChangeUrlStatus(ctx, coreDomain.UrlMapping{
		ShortUrlId: shortUrlId,
		Active:     true,
		UserId:     userId,
	})
}

// DisableUrl handles the request to disable a URL.
func (h *urlManageHandler) DisableUrl(ctx context.Context, shortUrlId string, userId string) error {
	return h.useCases.ChangeUrlStatus(ctx, coreDomain.UrlMapping{
		ShortUrlId: shortUrlId,
		Active:     false,
		UserId:     userId,
	})
}
