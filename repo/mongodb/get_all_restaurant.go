package mongodb

import (
	"context"
	"fmt"

	model "github.com/zephyrzth/wdiet-be/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) GetAllRestaurant(ctx context.Context) ([]model.Restaurants, error) {
	cur, err := r.mongoDB.Collection(collectionRestaurants).Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("[repo][GetAllRestaurant] error find data:", err)
		return []model.Restaurants{}, err
	}
	defer cur.Close(ctx)

	var data []model.Restaurants

	if err = cur.All(ctx, &data); err != nil {
		fmt.Println("[repo][GetAllRestaurant] error in cursor:", err)
		return []model.Restaurants{}, err
	}

	return data, nil
}
