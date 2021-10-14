package client

import (
	"context"
	"encoding/hex"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tendermint/tendermint/libs/log"
	"time"
)

type SyncedHandler interface {
	HandleGenesisTxs(txs []*GetTxResponse) error
	HandlePrevBlock(block *GetBlockByHeightResponse) error
	HandleBlock(block *GetBlockByHeightResponse, txs []*GetTxResponse) error
	Logger() log.Logger
}

// 同步数据
func (c *Client) Run(ctx context.Context, start int64, handler SyncedHandler) {
	var errCount int64
	sleepFun := func(err error, key string, value interface{}) {
		errCount++
		handler.Logger().Error("synced", key, value, "error", err)
		time.Sleep(time.Duration(errCount) * time.Second)
	}

	// 默认从1开始
	if start == 0 {
		start = 1
	}
	if start == 1 {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			if err := handler.HandleGenesisTxs(nil); err != nil {
				sleepFun(err, "HandleGenesisTxs", 0)
				continue
			}
			break
		}
	} else if start > 1 {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			block, err := c.GetBlockByHeight(start - 1)
			if err != nil {
				sleepFun(err, "GetBlockByHeight", start-1)
				continue
			}
			if err := handler.HandlePrevBlock(block); err != nil {
				panic(err)
			}
			break
		}
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		block, err := c.GetBlockByHeight(start)
		if err != nil {
			sleepFun(err, "GetBlockByHeight", start)
			continue
		}
		txs := make([]*GetTxResponse, len(block.Block.Data.Txs))
		for i, txBytes := range block.Block.Data.Txs {
			hash := hex.EncodeToString(tmhash.Sum(txBytes))
			tx, err := c.GetTx(hash)
			if err != nil {
				sleepFun(err, "GetTx", start)
				continue
			}
			txs[i] = tx
		}

		// 处理区块链
		if err := handler.HandleBlock(block, txs); err != nil {
			sleepFun(err, "HandleBlock", start)
			continue
		}

		// 处理世界状态
		start++
	}
}
