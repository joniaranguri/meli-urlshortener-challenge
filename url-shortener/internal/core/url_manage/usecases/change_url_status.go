package usecases

import (
	"context"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
)

// ChangeUrlStatus enables or disables a URL  in the repository.
func (u *urlManageUseCase) ChangeUrlStatus(ctx context.Context, updatedUrlMapping domain.UrlMapping) error {
	return u.UrlMappingRepository.UpdateStatus(ctx, updatedUrlMapping)
}
