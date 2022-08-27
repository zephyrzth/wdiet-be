package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const restaurantIDBlock = 2

func (h *handler) GetRestaurantByID(w http.ResponseWriter, r *http.Request) {
	// restaurantIDPathParam := strings.Split(r.URL.Path, "/")[restaurantIDBlock]
	// restaurantID, _ = strconv.ParseInt(restaurantIDPathParam, 10, 64)

	statusCode := http.StatusOK
	data, err := h.usecase.GetAllRestaurant(r.Context())
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
