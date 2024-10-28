package usecases

import (
	"context"
)

func (u *statisticsUseCase) AddClickCountStatistic(ctx context.Context, shortUrlId string) error {
	return u.StatisticsRepository.SaveStatistics(ctx, shortUrlId)
}
