package repository

import (
	"context"
	"github.com/joniaranguri/meli-urlshortener-challenge/url-shortener/internal/core/domain"
)

type urlMappingRepository struct {
	db any // TODO: Complete with corresponding database
}

type UrlMappingRepository interface {
	GetLongUrl(ctx context.Context, shortUrl string) (string, error)
	SaveUrlMapping(ctx context.Context, urlMapping domain.UrlMapping) error
}

// GetLongUrl implements repository.UrlMappingRepository
func (ur *urlMappingRepository) GetLongUrl(ctx context.Context, shortUrl string) (string, error) {
	// TODO: Complete with corresponding implementation
	return "https://articulo.mercadolibre.com.ar/MLA-1122519559-bicicleta-mtb-overtech-r29-acero-21v-freno-a-disco-pc-_JM#polycard_client=offers&deal_print_id=592cfeb6-8c12-4fb8-b7ea-e7b614346b80", nil
}

// SaveUrlMapping implements repository.UrlMappingRepository
func (ur *urlMappingRepository) SaveUrlMapping(ctx context.Context, urlMapping domain.UrlMapping) error {
	// TODO: Complete with corresponding implementation
	return nil
}

func NewUrlMappingRepository(db any) UrlMappingRepository {
	return &urlMappingRepository{
		db: db,
	}
}
