package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Timmy-Edibo/simple_bank/util"
	"github.com/stretchr/testify/require"
) 


func createRandomAccount(t *testing.T)Account{
	arg:= CreateAccountParams {
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account 

}
func TestCreateAccount(t *testing.T)  {
	createRandomAccount(t)
}


func TestGetAccount(t *testing.T){
	//create account
	account1:= createRandomAccount(t)
	res, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, account1.ID, res.ID)
	require.Equal(t, account1.Owner, res.Owner)
	require.Equal(t, account1.Balance, res.Balance)
	require.Equal(t, account1.Currency, res.Currency)
	require.WithinDuration(t, account1.CreatedAt, res.CreatedAt, time.Second)

		
}

func TestUpdateAccount(t *testing.T){
	// create account
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID: account1.ID,
		Balance: util.RandomMoney(),
	}

	res, err:= testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, account1.ID, res.ID)
	require.Equal(t, account1.Owner, res.Owner)
	require.Equal(t, arg.Balance, res.Balance)
	require.Equal(t, account1.Currency, res.Currency)
	require.WithinDuration(t, account1.CreatedAt, res.CreatedAt, time.Second)


}

func TestDeleteAccount(t *testing.T){
	// create account
	account1 := createRandomAccount(t)
	err:= testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	res, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, res)
	

}


func TestListAccount(t *testing.T){
	for i:=0; i<10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountParams {
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)	
	}
}