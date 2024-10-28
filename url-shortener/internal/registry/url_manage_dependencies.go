package registry

import (
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url_manage"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url_manage/usecases"
)

func (r *registry) NewUrlManageHandler() url_manage.Handler {
	return url_manage.NewUrlManageHandler(r.NewUrlManageUseCase())
}

func (r *registry) NewUrlManageUseCase() usecases.UrlManageUseCase {
	return usecases.NewUrlManageUseCase(r.NewUrlMappingRepository())
}
