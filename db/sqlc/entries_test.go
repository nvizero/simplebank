package db

import (
	"context"
	"testing"
	"udemy/util"

	"github.com/stretchr/testify/require"
)

func createRandomEntries(t *testing.T, account Accounts) Accounts {

	arg := CreateEntriesParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entries, err := testQueries.CreateEntries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entries)

	require.Equal(t, arg.AccountID, entries.AccountID)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateEntries(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntries(t, account)
}
