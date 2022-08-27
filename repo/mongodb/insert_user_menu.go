package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/zephyrzth/wdiet-be/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repo) InsertUserMenu(ctx context.Context, insertData model.InsertUserMenu) error {
	objectUserID, _ := primitive.ObjectIDFromHex(insertData.UserID)
	objectMenuID, _ := primitive.ObjectIDFromHex(insertData.MenuID)

	data := UserMenuBSON{
		MenuID:    objectMenuID,
		Quantity:  insertData.Quantity,
		Timestamp: time.Now(),
	}

	filter := bson.D{{Key: "_id", Value: objectUserID}}
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "menus", Value: data}}}}

	_, err := r.mongoDB.Collection(collectionUsers).UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println("[repo][InsertUserMenu] failed insert to user menu")
		return err
	}

	return nil
}
