package http

import (
	"encoding/json"
	"net/http"

	"github.com/zephyrzth/wdiet-be/model"
)

func (h *handler) InsertUserMenu(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK

	var requestData model.InsertUserMenu
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		statusCode = http.StatusBadRequest
		w.WriteHeader(statusCode)
		return
	}

	err = h.usecase.InsertUserMenu(r.Context(), requestData)
	if err != nil {
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
}
