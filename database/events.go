package database

import (
	"time"

	"github.com/ochom/event-utils/models"
)

func (r *impl) CreateOrUpdateEvent(data *models.Event) error {
	return r.db.Save(data).Error
}

func (r *impl) DeleteEvent(query *models.Event) error {
	return r.db.Where(query).Delete(&models.Event{}).Error
}

func (r *impl) GetEvent(query *models.Event) (*models.Event, error) {
	var data models.Event
	if err := r.db.Where(query).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *impl) GetEvents(query *models.Event) ([]*models.Event, error) {
	data := []*models.Event{}
	if err := r.db.Where(query).Order("created_at desc").Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *impl) GetActiveEvents() ([]*models.Event, error) {
	data := []*models.Event{}
	if err := r.db.Where("end_time > ?", time.Now()).Order("created_at desc").Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
