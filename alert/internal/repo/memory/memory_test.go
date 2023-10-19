package memory

import (
	"testing"

	"github.com/ovas33710/wm/alert/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestInMemory_Save(t *testing.T) {
	repo := New()

	service := models.Service{
		ServiceID:   "service1",
		ServiceName: "service_name1",
		Alerts:      []models.Alert{},
	}

	alert := &models.Alert{
		AlertID:   "alert1",
		Model:     "test_model",
		AlertType: "test_alert_type",
		AlertTS:   1,
		Severity:  "test_severity",
		TeamSlack: "test_slack_ch",
	}

	// Test saving a new alert
	err := repo.Save(service, alert)
	assert.Nil(t, err)

	// Validate the alert was saved
	savedService, err := repo.Get("service1", 0, 5)
	assert.Nil(t, err)
	assert.Equal(t, "service1", savedService.ServiceID)
	assert.Equal(t, "service_name1", savedService.ServiceName)
	assert.Equal(t, 1, len(savedService.Alerts))
	assert.Equal(t, "alert1", savedService.Alerts[0].AlertID)
	assert.Equal(t, "test_model", savedService.Alerts[0].Model)
	assert.Equal(t, "test_alert_type", savedService.Alerts[0].AlertType)
	assert.Equal(t, 1, savedService.Alerts[0].AlertTS)
	assert.Equal(t, "test_severity", savedService.Alerts[0].Severity)
	assert.Equal(t, "test_slack_ch", savedService.Alerts[0].TeamSlack)
}

func TestInMemory_Get(t *testing.T) {
	repo := New()

	service := models.Service{
		ServiceID: "service1",
		Alerts:    []models.Alert{},
	}

	alert := &models.Alert{
		AlertID: "alert1",
		AlertTS: 100,
	}

	err := repo.Save(service, alert)
	assert.Nil(t, err)

	// Test getting a non-existing service
	_, err = repo.Get("nonExistingService", 0, 0)
	assert.Equal(t, ErrServiceNotFound, err)

	// Test getting an existing service
	savedService, err := repo.Get("service1", 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, "service1", savedService.ServiceID)

	// Test getting alerts in a timeframe
	savedService, err = repo.Get("service1", 90, 110)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(savedService.Alerts))
	assert.Equal(t, "alert1", savedService.Alerts[0].AlertID)
}

func TestInMemory_GetByTS(t *testing.T) {
	repo := New()

	service := models.Service{
		ServiceID: "service1",
		Alerts:    []models.Alert{},
	}

	alert1 := &models.Alert{
		AlertID: "alert1",
		AlertTS: 50,
	}

	alert2 := &models.Alert{
		AlertID: "alert2",
		AlertTS: 100,
	}

	alert3 := &models.Alert{
		AlertID: "alert3",
		AlertTS: 150,
	}

	err := repo.Save(service, alert1)
	assert.Nil(t, err)
	err = repo.Save(service, alert2)
	assert.Nil(t, err)
	err = repo.Save(service, alert3)
	assert.Nil(t, err)

	// Test getting a non-existing service
	_, err = repo.Get("nonExistingService", 0, 0)
	assert.Equal(t, ErrServiceNotFound, err)

	// Test getting an existing service
	savedService, err := repo.Get("service1", 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, "service1", savedService.ServiceID)

	// Test getting alerts in a timeframe
	savedService, err = repo.Get("service1", 90, 110)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(savedService.Alerts))
	assert.Equal(t, "alert2", savedService.Alerts[0].AlertID)
}

func TestInMemory_AlertAlreadyExists(t *testing.T) {
	repo := New()

	service := models.Service{
		ServiceID: "service1",
		Alerts:    []models.Alert{},
	}

	alert1 := &models.Alert{
		AlertID: "alert1",
	}

	alert2 := &models.Alert{
		AlertID: "alert1",
	}

	err := repo.Save(service, alert1)
	assert.Nil(t, err)
	err = repo.Save(service, alert2)
	assert.Equal(t, ErrAlertAlreadyExists, err)
}
