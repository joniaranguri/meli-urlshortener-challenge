package statistics

import (
	"context"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/core/statistics/domain"
	"github.com/joniaranguri/meli-urlshortener-challenge/statistics/internal/core/statistics/usecases"
)

type statisticsHandler struct {
	userCases usecases.StatisticsUseCase
}

type Handler interface {
	GetClickStatistics(ctx context.Context, shortUrlId string) (domain.StatisticsResponse, error)
}

func NewStatisticsHandler(us usecases.StatisticsUseCase) Handler {
	return &statisticsHandler{us}
}

func (handler *statisticsHandler) GetClickStatistics(ctx context.Context, shortUrlId string) (res domain.StatisticsResponse, err error) {
	stats, err := handler.userCases.GetClickStatistics(ctx, shortUrlId)
	if err != nil {
		return res, err
	}
	res = domain.StatisticsResponse{
		Clicks:     stats.Clicks,
		ShortUrlId: shortUrlId,
	}
	return res, err
}
