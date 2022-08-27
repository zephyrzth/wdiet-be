package usecase

import (
	"context"

	"github.com/zephyrzth/wdiet-be/model"
)

func (uc *usecase) GetAllRestaurant(ctx context.Context) ([]model.Restaurants, error) {
	return uc.mongoRepo.GetAllRestaurant(ctx)
}
