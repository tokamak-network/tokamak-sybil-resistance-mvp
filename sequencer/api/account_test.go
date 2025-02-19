package api

import (
	"fmt"
	"testing"

	"tokamak-sybil-resistance/common"
	"tokamak-sybil-resistance/common/apitypes"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testAccount struct {
	ItemID  uint64              `json:"itemId"`
	Idx     apitypes.TonIdx     `json:"accountIndex"`
	EthAddr apitypes.TonEthAddr `json:"tonEthereumAddress"`
	Balance *apitypes.BigIntStr `json:"balance"`
}

func genTestAccounts(accounts []common.Account) []testAccount {
	tAccounts := []testAccount{}
	for x, account := range accounts {
		tAccount := testAccount{
			ItemID:  uint64(x + 1),
			Idx:     apitypes.TonIdx(common.IdxToTon(account.Idx)),
			EthAddr: apitypes.NewTonEthAddr(account.EthAddr),
			Balance: apitypes.NewBigIntStr(account.Balance),
		}
		tAccounts = append(tAccounts, tAccount)
	}
	return tAccounts
}

func TestGetAccounts(t *testing.T) {
	endpoint := apiURL + "accounts"

	// Test GetAccountByIndex
	path := fmt.Sprintf("%s/index/%v", endpoint, tc.accounts[2].Idx)
	account := testAccount{}
	require.NoError(t, doGoodReq("GET", path, nil, &account))
	assert.Equal(t, tc.accounts[2], account)

	// Test GetAccountByEthAddr
	path = fmt.Sprintf("%s/address/%v", endpoint, tc.accounts[2].EthAddr)
	require.NoError(t, doGoodReq("GET", path, nil, &account))
	assert.Equal(t, tc.accounts[2], account)

	// 400
	path = fmt.Sprintf("%s/index/ton:12345", endpoint)
	err := doBadReq("GET", path, nil, 404)
	require.NoError(t, err)
}
