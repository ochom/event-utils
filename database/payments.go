package database

import (
	"github.com/ochom/event-utils/models"
)

// EventGroup ...
type EventGroup struct {
	EventID    string `gorm:"column:event_id"`
	TicketName string `gorm:"column:ticket_name"`
}

func (r *impl) CreateOrUpdatePayment(data *models.Payment) error {
	err := r.db.Save(data).Error
	return err
}

func (r *impl) DeletePayment(query *models.Payment) error {
	err := r.db.Where(query).Delete(&models.Payment{}).Error
	return err
}

func (r *impl) GetPayment(query *models.Payment) (*models.Payment, error) {
	var data models.Payment
	if err := r.db.Where(query).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *impl) GetPayments(query *models.Payment) ([]*models.Payment, error) {
	data := []*models.Payment{}
	err := r.db.Where(query).Order("created_at desc").Find(&data).Error
	return data, err
}

func (r *impl) GetDistinctPayments() ([]*EventGroup, error) {
	data := []*EventGroup{}
	err := r.db.Table("payments").
		Select("event_id, ticket_name").
		Where("status = ?", models.Waiting).
		Group("event_id, ticket_name").
		Find(&data).Error

	return data, err
}
