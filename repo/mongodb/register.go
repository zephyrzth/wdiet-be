package mongodb

import (
	"context"
	"fmt"

	model "github.com/zephyrzth/wdiet-be/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repo) Register(ctx context.Context, user model.User) (string, error) {
	res, err := r.mongoDB.Collection(collectionUsers).InsertOne(ctx, user)
	if err != nil {
		fmt.Println("[repo][Register] error find data:", err)
	}

	userID, _ := res.InsertedID.(primitive.ObjectID)
	userIDStr := userID.Hex()

	return userIDStr, err
}
