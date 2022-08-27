package mongodb

import (
	"context"
	"fmt"

	model "github.com/zephyrzth/wdiet-be/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repo) GetMenuByID(ctx context.Context, menuID string) (model.Menu, error) {
	var result model.Menu
	objID, _ := primitive.ObjectIDFromHex(menuID)

	err := r.mongoDB.Collection(collectionMenus).FindOne(ctx, bson.M{"_id": objID}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Println("[repo][GetMenuByID] record does not exist")
	} else if err != nil {
		fmt.Println("[repo][GetMenuByID] error find data:", err)
	}

	return result, err
}
