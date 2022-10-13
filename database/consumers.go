package database

import (
	"context"
	"fmt"

	"github.com/ochom/event-utils/models"
)

func (r *impl) CreateOrUpdateConsumer(ctx context.Context, data models.Consumer) error {
	err := r.db.Save(&data).Error
	return err
}

func (r *impl) DeleteConsumer(ctx context.Context, query models.Consumer) error {
	err := r.db.Where(query).Delete(&models.Consumer{}).Error
	return err
}

func (r *impl) GetConsumer(ctx context.Context, query models.Consumer) (*models.Consumer, error) {
	data := []*models.Consumer{}
	err := r.db.Where(query).Find(&data).Error
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("user data matching query not found")
	}

	return data[0], nil
}

func (r *impl) GetConsumers(ctx context.Context, query models.Consumer) ([]*models.Consumer, error) {
	data := []*models.Consumer{}
	err := r.db.Where(query).Order("created_at desc").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
