package mongodb

import (
	"context"
	"fmt"

	model "github.com/zephyrzth/wdiet-be/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repo) GetUser(ctx context.Context, userID string) (model.User, error) {
	var result model.User
	objID, _ := primitive.ObjectIDFromHex(userID)

	err := r.mongoDB.Collection(collectionUsers).FindOne(ctx, bson.M{"_id": objID}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Println("[repo][GetUser] record does not exist")
	} else if err != nil {
		fmt.Println("[repo][GetUser] error find data:", err)
	}

	return result, err
}
