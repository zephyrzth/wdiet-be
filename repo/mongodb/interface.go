package mongodb

import (
	"context"

	model "github.com/zephyrzth/wdiet-be/model"
)

type RepoInterface interface {
	GetAllRestaurant(ctx context.Context) ([]model.Restaurants, error)
	GetRestaurantByID(ctx context.Context, id string) (model.Restaurants, error)
	Register(ctx context.Context, user model.User) (string, error)
	GetUser(ctx context.Context, userID string) (model.User, error)
	GetMenuByID(ctx context.Context, menuID string) (model.Menu, error)
}
