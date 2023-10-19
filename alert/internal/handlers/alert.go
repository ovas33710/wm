package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/ovas33710/wm/alert/internal/handlers/dto"
	"github.com/ovas33710/wm/alert/internal/models"
	alert "github.com/ovas33710/wm/alert/internal/service"

	"github.com/go-playground/validator"
)

var (
	ErrFailedToReadRequestBody     = errors.New("failed to read request body")
	ErrInvalidJSON                 = errors.New("invalid json")
	ErrInvalidInput                = errors.New("invalid input")
	ErrFailedToParseStartTS        = errors.New("failed to parse start time stamp")
	ErrFailedToParseEndTS          = errors.New("failed to parse start time stamp")
	ErrFailedToMarshalJSONResponse = errors.New("failed to marshal JSON response")
)

type AlertHandlerInterface interface {
	WriteAlert(w http.ResponseWriter, r *http.Request)
	ReadAlerts(w http.ResponseWriter, r *http.Request)
}

type AlertHandler struct {
	alertService alert.AlertServiceInterface
	validator    *validator.Validate
}

func NewAlertHandler() (AlertHandler, error) {
	alertService, err := alert.NewAlertService(alert.WithInMemoryRepo())
	if err != nil {
		return AlertHandler{}, nil
	}
	return AlertHandler{
		alertService: alertService,
		validator:    validator.New(),
	}, nil
}

// POST Request Handler (Write Alert)
func (ah AlertHandler) WriteAlert(w http.ResponseWriter, r *http.Request) {
	// 1. Parse the JSON request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, ErrFailedToReadRequestBody.Error(), http.StatusBadRequest)
		return
	}
	var req dto.WriteAlertRequest
	if err := json.Unmarshal(body, &req); err != nil {
		writeError(w, ErrInvalidJSON.Error(), http.StatusBadRequest)
		return
	}

	// validate request
	if err := ah.validator.Struct(req); err != nil {
		writeError(w, ErrInvalidInput.Error(), http.StatusBadRequest)
		return
	}

	// save alert to database
	alertTS, err := strconv.Atoi(req.AlertTS)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	service := models.Service{
		ServiceID:   req.ServiceID,
		ServiceName: req.ServiceName,
	}
	alert := models.Alert{
		AlertID:   req.AlertID,
		Model:     req.Model,
		AlertType: req.AlertType,
		AlertTS:   alertTS,
		Severity:  req.Severity,
	}
	if err := ah.alertService.WriteAlert(service, alert); err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// marhal response
	res, err := json.Marshal(dto.WriteAlertResponse{
		AlertID: alert.AlertID,
		Error:   "",
	})
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// GET Request Handler (Read Alerts)
func (ah AlertHandler) ReadAlerts(w http.ResponseWriter, r *http.Request) {
	// 1. Parse and validate query parameters
	params := r.URL.Query()
	query := dto.ReadAlertsRequest{
		ServiceID: params.Get("service_id"),
		StartTS:   params.Get("start_ts"),
		EndTS:     params.Get("end_ts"),
	}

	// validate
	if err := ah.validator.Struct(dto.ReadAlertsRequest{
		ServiceID: query.ServiceID,
		StartTS:   query.StartTS,
		EndTS:     query.EndTS,
	}); err != nil {
		writeError(w, ErrInvalidInput.Error(), http.StatusBadRequest)
		return
	}
	startTS, err := strconv.Atoi(query.StartTS)
	if err != nil {
		writeError(w, ErrFailedToParseStartTS.Error(), http.StatusInternalServerError)
		return
	}
	endTS, err := strconv.Atoi(query.EndTS)
	if err != nil {
		writeError(w, ErrFailedToParseEndTS.Error(), http.StatusInternalServerError)
		return
	}

	service, err := ah.alertService.GetAlerts(query.ServiceID, startTS, endTS)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(service)
	if err != nil {
		writeError(w, ErrFailedToMarshalJSONResponse.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func writeError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response, err := json.Marshal(dto.ErrorResponse{
		AlertID: "",
		Error:   message,
	})
	if err != nil {
		http.Error(w, ErrFailedToMarshalJSONResponse.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
