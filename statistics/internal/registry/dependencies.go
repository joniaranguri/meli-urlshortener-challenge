package registry

import (
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/core/statistics"
	"github.com/olebedev/config"
)

type AppContainer struct {
	StatisticsHandler statistics.Handler
}

type registry struct {
	conf config.Config
}

type Registry interface {
	NewAppContainer() AppContainer
}

func (r *registry) NewAppContainer() AppContainer {
	return AppContainer{
		StatisticsHandler: r.NewStatisticsHandler(),
	}
}

func NewRegistry(conf *config.Config) Registry {
	return &registry{
		conf: *conf,
	}
}
