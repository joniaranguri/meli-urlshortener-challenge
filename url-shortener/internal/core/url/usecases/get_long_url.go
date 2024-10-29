package usecases

import (
	"context"
	"log"
)

func (u *urlUseCase) GetLongUrl(ctx context.Context, shortUrlId string) (string, error) {
	longUrl, err := u.urlMappingRepository.GetLongUrl(ctx, shortUrlId)
	if err != nil {
		return "", err
	}

	go u.saveClickMetric(context.Background(), shortUrlId)

	return longUrl, nil
}

func (u *urlUseCase) saveClickMetric(ctx context.Context, shortUrlId string) {
	if err := u.urlMappingRepository.SaveClickCountMetrics(ctx, shortUrlId); err != nil {
		log.Printf("Error saving click metric for URL ID %s: %v", shortUrlId, err)
	}
}
