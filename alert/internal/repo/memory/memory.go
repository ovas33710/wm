package memory

import (
	"errors"

	"github.com/ovas33710/wm/alert/internal/models"
)

var (
	ErrServiceNotFound    = errors.New("service not found")
	ErrAlertAlreadyExists = errors.New("alert already exists")
)

type InMemory struct {
	db map[string]models.Service
}

func New() *InMemory {
	return &InMemory{
		db: make(map[string]models.Service),
	}
}

func (m *InMemory) Save(s models.Service, alert *models.Alert) error {
	service, ok := m.db[s.ServiceID]
	if !ok {
		s.Alerts = append(s.Alerts, *alert)
		m.db[s.ServiceID] = s
		return nil
	}
	for _, v := range service.Alerts {
		if v.AlertID == alert.AlertID {
			return ErrAlertAlreadyExists
		}
	}
	service.Alerts = append(service.Alerts, *alert)
	m.db[s.ServiceID] = service
	return nil
}

func (m *InMemory) Get(serviceID string, startTS, endTS int) (models.Service, error) {
	service, ok := m.db[serviceID]
	if !ok {
		return models.Service{}, ErrServiceNotFound
	}
	var alerts []models.Alert
	for _, v := range service.Alerts {
		if startTS <= v.AlertTS && v.AlertTS <= endTS {
			alerts = append(alerts, v)
		}
	}
	return models.Service{
		ServiceID:   service.ServiceID,
		ServiceName: service.ServiceName,
		Alerts:      alerts,
	}, nil
}
