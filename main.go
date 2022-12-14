package main

import (
	"context"
	"net/http"
	"time"

	mRepo "github.com/zephyrzth/wdiet-be/repo/mongodb"
	mUsecase "github.com/zephyrzth/wdiet-be/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// init config
	cfg := getConfig()

	// init mongo db client
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(cfg.MongoDB.URI).SetServerAPIOptions(serverAPIOptions)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// init repo
	mongoDBRepo := mRepo.New(client)

	// init usecase
	usecase := mUsecase.New(mongoDBRepo)

	r := newRouter(usecase)

	http.ListenAndServe(":8080", r)
}
