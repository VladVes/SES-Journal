package handler

import (
	"encoding/json"

	"net/http"

	"github.com/VladVes/SES-Journal/internal/logger"
	"github.com/VladVes/SES-Journal/internal/services"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	userService := services.NewUserServise()
	result, err := userService.GetUsersList()
	if err != nil {
		logger.Log.WithError(err).Fatal("error while getting users list")
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		logger.Log.WithError(err).Fatal("error while users list json serialize")
	}

	w.Write([]byte(jsonData))
}
