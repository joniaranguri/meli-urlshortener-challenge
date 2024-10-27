package registry

import (
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url"
	"github.com/olebedev/config"
)

type AppContainer struct {
	UrlHandler url.Handler
}

type registry struct {
	conf config.Config
}

type Registry interface {
	NewAppContainer() AppContainer
}

func (r *registry) NewAppContainer() AppContainer {
	return AppContainer{
		UrlHandler: r.NewUrlHandler(),
	}
}

func NewRegistry(conf *config.Config) Registry {
	return &registry{
		conf: *conf,
	}
}
