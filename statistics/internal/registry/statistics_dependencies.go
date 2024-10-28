package registry

import (
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/core/statistics"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/core/statistics/usecases"
)

func (r *registry) NewStatisticsHandler() statistics.Handler {
	return statistics.NewStatisticsHandler(r.NewStatisticsUseCase())
}

func (r *registry) NewStatisticsUseCase() usecases.StatisticsUseCase {
	return usecases.NewStatisticsUseCase(r.NewStatisticsRepository())
}
