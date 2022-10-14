package database_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ochom/event-utils/models"
	"github.com/stretchr/testify/require"
)

func Test_impl_CreateOrUpdateUser(t *testing.T) {
	r := initDB(t)
	data := &models.User{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateUser(data)
	require.NoError(t, err)
}

func Test_impl_DeleteUser(t *testing.T) {
	r := initDB(t)
	data := &models.User{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateUser(data)
	require.NoError(t, err)
	err = r.DeleteUser(data)
	require.NoError(t, err)
	data, err = r.GetUser(data)
	require.Error(t, err)
	require.Nil(t, data)
}

func Test_impl_GetUser(t *testing.T) {
	r := initDB(t)
	data := &models.User{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateUser(data)
	require.NoError(t, err)
	data, err = r.GetUser(data)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func Test_impl_GetUsers(t *testing.T) {
	r := initDB(t)
	data := &models.User{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateUser(data)
	require.NoError(t, err)
	many, err := r.GetUsers(data)
	require.NoError(t, err)
	require.NotNil(t, many)
	require.Len(t, many, 1)
}
