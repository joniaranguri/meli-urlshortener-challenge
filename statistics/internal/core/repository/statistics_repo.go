package repository

import (
	"context"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/core/domain"
)

type statisticsRepository struct {
	db any // TODO: Complete with corresponding database
}

type StatisticsRepository interface {
	GetClickStatistics(ctx context.Context, shortUrl string) (domain.ClickStatistics, error)
	SaveStatistics(ctx context.Context, shortUrlId string) error
}

// GetClickStatistics implements repository.StatisticsRepository
func (ur *statisticsRepository) GetClickStatistics(ctx context.Context, shortUrl string) (domain.ClickStatistics, error) {
	// TODO: Complete with corresponding implementation
	return domain.ClickStatistics{
		Clicks: uint64(8934853),
	}, nil
}

// SaveStatistics implements repository.StatisticsRepository
func (ur *statisticsRepository) SaveStatistics(ctx context.Context, shortUrlId string) error {
	// TODO: Complete with corresponding implementation
	return nil
}

func NewStatisticsRepository(db any) StatisticsRepository {
	return &statisticsRepository{
		db: db,
	}
}
