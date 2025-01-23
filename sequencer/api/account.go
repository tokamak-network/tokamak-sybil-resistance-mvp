package api

import (
	"fmt"
	"net/http"

	"tokamak-sybil-resistance/api/parsers"
	"tokamak-sybil-resistance/database/historydb"

	"github.com/gin-gonic/gin"
)

func (a *API) getAccounts(c *gin.Context) {
	for id := range c.Request.URL.Query() {
		if id != "tonEthereumAddress" && id != "fromItem" && id != "order" && id != "limit" {
			retBadReq(&apiError{
				Err:  fmt.Errorf("invalid Param: %s", id),
				Code: ErrParamValidationFailedCode,
				Type: ErrParamValidationFailedType,
			}, c)
			return
		}
	}

	accountsFilter, err := parsers.ParseAccountsFilters(c, a.validate)
	if err != nil {
		retBadReq(&apiError{
			Err:  err,
			Code: ErrParamValidationFailedCode,
			Type: ErrParamValidationFailedType,
		}, c)
		return
	}

	// Fetch Accounts from historyDB
	apiAccounts, pendingItems, err := a.historyDB.GetAccountsAPI(accountsFilter)
	if err != nil {
		retSQLErr(err, c)
		return
	}

	// Build successful response
	type accountResponse struct {
		Accounts     []historydb.AccountAPI `json:"accounts"`
		PendingItems uint64                 `json:"pendingItems"`
	}
	c.JSON(http.StatusOK, &accountResponse{
		Accounts:     apiAccounts,
		PendingItems: pendingItems,
	})
}
