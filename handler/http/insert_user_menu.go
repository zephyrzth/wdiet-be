package http

import (
	"encoding/json"
	"fmt"
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

	warning, err := h.usecase.InsertUserMenu(r.Context(), requestData)
	if err != nil {
		statusCode = http.StatusInternalServerError
		w.WriteHeader(statusCode)
		return
	}

	jsonData, err := json.Marshal(warning)
	if err != nil {
		fmt.Println("[handler][InsertUserMenu] fail parse json")
		statusCode = http.StatusInternalServerError
		w.WriteHeader(statusCode)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(jsonData)
}
