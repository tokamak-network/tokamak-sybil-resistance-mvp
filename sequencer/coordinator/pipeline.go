package coordinator

import (
	"context"
	"fmt"
	"sync"
	"time"
	"tokamak-sybil-resistance/batchbuilder"
	"tokamak-sybil-resistance/common"
	"tokamak-sybil-resistance/coordinator/prover"
	"tokamak-sybil-resistance/database/historydb"
	"tokamak-sybil-resistance/log"
	"tokamak-sybil-resistance/synchronizer"
	"tokamak-sybil-resistance/txselector"

	"github.com/hermeznetwork/tracerr"
)

type statsVars struct {
	Stats synchronizer.Stats
	Vars  common.SCVariablesPtr
}

type state struct {
	batchNum                     common.BatchNum
	lastScheduledL1BatchBlockNum int64
	lastForgeL1TxsNum            int64
	// lastSlotForged               int64
}

// Pipeline manages the forging of batches with parallel server proofs
type Pipeline struct {
	num    int
	cfg    Config
	consts common.SCConsts

	// state
	state         state
	started       bool
	rw            sync.RWMutex
	errAtBatchNum common.BatchNum
	lastForgeTime time.Time

	prover       prover.Client
	coord        *Coordinator
	txManager    *TxManager
	historyDB    *historydb.HistoryDB
	txSelector   *txselector.TxSelector
	batchBuilder *batchbuilder.BatchBuilder
	purger       *Purger

	stats       synchronizer.Stats
	vars        common.SCVariables
	statsVarsCh chan statsVars

	ctx    context.Context
	wg     sync.WaitGroup
	cancel context.CancelFunc
}

// reset pipeline state
func (p *Pipeline) reset(
	batchNum common.BatchNum,
	stats *synchronizer.Stats,
	vars *common.SCVariables,
) error {
	p.state = state{
		batchNum:                     batchNum,
		lastForgeL1TxsNum:            stats.Sync.LastForgeL1TxsNum,
		lastScheduledL1BatchBlockNum: 0,
		// lastSlotForged:               -1,
	}
	p.stats = *stats
	p.vars = *vars

	// Reset the StateDB in TxSelector and BatchBuilder from the
	// synchronizer only if the checkpoint we reset from either:
	// a. Doesn't exist in the TxSelector/BatchBuilder
	// b. The batch has already been synced by the synchronizer and has a
	//    different MTRoot than the BatchBuilder
	// Otherwise, reset from the local checkpoint.

	// First attempt to reset from local checkpoint if such checkpoint exists
	existsTxSelector, err := p.txSelector.LocalAccountsDB().CheckpointExists(p.state.batchNum)
	if err != nil {
		return common.Wrap(err)
	}
	fromSynchronizerTxSelector := !existsTxSelector
	if err := p.txSelector.Reset(p.state.batchNum, fromSynchronizerTxSelector); err != nil {
		return common.Wrap(err)
	}
	existsBatchBuilder, err := p.batchBuilder.LocalStateDB().CheckpointExists(p.state.batchNum)
	if err != nil {
		return common.Wrap(err)
	}
	fromSynchronizerBatchBuilder := !existsBatchBuilder
	if err := p.batchBuilder.Reset(p.state.batchNum, fromSynchronizerBatchBuilder); err != nil {
		return common.Wrap(err)
	}

	// TODO: implement
	// After reset, check that if the batch exists in the historyDB, the
	// stateRoot matches with the local one, if not, force a reset from
	// synchronizer
	// batch, err := p.historyDB.GetBatch(p.state.batchNum)
	// if common.Unwrap(err) == sql.ErrNoRows {
	// 	// nothing to do
	// } else if err != nil {
	// 	return common.Wrap(err)
	// } else {
	// 	localStateAccountRoot := p.batchBuilder.LocalStateDB().AccountTree.Root().BigInt()
	// 	localStateScoreRoot := p.batchBuilder.LocalStateDB().ScoreTree.Root().BigInt()
	// 	localStateVouchRoot := p.batchBuilder.LocalStateDB().VouchTree.Root().BigInt()

	// 	if batch.AccountStateRoot.Cmp(localStateAccountRoot) != 0 ||
	// 		batch.ScoreStateRoot.Cmp(localStateScoreRoot) != 0 ||
	// 		batch.VouchStateRoot.Cmp(localStateVouchRoot) != 0 {
	// 		log.Debugw(
	// 			"local state roots didn't match the historydb state roots:\n"+
	// 				"localStateAccountRoot: %v vs historydb stateAccountRoot: %v \n"+
	// 				"localStateVouchRoot: %v vs historydb stateVouchRoot: %v \n"+
	// 				"localStateScoreRoot: %v vs historydb stateScoreRoot: %v \n"+
	// 				"Forcing reset from Synchronizer",
	// 			localStateAccountRoot, batch.AccountStateRoot,
	// 			localStateVouchRoot, batch.VouchStateRoot,
	// 			localStateScoreRoot, batch.ScoreStateRoot,
	// 		)
	// 		// StateRoots from synchronizer doesn't match StateRoots
	// 		// from batchBuilder, force a reset from synchronizer
	// 		if err := p.txSelector.Reset(p.state.batchNum, true); err != nil {
	// 			return common.Wrap(err)
	// 		}
	// 		if err := p.batchBuilder.Reset(p.state.batchNum, true); err != nil {
	// 			return common.Wrap(err)
	// 		}
	// 	}
	// }
	return nil
}

func (p *Pipeline) syncSCVars(vars common.SCVariablesPtr) {
	updateSCVars(&p.vars, vars)
}

func (p *Pipeline) getErrAtBatchNum() common.BatchNum {
	p.rw.RLock()
	defer p.rw.RUnlock()
	return p.errAtBatchNum
}

// handleForgeBatch waits for an available proof server, calls p.forgeBatch to
// forge the batch and get the zkInputs, and then  sends the zkInputs to the
// selected proof server so that the proof computation begins.
func (p *Pipeline) handleForgeBatch(
	ctx context.Context,
	batchNum common.BatchNum,
) (batchInfo *BatchInfo, err error) {
	// 1. Wait for an available serverProof (blocking call)
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	// 2. Forge the batch internally (make a selection of txs and prepare
	// all the smart contract arguments)
	var skipReason *string
	batchInfo, skipReason, err = p.forgeBatch(batchNum)
	if ctx.Err() != nil {
		return nil, ctx.Err()
	} else if err != nil {
		log.Errorw("forgeBatch", "err", err)
		return nil, common.Wrap(err)
	} else if skipReason != nil {
		log.Debugw("skipping batch", "batch", batchNum, "reason", *skipReason)
		return nil, common.Wrap(errSkipBatchByPolicy)
	}

	// 3. Send the ZKInputs to the proof server
	batchInfo.ServerProof = p.prover
	batchInfo.ProofStart = time.Now()
	if err := p.sendServerProof(ctx, batchInfo); ctx.Err() != nil {
		return nil, ctx.Err()
	} else if err != nil {
		log.Errorw("sendServerProof", "err", err)
		return nil, common.Wrap(err)
	}
	return batchInfo, nil
}

// sendServerProof sends the circuit inputs to the proof server
func (p *Pipeline) sendServerProof(ctx context.Context, batchInfo *BatchInfo) error {
	// p.cfg.debugBatchStore(batchInfo)

	// Call the selected idle server proof with BatchBuilder output,
	// save server proof info for batchNum
	if err := batchInfo.ServerProof.CalculateProof(ctx, batchInfo.ZKInputs); err != nil {
		return tracerr.Wrap(err)
	}
	return nil
}

// forgeBatch forges the batchNum batch.
func (p *Pipeline) forgeBatch(batchNum common.BatchNum) (
	batchInfo *BatchInfo,
	skipReason *string,
	err error,
) {
	// TODO: investigate if we need this for L2Txs
	// remove transactions from the pool that have been there for too long
	// _, err = p.purger.InvalidateMaybe(
	// 	p.l2DB,
	// 	p.txSelector.LocalAccountsDB(),
	// 	p.stats.Sync.LastBlock.Num,
	// 	int64(batchNum),
	// )
	// if err != nil {
	// 	return nil, nil, tracerr.Wrap(err)
	// }
	// _, err = p.purger.PurgeMaybe(p.l2DB, p.stats.Sync.LastBlock.Num, int64(batchNum))
	// if err != nil {
	// 	return nil, nil, tracerr.Wrap(err)
	// }

	// Structure to accumulate data and metadata of the batch
	now := time.Now()
	batchInfo = &BatchInfo{PipelineNum: p.num, BatchNum: batchNum}
	batchInfo.Debug.StartTimestamp = now
	batchInfo.Debug.StartBlockNum = p.stats.Eth.LastBlock.Num + 1

	// var poolL2Txs []common.PoolL2Tx
	// var discardedL2Txs []common.PoolL2Tx
	var l1UserTxs, l1CoordTxs []common.L1Tx
	var auths [][]byte
	var coordIdxs []common.AccountIdx

	// 1. Decide if we forge L2Tx or L1+L2Tx
	// if p.shouldL1L2Batch(batchInfo) {
	batchInfo.L1Batch = true
	// 2a: L1+L2 txs
	_l1UserTxs, err := p.historyDB.GetUnforgedL1UserTxs(p.state.lastForgeL1TxsNum + 1)
	if err != nil {
		return nil, nil, tracerr.Wrap(err)
	}
	// l1UserFutureTxs are the l1UserTxs that are not being forged
	// in the next batch, but that are also in the queue for the
	// future batches
	l1UserFutureTxs, err := p.historyDB.GetUnforgedL1UserFutureTxs(p.state.lastForgeL1TxsNum + 1)
	if err != nil {
		return nil, nil, tracerr.Wrap(err)
	}

	coordIdxs, auths, l1UserTxs, l1CoordTxs, poolL2Txs, discardedL2Txs, err =
		p.txSelector.GetL1L2TxSelection(p.cfg.TxProcessorConfig, _l1UserTxs, l1UserFutureTxs)
	if err != nil {
		return nil, nil, tracerr.Wrap(err)
	}
	// }
	// else {
	// 	// get l1UserFutureTxs which are all the l1 pending in all the
	// 	// queues
	// 	l1UserFutureTxs, err := p.historyDB.GetUnforgedL1UserFutureTxs(p.state.lastForgeL1TxsNum) //nolint:gomnd
	// 	if err != nil {
	// 		return nil, nil, tracerr.Wrap(err)
	// 	}

	// 	// 2b: only L2 txs
	// 	coordIdxs, auths, l1CoordTxs, poolL2Txs, discardedL2Txs, err =
	// 		p.txSelector.GetL2TxSelection(p.cfg.TxProcessorConfig, l1UserFutureTxs)
	// 	if err != nil {
	// 		return nil, nil, tracerr.Wrap(err)
	// 	}
	// 	l1UserTxs = nil
	// }

	if skip, reason, err := p.forgePolicySkipPostSelection(now,
		l1UserTxs, l1CoordTxs, poolL2Txs, batchInfo); err != nil {
		return nil, nil, tracerr.Wrap(err)
	} else if skip {
		if err := p.txSelector.Reset(batchInfo.BatchNum-1, false); err != nil {
			return nil, nil, tracerr.Wrap(err)
		}
		return nil, &reason, tracerr.Wrap(err)
	}

	if batchInfo.L1Batch {
		p.state.lastScheduledL1BatchBlockNum = p.stats.Eth.LastBlock.Num + 1
		p.state.lastForgeL1TxsNum++
	}

	// 3.  Save metadata from TxSelector output for BatchNum
	batchInfo.L1UserTxs = l1UserTxs
	batchInfo.L1CoordTxs = l1CoordTxs
	batchInfo.L1CoordinatorTxsAuths = auths
	batchInfo.CoordIdxs = coordIdxs
	batchInfo.VerifierIdx = p.cfg.VerifierIdx

	if err := p.l2DB.StartForging(common.TxIDsFromPoolL2Txs(poolL2Txs),
		batchInfo.BatchNum); err != nil {
		return nil, nil, tracerr.Wrap(err)
	}
	if err := p.l2DB.UpdateTxsInfo(discardedL2Txs, batchInfo.BatchNum); err != nil {
		return nil, nil, tracerr.Wrap(err)
	}

	// Invalidate transactions that become invalid because of
	// the poolL2Txs selected.  Will mark as invalid the txs that have a
	// (fromIdx, nonce) which already appears in the selected txs (includes
	// all the nonces smaller than the current one)
	err = p.l2DB.InvalidateOldNonces(idxsNonceFromPoolL2Txs(poolL2Txs), batchInfo.BatchNum)
	if err != nil {
		return nil, nil, tracerr.Wrap(err)
	}

	// 4. Call BatchBuilder with TxSelector output
	configBatch := &batchbuilder.ConfigBatch{
		TxProcessorConfig: p.cfg.TxProcessorConfig,
	}
	zkInputs, err := p.batchBuilder.BuildBatch(coordIdxs, configBatch, l1UserTxs,
		l1CoordTxs, poolL2Txs)
	if err != nil {
		return nil, nil, tracerr.Wrap(err)
	}
	l2Txs, err := common.PoolL2TxsToL2Txs(poolL2Txs) // NOTE: This is a big uggly, find a better way
	if err != nil {
		return nil, nil, tracerr.Wrap(err)
	}
	batchInfo.L2Txs = l2Txs

	// 5. Save metadata from BatchBuilder output for BatchNum
	batchInfo.ZKInputs = zkInputs
	batchInfo.Debug.Status = StatusForged
	p.cfg.debugBatchStore(batchInfo)
	log.Infow("Pipeline: batch forged internally", "batch", batchInfo.BatchNum)

	p.state.lastSlotForged = p.stats.Sync.Auction.CurrentSlot.SlotNum

	return batchInfo, nil, nil
}

func (p *Pipeline) setErrAtBatchNum(batchNum common.BatchNum) {
	p.rw.Lock()
	defer p.rw.Unlock()
	p.errAtBatchNum = batchNum
}

// TODO: implement
// waitServerProof gets the generated zkProof & sends it to the SmartContract
func (p *Pipeline) waitServerProof(ctx context.Context, batchInfo *BatchInfo) error {
	// defer metric.MeasureDuration(metric.WaitServerProof, batchInfo.ProofStart,
	// 	batchInfo.BatchNum.BigInt().String(), strconv.Itoa(batchInfo.PipelineNum))

	// proof, pubInputs, err := batchInfo.ServerProof.GetProof(ctx) // blocking call,
	// // until not resolved don't continue. Returns when the proof server has calculated the proof
	// if err != nil {
	// 	return common.Wrap(err)
	// }
	// batchInfo.Proof = proof
	// batchInfo.PublicInputs = pubInputs
	// batchInfo.ForgeBatchArgs = prepareForgeBatchArgs(batchInfo)
	// batchInfo.Debug.Status = StatusProof
	// p.cfg.debugBatchStore(batchInfo)
	// log.Infow("Pipeline: batch proof calculated", "batch", batchInfo.BatchNum)
	return nil
}

// Start the forging pipeline
func (p *Pipeline) Start(
	batchNum common.BatchNum,
	stats *synchronizer.Stats,
	vars *common.SCVariables,
) error {
	if p.started {
		log.Fatal("Pipeline already started")
	}
	p.started = true

	if err := p.reset(batchNum, stats, vars); err != nil {
		return common.Wrap(err)
	}
	p.ctx, p.cancel = context.WithCancel(context.Background())

	queueSize := 1
	batchChSentServerProof := make(chan *BatchInfo, queueSize)

	p.wg.Add(1)
	go func() {
		timer := time.NewTimer(zeroDuration)
		for {
			select {
			case <-p.ctx.Done():
				log.Infow("Pipeline forgeBatch loop done")
				p.wg.Done()
				return
			case statsVars := <-p.statsVarsCh:
				p.stats = statsVars.Stats
				p.syncSCVars(statsVars.Vars)
			case <-timer.C:
				timer.Reset(p.cfg.ForgeRetryInterval)
				// Once errAtBatchNum != 0, we stop forging
				// batches because there's been an error and we
				// wait for the pipeline to be stopped.
				if p.getErrAtBatchNum() != 0 {
					continue
				}
				batchNum = p.state.batchNum + 1
				batchInfo, err := p.handleForgeBatch(p.ctx, batchNum)
				if p.ctx.Err() != nil {
					// p.revertPoolChanges(batchNum)
					continue
				} else if common.Unwrap(err) == errSkipBatchByPolicy {
					// p.revertPoolChanges(batchNum)
					continue
				} else if err != nil {
					p.setErrAtBatchNum(batchNum)
					p.coord.SendMsg(p.ctx, MsgStopPipeline{
						Reason: fmt.Sprintf(
							"Pipeline.handleForgBatch: %v", err),
						FailedBatchNum: batchNum,
					})
					// p.revertPoolChanges(batchNum)
					continue
				}
				p.lastForgeTime = time.Now()

				p.state.batchNum = batchNum
				select {
				case batchChSentServerProof <- batchInfo:
				case <-p.ctx.Done():
				}
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(zeroDuration)
			}
		}
	}()

	p.wg.Add(1)
	go func() {
		for {
			select {
			case <-p.ctx.Done():
				log.Info("Pipeline waitServerProofSendEth loop done")
				p.wg.Done()
				return
			case batchInfo := <-batchChSentServerProof:
				go func(p *Pipeline, batchInfo *BatchInfo, batchNum common.BatchNum) {
					// Once errAtBatchNum != 0, we stop forging
					// batches because there's been an error and we
					// wait for the pipeline to be stopped.
					if p.getErrAtBatchNum() != 0 {
						// p.revertPoolChanges(batchNum)
						return
					}
					err := p.waitServerProof(p.ctx, batchInfo)
					if p.ctx.Err() != nil {
						// p.revertPoolChanges(batchNum)
						return
					} else if err != nil {
						log.Errorw("waitServerProof", "err", err)
						p.setErrAtBatchNum(batchInfo.BatchNum)
						p.coord.SendMsg(p.ctx, MsgStopPipeline{
							Reason: fmt.Sprintf(
								"Pipeline.waitServerProof: %v", err),
							FailedBatchNum: batchInfo.BatchNum,
						})
						// p.revertPoolChanges(batchNum)
						return
					}
					// We are done with this serverProof, add it back to the pool
					// p.proversPool.Add(p.ctx, batchInfo.ServerProof)
					p.txManager.AddBatch(p.ctx, batchInfo)
				}(p, batchInfo, batchNum)
			}
		}
	}()
	return nil
}

// SetSyncStatsVars is a thread safe method to set the synchronizer Stats
func (p *Pipeline) SetSyncStatsVars(
	ctx context.Context,
	stats *synchronizer.Stats,
	vars *common.SCVariablesPtr,
) {
	select {
	case p.statsVarsCh <- statsVars{Stats: *stats, Vars: *vars}:
	case <-ctx.Done():
	}
}
