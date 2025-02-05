package api

import (
	"net/http"

	"tokamak-sybil-resistance/api/parsers"
	"tokamak-sybil-resistance/database/historydb"

	"github.com/gin-gonic/gin"
)

func (a *API) getHistoryTxs(c *gin.Context) {
	txFilters, err := parsers.ParseHistoryTxsFilters(c, a.validate)
	if err != nil {
		retBadReq(&apiError{
			Err:  err,
			Code: ErrParamValidationFailedCode,
			Type: ErrParamValidationFailedType,
		}, c)
		return
	}
	// Fetch txs from historyDB
	txs, pendingItems, err := a.historyDB.GetTxsAPI(txFilters)
	if err != nil {
		retSQLErr(err, c)
		return
	}

	// Build successful response
	type txsResponse struct {
		Txs          []historydb.TxAPI `json:"transactions"`
		PendingItems uint64            `json:"pendingItems"`
	}
	c.JSON(http.StatusOK, &txsResponse{
		Txs:          txs,
		PendingItems: pendingItems,
	})
}

func (a *API) getHistoryTx(c *gin.Context) {
	// Get TxID
	txID, err := parsers.ParseHistoryTxFilter(c)
	if err != nil {
		retBadReq(&apiError{
			Err:  err,
			Code: ErrParamValidationFailedCode,
			Type: ErrParamValidationFailedType,
		}, c)
		return
	}
	// Fetch tx from historyDB
	tx, err := a.historyDB.GetTxAPI(txID)
	if err != nil {
		retSQLErr(err, c)
		return
	}
	// Build successful response
	c.JSON(http.StatusOK, tx)
}
