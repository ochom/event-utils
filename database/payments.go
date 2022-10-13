package database

import (
	"context"
	"fmt"

	"github.com/ochom/event-utils/models"
)

func (r *impl) CreateOrUpdatePayment(ctx context.Context, data models.Payment) error {
	err := r.db.Save(&data).Error
	return err
}

func (r *impl) DeletePayment(ctx context.Context, query models.Payment) error {
	err := r.db.Where(query).Delete(&models.Payment{}).Error
	return err
}

func (r *impl) GetPayment(ctx context.Context, query models.Payment) (*models.Payment, error) {
	data := []*models.Payment{}
	err := r.db.Where(query).Find(&data).Error
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("payment data matching query not found")
	}

	return data[0], nil
}

func (r *impl) GetPayments(ctx context.Context, query models.Payment) ([]*models.Payment, error) {
	data := []*models.Payment{}
	err := r.db.Where(query).Order("created_at desc").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
