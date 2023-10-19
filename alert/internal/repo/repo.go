package repo

import "github.com/ovas33710/wm/alert/internal/models"

type Repo interface {
	Save(service models.Service, alert *models.Alert) error
	Get(serviceID string, startTS, endTS int) (models.Service, error)
}
