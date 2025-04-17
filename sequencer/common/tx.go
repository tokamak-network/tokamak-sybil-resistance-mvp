package common

import (
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethCommon "github.com/ethereum/go-ethereum/common"
)

type TxID [TxIDLen]byte

// Scan implements Scanner for database/sql.
func (txid *TxID) Scan(src interface{}) error {
	srcB, ok := src.([]byte)
	if !ok {
		return Wrap(fmt.Errorf("can't scan %T into TxID", src))
	}
	if len(srcB) != TxIDLen {
		return Wrap(fmt.Errorf("can't scan []byte of len %d into TxID, need %d",
			len(srcB), TxIDLen))
	}
	copy(txid[:], srcB)
	return nil
}

// Value implements valuer for database/sql.
func (txid TxID) Value() (driver.Value, error) {
	return txid[:], nil
}

// String returns a string hexadecimal representation of the TxID
func (txid TxID) String() string {
	return "0x" + hex.EncodeToString(txid[:])
}

// NewTxIDFromString returns a string hexadecimal representation of the TxID
func NewTxIDFromString(idStr string) (TxID, error) {
	txid := TxID{}
	idStr = strings.TrimPrefix(idStr, "0x")
	decoded, err := hex.DecodeString(idStr)
	if err != nil {
		return TxID{}, Wrap(err)
	}
	if len(decoded) != TxIDLen {
		return txid, Wrap(errors.New("invalid idStr"))
	}
	copy(txid[:], decoded)
	return txid, nil
}

// MarshalText marshals a TxID
func (txid TxID) MarshalText() ([]byte, error) {
	return []byte(txid.String()), nil
}

// UnmarshalText unmarshalls a TxID
func (txid *TxID) UnmarshalText(data []byte) error {
	idStr := string(data)
	id, err := NewTxIDFromString(idStr)
	if err != nil {
		return Wrap(err)
	}
	*txid = id
	return nil
}

type TxType string

const (
	// TxTypeDeposit represents L1->L2 transfer
	TxTypeDeposit TxType = "Deposit"
	// TxTypeCreateAccountDeposit represents creation of a new leaf in the state tree
	// (newAcconut) + L1->L2 transfer
	TxTypeCreateAccountDeposit TxType = "CreateAccountDeposit"
	// TxTypeWithdraw
	TxTypeWithdraw TxType = "Withdraw"
	// TxTypeCreateVouch
	TxTypeCreateVouch TxType = "CreateVouch"
	// TxTypeDeleteVouch
	TxTypeDeleteVouch TxType = "DeleteVouch"
	// TxTypeExplode
	TxTypeExplode TxType = "Explode"
	// TxTypeUnknown
	TxTypeUnknown TxType = ""
)

// Tx is a struct used by the TxSelector & BatchBuilder as a generic type generated from L1Tx
type Tx struct {
	// Generic
	IsL1        bool       `meddler:"is_l1"`
	TxID        TxID       `meddler:"id"`
	Type        TxType     `meddler:"type"`
	Position    int        `meddler:"position"`
	FromIdx     AccountIdx `meddler:"from_idx"`
	ToIdx       AccountIdx `meddler:"to_idx"`
	Amount      *big.Int   `meddler:"amount,bigint"`
	AmountFloat float64    `meddler:"amount_f"`
	// TokenID     TokenID    `meddler:"token_id"`
	// USD         *float64   `meddler:"amount_usd"`
	// BatchNum in which this tx was forged.
	BatchNum *BatchNum `meddler:"batch_num"`
	// Ethereum Block Number in which this L1Tx was added to the queue
	EthBlockNum int64 `meddler:"eth_block_num"`
	// L1
	// ToForgeL1TxsNum in which the tx was forged / will be forged
	ToForgeL1TxsNum *int64            `meddler:"to_forge_l1_txs_num"`
	FromEthAddr     ethCommon.Address `meddler:"from_eth_addr"`
	DepositAmount   *big.Int          `meddler:"deposit_amount,bigintnull"`
	Nonce           *Nonce            `meddler:"nonce"`
}

const (
	// TxIDPrefixL1UserTx is the prefix that determines that the TxID is for
	// a CreateAccountDeposit
	//nolinter:gomnd
	TxIDPrefixCreateAccountDeposit = byte(0)

	// TxIDPrefixL1CoordTx is the prefix that determines that the TxID is
	// for a Deposit
	//nolinter:gomnd
	TxIDPrefixDeposit = byte(1)

	// TxIDPrefixL1CoordTx is the prefix that determines that the TxID is
	// for a Withdraw
	//nolinter:gomnd
	TxIDPrefixWithdraw = byte(2)

	// TxIDPrefixL1CoordTx is the prefix that determines that the TxID is
	// for a Vouch
	//nolinter:gomnd
	TxIDPrefixVouch = byte(3)

	// TxIDPrefixL1CoordTx is the prefix that determines that the TxID is
	// for a Unvouch
	//nolinter:gomnd
	TxIDPrefixUnvouch = byte(4)

	// TxIDPrefixL1CoordTx is the prefix that determines that the TxID is
	// for a Explode
	//nolinter:gomnd
	TxIDPrefixExplode = byte(5)

	// TxIDLen is the length of the TxID byte array
	TxIDLen = 33
)

var (
	// SignatureConstantBytes contains the SignatureConstant in byte array
	// format, which is equivalent to 3322668559 as uint32 in byte array in
	// big endian representation.
	SignatureConstantBytes = []byte{198, 11, 230, 15}

	// EmptyTxID is used to check if a TxID is 0
	EmptyTxID = TxID([TxIDLen]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
)
