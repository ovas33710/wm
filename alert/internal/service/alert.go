package service

import (
	"github.com/ovas33710/wm/alert/internal/models"
	"github.com/ovas33710/wm/alert/internal/repo"
	"github.com/ovas33710/wm/alert/internal/repo/memory"
)

type AlertServiceInterface interface {
	WriteAlert(service models.Service, alert models.Alert) error
	GetAlerts(serviceId string, startTS, endTS int) (models.Service, error)
}

type AlertServiceConfiguration func(as *AlertService) error

type AlertService struct {
	db repo.Repo
}

func NewAlertService(cfgs ...AlertServiceConfiguration) (*AlertService, error) {
	as := &AlertService{}
	for _, cfg := range cfgs {
		err := cfg(as)
		if err != nil {
			return nil, err
		}
	}
	return as, nil
}

func WithRepo(repo repo.Repo) AlertServiceConfiguration {
	return func(as *AlertService) error {
		as.db = repo
		return nil
	}
}

func WithInMemoryRepo() AlertServiceConfiguration {
	imr := memory.New()
	return WithRepo(imr)
}

func (s *AlertService) WriteAlert(service models.Service, alert models.Alert) error {
	if err := s.db.Save(service, &alert); err != nil {
		return err
	}
	return nil
}

func (s *AlertService) GetAlerts(serviceId string, startTS, endTS int) (models.Service, error) {
	service, err := s.db.Get(serviceId, startTS, endTS)
	if err != nil {
		return models.Service{}, err
	}
	return service, nil
}
