package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/SameerJahagirdar/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  int64(util.RandomBalance()),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)

	require.NotZero(t, account.ID)

}

func CreateRandomAccount() *Accounts {
	account, err := testQueries.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  int64(util.RandomBalance()),
		Currency: util.RandomCurrency(),
	})
	if err != nil {
		fmt.Print(err)
	}
	return &account
}
