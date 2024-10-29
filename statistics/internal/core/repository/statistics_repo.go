package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/core/domain"
	"strconv"
)

type statisticsRepository struct {
	metricsDb *redis.Client
}

type StatisticsRepository interface {
	GetClickStatistics(ctx context.Context, shortUrl string) (domain.ClickStatistics, error)
}

// GetClickStatistics fetches the click count for a given shortUrl from Redis.
func (ur *statisticsRepository) GetClickStatistics(ctx context.Context, shortUrl string) (domain.ClickStatistics, error) {
	// Fetch the click count from Redis
	clicks, err := ur.metricsDb.Get(ctx, shortUrl).Result()
	if err == redis.Nil {
		return domain.ClickStatistics{Clicks: 0}, nil
	} else if err != nil {
		return domain.ClickStatistics{}, err
	}

	clickCount, err := strconv.ParseUint(clicks, 10, 64)
	if err != nil {
		return domain.ClickStatistics{}, err
	}

	return domain.ClickStatistics{Clicks: clickCount}, nil
}

// NewStatisticsRepository initializes a new StatisticsRepository with the provided Redis client.
func NewStatisticsRepository(db *redis.Client) StatisticsRepository {
	return &statisticsRepository{
		metricsDb: db,
	}
}
