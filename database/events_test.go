package database_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ochom/event-utils/models"
	"github.com/stretchr/testify/require"
)

func Test_impl_CreateOrUpdateEvent(t *testing.T) {
	r := initDB(t)
	data := &models.Event{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateEvent(data)
	require.NoError(t, err)
}

func Test_impl_DeleteEvent(t *testing.T) {
	r := initDB(t)
	data := &models.Event{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateEvent(data)
	require.NoError(t, err)
	err = r.DeleteEvent(data)
	require.NoError(t, err)
	data, err = r.GetEvent(data)
	require.Error(t, err)
	require.Nil(t, data)
}

func Test_impl_GetEvent(t *testing.T) {
	r := initDB(t)
	data := &models.Event{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateEvent(data)
	require.NoError(t, err)
	data, err = r.GetEvent(data)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func Test_impl_GetEvents(t *testing.T) {
	r := initDB(t)
	data := &models.Event{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateEvent(data)
	require.NoError(t, err)
	many, err := r.GetEvents(data)
	require.NoError(t, err)
	require.NotNil(t, many)
	require.Len(t, many, 1)
}
