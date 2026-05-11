package handler

import (
	"net/http"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ses journal and events log"))
}
