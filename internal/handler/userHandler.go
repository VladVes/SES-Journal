package handler

import (
	"encoding/json"

	"net/http"

	"github.com/VladVes/SES-Journal/internal/logger"
	"github.com/VladVes/SES-Journal/internal/services"
	"github.com/VladVes/SES-Journal/internal/dto"
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
	w.Write(jsonData)
}

func (h *UserHandler) HandleCreatUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	var req dto.CreateUserRequest 
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	defer r.Body.Close()

	if err := dec.Decode(&req); err != nil {
		logger.Log.WithError(err).Warn("invalid json body for create user")
		http.Error(w, "bad reaquest", http.StatusBadRequest)
		return
	}

	// TODO: implement CreateUser
	// newUser, err := h.userService.CreateUser(req)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}