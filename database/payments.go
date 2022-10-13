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

func (r *impl) GetDistinctPayments(ctx context.Context) ([][]string, error) {

	type pair struct {
		EventID    string `gorm:"column:event_id"`
		TicketName string `gorm:"column:ticket_name"`
	}

	data := []pair{}

	err := r.db.Raw("SELECT event_id, ticket_name FROM payments WHERE status = ? GROUP BY event_id, ticket_name", models.Waiting).Scan(&data).Error
	if err != nil {
		return nil, err
	}

	resp := [][]string{}
	for _, v := range data {
		resp = append(resp, []string{v.EventID, v.TicketName})
	}

	return resp, nil
}
