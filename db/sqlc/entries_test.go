package db

import (
	"context"
	"testing"
	"udemy/util"

	"github.com/stretchr/testify/require"
)

func createRandomEntries(t *testing.T, account Account) Account {

	arg := CreateEntriesParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entries, err := testQueries.CreateEntries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entries)

	require.Equal(t, arg.AccountID, entries.AccountID)

	require.NotZero(t, account.ID)

	return account
}

func TestCreateEntries(t *testing.T) {
	account := CreateRandomAccount(t)
	createRandomEntries(t, account)
}

func TestListEntries(t *testing.T) {

	account := CreateRandomAccount(t)
	createRandomEntries(t, account)
	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     1,
		Offset:    0,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, entries[0].ID)
	require.NotZero(t, entries[0].Amount)

	ent, err := testQueries.GetEntry(context.Background(), entries[0].ID)
	require.NoError(t, err)
	require.NotZero(t, ent.ID)
}
