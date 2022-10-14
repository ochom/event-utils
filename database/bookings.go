package database

import (
	"github.com/ochom/event-utils/models"
)

func (r *impl) CreateOrUpdateBooking(data *models.Booking) error {
	err := r.db.Save(data).Error
	return err
}

func (r *impl) DeleteBooking(query *models.Booking) error {
	err := r.db.Where(query).Delete(&models.Booking{}).Error
	return err
}

func (r *impl) GetBooking(query *models.Booking) (*models.Booking, error) {
	var data models.Booking
	if err := r.db.Where(query).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *impl) GetBookings(query *models.Booking) ([]*models.Booking, error) {
	data := []*models.Booking{}
	err := r.db.Where(query).Order("created_at desc").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
