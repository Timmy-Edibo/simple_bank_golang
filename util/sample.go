// package db

// import (
// 	"context"
// 	"database/sql"
// 	"testing"
// 	"time"

// 	"github.com/Timmy-Edibo/simple_bank/util"
// 	"github.com/stretchr/testify/require"
// )

// func createRandomTransfer(t *testing.T) Transfer {
// 	account1 := createRandomAccount(t)
// 	account2 := createRandomAccount(t)

// 	arg := CreateTransferParams{
// 		FromAccountID: account1.ID,
// 		ToAccountID:   account2.ID,
// 		Amount:        util.RandomMoney(),
// 	}

// 	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, transfer)

// 	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
// 	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
// 	require.Equal(t, arg.Amount, transfer.Amount)

// 	require.NotZero(t, transfer.ID)
// 	require.NotZero(t, transfer.CreatedAt)

// 	return transfer

// }

// func TestCreateTransfer(t *testing.T) {
// 	createRandomTransfer(t)
// }

// func TestGetTransfer(t *testing.T) {
// 	transfer := createRandomTransfer(t)

// 	res, err := testQueries.GetTransfer(context.Background(), transfer.ID)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, res)

// 	require.Equal(t, transfer.ID, res.ID)
// 	require.Equal(t, transfer.FromAccountID, res.FromAccountID)
// 	require.Equal(t, transfer.ToAccountID, res.ToAccountID)
// 	require.Equal(t, transfer.Amount, res.Amount)
// 	require.WithinDuration(t, transfer.CreatedAt, res.CreatedAt, time.Second)

// }

// func TestUpdateTransfer(t *testing.T) {
// 	transfer := createRandomTransfer(t)

// 	arg := UpdateTransferParams{
// 		ID:     transfer.ID,
// 		Amount: util.RandomMoney(),
// 	}

// 	res, err := testQueries.UpdateTransfer(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, res)

// 	require.Equal(t, transfer.ID, res.ID)
// 	require.Equal(t, transfer.FromAccountID, res.FromAccountID)
// 	require.Equal(t, transfer.ToAccountID, res.ToAccountID)
// 	require.Equal(t, arg.Amount, res.Amount)
// 	require.WithinDuration(t, transfer.CreatedAt, res.CreatedAt, time.Second)

// }

// func TestDeleteTransfer(t *testing.T) {
// 	transfer := createRandomTransfer(t)

// 	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
// 	require.NoError(t, err)

// 	res, err := testQueries.GetTransfer(context.Background(), transfer.ID)
// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, res)

// }
