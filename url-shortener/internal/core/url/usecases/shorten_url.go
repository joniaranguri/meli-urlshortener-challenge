package usecases

import (
	"context"
	coreDomain "github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url/domain"
)

func (u *urlUseCase) ShortenUrl(ctx context.Context, shortenUrlRequest domain.ShortenUrlRequest) (res domain.ShortenUrlResponse, err error) {
	shortUrlId, err := u.urlMappingRepository.GetNewUniqueId(ctx)
	if err != nil {
		return res, err
	}
	err = u.urlMappingRepository.SaveUrlMapping(ctx, coreDomain.UrlMapping{
		ShortUrlId: shortUrlId,
		LongUrl:    shortenUrlRequest.LongUrl,
		UserId:     "user", // TODO: Complete with userId
		Active:     true,
	})
	if err != nil {
		return res, err
	}
	res.ShortUrl = buildShortUrl(shortUrlId)
	return res, err
}

func buildShortUrl(urlId string) string {
	return "http://localhost:80/" + urlId
}
