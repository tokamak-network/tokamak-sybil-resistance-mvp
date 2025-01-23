package parsers

import (
	"fmt"
	"strings"
	"tokamak-sybil-resistance/common"
	"tokamak-sybil-resistance/database/historydb"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// AccountFilter for parsing /accounts/{accountIndex} request to struct
type AccountFilter struct {
	AccountIndex string `uri:"accountIndex" binding:"required"`
}

// AccountsFilters for parsing /accounts query params to struct
type AccountsFilters struct {
	Addr string `form:"tonEthereumAddress"`

	Pagination
}

// ParseAccountsFilters parsing /accounts query params to GetAccountsAPIRequest
func ParseAccountsFilters(c *gin.Context, v *validator.Validate) (historydb.GetAccountsAPIRequest, error) {
	var accountsFilter AccountsFilters

	if err := c.BindQuery(&accountsFilter); err != nil {
		return historydb.GetAccountsAPIRequest{}, err
	}

	fmt.Println(accountsFilter.FromItem)

	// if err := v.Struct(accountsFilter); err != nil {
	// 	return historydb.GetAccountsAPIRequest{}, common.Wrap(err)
	// }

	addr, err := tonStringToEthAddr(accountsFilter.Addr, "tonEthereumAddress")
	if err != nil {
		return historydb.GetAccountsAPIRequest{}, err
	}

	return historydb.GetAccountsAPIRequest{
		EthAddr:  addr,
		FromItem: accountsFilter.FromItem,
		Order:    *accountsFilter.Order,
		Limit:    accountsFilter.Limit,
	}, nil
}

// tonStringToEthAddr converts ton ethereum address to ethereum address
func tonStringToEthAddr(addrStr, name string) (*ethCommon.Address, error) {
	if addrStr == "" {
		return nil, nil
	}
	splitted := strings.Split(addrStr, "ton:")
	if len(splitted) != 2 || len(splitted[1]) != 42 {
		return nil, common.Wrap(fmt.Errorf(
			"Invalid %s, must follow this regex: ^hez:0x[a-fA-F0-9]{40}$", name))
	}
	var addr ethCommon.Address
	err := addr.UnmarshalText([]byte(splitted[1]))
	return &addr, common.Wrap(err)
}
