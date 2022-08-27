package usecase

import (
	"context"

	"github.com/zephyrzth/wdiet-be/model"
)

type UsecaseInterface interface {
	GetRestaurantByID(ctx context.Context, id string) (model.Restaurants, error)
	GetAllRestaurant(ctx context.Context) ([]model.Restaurants, error)
	Register(ctx context.Context, user model.User) (string, error)
	GetProfile(ctx context.Context, userID string) (model.User, error)
}
