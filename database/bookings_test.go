package database_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ochom/event-utils/models"
	"github.com/stretchr/testify/require"
)

func Test_impl_CreateOrUpdateBooking(t *testing.T) {
	r := initDB(t)
	data := &models.Booking{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateBooking(data)
	require.NoError(t, err)
}

func Test_impl_DeleteBooking(t *testing.T) {
	r := initDB(t)
	data := &models.Booking{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateBooking(data)
	require.NoError(t, err)
	err = r.DeleteBooking(data)
	require.NoError(t, err)
	data, err = r.GetBooking(data)
	require.Error(t, err)
	require.Nil(t, data)
}

func Test_impl_GetBooking(t *testing.T) {
	r := initDB(t)
	data := &models.Booking{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateBooking(data)
	require.NoError(t, err)
	data, err = r.GetBooking(data)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func Test_impl_GetBookings(t *testing.T) {
	r := initDB(t)
	data := &models.Booking{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdateBooking(data)
	require.NoError(t, err)
	many, err := r.GetBookings(data)
	require.NoError(t, err)
	require.NotNil(t, many)
	require.Len(t, many, 1)
}
