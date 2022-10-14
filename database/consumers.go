package database

import (
	"github.com/ochom/event-utils/models"
)

func (r *impl) CreateOrUpdateConsumer(data *models.Consumer) error {
	err := r.db.Save(data).Error
	return err
}

func (r *impl) DeleteConsumer(query *models.Consumer) error {
	err := r.db.Where(query).Delete(&models.Consumer{}).Error
	return err
}

func (r *impl) GetConsumer(query *models.Consumer) (*models.Consumer, error) {
	var data models.Consumer
	if err := r.db.Where(query).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *impl) GetConsumers(query *models.Consumer) ([]*models.Consumer, error) {
	data := []*models.Consumer{}
	err := r.db.Where(query).Order("created_at desc").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
