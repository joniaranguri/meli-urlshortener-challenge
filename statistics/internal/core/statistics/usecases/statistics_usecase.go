package usecases

import (
	"context"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/core/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/core/repository"
)

type statisticsUseCase struct {
	StatisticsRepository repository.StatisticsRepository
}

type StatisticsUseCase interface {
	GetClickStatistics(ctx context.Context, shortUrlId string) (domain.ClickStatistics, error)
}

func NewStatisticsUseCase(repo repository.StatisticsRepository) StatisticsUseCase {
	return &statisticsUseCase{
		StatisticsRepository: repo,
	}
}
