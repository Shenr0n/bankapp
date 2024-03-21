package db

import (
	"context"
	"testing"
	"time"

	"github.com/Shenr0n/bankapp/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer

}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}
func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}
func TestListTransfers(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	for i := 0; i < 6; i++ {
		arg := CreateTransferParams{
			FromAccountID: fromAccount.ID,
			ToAccountID:   toAccount.ID,
			Amount:        util.RandomMoney(),
		}

		transfer, err := testQueries.CreateTransfer(context.Background(), arg)

		require.NoError(t, err)
		require.NotEmpty(t, transfer)
	}

	arg := ListTransfersParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Limit:         3,
		Offset:        3,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 3)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
