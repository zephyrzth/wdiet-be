package mongodb

import (
	"context"
	"fmt"

	model "github.com/zephyrzth/wdiet-be/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repo) CheckUserCredentials(ctx context.Context, user model.User) (bool, string, error) {
	var result model.User

	err := r.mongoDB.Collection(collectionUsers).FindOne(ctx, bson.M{"email": user.Email, "password": user.Password}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false, "", nil
	} else if err != nil {
		fmt.Println("[repo][GetUser] error find data:", err)
		return false, "", err
	}

	return true, result.ID, nil
}
