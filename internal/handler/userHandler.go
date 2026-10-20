package handler

import (
	"encoding/json"

	"net/http"

	"github.com/VladVes/SES-Journal/internal/logger"
	"github.com/VladVes/SES-Journal/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userv *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userv,
	}
}

func (h *UserHandler) HandleUsersList(w http.ResponseWriter, r *http.Request) {
	result, err := h.userService.GetUsersList()
	if err != nil {
		logger.Log.WithError(err).Error("error while getting users list")
		http.Error(w, "internal server error while getting users list", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		logger.Log.WithError(err).Error("error while users list json serialize")
		http.Error(w, "internal server error while users list json serialize", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(jsonData))
	w.Write(jsonData)
}
