package usecases

import (
	"context"
	coreDomain "github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/utils/constants"
)

func (u *urlUseCase) ShortenUrl(ctx context.Context, urlMapping coreDomain.UrlMapping) (string, error) {
	shortUrlId, err := u.urlMappingRepository.GetNewUniqueId(ctx)
	if err != nil {
		return "", err
	}
	urlMapping.ShortUrlId = shortUrlId

	err = u.urlMappingRepository.SaveUrlMapping(ctx, urlMapping)
	if err != nil {
		return "", err
	}
	return buildShortUrl(shortUrlId), err
}

func buildShortUrl(urlId string) string {
	return constants.BaseUrl + urlId
}
