package repository

import (
	"context"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/url/domain"
)

type urlRepository struct {
	db any // TODO: Complete with corresponding database
}

type UrlRepository interface {
	GetNewUniqueId(ctx context.Context) (string, error)
	GetLongUrl(ctx context.Context, shortUrl string) (string, error)
	SaveUrlMapping(ctx context.Context, urlMapping domain.UrlMapping) error
}

// GetNewUniqueId implements repository.UrlRepository
func (ur *urlRepository) GetNewUniqueId(ctx context.Context) (string, error) {
	// TODO: Complete with corresponding implementation
	return "mocked_url", nil
}

// GetLongUrl implements repository.UrlRepository
func (ur *urlRepository) GetLongUrl(ctx context.Context, shortUrl string) (string, error) {
	// TODO: Complete with corresponding implementation
	return "https://articulo.mercadolibre.com.ar/MLA-1122519559-bicicleta-mtb-overtech-r29-acero-21v-freno-a-disco-pc-_JM#polycard_client=offers&deal_print_id=592cfeb6-8c12-4fb8-b7ea-e7b614346b80", nil
}

// SaveUrlMapping implements repository.UrlRepository
func (ur *urlRepository) SaveUrlMapping(ctx context.Context, urlMapping domain.UrlMapping) error {
	// TODO: Complete with corresponding implementation
	return nil
}

func NewUrlRepository() UrlRepository {
	return &urlRepository{}
}
