package usecases

import (
	"context"
	coreDomain "github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url_manage/domain"
)

func (u *urlManageUseCase) UpdateUrl(ctx context.Context, request domain.PatchUrlRequest) (res domain.PatchUrlResponse, err error) {
	updatedUrlMapping := coreDomain.UrlMapping{
		ShortUrlId: request.ShortUrlId,
		LongUrl:    request.LongUrl,
	}
	err = u.UrlMappingRepository.SaveUrlMapping(ctx, updatedUrlMapping)
	if err != nil {
		return res, err
	}
	res = domain.PatchUrlResponse{
		LongUrl: updatedUrlMapping.LongUrl,
	}
	return res, err
}
