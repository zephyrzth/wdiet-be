package http

import (
	"encoding/json"
	"net/http"

	"github.com/zephyrzth/wdiet-be/model"
)

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK
	var p model.User
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		statusCode = http.StatusBadRequest
		w.WriteHeader(statusCode)
		return
	}

	userID, err := h.usecase.Register(r.Context(), p)
	if err != nil {
		statusCode = http.StatusInternalServerError
	}
	jsonData, _ := json.Marshal(map[string]string{"id": userID})

	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
