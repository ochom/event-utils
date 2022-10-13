package database

import (
	"context"
	"fmt"

	"github.com/ochom/event-utils/models"
)

func (r *impl) CreateOrUpdateBooking(ctx context.Context, data models.Booking) error {
	err := r.db.Save(&data).Error
	return err
}

func (r *impl) DeleteBooking(ctx context.Context, query models.Booking) error {
	err := r.db.Where(query).Delete(&models.Booking{}).Error
	return err
}

func (r *impl) GetBooking(ctx context.Context, query models.Booking) (*models.Booking, error) {
	data := []*models.Booking{}
	err := r.db.Where(query).Find(&data).Error
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("booking data matching query not found")
	}

	return data[0], nil
}

func (r *impl) GetBookings(ctx context.Context, query models.Booking) ([]*models.Booking, error) {
	data := []*models.Booking{}
	err := r.db.Where(query).Order("created_at desc").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
