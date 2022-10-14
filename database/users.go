package database

import (
	"github.com/ochom/event-utils/models"
)

func (r *impl) CreateOrUpdateUser(data *models.User) error {
	err := r.db.Save(data).Error
	return err
}

func (r *impl) DeleteUser(query *models.User) error {
	err := r.db.Where(query).Delete(&models.User{}).Error
	return err
}

func (r *impl) GetUser(query *models.User) (*models.User, error) {
	var data models.User
	if err := r.db.Where(query).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *impl) GetUsers(query *models.User) ([]*models.User, error) {
	data := []*models.User{}
	err := r.db.Where(query).Order("created_at desc").Find(&data).Error
	return data, err
}
