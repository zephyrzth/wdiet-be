package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	mRepo "github.com/zephyrzth/wdiet-be/repo/mongodb"
	mUsecase "github.com/zephyrzth/wdiet-be/usecase"
	"github.com/zephyrzth/wdiet-be/utils"
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

	restoID := "6309cf4b76fa949ae0b554ab"
	data, err := mongoDBRepo.GetRestaurantByID(ctx, restoID)

	fmt.Println("THE DATA", utils.PrettyPrint(data), err)

	r := newRouter(usecase)
	http.ListenAndServe(":8080", r)
}
