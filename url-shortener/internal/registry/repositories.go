package registry

import "github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/repository"

func (r *registry) NewUrlMappingRepository() repository.UrlMappingRepository {
	dbClient, err := r.NewUrlMappingDatabaseClient()

	if err != nil {
		panic(err)
	}
	cacheClient, err := r.NewUrlMappingCacheClient()

	if err != nil {
		panic(err)
	}
	return repository.NewUrlMappingRepository(dbClient, cacheClient)
}
