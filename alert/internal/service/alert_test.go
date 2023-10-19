package service

import (
	"testing"

	"github.com/ovas33710/wm/alert/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Save(s models.Service, alert *models.Alert) error {
	args := m.Called(s, alert)
	return args.Error(0)
}

func (m *MockRepo) Get(serviceID string, startTS, endTS int) (models.Service, error) {
	args := m.Called(serviceID, startTS, endTS)
	return args.Get(0).(models.Service), args.Error(1)
}

func TestAlertService_WriteAlert(t *testing.T) {
	mockRepo := new(MockRepo)
	a := AlertService{
		db: mockRepo,
	}

	service := models.Service{
		ServiceID: "service1",
	}

	alert := models.Alert{
		AlertID: "alert1",
	}

	mockRepo.On("Save", service, &alert).Return(nil)

	err := a.WriteAlert(service, alert)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAlertService_GetAlerts(t *testing.T) {
	mockRepo := new(MockRepo)
	a := AlertService{
		db: mockRepo,
	}

	service := models.Service{
		ServiceID: "service1",
	}

	alert := models.Alert{
		AlertID: "alert1",
	}

	service.Alerts = append(service.Alerts, alert)

	mockRepo.On("Get", "service1", 0, 0).Return(service, nil)

	savedService, err := a.GetAlerts("service1", 0, 0)

	assert.Nil(t, err)
	assert.Equal(t, "service1", savedService.ServiceID)
	assert.Equal(t, 1, len(savedService.Alerts))
	assert.Equal(t, "alert1", savedService.Alerts[0].AlertID)
	mockRepo.AssertExpectations(t)
}
