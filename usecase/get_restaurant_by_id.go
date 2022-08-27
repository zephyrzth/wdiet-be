package usecase

import (
	"context"

	"github.com/zephyrzth/wdiet-be/model"
)

func (uc *usecase) GetRestaurantByID(ctx context.Context, id string) (model.Restaurants, error) {
	return uc.mongoRepo.GetRestaurantByID(ctx, id)
}
