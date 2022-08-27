package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserMenuBSON struct {
	MenuID    primitive.ObjectID `bson:"id"`
	Quantity  int                `bson:"quantity"`
	Timestamp time.Time          `bson:"timestamp"`
}
