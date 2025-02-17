package api

import (
	"net/http"

	"tokamak-sybil-resistance/api/parsers"

	"github.com/gin-gonic/gin"
)

func (a *API) getAccountByIndex(c *gin.Context) {
	// Get Addr
	account, err := parsers.ParseAccountFilterByIndex(c)
	if err != nil {
		retBadReq(&apiError{
			Err:  err,
			Code: ErrParamValidationFailedCode,
			Type: ErrParamValidationFailedType,
		}, c)
		return
	}
	apiAccount, err := a.historyDB.GetAccountAPIByIndex(*account)
	if err != nil {
		retSQLErr(err, c)
		return
	}
	c.JSON(http.StatusOK, apiAccount)
}

func (a *API) getAccountByEthAddr(c *gin.Context) {
	// Get Addr
	account, err := parsers.ParseAccountFilterByEthAddr(c)
	if err != nil {
		retBadReq(&apiError{
			Err:  err,
			Code: ErrParamValidationFailedCode,
			Type: ErrParamValidationFailedType,
		}, c)
		return
	}
	apiAccount, err := a.historyDB.GetAccountAPIByEthAddr(*account)
	if err != nil {
		retSQLErr(err, c)
		return
	}
	c.JSON(http.StatusOK, apiAccount)
}
