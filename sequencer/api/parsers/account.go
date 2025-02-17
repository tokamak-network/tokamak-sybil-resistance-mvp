package parsers

import (
	"fmt"
	"strconv"
	"strings"
	"tokamak-sybil-resistance/common"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// AccountFilter for parsing /accounts/{accountIndex} request to struct
type AccountFilterByIndex struct {
	AccountIndex string `uri:"accountIndex" binding:"required"`
}

// AccountFilterByEthAddr for parsing /accounts/{ethAddr} request to struct
type AccountFilterByEthAddr struct {
	EthAddr string `uri:"ethAddr" binding:"required"`
}

// ParseAccountFilterByIndex parses account filter to the account index
func ParseAccountFilterByIndex(c *gin.Context) (*common.AccountIdx, error) {
	var accountFilter AccountFilterByIndex
	if err := c.ShouldBindUri(&accountFilter); err != nil {
		return nil, common.Wrap(err)
	}
	return stringToAccountIdx(accountFilter.AccountIndex)
}

// ParseAccountFilterByEthAddr parses account filter to the ethAddress
func ParseAccountFilterByEthAddr(c *gin.Context) (*ethCommon.Address, error) {
	var accountFilter AccountFilterByEthAddr
	if err := c.ShouldBindUri(&accountFilter); err != nil {
		return nil, common.Wrap(err)
	}
	return common.TonStringToEthAddr(accountFilter.EthAddr, "ethAddr")
}

// StringToIdx converts string to account index
func stringToAccountIdx(idxStr string) (*common.AccountIdx, error) {
	if idxStr == "" {
		return nil, nil
	}
	splitted := strings.Split(idxStr, ":")
	const expectedLen = 2
	const tonIndex = "ton"
	if len(splitted) != expectedLen || splitted[0] != tonIndex {
		return nil, common.Wrap(fmt.Errorf(
			"invalid format, must follow this: ton:index"))
	}
	idxInt, err := strconv.Atoi(splitted[1])
	idx := common.AccountIdx(idxInt)
	return &idx, common.Wrap(err)
}
