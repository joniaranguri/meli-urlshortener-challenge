package registry

import "github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/repository"

func (r *registry) NewUrlMappingRepository() repository.UrlMappingRepository {
	client, err := r.NewUrlMappingDatabaseClient()

	if err != nil {
		panic(err)
	}
	return repository.NewUrlMappingRepository(client)
}

func (r *registry) NewUrlIdsRepository() repository.UrlIdsRepository {
	client, err := r.NewUrlIdsDatabaseClient()

	if err != nil {
		panic(err)
	}
	return repository.NewUrlIdsRepository(client)
}
