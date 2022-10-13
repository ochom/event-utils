package database

import (
	"context"
	"fmt"
	"time"

	"github.com/ochom/event-utils/models"
)

func (r *impl) CreateOrUpdateEvent(ctx context.Context, data models.Event) error {
	return r.db.Save(&data).Error
}

func (r *impl) DeleteEvent(ctx context.Context, query models.Event) error {
	return r.db.Where(query).Delete(&models.Event{}).Error
}

func (r *impl) GetEvent(ctx context.Context, query models.Event) (*models.Event, error) {
	data := []*models.Event{}
	if err := r.db.Where(query).Find(&data).Error; err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("events data matching query not found")
	}

	return data[0], nil
}

func (r *impl) GetEvents(ctx context.Context, query models.Event) ([]*models.Event, error) {
	data := []*models.Event{}
	if err := r.db.Where(query).Order("created_at desc").Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *impl) GetActiveEvents(ctx context.Context) ([]*models.Event, error) {
	data := []*models.Event{}
	if err := r.db.Where("end_time > ?", time.Now()).Order("created_at desc").Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
