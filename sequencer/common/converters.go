package common

import (
	"fmt"
	"strconv"
	"strings"

	ethCommon "github.com/ethereum/go-ethereum/common"
)

// QueryAccount is a representation of an account with accountIndex and its token symbol
type QueryAccount struct {
	AccountIndex *AccountIdx
}

// IdxToTon converts account index to ton account index with token symbol
func IdxToTon(idx AccountIdx) string {
	return "ton:" + strconv.Itoa(int(idx))
}

// StringToIdx converts string to account index
func StringToIdx(idxStr, name string) (QueryAccount, error) {
	if idxStr == "" {
		return QueryAccount{}, nil
	}
	splitted := strings.Split(idxStr, ":")
	const expectedLen = 2
	if len(splitted) != expectedLen || splitted[0] != "ton" {
		return QueryAccount{}, Wrap(fmt.Errorf(
			"invalid %s, must follow this: ton:index", name))
	}
	idxInt, err := strconv.Atoi(splitted[1])
	idx := AccountIdx(idxInt)
	return QueryAccount{AccountIndex: &idx}, Wrap(err)
}

// TonStringToEthAddr converts ton ethereum address to ethereum address
func TonStringToEthAddr(addrStr, name string) (*ethCommon.Address, error) {
	if addrStr == "" {
		return nil, nil
	}
	splitted := strings.Split(addrStr, "ton:")
	if len(splitted) != 2 || len(splitted[1]) != 42 {
		return nil, Wrap(fmt.Errorf(
			"Invalid %s, must follow this regex: ^ton:0x[a-fA-F0-9]{40}$", name))
	}
	var addr ethCommon.Address
	err := addr.UnmarshalText([]byte(splitted[1]))
	return &addr, Wrap(err)
}

// StringToTxType converts string to transaction type
func StringToTxType(txType string) (*TxType, error) {
	if txType == "" {
		return nil, nil
	}
	txTypeCasted := TxType(txType)
	switch txTypeCasted {
	case TxTypeDeposit, TxTypeCreateAccountDeposit, TxTypeWithdraw,
		TxTypeCreateVouch, TxTypeDeleteVouch:
		return &txTypeCasted, nil
	default:
		return nil, Wrap(fmt.Errorf(
			"invalid %s, %s is not a valid option. Check the valid options in the documentation",
			"type", txType,
		))
	}
}

// StringToEthAddr converts string to ethereum address
func StringToEthAddr(ethAddrStr string) (*ethCommon.Address, error) {
	if ethAddrStr == "" {
		return nil, nil
	}
	var addr ethCommon.Address
	err := addr.UnmarshalText([]byte(ethAddrStr))
	return &addr, Wrap(err)
}
