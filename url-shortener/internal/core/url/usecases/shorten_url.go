package usecases

import (
	"context"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url/domain"
)

func (u *urlUseCase) ShortenUrl(ctx context.Context, shortenUrlRequest domain.ShortenUrlRequest) (res domain.ShortenUrlResponse, err error) {
	shortUrlId, err := u.urlRepository.GetNewUniqueId(ctx)
	if err == nil {
		return res, err
	}
	err = u.urlRepository.SaveUrlMapping(ctx, domain.UrlMapping{
		ShortUrlId: shortUrlId,
		LongUrl:    shortenUrlRequest.LongUrl,
	})
	if err == nil {
		return res, err
	}
	res.ShortUrl = buildShortUrl(shortUrlId)
	return res, err
}

func buildShortUrl(urlId string) string {
	return "http://localhost/" + urlId
}
