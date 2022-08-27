package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *handler) GetAllRestaurant(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK
	data, err := h.usecase.GetAllRestaurant(r.Context())
	if err != nil {
		statusCode = http.StatusInternalServerError
		w.WriteHeader(statusCode)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("[handler][GetAllRestaurant] fail parse json")
		statusCode = http.StatusInternalServerError
		w.WriteHeader(statusCode)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
