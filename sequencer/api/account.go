package api

import (
	"net/http"

	"tokamak-sybil-resistance/api/parsers"

	"github.com/gin-gonic/gin"
)

func (a *API) getAccount(c *gin.Context) {
	// Get Addr
	account, err := parsers.ParseAccountFilter(c)
	if err != nil {
		retBadReq(&apiError{
			Err:  err,
			Code: ErrParamValidationFailedCode,
			Type: ErrParamValidationFailedType,
		}, c)
		return
	}
	apiAccount, err := a.historyDB.GetAccountAPI(*account)
	if err != nil {
		retSQLErr(err, c)
		return
	}
	c.JSON(http.StatusOK, apiAccount)
}
