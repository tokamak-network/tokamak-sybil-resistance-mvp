package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
	"tokamak-sybil-resistance/api/stateapiupdater"
	"tokamak-sybil-resistance/common"
	"tokamak-sybil-resistance/database"
	"tokamak-sybil-resistance/database/historydb"
	"tokamak-sybil-resistance/log"
	"tokamak-sybil-resistance/test"
	"tokamak-sybil-resistance/test/til"

	swagger "github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-gonic/gin"
	"github.com/iden3/go-merkletree"

	ethCommon "github.com/ethereum/go-ethereum/common"
)

// Pendinger is an interface that allows getting last returned item ID and PendingItems to be used for building fromItem
// when testing paginated endpoints.
type Pendinger interface {
	GetPending() (pendingItems, lastItemID uint64)
	Len() int
	New() Pendinger
}

const (
	apiPort = "4010"
	apiIP   = "http://localhost:"
	apiURL  = apiIP + apiPort + "/v1/"
)

var SetBlockchain = `
Type: Blockchain

CreateAccountDeposit A: 2000 // Idx=256+0=256
CreateAccountDeposit B: 500  // Idx=256+1=257
CreateAccountDeposit C: 2000 // Idx=256+2=258
CreateAccountDeposit D: 500  // Idx=256+3=259

> batchL1 // forge L1UserTxs{nil}, freeze defined L1UserTxs{4}
> batchL1 // forge defined L1UserTxs{4}, freeze L1UserTxs{nil}
> block // blockNum=2

ForceExit A: 100
ForceExit B: 80

CreateVouch C-A
CreateVouch C-B
CreateVouch C-D
// Exit C: 50
// Exit D: 30

> batchL1 // forge L1UserTxs{nil}, freeze defined L1UserTxs{2}
> batchL1 // forge L1UserTxs{2}, freeze defined L1UserTxs{nil}
> block // blockNum=3
`

type testCommon struct {
	blocks []common.Block
	// batches     []testBatch
	// fullBatches []testFullBatch
	accounts []testAccount
	// txs         []testTx
	router     *swagger.Router
	rollupVars common.RollupVariables
}

var tc testCommon
var config configAPI
var api *API
var stateAPIUpdater *stateapiupdater.Updater

// TestMain initializes the API server, and fill HistoryDB and StateDB with fake data,
// emulating the task of the synchronizer in order to have data to be returned
// by the API endpoints that will be tested
func TestMain(m *testing.M) {
	// Initializations
	// Swagger
	router := swagger.NewRouter().WithSwaggerFromFile("./swagger.yml")
	// HistoryDB
	db, err := database.InitSQLDB()
	if err != nil {
		panic(err)
	}
	apiConncon := database.NewAPIConnectionController(1, time.Second)
	hdb := historydb.NewHistoryDB(db, db, apiConncon)
	nodeConfig := &historydb.NodeConfig{}
	// Config (smart contract constants)
	chainID := uint64(0)
	_config := getConfigTest(chainID)
	config = configAPI{
		ChainID:         chainID,
		RollupConstants: *newRollupConstants(_config.RollupConstants),
	}

	// API
	apiGin := gin.Default()
	// Reset DB
	test.WipeDB(hdb.DB())

	constants := &historydb.Constants{
		SCConsts: common.SCConsts{
			Rollup: _config.RollupConstants,
		},
		ChainID:    chainID,
		TonAddress: _config.TonAddress,
	}
	if err := hdb.SetConstants(constants); err != nil {
		panic(err)
	}
	if err := hdb.SetNodeConfig(nodeConfig); err != nil {
		panic(err)
	}

	api, err = NewAPI(Config{
		Version:           "test",
		ExplorerEndpoints: true,
		Server:            apiGin,
		HistoryDB:         hdb,
		StateDB:           nil,
		EthClient:         nil,
		ForgerAddress:     nil,
	})
	if err != nil {
		log.Error(err)
		panic(err)
	}

	// Start server
	listener, err := net.Listen("tcp", ":"+apiPort)
	if err != nil {
		panic(err)
	}
	server := &http.Server{Handler: apiGin}
	go func() {
		if err := server.Serve(listener); err != nil &&
			common.Unwrap(err) != http.ErrServerClosed {
			panic(err)
		}
	}()

	// Generate blockchain data with til
	tcc := til.NewContext(chainID, common.RollupConstMaxL1Tx)
	tilCfgExtra := til.ConfigExtra{
		BootCoordAddr: ethCommon.HexToAddress("0xE39fEc6224708f0772D2A74fd3f9055A90E0A9f2"),
		CoordUser:     "A",
	}
	blocks, err := tcc.GenerateBlocks(SetBlockchain)
	if err != nil {
		panic(err)
	}
	err = tcc.FillBlocksExtra(blocks, &tilCfgExtra)
	if err != nil {
		panic(err)
	}
	err = tcc.FillBlocksForgedL1UserTxs(blocks)
	if err != nil {
		panic(err)
	}
	AddAdditionalInformation(blocks)

	// Extract til generated data, and add it to HistoryDB
	var commonBlocks []common.Block
	var commonAccounts []common.Account
	var commonL1Txs []common.L1Tx

	for _, block := range blocks {
		// Insert block into HistoryDB
		if err := api.historyDB.AddBlockSCData(&block); err != nil {
			log.Error(err)
			panic(err)
		}
		// Extract data
		commonBlocks = append(commonBlocks, block.Block)
		for _, batch := range block.Rollup.Batches {
			for i := range batch.CreatedAccounts {
				batch.CreatedAccounts[i].Nonce = common.Nonce(i)
				commonAccounts = append(commonAccounts, batch.CreatedAccounts[i])
			}
			// commonBatches = append(commonBatches, batch.Batch)
			commonL1Txs = append(commonL1Txs, batch.L1UserTxs...)
		}
	}
	// Add unforged L1 Tx
	unforgedTx := blocks[len(blocks)-1].Rollup.L1UserTxs[0]
	if unforgedTx.BatchNum != nil {
		panic("Unforged tx batch num should be nil")
	}
	commonL1Txs = append(commonL1Txs, unforgedTx)

	// Generate SC vars and add them to HistoryDB (if needed)
	rollupVars := common.RollupVariables{
		EthBlockNum:           int64(3),
		ForgeL1L2BatchTimeout: int64(44),
		SafeMode:              false,
	}

	stateAPIUpdater, err = stateapiupdater.NewUpdater(hdb, nodeConfig, &common.SCVariables{
		Rollup: rollupVars,
	}, constants, 400)
	if err != nil {
		panic(err)
	}

	// Generate test data, as expected to be received/sended from/to the API
	// testTxs := genTestTxs(commonL1Txs, commonAccounts, commonBlocks)
	// testBatches, testFullBatches := genTestBatches(commonBlocks, commonBatches, testTxs)
	// Add balance and nonce to historyDB
	accounts := genTestAccounts(commonAccounts)
	accUpdates := []common.AccountUpdate{}
	for i := 0; i < len(accounts); i++ {
		balance := new(big.Int)
		balance.SetString(string(*accounts[i].Balance), 10)
		accUpdates = append(accUpdates, common.AccountUpdate{
			EthBlockNum: 0,
			BatchNum:    1,
			Idx:         accounts[i].Idx,
			Nonce:       0,
			Balance:     balance,
		})
		accUpdates = append(accUpdates, common.AccountUpdate{
			EthBlockNum: 0,
			BatchNum:    1,
			Idx:         accounts[i].Idx,
			Nonce:       accounts[i].Nonce,
			Balance:     balance,
		})
	}
	if err := api.historyDB.AddAccountUpdates(accUpdates); err != nil {
		panic(err)
	}
	tc = testCommon{
		blocks: commonBlocks,
		// batches:     testBatches,
		// fullBatches: testFullBatches,
		accounts: accounts,
		// txs:         testTxs,
		router:     router,
		rollupVars: rollupVars,
	}
	if err := api.historyDB.AddAccountUpdates(accUpdates); err != nil {
		panic(err)
	}
	tc = testCommon{
		blocks: commonBlocks,
		// batches:     testBatches,
		// fullBatches: testFullBatches,
		accounts: accounts,
		// txs:         testTxs,
		router:     router,
		rollupVars: rollupVars,
	}

	// Run tests
	result := m.Run()
	// Fake server
	if os.Getenv("FAKE_SERVER") == "yes" {
		for {
			log.Info("Running fake server at " + apiURL + " until ^C is received")
			time.Sleep(30 * time.Second)
		}
	}
	// Stop server
	if err := server.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	if err := db.Close(); err != nil {
		panic(err)
	}
	os.Exit(result)
}

func AddAdditionalInformation(blocks []common.BlockData) {
	for i := range blocks {
		blocks[i].Block.Timestamp = time.Now().Add(time.Second * 13).UTC()
		blocks[i].Block.Hash = ethCommon.BigToHash(big.NewInt(blocks[i].Block.Num))
		for x := range blocks[i].Rollup.Batches {
			for q := range blocks[i].Rollup.Batches[x].CreatedAccounts {
				blocks[i].Rollup.Batches[x].CreatedAccounts[q].Balance =
					big.NewInt(int64(blocks[i].Rollup.Batches[x].CreatedAccounts[q].Idx * 10000000))
			}
			for y := range blocks[i].Rollup.Batches[x].ExitTree {
				blocks[i].Rollup.Batches[x].ExitTree[y].MerkleProof =
					&merkletree.CircomVerifierProof{
						Root: &merkletree.Hash{byte(y), byte(y + 1)},
						Siblings: []*merkletree.Hash{
							merkletree.NewHashFromBigInt(big.NewInt(int64(y) * 10)),
							merkletree.NewHashFromBigInt(big.NewInt(int64(y)*100 + 1)),
							merkletree.NewHashFromBigInt(big.NewInt(int64(y)*1000 + 2))},
						OldKey:   &merkletree.Hash{byte(y * 1), byte(y*1 + 1)},
						OldValue: &merkletree.Hash{byte(y * 2), byte(y*2 + 1)},
						IsOld0:   y%2 == 0,
						Key:      &merkletree.Hash{byte(y * 3), byte(y*3 + 1)},
						Value:    &merkletree.Hash{byte(y * 4), byte(y*4 + 1)},
						Fnc:      y % 2,
					}
			}
		}
	}
}

func doGoodReqPaginated(
	path, order string,
	iterStruct Pendinger,
	appendIter func(res interface{}),
) error {
	var next uint64
	firstIte := true
	expectedTotal := 0
	totalReceived := 0
	for {
		// Calculate fromItem
		iterPath := path
		if !firstIte {
			iterPath += "&fromItem=" + strconv.Itoa(int(next))
		}
		// Call API to get this iteration items
		iterStruct = iterStruct.New()
		if err := doGoodReq(
			"GET", iterPath+"&order="+order, nil,
			iterStruct,
		); err != nil {
			return common.Wrap(err)
		}
		appendIter(iterStruct)
		// Keep iterating?
		remaining, lastID := iterStruct.GetPending()
		if remaining == 0 {
			break
		}
		if order == "DESC" {
			next = lastID - 1
		} else {
			next = lastID + 1
		}
		// Check that the expected amount of items is consistent across iterations
		totalReceived += iterStruct.Len()
		if firstIte {
			firstIte = false
			expectedTotal = totalReceived + int(remaining)
		}
		if expectedTotal != totalReceived+int(remaining) {
			panic(fmt.Sprintf(
				"pagination error, totalReceived + remaining should be %d, but is %d",
				expectedTotal, totalReceived+int(remaining),
			))
		}
	}
	return nil
}

func doGoodReq(method, path string, reqBody io.Reader, returnStruct interface{}) error {
	ctx := context.Background()
	client := &http.Client{}
	httpReq, err := http.NewRequest(method, path, reqBody)
	if err != nil {
		return common.Wrap(err)
	}
	if reqBody != nil {
		httpReq.Header.Add("Content-Type", "application/json")
	}
	route, pathParams, err := tc.router.FindRoute(httpReq.Method, httpReq.URL)
	if err != nil {
		return common.Wrap(err)
	}
	// Validate request against swagger spec
	requestValidationInput := &swagger.RequestValidationInput{
		Request:    httpReq,
		PathParams: pathParams,
		Route:      route,
	}
	if err := swagger.ValidateRequest(ctx, requestValidationInput); err != nil {
		return common.Wrap(err)
	}
	// Do API call
	resp, err := client.Do(httpReq)
	if err != nil {
		return common.Wrap(err)
	}
	if resp.Body == nil && returnStruct != nil {
		return common.Wrap(errors.New("Nil body"))
	}
	//nolint
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return common.Wrap(err)
	}
	if resp.StatusCode != 200 {
		return common.Wrap(fmt.Errorf("%d response. Body: %s", resp.StatusCode, string(body)))
	}
	if returnStruct == nil {
		return nil
	}
	// Unmarshal body into return struct
	if err := json.Unmarshal(body, returnStruct); err != nil {
		log.Error("invalid json: " + string(body))
		return common.Wrap(err)
	}
	// log.Info(string(body))
	// Validate response against swagger spec
	responseValidationInput := &swagger.ResponseValidationInput{
		RequestValidationInput: requestValidationInput,
		Status:                 resp.StatusCode,
		Header:                 resp.Header,
	}
	responseValidationInput = responseValidationInput.SetBodyBytes(body)
	return swagger.ValidateResponse(ctx, responseValidationInput)
}

func doBadReq(method, path string, reqBody io.Reader, expectedResponseCode int) error {
	ctx := context.Background()
	client := &http.Client{}
	httpReq, _ := http.NewRequest(method, path, reqBody)
	httpReq.Header.Add("Content-Type", "application/json")
	route, pathParams, err := tc.router.FindRoute(httpReq.Method, httpReq.URL)
	if err != nil {
		return common.Wrap(err)
	}
	// Validate request against swagger spec
	requestValidationInput := &swagger.RequestValidationInput{
		Request:    httpReq,
		PathParams: pathParams,
		Route:      route,
	}
	if err := swagger.ValidateRequest(ctx, requestValidationInput); err != nil {
		if expectedResponseCode != 400 {
			return common.Wrap(err)
		}
		log.Warn("The request does not match the API spec")
	}
	// Do API call
	resp, err := client.Do(httpReq)
	if err != nil {
		return common.Wrap(err)
	}
	if resp.Body == nil {
		return common.Wrap(errors.New("Nil body"))
	}
	//nolint
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return common.Wrap(err)
	}
	if resp.StatusCode != expectedResponseCode {
		return common.Wrap(fmt.Errorf("Unexpected response code: %d. Body: %s", resp.StatusCode, string(body)))
	}
	// Validate response against swagger spec
	responseValidationInput := &swagger.ResponseValidationInput{
		RequestValidationInput: requestValidationInput,
		Status:                 resp.StatusCode,
		Header:                 resp.Header,
	}
	responseValidationInput = responseValidationInput.SetBodyBytes(body)
	return swagger.ValidateResponse(ctx, responseValidationInput)
}

func doSimpleReq(method, endpoint string) (string, error) {
	client := &http.Client{}
	httpReq, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return "", common.Wrap(err)
	}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", common.Wrap(err)
	}
	//nolint
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", common.Wrap(err)
	}
	return string(body), nil
}

// test helpers

func getTimestamp(blockNum int64, blocks []common.Block) time.Time {
	for i := 0; i < len(blocks); i++ {
		if blocks[i].Num == blockNum {
			return blocks[i].Timestamp
		}
	}
	panic("timesamp not found")
}

func getAccountByIdx(idx common.AccountIdx, accs []common.Account) *common.Account {
	for _, acc := range accs {
		if acc.Idx == idx {
			return &acc
		}
	}
	panic("account not found")
}

func getBlockByNum(ethBlockNum int64, blocks []common.Block) common.Block {
	for _, b := range blocks {
		if b.Num == ethBlockNum {
			return b
		}
	}
	panic("block not found")
}
