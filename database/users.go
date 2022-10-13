package database

import (
	"context"
	"fmt"

	"github.com/ochom/event-utils/models"
)

func (r *impl) CreateOrUpdateUser(ctx context.Context, data models.User) error {
	err := r.db.Save(&data).Error
	return err
}

func (r *impl) DeleteUser(ctx context.Context, query models.User) error {
	err := r.db.Where(query).Delete(&models.User{}).Error
	return err
}

func (r *impl) GetUser(ctx context.Context, query models.User) (*models.User, error) {
	data := []*models.User{}
	err := r.db.Where(query).Find(&data).Error
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("user data matching query not found")
	}

	return data[0], nil
}

func (r *impl) GetUsers(ctx context.Context, query models.User) ([]*models.User, error) {
	data := []*models.User{}
	err := r.db.Where(query).Order("created_at desc").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
