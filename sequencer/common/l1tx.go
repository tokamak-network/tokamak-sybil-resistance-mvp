package common

import (
	"encoding/binary"
	"fmt"
	"math/big"

	ethCommon "github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/iden3/go-iden3-crypto/babyjub"
)

// EmptyBJJComp contains the 32 byte array of a empty BabyJubJub PublicKey
// Compressed. It is a valid point in the BabyJubJub curve, so does not give
// errors when being decompressed.
var EmptyBJJComp = babyjub.PublicKeyComp([32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

// L1Tx is a struct that represents a L1 tx
type L1Tx struct {
	// Stored in DB: mandatory fileds

	// TxID (32 bytes) for L1Tx is the Keccak256 (ethereum) hash of:
	// bytes:  |  1   |        8        |    2     |      1      |
	// values: | type | ToForgeL1TxsNum | Position | 0 (padding) |
	// where type:
	// 	- CreateAccountDeposit: 0
	// 	- Deposit: 1
	// 	- Withdraw: 2
	// 	- Vouch: 3
	// 	- Unvouch: 4
	// 	- Explode: 5
	TxID TxID `meddler:"id"`
	// ToForgeL1TxsNum indicates in which L1UserTx queue the tx was forged / will be forged
	ToForgeL1TxsNum *int64 `meddler:"to_forge_l1_txs_num"`
	Position        int    `meddler:"position"`
	// FromIdx is used by L1Tx/Deposit to indicate the Idx receiver of the L1Tx.DepositAmount
	// (deposit)
	FromIdx AccountIdx `meddler:"from_idx,zeroisnull"`
	// EffectiveFromIdx AccountIdx            `meddler:"effective_from_idx,zeroisnull"`
	FromEthAddr ethCommon.Address `meddler:"from_eth_addr,zeroisnull"`
	// FromBJJ          babyjub.PublicKeyComp `meddler:"from_bjj,zeroisnull"`
	// ToIdx is ignored in L1Tx/Deposit, but used in the L1Tx/DepositAndTransfer
	ToEthAddr ethCommon.Address `meddler:"to_eth_addr,zeroisnull"`
	ToIdx     AccountIdx        `meddler:"to_idx"`
	Amount    *big.Int          `meddler:"amount,bigint"`
	// EffectiveAmount only applies to L1UserTx.
	EffectiveAmount *big.Int `meddler:"effective_amount,bigintnull"`
	DepositAmount   *big.Int `meddler:"deposit_amount,bigint"`
	// EffectiveDepositAmount only applies to L1UserTx.
	EffectiveDepositAmount *big.Int `meddler:"effective_deposit_amount,bigintnull"`
	// Ethereum Block Number in which this L1Tx was added to the queue
	EthBlockNum int64          `meddler:"eth_block_num"`
	EthTxHash   ethCommon.Hash `meddler:"eth_tx_hash,zeroisnull"`
	// L1Fee       *big.Int       `meddler:"l1_fee,bigintnull"`
	Type     TxType    `meddler:"type"`
	BatchNum *BatchNum `meddler:"batch_num"`
}

// NewL1Tx returns the given L1Tx with the TxId & Type parameters calculated
// from the L1Tx values
func NewL1Tx(tx *L1Tx) (*L1Tx, error) {
	txIDOld := tx.TxID
	if err := tx.SetID(); err != nil {
		return nil, Wrap(err)
	}
	// If original TxID doesn't match the correct one, return error
	if txIDOld != (TxID{}) && txIDOld != tx.TxID {
		return tx, Wrap(fmt.Errorf("L1Tx.TxID: %s, should be: %s",
			tx.TxID.String(), txIDOld.String()))
	}

	return tx, nil
}

// SetID sets the ID of the transaction.  For L1UserTx uses (ToForgeL1TxsNum,
// Position), for L1CoordinatorTx uses (BatchNum, Position).
func (tx *L1Tx) SetID() error {
	var b []byte
	if tx.ToForgeL1TxsNum == nil {
		return Wrap(fmt.Errorf("L1Tx.UserOrigin == true && L1Tx.ToForgeL1TxsNum == nil"))
	}
	tx.TxID[0] = ByteFromType(tx.Type)

	var toForgeL1TxsNumBytes [8]byte
	binary.BigEndian.PutUint64(toForgeL1TxsNumBytes[:], uint64(*tx.ToForgeL1TxsNum))
	b = append(b, toForgeL1TxsNumBytes[:]...)

	var positionBytes [2]byte
	binary.BigEndian.PutUint16(positionBytes[:], uint16(tx.Position))
	b = append(b, positionBytes[:]...)

	// calculate hash
	h := ethCrypto.Keccak256Hash(b).Bytes()

	copy(tx.TxID[1:], h)

	return nil
}

// Tx returns a *Tx from the L1Tx
func (tx L1Tx) Tx() Tx {
	f := new(big.Float).SetInt(tx.EffectiveAmount)
	amountFloat, _ := f.Float64()
	genericTx := Tx{
		IsL1:     true,
		TxID:     tx.TxID,
		Type:     tx.Type,
		Position: tx.Position,
		// FromIdx:         tx.FromIdx,
		ToIdx:           tx.ToIdx,
		Amount:          tx.EffectiveAmount,
		AmountFloat:     amountFloat,
		ToForgeL1TxsNum: tx.ToForgeL1TxsNum,
		FromEthAddr:     tx.FromEthAddr,
		DepositAmount:   tx.EffectiveDepositAmount,
		EthBlockNum:     tx.EthBlockNum,
	}
	return genericTx
}

// TxCompressedData spec:
// [ 1 bits  ] empty (toBJJSign) // 1 byte
// [ 8 bits  ] empty (userFee) // 1 byte
// [ 40 bits ] empty (nonce) // 5 bytes
// [ 24 bits ] toIdx // 3 bytes
// [ 24 bits ] fromIdx // 3 bytes
// [ 16 bits ] chainId // 2 bytes
// [ 32 bits ] empty (signatureConstant) // 4 bytes
// Total bits compressed data:  225 bits // 29 bytes in *big.Int representation
func (tx L1Tx) TxCompressedData(chainID uint64) (*big.Int, error) {
	var b [29]byte
	// b[0:11] empty: no ToBJJSign, no fee, no nonce
	toIdxBytes, err := tx.ToIdx.Bytes()
	if err != nil {
		return nil, Wrap(err)
	}
	copy(b[11:17], toIdxBytes[:])
	fromIdxBytes, err := tx.FromIdx.Bytes()
	if err != nil {
		return nil, Wrap(err)
	}
	copy(b[17:23], fromIdxBytes[:])
	binary.BigEndian.PutUint64(b[23:25], chainID)
	copy(b[25:29], SignatureConstantBytes[:])

	bi := new(big.Int).SetBytes(b[:])
	return bi, nil
}

// L1UserTxFromEvent decodes L1UserTxEvent data
type L1UserTxEvent struct {
	QueueIndex uint32
	Position   uint8
	L1UserTx   *L1Tx
}

// L1UserTxFromEventData decodes the event data from L1UserTxEvent
func L1UserTxFromEventData(queueIndex uint32, position uint8, txData []byte) (*L1UserTxEvent, error) {
	// Parse the l1UserTx bytes
	tx, err := L1UserTxFromBytes(txData)
	if err != nil {
		return nil, Wrap(err)
	}

	return &L1UserTxEvent{
		QueueIndex: queueIndex,
		Position:   position,
		L1UserTx:   tx,
	}, nil
}

// L1UserTxFromBytes decodes the l1UserTx bytes from the event
func L1UserTxFromBytes(b []byte) (*L1Tx, error) {
	// The bytes should contain: ethAddress(20) + fromIdx(6) + loadAmountF(5) + amountF(5) + toIdx(6)
	if len(b) != RollupConstL1UserTotalBytes {
		return nil, Wrap(fmt.Errorf("invalid L1UserTx length: got %d, want %d", len(b), RollupConstL1UserTotalBytes))
	}

	tx := &L1Tx{}

	// Parse txType (1 byte)
	tx.Type = TypeFromByte(b[0])

	// Parse fromEthAddress (20 bytes)
	tx.FromEthAddr = ethCommon.BytesToAddress(b[1:21])

	// Parse toEthAddress (20 bytes)
	tx.ToEthAddr = ethCommon.BytesToAddress(b[21:41])

	// Parse amountF (5 bytes)
	tx.Amount = new(big.Int).SetBytes(b[41:73])

	// // Parse toIdx (6 bytes)
	// tx.ToIdx, err = AccountIdxFromBytes(b[36:42])
	// if err != nil {
	// 	return nil, Wrap(err)
	// }

	return tx, nil
}

// L1TxFromDataAvailability decodes a L1Tx from []byte (Data Availability)
// TODO: restruct based on L1Tx data field on contract
func L1TxFromDataAvailability(b []byte, nLevels uint32) (*L1Tx, error) {
	idxLen := nLevels / 8 //nolint:gomnd

	fromIdxBytes := b[0:idxLen]
	toIdxBytes := b[idxLen : idxLen*2]
	// amountBytes := b[idxLen*2 : idxLen*2+Float40BytesLength]

	l1tx := L1Tx{}
	fromIdx, err := AccountIdxFromBytes(ethCommon.LeftPadBytes(fromIdxBytes, 6))
	if err != nil {
		return nil, Wrap(err)
	}
	l1tx.FromIdx = fromIdx
	toIdx, err := AccountIdxFromBytes(ethCommon.LeftPadBytes(toIdxBytes, 6))
	if err != nil {
		return nil, Wrap(err)
	}
	l1tx.ToIdx = toIdx
	// l1tx.Amount = big.NewInt(binary.BigEndian.Uint64(amountBytes))
	return &l1tx, Wrap(err)
}

// // BytesDataAvailability encodes a L1Tx into []byte for the Data Availability
// // [ fromIdx | toIdx | amountFloat40 | Fee ]
func (tx *L1Tx) BytesDataAvailability(nLevels uint32) ([]byte, error) {
	idxLen := nLevels / 8 //nolint:gomnd

	b := make([]byte, ((nLevels*2)+40+8)/8) //nolint:gomnd

	fromIdxBytes, err := tx.FromIdx.Bytes()
	if err != nil {
		return nil, Wrap(err)
	}
	copy(b[0:idxLen], fromIdxBytes[6-idxLen:])
	toIdxBytes, err := tx.ToIdx.Bytes()
	if err != nil {
		return nil, Wrap(err)
	}
	copy(b[idxLen:idxLen*2], toIdxBytes[6-idxLen:])

	if tx.EffectiveAmount != nil {
		amountFloat40, err := NewFloat40(tx.EffectiveAmount)
		if err != nil {
			return nil, Wrap(err)
		}
		amountFloat40Bytes, err := amountFloat40.Bytes()
		if err != nil {
			return nil, Wrap(err)
		}
		copy(b[idxLen*2:idxLen*2+Float40BytesLength], amountFloat40Bytes)
	}
	// fee = 0 (as is L1Tx)
	return b[:], nil
}

// TypeFromByte converts a byte representation of txType to the corresponding TxType
func TypeFromByte(txType byte) TxType {
	switch txType {
	case 0:
		return TxTypeCreateAccountDeposit
	case 1:
		return TxTypeDeposit
	case 2:
		return TxTypeWithdraw
	case 3:
		return TxTypeCreateVouch
	case 4:
		return TxTypeDeleteVouch
	case 5:
		return TxTypeExplode
	default:
		return TxTypeUnknown
	}
}

func ByteFromType(txType TxType) byte {
	switch txType {
	case TxTypeCreateAccountDeposit:
		return 0
	case TxTypeDeposit:
		return 1
	case TxTypeWithdraw:
		return 2
	case TxTypeCreateVouch:
		return 3
	case TxTypeDeleteVouch:
		return 4
	case TxTypeExplode:
		return 5
	case TxTypeUnknown:
		return 6
	default:
		return 255
	}
}
