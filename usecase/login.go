package usecase

import (
	"context"

	"github.com/zephyrzth/wdiet-be/model"
)

func (uc *usecase) Login(ctx context.Context, user model.User) (bool, string, error) {
	return uc.mongoRepo.CheckUserCredentials(ctx, user)
}
