package main

import (
	"net/http"

	"github.com/gorilla/mux"
	handlerHttp "github.com/zephyrzth/wdiet-be/handler/http"
	"github.com/zephyrzth/wdiet-be/usecase"
)

func newRouter(usecase usecase.UsecaseInterface) *mux.Router {
	handlerHttp := handlerHttp.New(usecase)

	r := mux.NewRouter()
	r.HandleFunc("/restaurants", wrapHandler(handlerHttp.GetAllRestaurant)).Methods("GET", "OPTIONS")
	r.HandleFunc("/restaurants/{id:[a-zA-Z0-9]+}", wrapHandler(handlerHttp.GetRestaurantByID)).Methods("GET", "OPTIONS")
	r.HandleFunc("/register", wrapHandler(handlerHttp.Register)).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", wrapHandler(handlerHttp.Login)).Methods("POST", "OPTIONS")
	r.HandleFunc("/profile/{id:[a-zA-Z0-9]+}", wrapHandler(handlerHttp.GetProfile)).Methods("GET", "OPTIONS")
	r.HandleFunc("/usermenu", wrapHandler(handlerHttp.InsertUserMenu)).Methods("POST", "OPTIONS")

	return r
}

func wrapHandler(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if (*r).Method == "OPTIONS" {
			return
		}
		handler(w, r)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	(*w).Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	(*w).Header().Set("Access-Control-Expose-Headers", "Authorization")
}
