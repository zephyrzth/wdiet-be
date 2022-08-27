package mongodb

import (
	"context"
	"fmt"

	"github.com/zephyrzth/wdiet-be/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repo) GetRestaurantByID(ctx context.Context, id string) (model.Restaurants, error) {
	var result model.Restaurants
	objID, _ := primitive.ObjectIDFromHex(id)

	err := r.mongoDB.Collection(collectionRestaurants).FindOne(ctx, bson.M{"_id": objID}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Println("[repo][GetRestaurantByID] record does not exist")
	} else if err != nil {
		fmt.Println("[repo][GetRestaurantByID] error find data:", err)
	}

	return result, err
}
