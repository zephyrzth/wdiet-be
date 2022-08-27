package main

import (
	"github.com/gorilla/mux"
	handlerHttp "github.com/zephyrzth/wdiet-be/handler/http"
	"github.com/zephyrzth/wdiet-be/usecase"
)

func newRouter(usecase usecase.UsecaseInterface) *mux.Router {
	handlerHttp := handlerHttp.New(usecase)

	r := mux.NewRouter()
	r.HandleFunc("/restaurants", handlerHttp.GetRestaurantByID).Methods("GET")
	return r
}
