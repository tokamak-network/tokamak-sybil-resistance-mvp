package api

import (
	"fmt"
	"testing"

	"tokamak-sybil-resistance/common"
	"tokamak-sybil-resistance/common/apitypes"

	"github.com/mitchellh/copystructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testAccount struct {
	ItemID   uint64              `json:"itemId"`
	Idx      common.AccountIdx   `json:"accountIndex"`
	BatchNum common.BatchNum     `json:"batchNum"`
	EthAddr  apitypes.TonEthAddr `json:"tonEthereumAddress"`
	Nonce    common.Nonce        `json:"nonce"`
	Balance  *apitypes.BigIntStr `json:"balance"`
}

type testAccountsResponse struct {
	Accounts     []testAccount `json:"accounts"`
	PendingItems uint64        `json:"pendingItems"`
}

func (t testAccountsResponse) GetPending() (pendingItems, lastItemID uint64) {
	pendingItems = t.PendingItems
	lastItemID = t.Accounts[len(t.Accounts)-1].ItemID
	return pendingItems, lastItemID
}

func (t *testAccountsResponse) Len() int { return len(t.Accounts) }

func (t testAccountsResponse) New() Pendinger { return &testAccountsResponse{} }

func genTestAccounts(accounts []common.Account) []testAccount {
	tAccounts := []testAccount{}
	for x, account := range accounts {
		tAccount := testAccount{
			ItemID:  uint64(x + 1),
			Idx:     account.Idx,
			EthAddr: apitypes.NewTonEthAddr(account.EthAddr),
			Nonce:   account.Nonce,
			Balance: apitypes.NewBigIntStr(account.Balance),
		}
		tAccounts = append(tAccounts, tAccount)
	}
	return tAccounts
}

func TestGetAccounts(t *testing.T) {
	endpoint := apiURL + "accounts"
	fetchedAccounts := []testAccount{}

	appendIter := func(intr interface{}) {
		for i := 0; i < len(intr.(*testAccountsResponse).Accounts); i++ {
			tmp, err := copystructure.Copy(intr.(*testAccountsResponse).Accounts[i])
			if err != nil {
				panic(err)
			}
			fetchedAccounts = append(fetchedAccounts, tmp.(testAccount))
		}
	}

	limit := 5

	// Filter by ethAddr
	path := fmt.Sprintf("%s?tonEthereumAddress=%s&limit=%d", endpoint, tc.accounts[3].EthAddr, limit)
	err := doGoodReqPaginated(path, "ASC", &testAccountsResponse{}, appendIter)
	require.NoError(t, err)
	assert.Greater(t, len(fetchedAccounts), 0)
	assert.LessOrEqual(t, len(fetchedAccounts), len(tc.accounts))
	fetchedAccounts = []testAccount{}
	// No filters (checks response content)
	path = fmt.Sprintf("%s?limit=%d", endpoint, limit)
	err = doGoodReqPaginated(path, "ASC", &testAccountsResponse{}, appendIter)
	require.NoError(t, err)
	assert.Equal(t, len(tc.accounts), len(fetchedAccounts))
	for i := 0; i < len(fetchedAccounts); i++ {
		assert.Equal(t, tc.accounts[i], fetchedAccounts[i])
	}

	// No filters  Reverse Order (checks response content)
	reversedAccounts := []testAccount{}
	appendIter = func(intr interface{}) {
		for i := 0; i < len(intr.(*testAccountsResponse).Accounts); i++ {
			tmp, err := copystructure.Copy(intr.(*testAccountsResponse).Accounts[i])
			if err != nil {
				panic(err)
			}
			reversedAccounts = append(reversedAccounts, tmp.(testAccount))
		}
	}
	err = doGoodReqPaginated(path, "ASC", &testAccountsResponse{}, appendIter)
	require.NoError(t, err)
	assert.Equal(t, len(reversedAccounts), len(fetchedAccounts))
	for i := 0; i < len(fetchedAccounts); i++ {
		assert.Equal(t, reversedAccounts[i], fetchedAccounts[len(fetchedAccounts)-1-i])
	}

	// 400
	path = fmt.Sprintf("%s?tonEthereumAddress=ton:0x123456", endpoint)
	err = doBadReq("GET", path, nil, 400)
	require.NoError(t, err)

	// Test GetAccount
	path = fmt.Sprintf("%s/%v", endpoint, fetchedAccounts[2].Idx)
	account := testAccount{}
	require.NoError(t, doGoodReq("GET", path, nil, &account))
	assert.Equal(t, fetchedAccounts[2], account)

	// 400
	path = fmt.Sprintf("%s/ton:12345", endpoint)
	err = doBadReq("GET", path, nil, 400)
	require.NoError(t, err)

	// 404
	path = fmt.Sprintf("%s/ton:10:12345", endpoint)
	err = doBadReq("GET", path, nil, 404)
	require.NoError(t, err)

	// 400
	path = fmt.Sprintf("%s?ton:ton:25641", endpoint)
	err = doBadReq("GET", path, nil, 400)
	require.NoError(t, err)

	// 400
	path = fmt.Sprintf("%s?ton:ton:0xb4A2333993a70fD103b7cC39883797Aa209bAa21", endpoint)
	err = doBadReq("GET", path, nil, 400)
	require.NoError(t, err)
}
