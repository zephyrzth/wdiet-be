package usecase

import (
	"context"

	"github.com/zephyrzth/wdiet-be/model"
)

func (uc *usecase) InsertUserMenu(ctx context.Context, insertData model.InsertUserMenu) error {
	return uc.mongoRepo.InsertUserMenu(ctx, insertData)
}
