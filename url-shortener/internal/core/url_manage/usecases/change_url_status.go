package usecases

import (
	"context"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
	coreDomain "github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
)

func (u *urlManageUseCase) ChangeUrlStatus(ctx context.Context, urlStatusInfo domain.UrlStatusInfo) error {
	updatedUrlMapping := coreDomain.UrlMapping{
		ShortUrlId: urlStatusInfo.ShortUrlId,
		Enabled:    urlStatusInfo.Enabled,
	}
	return u.UrlMappingRepository.SaveUrlMapping(ctx, updatedUrlMapping)
}
