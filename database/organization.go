package database

import (
	"context"
	"fmt"

	"github.com/ochom/event-utils/models"
)

func (r *impl) CreateOrUpdateOrganization(ctx context.Context, data models.Organization) error {
	err := r.db.Save(&data).Error
	return err
}

func (r *impl) DeleteOrganization(ctx context.Context, query models.Organization) error {
	err := r.db.Where(query).Delete(&models.Organization{}).Error
	return err
}

func (r *impl) GetOrganization(ctx context.Context, query models.Organization) (*models.Organization, error) {
	data := []*models.Organization{}
	err := r.db.Where(query).Find(&data).Error
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("organization data matching query not found")
	}

	return data[0], nil
}

func (r *impl) GetOrganizations(ctx context.Context, query models.Organization) ([]*models.Organization, error) {
	data := []*models.Organization{}
	err := r.db.Where(query).Order("created_at desc").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
