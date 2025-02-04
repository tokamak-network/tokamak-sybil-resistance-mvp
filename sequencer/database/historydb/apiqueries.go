package historydb

import (
	"errors"
	"tokamak-sybil-resistance/common"
	"tokamak-sybil-resistance/database"

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
