package dto

type WriteAlertRequest struct {
	AlertID     string `json:"alert_id" validate:"required,alphanum"`
	ServiceID   string `json:"service_id" validate:"required"`
	ServiceName string `json:"service_name" validate:"required"`
	Model       string `json:"model" validate:"required"`
	AlertType   string `json:"alert_type" validate:"required"`
	AlertTS     string `json:"alert_ts" validate:"required,numeric"`
	Severity    string `json:"severity" validate:"required"`
	TeamSlack   string `json:"team_slack" validate:"required"`
}

type WriteAlertResponse struct {
	AlertID string `json:"alert_id"`
	Error   string `json:"error"`
}

type ReadAlertsRequest struct {
	ServiceID string `json:"service_id" validate:"required"`
	StartTS   string `json:"start_ts" validate:"required"`
	EndTS     string `json:"end_ts"  validate:"required"`
}

type ReadAlertsResponse struct {
	ServiceID   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	Alerts      []struct {
		AlertID   string `json:"alert_id"`
		Model     string `json:"model"`
		AlertType string `json:"alert_type"`
		AlertTS   string `json:"alert_ts"`
		Severity  string `json:"severity"`
		TeamSlack string `json:"team_slack"`
	} `json:"alerts"`
}

type ErrorResponse struct {
	AlertID string `json:"alert_id"`
	Error   string `json:"error"`
}
