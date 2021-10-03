package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/AirellJordan98/simplebank/util"
	"github.com/stretchr/testify/require"
)

// This variable will be reassigned eventually with the value from the response of 'TestCreateAccount' function

func createRandomAccount(t *testing.T) Account {

	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err, "Error has occured")
	require.NotEmpty(t, account, "account should no be empty")

	require.Equal(t, account.Owner, arg.Owner, "account owner should be the same as arg owner")
	require.Equal(t, account.Balance, arg.Balance, "account balance should be the same as arg balance")
	require.Equal(t, account.Currency, arg.Currency, "account currency should be the same as arg currency")

	require.NotZero(t, account.ID, "account id should not be a zero value")

	return account
}

func TestCreateAccount(t *testing.T) {
	// t.Skip()
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	// t.Skip()
	newAccount := createRandomAccount(t)

	account, err := testQueries.GetAccount(context.Background(), newAccount.ID)

	require.NoError(t, err, "get one account currently facing an error")
	require.NotEmpty(t, account, "account should not be empty")

	require.Equal(t, account.ID, newAccount.ID, fmt.Sprintf("account id should have a value of id %d", newAccount.ID))
	require.Equal(t, account.Owner, newAccount.Owner, fmt.Sprintf("account owner should be the same as new account owner which has the name of %s", newAccount.Owner))
	require.Equal(t, account.Balance, newAccount.Balance, "account balance should be the same as new account balance")
	require.Equal(t, account.Currency, newAccount.Currency, "account currency should be the same as new account currency")
	require.Equal(t, account.CreatedAt, newAccount.CreatedAt, "account created at date should be the same as new account created at date")
}

func TestUpdateAccount(t *testing.T) {
	// t.Skip()
	newAccount := createRandomAccount(t)

	args := UpdateAccountParams{
		ID:      newAccount.ID,
		Balance: newAccount.Balance,
	}
	updatedAccount, err := testQueries.UpdateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, updatedAccount.ID, newAccount.ID, fmt.Sprintf("account id should have a value of id %d", newAccount.ID))
	require.Equal(t, updatedAccount.Owner, newAccount.Owner, fmt.Sprintf("account owner should be the same as new account owner which has the name of %s", newAccount.Owner))
	require.Equal(t, updatedAccount.Balance, newAccount.Balance, "account balance should be the same as args balance")
	require.Equal(t, updatedAccount.Currency, newAccount.Currency, "account currency should be the same as new account currency")
	require.Equal(t, updatedAccount.CreatedAt, newAccount.CreatedAt, "account created at date should be the same as new account created at date")
}

func TestListAccounts(t *testing.T) {
	// t.Skip()
	createRandomAccount(t)

	args := ListAccountsParams{
		Offset: 0,
		Limit:  3,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, accounts, int(args.Limit))

	for i := 0; i < len(accounts); i++ {
		require.NotEmpty(t, accounts)
	}
}

func TestDeleteAccount(t *testing.T) {
	// t.Skip()
	newAccount := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), newAccount.ID)

	require.NoError(t, err)
}
