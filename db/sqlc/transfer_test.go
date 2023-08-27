package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func DefAccounts(t *testing.T) []Accounts {

	arg := ListAccountsParams{
		Limit:  2,
		Offset: 9,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	return accounts
}

func createRandomTransfer(t *testing.T) Transfers {

	accounts := DefAccounts(t)
	arg2 := CreateTransferParams{
		FromAccountID: accounts[0].ID,
		ToAccountID:   accounts[1].ID,
		Amount:        1,
	}

	transfers, err := testQueries.CreateTransfer(context.Background(), arg2)
	require.NoError(t, err)
	return transfers
}

func TestCreateTransfer(t *testing.T) {
	tran := createRandomTransfer(t)
	require.NotEmpty(t, tran)
}

func TestListTransfer(t *testing.T) {
	accounts := DefAccounts(t)
	arg := ListTransfersParams{
		FromAccountID: accounts[0].ID,
		ToAccountID:   accounts[1].ID,
		Limit:         2,
		Offset:        0,
	}
	trans, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, trans)
}
