package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	statusCode := http.StatusOK
	data, err := h.usecase.GetProfile(r.Context(), userID)
	if err != nil {
		statusCode = http.StatusInternalServerError
		w.WriteHeader(statusCode)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("[handler][GetProfile] fail parse json")
		statusCode = http.StatusInternalServerError
		w.WriteHeader(statusCode)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
