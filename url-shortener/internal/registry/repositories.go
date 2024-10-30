package registry

import (
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/repository"
	"os"
)

func (r *registry) NewUrlMappingRepository() repository.UrlMappingRepository {
	dbClient, err := r.NewUrlMappingDatabaseClient()

	if err != nil {
		panic(err)
	}
	cacheClient, err := r.NewUrlMappingCacheClient()

	scope := os.Getenv("SCOPE")

	if err != nil && scope == "DEMO" {
		panic(err)
	}

	statisticsDbClient, err := r.NewStatisticsDbClient()

	if err != nil && scope == "DEMO" {
		panic(err)
	}
	return repository.NewUrlMappingRepository(dbClient, cacheClient, statisticsDbClient)
}
