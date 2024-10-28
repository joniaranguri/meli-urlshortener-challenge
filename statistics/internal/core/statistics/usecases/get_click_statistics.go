package usecases

import (
	"context"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/core/domain"
)

func (u *statisticsUseCase) GetClickStatistics(ctx context.Context, shortUrlId string) (res domain.ClickStatistics, err error) {
	return u.StatisticsRepository.GetClickStatistics(ctx, shortUrlId)
}
