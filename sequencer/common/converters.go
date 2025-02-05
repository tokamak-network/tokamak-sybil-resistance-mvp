package common

import (
	"fmt"
	"strconv"
	"strings"
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
func StringToIdx(idxStr string) (QueryAccount, error) {
	if idxStr == "" {
		return QueryAccount{}, nil
	}
	splitted := strings.Split(idxStr, ":")
	const expectedLen = 2
	if len(splitted) != expectedLen || splitted[0] != "ton" {
		return QueryAccount{}, Wrap(fmt.Errorf(
			"invalid format, must follow this: ton:index"))
	}
	idxInt, err := strconv.Atoi(splitted[1])
	idx := AccountIdx(idxInt)
	return QueryAccount{AccountIndex: &idx}, Wrap(err)
}
