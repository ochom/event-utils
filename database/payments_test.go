package database_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ochom/event-utils/models"
	"github.com/stretchr/testify/require"
)

func Test_impl_CreateOrUpdatePayment(t *testing.T) {
	r := initDB(t)
	data := &models.Payment{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdatePayment(data)
	require.NoError(t, err)
}

func Test_impl_DeletePayment(t *testing.T) {
	r := initDB(t)
	data := &models.Payment{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdatePayment(data)
	require.NoError(t, err)
	err = r.DeletePayment(data)
	require.NoError(t, err)
	data, err = r.GetPayment(data)
	require.Error(t, err)
	require.Nil(t, data)
}

func Test_impl_GetPayment(t *testing.T) {
	r := initDB(t)
	data := &models.Payment{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdatePayment(data)
	require.NoError(t, err)
	data, err = r.GetPayment(data)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func Test_impl_GetPayments(t *testing.T) {
	r := initDB(t)
	data := &models.Payment{
		ID: uuid.NewString(),
	}
	err := r.CreateOrUpdatePayment(data)
	require.NoError(t, err)
	many, err := r.GetPayments(data)
	require.NoError(t, err)
	require.NotNil(t, many)
	require.Len(t, many, 1)
}

func Test_impl_GetDistinctPayments(t *testing.T) {
	r := initDB(t)
	eventID1, eventID2, eventID3 := uuid.NewString(), uuid.NewString(), uuid.NewString()
	payments := []*models.Payment{
		{
			ID:         uuid.NewString(),
			EventID:    eventID1,
			Status:     models.Waiting,
			TicketName: "ticket1",
		},
		{
			ID:         uuid.NewString(),
			EventID:    eventID1,
			Status:     models.Waiting,
			TicketName: "ticket1",
		},
		{
			ID:         uuid.NewString(),
			EventID:    eventID1,
			Status:     models.Completed,
			TicketName: "ticket1",
		},
		{
			ID:         uuid.NewString(),
			EventID:    eventID2,
			Status:     models.Waiting,
			TicketName: "ticket1",
		},
		{
			ID:         uuid.NewString(),
			EventID:    eventID2,
			Status:     models.Waiting,
			TicketName: "ticket1",
		},
		{
			ID:         uuid.NewString(),
			EventID:    eventID2,
			Status:     models.Waiting,
			TicketName: "ticket2",
		},
		{
			ID:         uuid.NewString(),
			EventID:    eventID3,
			Status:     models.Waiting,
			TicketName: "ticket1",
		},
		{
			ID:         uuid.NewString(),
			EventID:    eventID3,
			Status:     models.Waiting,
			TicketName: "ticket2",
		},
		{
			ID:         uuid.NewString(),
			EventID:    eventID3,
			Status:     models.Waiting,
			TicketName: "ticket3",
		},
	}

	for _, p := range payments {
		err := r.CreateOrUpdatePayment(p)
		require.NoError(t, err)
	}

	many, err := r.GetDistinctPayments()
	require.NoError(t, err)
	require.NotNil(t, many)
	require.Len(t, many, 6)
}
