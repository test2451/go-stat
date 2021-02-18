package observer

import (
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/pieswap/pie-statas/common"
	"github.com/pieswap/pie-statas/executor"
	"github.com/pieswap/pie-statas/model"
	"github.com/pieswap/pie-statas/util"
)

type Observer struct {
	StatasDB    *gorm.DB
	StartHeight int64
	ConfirmNum  int64

	Config   *util.Config
	Executor executor.Executor

	FetchInterval time.Duration
}

// NewObserver returns the observer instance
func NewObserver(stataDB *gorm.DB, cfg *util.Config, executor executor.Executor) *Observer {
	return &Observer{
		StatasDB: stataDB,

		StartHeight: cfg.ChainConfig.BSCStartHeight,
		ConfirmNum:  cfg.ChainConfig.BSCConfirmNum,

		Config:        cfg,
		FetchInterval: time.Duration(cfg.ChainConfig.BSCFetchInterval) * time.Millisecond,
		Executor:      executor,
	}
}

// Start starts the routines of observer
func (ob *Observer) Start() {
	go ob.Fetch(ob.StartHeight)
	go ob.Prune()
	go ob.Alert()
}

// Fetch starts the main routine for fetching blocks of BSC
func (ob *Observer) Fetch(startHeight int64) {
	for {
		curBlockLog, err := ob.GetCurrentBlockLog()
		if err != nil {
			util.Logger.Errorf("get current block log error, err=%s, sleep %d second", err.Error(), ob.FetchInterval.Seconds())
			time.Sleep(ob.FetchInterval)
			continue
		}

		nextHeight := curBlockLog.Height + 1
		if curBlockLog.Height == 0 && startHeight != 0 {
			nextHeight = startHeight
		}

		util.Logger.Infof("fetching block, height=%d", nextHeight)
		err = ob.fetchBlock(curBlockLog.Height, nextHeight, curBlockLog.BlockHash)
		if err != nil {
			util.Logger.Errorf("fetch block error, err=%s", err.Error())
			time.Sleep(ob.FetchInterval)
		}
	}
}

// fetchBlock fetches the next block of BSC and saves it to database. if the next block hash
// does not match to the parent hash, the current block will be deleted for there is a fork.
func (ob *Observer) fetchBlock(curHeight, nextHeight int64, curBlockHash string) error {
	blockAndEventLogs, err := ob.Executor.GetBlockAndTxEvents(nextHeight)
	if err != nil {
		return fmt.Errorf("get block info error, height=%d, err=%s", nextHeight, err.Error())
	}

	parentHash := blockAndEventLogs.ParentBlockHash
	if curHeight != 0 && parentHash != curBlockHash {
		return ob.DeleteBlockAndTxEvents(curHeight)
	} else {
		nextBlockLog := model.BlockLog{
			BlockHash:  blockAndEventLogs.BlockHash,
			ParentHash: parentHash,
			Height:     blockAndEventLogs.Height,
			BlockTime:  blockAndEventLogs.BlockTime,
		}

		err := ob.SaveBlockAndTxEvents(&nextBlockLog, blockAndEventLogs.Events)
		if err != nil {
			return err
		}

		err = ob.UpdateConfirmedNum(nextBlockLog.Height)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteBlockAndTxEvents deletes the block and txs of the given height
func (ob *Observer) DeleteBlockAndTxEvents(height int64) error {
	tx := ob.StatasDB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Where("height = ?", height).Delete(model.BlockLog{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("height = ? and status = ?", height, model.TxStatusInit).Delete(model.TxEventLog{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (ob *Observer) UpdateConfirmedNum(height int64) error {
	err := ob.StatasDB.Model(model.TxEventLog{}).Where("status = ?", model.TxStatusInit).Updates(
		map[string]interface{}{
			"confirmed_num": gorm.Expr("? - height", height+1),
		}).Error
	if err != nil {
		return err
	}

	err = ob.StatasDB.Model(model.TxEventLog{}).Where("status = ? and confirmed_num >= ?",
		model.TxStatusInit, ob.ConfirmNum).Updates(
		map[string]interface{}{
			"status": model.TxStatusConfirmed,
		}).Error
	if err != nil {
		return err
	}

	return nil
}

// Prune prunes the outdated blocks
func (ob *Observer) Prune() {
	for {
		curBlockLog, err := ob.GetCurrentBlockLog()
		if err != nil {
			util.Logger.Errorf("get current block log error, err=%s", err.Error())
			time.Sleep(common.ObserverPruneInterval)

			continue
		}
		err = ob.StatasDB.Where("height < ?", curBlockLog.Height-common.ObserverMaxBlockNumber).Delete(model.BlockLog{}).Error
		if err != nil {
			util.Logger.Infof("prune block logs error, err=%s", err.Error())
		}
		err = ob.StatasDB.Where("height < ?", curBlockLog.Height-common.ObservceMaxTxNumber).Delete(model.TxEventLog{}).Error
		if err != nil {
			util.Logger.Infof("prune block logs error, err=%s", err.Error())
		}
		time.Sleep(common.ObserverPruneInterval)
	}
}

func (ob *Observer) SaveBlockAndTxEvents(blockLog *model.BlockLog, packages []interface{}) error {
	tx := ob.StatasDB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(blockLog).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, pack := range packages {
		if err := tx.Create(pack).Error; err != nil {
			if strings.Contains(err.Error(),"Out of range value"){
				continue
			}else{
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}

// GetCurrentBlockLog returns the highest block log
func (ob *Observer) GetCurrentBlockLog() (*model.BlockLog, error) {
	blockLog := model.BlockLog{}
	err := ob.StatasDB.Order("height desc").First(&blockLog).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &blockLog, nil
}

// Alert sends alerts to tg group if there is no new block fetched in a specific time
func (ob *Observer) Alert() {
	for {
		curChainBlockLog, err := ob.GetCurrentBlockLog()
		if err != nil {
			util.Logger.Errorf("get current block log error, err=%s", err.Error())
			time.Sleep(common.ObserverAlertInterval)

			continue
		}
		if curChainBlockLog.Height > 0 {
			if time.Now().Unix()-curChainBlockLog.CreateTime > ob.Config.AlertConfig.BlockUpdateTimeout {
				msg := fmt.Sprintf("Statas Service: big lagger now, last block fetched at %s, height=%d",
					time.Unix(curChainBlockLog.CreateTime, 0).String(), curChainBlockLog.Height)
				util.SendTelegramMessage(msg)
			}
		}

		time.Sleep(common.ObserverAlertInterval)
	}
}
