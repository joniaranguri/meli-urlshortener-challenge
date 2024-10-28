package usecases

import (
	"context"
	coreDomain "github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/repository"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url_manage/domain"
)

type urlManageUseCase struct {
	UrlMappingRepository repository.UrlMappingRepository
}

type UrlManageUseCase interface {
	UpdateUrl(ctx context.Context, request domain.PatchUrlRequest) (domain.PatchUrlResponse, error)
	ChangeUrlStatus(ctx context.Context, urlStatusInfo coreDomain.UrlStatusInfo) error
}

func NewUrlManageUseCase(repo repository.UrlMappingRepository) UrlManageUseCase {
	return &urlManageUseCase{
		UrlMappingRepository: repo,
	}
}
