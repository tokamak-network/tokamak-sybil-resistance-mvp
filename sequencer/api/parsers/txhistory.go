package parsers

import (
	"fmt"

	"tokamak-sybil-resistance/common"
	"tokamak-sybil-resistance/database/historydb"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// HistoryTxFilter struct to get history tx uri param from /transaction-history/:id request
type HistoryTxFilter struct {
	TxID string `uri:"id" binding:"required"`
}

// ParseHistoryTxFilter function for parsing history tx filter to the txID
func ParseHistoryTxFilter(c *gin.Context) (common.TxID, error) {
	var historyTxFilter HistoryTxFilter
	if err := c.ShouldBindUri(&historyTxFilter); err != nil {
		return common.TxID{}, common.Wrap(err)
	}
	txID, err := common.NewTxIDFromString(historyTxFilter.TxID)
	if err != nil {
		return common.TxID{}, common.Wrap(fmt.Errorf("invalid txID"))
	}
	return txID, nil
}

// HistoryTxsFilters struct for holding filters from the /transaction-history request
type HistoryTxsFilters struct {
	TonEthereumAddr     string `form:"tonEthereumAddress"`
	FromTonEthereumAddr string `form:"fromTonEthereumAddress"`
	ToTonEthereumAddr   string `form:"toTonEthereumAddress"`
	AccountIndex        string `form:"accountIndex"`
	FromAccountIndex    string `form:"fromAccountIndex"`
	ToAccountIndex      string `form:"toAccountIndex"`
	BatchNum            *uint  `form:"batchNum"`
	TxType              string `form:"type"`
	IncludePendingTxs   *bool  `form:"includePendingL1s"`

	Pagination
}

// ParseHistoryTxsFilters func to parse history txs filters from query to the GetTxsAPIRequest
func ParseHistoryTxsFilters(c *gin.Context, v *validator.Validate) (historydb.GetTxsAPIRequest, error) {
	var historyTxsFilters HistoryTxsFilters
	if err := c.ShouldBindQuery(&historyTxsFilters); err != nil {
		return historydb.GetTxsAPIRequest{}, common.Wrap(err)
	}

	addr, err := common.TonStringToEthAddr(historyTxsFilters.TonEthereumAddr, "tonEthereumAddress")
	if err != nil {
		return historydb.GetTxsAPIRequest{}, common.Wrap(err)
	}

	fromAddr, err := common.TonStringToEthAddr(historyTxsFilters.FromTonEthereumAddr, "fromTonEthereumAddress")
	if err != nil {
		return historydb.GetTxsAPIRequest{}, common.Wrap(err)
	}

	toAddr, err := common.TonStringToEthAddr(historyTxsFilters.ToTonEthereumAddr, "toTonEthereumAddress")
	if err != nil {
		return historydb.GetTxsAPIRequest{}, common.Wrap(err)
	}

	// Idx
	queryAccount, err := common.StringToIdx(historyTxsFilters.AccountIndex, "accountIndex")
	if err != nil {
		return historydb.GetTxsAPIRequest{}, common.Wrap(err)
	}

	fromQueryAccount, err := common.StringToIdx(historyTxsFilters.FromAccountIndex, "fromAccountIndex")
	if err != nil {
		return historydb.GetTxsAPIRequest{}, common.Wrap(err)
	}

	toQueryAccount, err := common.StringToIdx(historyTxsFilters.ToAccountIndex, "toAccountIndex")
	if err != nil {
		return historydb.GetTxsAPIRequest{}, common.Wrap(err)
	}

	txType, err := common.StringToTxType(historyTxsFilters.TxType)
	if err != nil {
		return historydb.GetTxsAPIRequest{}, common.Wrap(err)
	}

	return historydb.GetTxsAPIRequest{
		EthAddr:           addr,
		FromEthAddr:       fromAddr,
		ToEthAddr:         toAddr,
		Idx:               queryAccount.AccountIndex,
		FromIdx:           fromQueryAccount.AccountIndex,
		ToIdx:             toQueryAccount.AccountIndex,
		BatchNum:          historyTxsFilters.BatchNum,
		TxType:            txType,
		IncludePendingL1s: historyTxsFilters.IncludePendingTxs,
		FromItem:          historyTxsFilters.FromItem,
		Limit:             historyTxsFilters.Limit,
		Order:             *historyTxsFilters.Order,
	}, nil
}
