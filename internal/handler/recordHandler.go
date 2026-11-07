package handler

import (
	"encoding/json"
	"net/http"

	"github.com/VladVes/SES-Journal/internal/logger"
	"github.com/VladVes/SES-Journal/internal/services"
)

type RecordHandler struct {
	service *services.RecordService
}

func NewRecordHandler(rs *services.RecordService) *RecordHandler {
	return &RecordHandler{
		service: rs,
	}
}

func (rh *RecordHandler) HandleRecordsList(w http.ResponseWriter, r *http.Request) {

	recordsList, err := rh.service.GetRecordsList()
	if err != nil {
		logger.Log.WithError(err).Error("handler failed to get records list")
		http.Error(w, "internal server error while getting records list", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type:", "application/json")
	w.Header().Set("Location", "/records")

	if err := json.NewEncoder(w).Encode(&recordsList); err != nil {
		logger.Log.WithError(err).Error("failed to encode records list")
		http.Error(w, "internal server error while json encodind records list", http.StatusInternalServerError)
		return
	}
}
