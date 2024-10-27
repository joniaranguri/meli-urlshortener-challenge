package usecases

import (
	"context"
)

func (u *urlUseCase) GetLongUrl(ctx context.Context, shortUrlId string) (string, error) {
	// TODO: Implement metrics logic
	return u.urlRepository.GetLongUrl(ctx, shortUrlId)
}
