package database_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ochom/event-utils/models"
	"github.com/stretchr/testify/require"
)

func Test_impl_CreateOrUpdateConsumer(t *testing.T) {
	r := initDB(t)
	data := &models.Consumer{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateConsumer(data)
	require.NoError(t, err)
}

func Test_impl_DeleteConsumer(t *testing.T) {
	r := initDB(t)
	data := &models.Consumer{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateConsumer(data)
	require.NoError(t, err)
	err = r.DeleteConsumer(data)
	require.NoError(t, err)
	data, err = r.GetConsumer(data)
	require.Error(t, err)
	require.Nil(t, data)
}

func Test_impl_GetConsumer(t *testing.T) {
	r := initDB(t)
	data := &models.Consumer{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateConsumer(data)
	require.NoError(t, err)
	data, err = r.GetConsumer(data)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func Test_impl_GetConsumers(t *testing.T) {
	r := initDB(t)
	data := &models.Consumer{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateConsumer(data)
	require.NoError(t, err)
	many, err := r.GetConsumers(data)
	require.NoError(t, err)
	require.NotNil(t, many)
	require.Len(t, many, 1)
}
