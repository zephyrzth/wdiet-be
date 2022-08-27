package mongodb

import "go.mongodb.org/mongo-driver/mongo"

const DB_NAME = "healthfood"

type repo struct {
	mongoDB *mongo.Database
}

func New(mongoClient *mongo.Client) *repo {
	return &repo{
		mongoDB: mongoClient.Database(DB_NAME),
	}
}
