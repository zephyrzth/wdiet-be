package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/zephyrzth/wdiet-be/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repo) InsertUserMenu(ctx context.Context, userMenu model.UserMenu) error {
	objectUserID, _ := primitive.ObjectIDFromHex(userMenu.UserID)
	objectMenuID, _ := primitive.ObjectIDFromHex(userMenu.MenuID)

	data := UserMenuBSON{
		MenuID:    objectMenuID,
		Quantity:  userMenu.Quantity,
		Timestamp: time.Now(),
	}

	filter := bson.D{{"_id", objectUserID}}
	update := bson.D{{"$push", bson.D{{"menus", data}}}}

	_, err := r.mongoDB.Collection(collectionUsers).UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println("[repo][InsertUserMenu] failed insert to user menu")
		return err
	}

	return nil
}
