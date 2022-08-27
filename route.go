package main

import (
	"github.com/gorilla/mux"
	handlerHttp "github.com/zephyrzth/wdiet-be/handler/http"
	"github.com/zephyrzth/wdiet-be/usecase"
)

func newRouter(usecase usecase.UsecaseInterface) *mux.Router {
	handlerHttp := handlerHttp.New(usecase)

	r := mux.NewRouter()
	r.HandleFunc("/restaurants", handlerHttp.GetAllRestaurant).Methods("GET")
	r.HandleFunc("/restaurants/{id:[a-zA-Z0-9]+}", handlerHttp.GetRestaurantByID).Methods("GET")
	r.HandleFunc("/register", handlerHttp.Register).Methods("POST")
	r.HandleFunc("/login", handlerHttp.Login).Methods("POST")
	r.HandleFunc("/profile/{id:[a-zA-Z0-9]+}", handlerHttp.GetProfile).Methods("GET")
	r.HandleFunc("/usermenu", handlerHttp.InsertUserMenu).Methods("POST")

	return r
}
