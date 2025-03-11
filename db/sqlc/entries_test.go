package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/musthafa-vakkayil/swiss_bank_server/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)

	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry := createRandomEntry(t)

	entry1, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)

	require.NotEmpty(t, entry1)

	require.Equal(t, entry.ID, entry.ID)
	require.WithinDuration(t, entry.CreatedAt, entry.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T) {
	entry := createRandomEntry(t)

	updatedAmount := util.RandomMoney()

	entry1, err := testQueries.UpdateEntry(context.Background(), UpdateEntryParams{
		ID:     entry.ID,
		Amount: updatedAmount,
	})

	require.NoError(t, err)

	require.NotEmpty(t, entry1)

	require.Equal(t, entry.ID, entry1.ID)

	require.Equal(t, updatedAmount, entry1.Amount)
}

func TestDeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entry.ID)

	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)

	err = testQueries.DeleteAccount(context.Background(), entry.AccountID)
	require.NoError(t, err)
}

func TestListEntriies(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, account := range entries {
		require.NotEmpty(t, account)
	}
}
