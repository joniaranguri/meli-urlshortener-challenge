package registry

import (
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url/usecases"
)

func (r *registry) NewUrlHandler() url.Handler {
	return url.NewUrlHandler(r.NewUrlUseCase())
}

func (r *registry) NewUrlUseCase() usecases.UrlUseCase {
	return usecases.NewUrlUseCase(r.NewUrlMappingRepository(), r.NewUrlIdsRepository())
}
