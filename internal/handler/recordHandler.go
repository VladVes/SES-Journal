package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/VladVes/SES-Journal/internal/dto"
	"github.com/VladVes/SES-Journal/internal/logger"
	"github.com/VladVes/SES-Journal/internal/repository"
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

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", "/records")

	if err := json.NewEncoder(w).Encode(&recordsList); err != nil {
		logger.Log.WithError(err).Error("failed to encode records list")
		http.Error(w, "internal server error while json encodind records list", http.StatusInternalServerError)
		return
	}
}

func (rh *RecordHandler) HandleGetRecordByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSpace(r.PathValue("id"))
	if id == "" {
		logger.Log.WithField("id", id).Warn("failed to get record by id: id is missing or empty")
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	// добавить проверку формата id, соответствующего тому, как он генерируется
	// (если ID — UUID, использовать uuid.Parse; если произвольная строка из БД — хотя бы разумный лимит длины, например len(id) > 128 → 400).

	record, err := rh.service.GetRecordByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFoundByID) {
			http.Error(w, "no record by given id found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed get record by id", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	recordResponse := dto.GetRecordResponse{
		ID:        record.ID,
		RecNumber: record.RecNumber,
		RecType:   record.RecType,

		VisitorName:      record.VisitorName,
		VisitPurpose:     record.VisitPurpose,
		OrganizationName: record.OrganizationName,

		MalfunctionDesc:       record.MalfunctionDesc,
		MalfunctionCause:      record.MalfunctionCause,
		MalfunctionResolution: record.MalfunctionResolution,
		MalfunctionStatus:     record.MalfunctionStatus,

		UserID:   record.UserID,
		UserName: record.User.Name,

		StartTime: record.StartTime,
		EndTime:   record.EndTime,
	}

	if err := json.NewEncoder(w).Encode(&recordResponse); err != nil {
		logger.Log.WithError(err).Error("record json encoding failed")
		http.Error(w, "record json encoding failed", http.StatusInternalServerError)
	}
}
