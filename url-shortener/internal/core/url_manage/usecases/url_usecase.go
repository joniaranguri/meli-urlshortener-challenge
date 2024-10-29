package usecases

import (
	"context"
	coreDomain "github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/repository"
)

type urlManageUseCase struct {
	UrlMappingRepository repository.UrlMappingRepository
}

type UrlManageUseCase interface {
	UpdateUrl(ctx context.Context, urlMapping coreDomain.UrlMapping) (string, error)
	ChangeUrlStatus(ctx context.Context, urlStatusInfo coreDomain.UrlMapping) error
}

func NewUrlManageUseCase(repo repository.UrlMappingRepository) UrlManageUseCase {
	return &urlManageUseCase{
		UrlMappingRepository: repo,
	}
}
