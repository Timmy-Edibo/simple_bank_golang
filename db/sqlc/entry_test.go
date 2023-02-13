package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Timmy-Edibo/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry

}


func TestCreateEntry(t *testing.T){
	createRandomEntry(t)

}

func TestGetEntry(t *testing.T){
	// create an entry
	entry := createRandomEntry(t)
	res, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, entry.ID, res.ID)
	require.Equal(t, entry.Amount, res.Amount)
	require.Equal(t, entry.AccountID, res.AccountID)
	require.WithinDuration(t, entry.CreatedAt, res.CreatedAt, time.Second)


}

func TestUpdateEntry(t *testing.T){
	entry := createRandomEntry(t)

	arg := UpdateEntryParams{
		ID: entry.ID,
		Amount: util.RandomMoney(),
	}	

	res, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, entry.ID, res.ID)
	require.Equal(t, entry.AccountID, res.AccountID)
	require.Equal(t, arg.Amount, res.Amount)
	require.WithinDuration(t, entry.CreatedAt, res.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T){
	entry := createRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	
	res, err:= testQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, res)
}

func TestListEntry(t *testing.T){
	for i:=0; i<=10; i++{
		createRandomEntry(t)
	}

	arg := ListEntryParams{
		Limit: 5,
		Offset: 5,	
	}

	entries, err := testQueries.ListEntry(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry:= range entries {
		require.NotEmpty(t, entry)
	}

}