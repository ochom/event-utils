package database

import (
	"github.com/ochom/event-utils/models"
)

func (r *impl) CreateOrUpdateOrganization(data *models.Organization) error {
	err := r.db.Save(data).Error
	return err
}

func (r *impl) DeleteOrganization(query *models.Organization) error {
	err := r.db.Where(query).Delete(&models.Organization{}).Error
	return err
}

func (r *impl) GetOrganization(query *models.Organization) (*models.Organization, error) {
	var data models.Organization
	if err := r.db.Where(query).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *impl) GetOrganizations(query *models.Organization) ([]*models.Organization, error) {
	data := []*models.Organization{}
	err := r.db.Where(query).Order("created_at desc").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
