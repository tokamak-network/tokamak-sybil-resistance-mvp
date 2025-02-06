package historydb

import (
	"errors"
	"fmt"
	"tokamak-sybil-resistance/common"
	"tokamak-sybil-resistance/database"
	"tokamak-sybil-resistance/log"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/jmoiron/sqlx"
	"github.com/russross/meddler"
)

// GetBatchInternalAPI return the batch with the given batchNum
func (hdb *HistoryDB) GetBatchInternalAPI(batchNum common.BatchNum) (*BatchAPI, error) {
	return hdb.getBatchAPI(hdb.dbRead, batchNum)
}

func (hdb *HistoryDB) getBatchAPI(d meddler.DB, batchNum common.BatchNum) (*BatchAPI, error) {
	batch := &BatchAPI{}
	if err := meddler.QueryRow(
		d, batch,
		`SELECT batch.item_id, batch.batch_num, batch.eth_block_num,
		batch.forger_addr, batch.account_root, batch.score_root, batch.vouch_root, batch.num_accounts, batch.exit_root, batch.forge_l1_txs_num,
		COALESCE(batch.eth_tx_hash, DECODE('0000000000000000000000000000000000000000000000000000000000000000', 'hex')) as eth_tx_hash,
		block.timestamp, block.hash, COALESCE ((SELECT COUNT(*) FROM tx WHERE batch_num = batch.batch_num), 0) AS forged_txs
	    FROM batch INNER JOIN block ON batch.eth_block_num = block.eth_block_num
	 	WHERE batch_num = $1;`, batchNum,
	); err != nil {
		return nil, common.Wrap(err)
	}
	return batch, nil
}

// GetAccountAPI returns an account by its index
func (hdb *HistoryDB) GetAccountAPI(idx common.AccountIdx) (*AccountAPI, error) {
	cancel, err := hdb.apiConnCon.Acquire()
	defer cancel()
	if err != nil {
		return nil, common.Wrap(err)
	}
	defer hdb.apiConnCon.Release()
	account := &AccountAPI{}
	err = meddler.QueryRow(hdb.dbRead, account, `SELECT account.item_id, ton_idx(account.idx) as idx,
		account.batch_num, account.eth_addr, account_update.nonce, account_update.balance 
		FROM account inner JOIN (
			SELECT idx, nonce, balance 
			FROM account_update
			WHERE idx = $1
			ORDER BY item_id DESC LIMIT 1
		) AS account_update ON account_update.idx = account.idx
		WHERE account.idx = $1;`, idx)

	if err != nil {
		return nil, common.Wrap(err)
	}

	return account, nil
}

// GetAccountsAPIRequest is an API request struct for getting accounts
type GetAccountsAPIRequest struct {
	EthAddr *ethCommon.Address
}

// GetAccountsAPI returns a list of accounts from the DB and pagination info
func (hdb *HistoryDB) GetAccountsAPI(
	request GetAccountsAPIRequest,
) ([]AccountAPI, error) {
	if request.EthAddr == nil {
		return nil, common.Wrap(errors.New("ethAddr is required"))
	}
	cancel, err := hdb.apiConnCon.Acquire()
	defer cancel()
	if err != nil {
		return nil, common.Wrap(err)
	}
	defer hdb.apiConnCon.Release()
	var query string
	var args []interface{}
	queryStr := `SELECT account.item_id, ton_idx(account.idx) as idx, account.batch_num, 
	account.eth_addr, 
	account_update.nonce, account_update.balance, COUNT(*) OVER() AS total_items
	FROM account INNER JOIN (
		SELECT DISTINCT idx,
		first_value(nonce) OVER w AS nonce,
		first_value(balance) OVER w AS balance
		FROM account_update
		WINDOW w as (PARTITION BY idx ORDER BY item_id DESC)
	) AS account_update ON account_update.idx = account.idx `
	// ethAddr filter
	if request.EthAddr != nil {
		queryStr += "WHERE account.eth_addr = ? "
		args = append(args, request.EthAddr)
	}
	query, argsQ, err := sqlx.In(queryStr, args...)
	if err != nil {
		return nil, common.Wrap(err)
	}
	query = hdb.dbRead.Rebind(query)

	accounts := []*AccountAPI{}
	if err := meddler.QueryAll(hdb.dbRead, &accounts, query, argsQ...); err != nil {
		return nil, common.Wrap(err)
	}
	if len(accounts) == 0 {
		return []AccountAPI{}, nil
	}

	return database.SlicePtrsToSlice(accounts).([]AccountAPI), nil
}

// GetTxAPI returns a tx from the DB given a TxID
func (hdb *HistoryDB) GetTxAPI(txID common.TxID) (*TxAPI, error) {
	// Warning: amount_success and deposit_amount_success have true as default for
	// performance reasons. The expected default value is false (when txs are unforged)
	// this case is handled at the function func (tx TxAPI) MarshalJSON() ([]byte, error)
	cancel, err := hdb.apiConnCon.Acquire()
	defer cancel()
	if err != nil {
		return nil, common.Wrap(err)
	}
	defer hdb.apiConnCon.Release()
	tx := &TxAPI{}
	err = meddler.QueryRow(
		hdb.dbRead, tx, `SELECT tx.item_id, tx.is_l1, tx.id, tx.type, tx.position, 
		ton_idx(tx.effective_from_idx) AS from_idx, tx.from_eth_addr,
		ton_idx(tx.to_idx) AS to_idx, tx.to_eth_addr,
		tx.amount, tx.amount_success,
		tx.batch_num, tx.eth_block_num, tx.to_forge_l1_txs_num, tx.user_origin, tx.eth_tx_hash, tx.l1_fee,
		tx.deposit_amount, tx.deposit_amount_success,
		block.timestamp
		FROM tx INNER JOIN block ON tx.eth_block_num = block.eth_block_num 
		WHERE tx.id = $1;`, txID,
	)
	return tx, common.Wrap(err)
}

// GetTxsAPIRequest is an API request struct for getting txs
type GetTxsAPIRequest struct {
	EthAddr           *ethCommon.Address
	FromEthAddr       *ethCommon.Address
	ToEthAddr         *ethCommon.Address
	Idx               *common.AccountIdx
	FromIdx           *common.AccountIdx
	ToIdx             *common.AccountIdx
	BatchNum          *uint
	TxType            *common.TxType
	IncludePendingL1s *bool

	FromItem *uint
	Limit    *uint
	Order    string
}

// GetTxsAPI returns a list of txs from the DB using the HistoryTx struct
// and pagination info
func (hdb *HistoryDB) GetTxsAPI(
	request GetTxsAPIRequest,
) ([]TxAPI, uint64, error) {
	// Warning: amount_success and deposit_amount_success have true as default for
	// performance reasons. The expected default value is false (when txs are unforged)
	// this case is handled at the function func (tx TxAPI) MarshalJSON() ([]byte, error)
	cancel, err := hdb.apiConnCon.Acquire()
	defer cancel()
	if err != nil {
		return nil, 0, common.Wrap(err)
	}
	defer hdb.apiConnCon.Release()
	var query string
	var args []interface{}
	queryStr := `SELECT tx.item_id, tx.is_l1, tx.id, tx.type, tx.position, 
	ton_idx(tx.effective_from_idx) AS from_idx, tx.from_eth_addr,
	ton_idx(tx.to_idx) AS to_idx, tx.to_eth_addr,
	tx.amount, tx.amount_success,
	tx.batch_num, tx.eth_block_num, tx.to_forge_l1_txs_num, tx.user_origin, tx.eth_tx_hash, tx.l1_fee,
	tx.deposit_amount, tx.deposit_amount_success,
	block.timestamp, count(*) OVER() AS total_items 
	FROM tx INNER JOIN block ON tx.eth_block_num = block.eth_block_num `
	// Apply filters
	nextIsAnd := false
	// ethAddr filter
	if request.EthAddr != nil {
		queryStr += "WHERE (tx.from_eth_addr = ? OR tx.to_eth_addr = ?) "
		nextIsAnd = true
		args = append(args, request.EthAddr, request.EthAddr)
	} else if request.FromEthAddr != nil && request.ToEthAddr != nil {
		queryStr += "WHERE (tx.from_eth_addr = ? AND tx.to_eth_addr = ?) "
		nextIsAnd = true
		args = append(args, request.FromEthAddr, request.ToEthAddr)
	} else if request.FromEthAddr != nil {
		queryStr += "WHERE tx.from_eth_addr = ? "
		nextIsAnd = true
		args = append(args, request.FromEthAddr)
	} else if request.ToEthAddr != nil {
		queryStr += "WHERE tx.to_eth_addr = ? "
		nextIsAnd = true
		args = append(args, request.ToEthAddr)
	}
	// idx filter
	if request.Idx != nil {
		if nextIsAnd {
			queryStr += "AND "
		} else {
			queryStr += "WHERE "
		}
		queryStr += "(tx.effective_from_idx = ? OR tx.to_idx = ?) "
		args = append(args, request.Idx, request.Idx)
		nextIsAnd = true
	} else if request.FromIdx != nil && request.ToIdx != nil {
		if nextIsAnd {
			queryStr += "AND "
		} else {
			queryStr += "WHERE "
		}
		queryStr += "(tx.effective_from_idx = ? AND tx.to_idx = ?) "
		args = append(args, request.FromIdx, request.ToIdx)
		nextIsAnd = true
	} else if request.FromIdx != nil {
		if nextIsAnd {
			queryStr += "AND "
		} else {
			queryStr += "WHERE "
		}
		queryStr += "tx.effective_from_idx = ? "
		args = append(args, request.FromIdx)
		nextIsAnd = true
	} else if request.ToIdx != nil {
		if nextIsAnd {
			queryStr += "AND "
		} else {
			queryStr += "WHERE "
		}
		queryStr += "tx.to_idx = ? "
		args = append(args, request.ToIdx)
		nextIsAnd = true
	}
	// batchNum filter
	if request.BatchNum != nil {
		if nextIsAnd {
			queryStr += "AND "
		} else {
			queryStr += "WHERE "
		}
		queryStr += "tx.batch_num = ? "
		args = append(args, request.BatchNum)
		nextIsAnd = true
	}
	// txType filter
	if request.TxType != nil {
		if nextIsAnd {
			queryStr += "AND "
		} else {
			queryStr += "WHERE "
		}
		queryStr += "tx.type = ? "
		args = append(args, request.TxType)
		nextIsAnd = true
	}
	if request.FromItem != nil {
		if nextIsAnd {
			queryStr += "AND "
		} else {
			queryStr += "WHERE "
		}
		if request.Order == "ASC" {
			queryStr += "tx.item_id >= ? "
		} else {
			queryStr += "tx.item_id <= ? "
		}
		args = append(args, request.FromItem)
		nextIsAnd = true
	}

	// Include pending L1 txs? (default false)
	if request.IncludePendingL1s == nil || (request.IncludePendingL1s != nil && !*request.IncludePendingL1s) {
		if nextIsAnd {
			queryStr += "AND "
		} else {
			queryStr += "WHERE "
		}
		queryStr += "tx.batch_num IS NOT NULL "
	}

	// pagination
	queryStr += "ORDER BY tx.item_id "
	if request.Order == "ASC" {
		queryStr += " ASC "
	} else {
		queryStr += " DESC "
	}
	queryStr += fmt.Sprintf("LIMIT %d;", *request.Limit)
	query = hdb.dbRead.Rebind(queryStr)
	log.Debug(query)
	txsPtrs := []*TxAPI{}
	if err := meddler.QueryAll(hdb.dbRead, &txsPtrs, query, args...); err != nil {
		return nil, 0, common.Wrap(err)
	}
	txs := database.SlicePtrsToSlice(txsPtrs).([]TxAPI)
	if len(txs) == 0 {
		return txs, 0, nil
	}
	return txs, txs[0].TotalItems - uint64(len(txs)), nil
}

// GetBatchAPI return the batch with the given batchNum
func (hdb *HistoryDB) GetBatchAPI(batchNum common.BatchNum) (*BatchAPI, error) {
	cancel, err := hdb.apiConnCon.Acquire()
	defer cancel()
	if err != nil {
		return nil, common.Wrap(err)
	}
	defer hdb.apiConnCon.Release()
	return hdb.getBatchAPI(hdb.dbRead, batchNum)
}

// GetBatchesAPIRequest is an API request struct for getting batches
type GetBatchesAPIRequest struct {
	MinBatchNum *uint
	MaxBatchNum *uint
	ForgerAddr  *ethCommon.Address

	FromItem *uint
	Limit    *uint
	Order    string
}

// GetBatchesAPI return the batches applying the given filters
func (hdb *HistoryDB) GetBatchesAPI(
	request GetBatchesAPIRequest,
) ([]BatchAPI, uint64, error) {
	cancel, err := hdb.apiConnCon.Acquire()
	defer cancel()
	if err != nil {
		return nil, 0, common.Wrap(err)
	}
	defer hdb.apiConnCon.Release()
	var query string
	var args []interface{}
	queryStr := `SELECT batch.item_id, batch.batch_num, batch.eth_block_num,
	batch.forger_addr, batch.num_accounts, batch.exit_root, batch.forge_l1_txs_num,
	batch.eth_tx_hash, block.timestamp, block.hash,
	COALESCE ((SELECT COUNT(*) FROM tx WHERE batch_num = batch.batch_num), 0) AS forged_txs,
	count(*) OVER() AS total_items
	FROM batch INNER JOIN block ON batch.eth_block_num = block.eth_block_num `
	// Apply filters
	nextIsAnd := false
	// minBatchNum filter
	if request.MinBatchNum != nil {
		if nextIsAnd {
			queryStr += "AND "
		} else {
			queryStr += "WHERE "
		}
		queryStr += "batch.batch_num > ? "
		args = append(args, request.MinBatchNum)
		nextIsAnd = true
	}
	// maxBatchNum filter
	if request.MaxBatchNum != nil {
		if nextIsAnd {
			queryStr += "AND "
		} else {
			queryStr += "WHERE "
		}
		queryStr += "batch.batch_num < ? "
		args = append(args, request.MaxBatchNum)
		nextIsAnd = true
	}
	// forgerAddr filter
	if request.ForgerAddr != nil {
		if nextIsAnd {
			queryStr += "AND "
		} else {
			queryStr += "WHERE "
		}
		queryStr += "batch.forger_addr = ? "
		args = append(args, request.ForgerAddr)
		nextIsAnd = true
	}
	// pagination
	if request.FromItem != nil {
		if nextIsAnd {
			queryStr += "AND "
		} else {
			queryStr += "WHERE "
		}
		if request.Order == "ASC" {
			queryStr += "batch.item_id >= ? "
		} else {
			queryStr += "batch.item_id <= ? "
		}
		args = append(args, request.FromItem)
	}
	queryStr += "ORDER BY batch.item_id "
	if request.Order == "ASC" {
		queryStr += " ASC "
	} else {
		queryStr += " DESC "
	}
	queryStr += fmt.Sprintf("LIMIT %d;", *request.Limit)
	query = hdb.dbRead.Rebind(queryStr)
	// log.Debug(query)
	batchPtrs := []*BatchAPI{}
	if err := meddler.QueryAll(hdb.dbRead, &batchPtrs, query, args...); err != nil {
		return nil, 0, common.Wrap(err)
	}
	batches := database.SlicePtrsToSlice(batchPtrs).([]BatchAPI)
	if len(batches) == 0 {
		return batches, 0, nil
	}
	return batches, batches[0].TotalItems - uint64(len(batches)), nil
}
