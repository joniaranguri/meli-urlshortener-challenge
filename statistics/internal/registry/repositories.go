package registry

import "github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/core/repository"

func (r *registry) NewStatisticsRepository() repository.StatisticsRepository {
	client, err := r.NewStatisticsDatabaseClient()

	if err != nil {
		panic(err)
	}
	return repository.NewStatisticsRepository(client)
}