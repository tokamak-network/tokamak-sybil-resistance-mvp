package api

import (
	"errors"
	"tokamak-sybil-resistance/common"
	"tokamak-sybil-resistance/database/historydb"
	"tokamak-sybil-resistance/database/statedb"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// API serves HTTP requests to allow external interaction with the Hermez node
type API struct {
	historyDB     *historydb.HistoryDB
	config        *configAPI
	stateDB       *statedb.StateDB
	hermezAddress ethCommon.Address
	validate      *validator.Validate
}

// Config wraps the parameters needed to start the API
type Config struct {
	Version           string
	ExplorerEndpoints bool
	Server            *gin.Engine
	HistoryDB         *historydb.HistoryDB
	StateDB           *statedb.StateDB
	EthClient         *ethclient.Client
	ForgerAddress     *ethCommon.Address
}

// NewAPI sets the endpoints and the appropriate handlers, but doesn't start the server
func NewAPI(setup Config) (*API, error) {
	// Check input
	if setup.ExplorerEndpoints && setup.HistoryDB == nil {
		return nil, common.Wrap(errors.New("cannot serve Explorer endpoints without HistoryDB"))
	}
	consts, err := setup.HistoryDB.GetConstants()
	if err != nil {
		return nil, err
	}

	a := &API{
		historyDB: setup.HistoryDB,
		config: &configAPI{
			RollupConstants: *newRollupConstants(consts.Rollup),
			ChainID:         consts.ChainID,
		},
		stateDB:       setup.StateDB,
		hermezAddress: consts.TonAddress,
		validate:      nil, //TODO: Add validations
	}

	// Setup http interface
	// middleware, err := metric.PrometheusMiddleware()
	// if err != nil {
	// 	return nil, err
	// }
	// setup.Server.Use(middleware)

	// setup.Server.NoRoute(a.noRoute)

	v1 := setup.Server.Group("/v1")

	v1.GET("/health", gin.WrapH(a.healthRoute(setup.Version, setup.EthClient, setup.ForgerAddress)))

	// Add explorer endpoints
	if setup.ExplorerEndpoints {
		// Account
		v1.GET("/accounts", a.getAccounts)
		// v1.GET("/accounts/:accountIndex", a.getAccount)
		// // Transaction
		// v1.GET("/transactions-history", a.getHistoryTxs)
		// v1.GET("/transactions-history/:id", a.getHistoryTx)
		// // Batches
		// v1.GET("/batches", a.getBatches)
		// v1.GET("/batches/:batchNum", a.getBatch)
		// v1.GET("/full-batches/:batchNum", a.getFullBatch)
	}

	return a, nil
}
