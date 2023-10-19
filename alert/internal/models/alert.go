package models

type Service struct {
	ServiceID   string  `json:"service_id"`
	ServiceName string  `json:"service_name"`
	Alerts      []Alert `json:"alerts"`
}

type Alert struct {
	AlertID   string `json:"alert_id"`
	Model     string `json:"model"`
	AlertType string `json:"alert_type"`
	AlertTS   int    `json:"alert_ts"`
	Severity  string `json:"severity"`
	TeamSlack string `json:"team_slack"`
}
