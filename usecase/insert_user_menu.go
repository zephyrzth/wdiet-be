package usecase

import (
	"context"

	"github.com/zephyrzth/wdiet-be/model"
)

func (uc *usecase) InsertUserMenu(ctx context.Context, userMenu model.UserMenu) error {
	return uc.mongoRepo.InsertUserMenu(ctx, userMenu)
}
