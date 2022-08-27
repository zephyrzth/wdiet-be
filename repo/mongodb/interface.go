package mongodb

import (
	"context"

	model "github.com/zephyrzth/wdiet-be/model"
)

type RepoInterface interface {
	GetAllRestaurant(ctx context.Context) ([]model.Restaurants, error)
	GetRestaurantByID(ctx context.Context, id string) (model.Restaurants, error)
}
