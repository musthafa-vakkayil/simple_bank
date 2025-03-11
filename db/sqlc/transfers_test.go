package db

import (
	"context"
	"testing"

	"github.com/musthafa-vakkayil/swiss_bank_server/util"
	"github.com/stretchr/testify/require"
)

func createRandomtransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	require.Equal(t, args.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomtransfer(t)
}
