package repository

import (
	"context"
)

type urlIdsRepository struct {
	db any // TODO: Complete with corresponding database
}

type UrlIdsRepository interface {
	GetNewUniqueId(ctx context.Context) (string, error)
}

// GetNewUniqueId implements repository.UrlMappingRepository
func (ur *urlIdsRepository) GetNewUniqueId(ctx context.Context) (string, error) {
	// TODO: Complete with corresponding implementation
	return "mocked_url", nil
}

func NewUrlIdsRepository(db any) UrlIdsRepository {
	return &urlIdsRepository{
		db: db,
	}
}
