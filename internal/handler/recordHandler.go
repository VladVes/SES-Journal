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

func (rh *RecordHandler) HandleCreateRecord(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var createRecReq dto.CreateRecordRequest

	if err := dec.Decode(&createRecReq); err != nil {
		logger.Log.Warning("failed to decode create record req")
		http.Error(w, "failed to decode create record request", http.StatusBadRequest)
		return
	}

	createRecReq.Normalize()
	if err := createRecReq.Validate(); err != nil {
		logger.Log.WithError(err).Warning("failed to validate create record req")
		http.Error(w, "failed to validate create record request", http.StatusBadRequest)
		return
	}

	newRecord, err := rh.service.CreateNewRecord(&createRecReq)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			logger.Log.WithError(err).WithField("UserID", createRecReq.UserID).Warning("failde to create a new record")
			http.Error(w, "user with id in UserID field of record request not found", http.StatusUnprocessableEntity)
			return
		}

		logger.Log.WithError(err).Error("failde to check is user exists")
		http.Error(w, "failde to check is user exists", http.StatusInternalServerError)
		return
	}

	resp := dto.CreateRecordResponse{

		ID:        newRecord.ID,
		RecNumber: newRecord.RecNumber,
		RecType:   newRecord.RecType,

		VisitorName:      newRecord.VisitorName,
		VisitPurpose:     newRecord.VisitPurpose,
		OrganizationName: newRecord.OrganizationName,

		MalfunctionDesc:       newRecord.MalfunctionDesc,
		MalfunctionCause:      newRecord.MalfunctionCause,
		MalfunctionResolution: newRecord.MalfunctionResolution,
		MalfunctionStatus:     newRecord.MalfunctionStatus,

		UserID: newRecord.UserID,

		StartTime: newRecord.StartTime,
		EndTime:   newRecord.StartTime,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", "records/"+newRecord.ID)
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Log.WithError(err).Error("failed to encode new record response")
		http.Error(w, "failed to encode new record response", http.StatusInternalServerError)
		return
	}

}
