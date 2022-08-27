package usecase

import (
	"context"

	"github.com/zephyrzth/wdiet-be/model"
)

func (uc *usecase) Register(ctx context.Context, user model.User) (string, error) {
	return uc.mongoRepo.Register(ctx, user)
}
