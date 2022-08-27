package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *handler) GetRestaurantByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	restaurantID := vars["id"]

	statusCode := http.StatusOK
	data, err := h.usecase.GetRestaurantByID(r.Context(), restaurantID)
	if err != nil {
		statusCode = http.StatusInternalServerError
		w.WriteHeader(statusCode)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("[handler][GetRestaurantByID] fail parse json")
		statusCode = http.StatusInternalServerError
		w.WriteHeader(statusCode)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(jsonData)
}
