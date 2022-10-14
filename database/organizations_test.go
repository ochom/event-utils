package database_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ochom/event-utils/models"
	"github.com/stretchr/testify/require"
)

func Test_impl_CreateOrUpdateOrganization(t *testing.T) {
	r := initDB(t)
	data := &models.Organization{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateOrganization(data)
	require.NoError(t, err)
}

func Test_impl_DeleteOrganization(t *testing.T) {
	r := initDB(t)
	data := &models.Organization{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateOrganization(data)
	require.NoError(t, err)
	err = r.DeleteOrganization(data)
	require.NoError(t, err)
	data, err = r.GetOrganization(data)
	require.Error(t, err)
	require.Nil(t, data)
}

func Test_impl_GetOrganization(t *testing.T) {
	r := initDB(t)
	data := &models.Organization{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateOrganization(data)
	require.NoError(t, err)
	data, err = r.GetOrganization(data)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func Test_impl_GetOrganizations(t *testing.T) {
	r := initDB(t)
	data := &models.Organization{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateOrganization(data)
	require.NoError(t, err)
	many, err := r.GetOrganizations(data)
	require.NoError(t, err)
	require.NotNil(t, many)
	require.Len(t, many, 1)
}
