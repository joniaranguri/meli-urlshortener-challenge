package url_manage

import (
	"context"
	coreDomain "github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url_manage/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url_manage/usecases"
)

type urlManageHandler struct {
	userCases usecases.UrlManageUseCase
}

type Handler interface {
	UpdateUrl(ctx context.Context, request domain.PatchUrlRequest) (domain.PatchUrlResponse, error)
	EnableUrl(ctx context.Context, shortUrlId string) error
	DisableUrl(ctx context.Context, shortUrlId string) error
}

func NewUrlManageHandler(us usecases.UrlManageUseCase) Handler {
	return &urlManageHandler{us}
}

func (handler *urlManageHandler) UpdateUrl(ctx context.Context, request domain.PatchUrlRequest) (domain.PatchUrlResponse, error) {
	return handler.userCases.UpdateUrl(ctx, request)
}

func (handler *urlManageHandler) EnableUrl(ctx context.Context, shortUrlId string) error {
	return handler.userCases.ChangeUrlStatus(ctx, coreDomain.UrlStatusInfo{
		ShortUrlId: shortUrlId,
		Active:     true,
	})
}

func (handler *urlManageHandler) DisableUrl(ctx context.Context, shortUrlId string) error {
	return handler.userCases.ChangeUrlStatus(ctx, coreDomain.UrlStatusInfo{
		ShortUrlId: shortUrlId,
		Active:     false,
	})
}
