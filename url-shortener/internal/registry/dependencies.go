package registry

import (
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url_manage"
	"github.com/olebedev/config"
)

type AppContainer struct {
	UrlHandler       url.Handler
	UrlManageHandler url_manage.Handler
}

type registry struct {
	conf config.Config
}

type Registry interface {
	NewAppContainer() AppContainer
}

func (r *registry) NewAppContainer() AppContainer {
	return AppContainer{
		UrlHandler:       r.NewUrlHandler(),
		UrlManageHandler: r.NewUrlManageHandler(),
	}
}

func NewRegistry(conf *config.Config) Registry {
	return &registry{
		conf: *conf,
	}
}
