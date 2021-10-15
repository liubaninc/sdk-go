package client

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/golang/protobuf/proto"
)

func (c Client) GetGenesisTxs() ([]*tx.GetTxResponse, error) {
	resultGenesis, err := c.conn.Genesis(context.Background())
	if err != nil {
		return nil, err
	}
	appState, err := types.GenesisStateFromGenDoc(*resultGenesis.Genesis)
	if err != nil {
		return nil, err
	}

	genState := types.GetGenesisStateFromAppState(encodingConfig.Marshaler, appState)
	txs := make([]*tx.GetTxResponse, len(genState.GenTxs))
	for i, gtx := range genState.GenTxs {
		txr, err := encodingConfig.TxConfig.TxDecoder()(gtx)
		if err != nil {
			panic(err)
		}
		txs[i].Tx = txr.(*tx.Tx)
	}
	return txs, nil
}

func (c Client) GetNodeInfo() (*tmservice.GetNodeInfoResponse, error) {
	req := &tmservice.GetNodeInfoRequest{}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/cosmos.base.tendermint.v1beta1.Service/GetNodeInfo", bytes)
	if err != nil {
		return nil, err
	}

	var resp tmservice.GetNodeInfoResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetSyncing() (*tmservice.GetSyncingResponse, error) {
	req := &tmservice.GetSyncingRequest{}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/cosmos.base.tendermint.v1beta1.Service/GetSyncing", bytes)
	if err != nil {
		return nil, err
	}

	var resp tmservice.GetSyncingResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetLatestBlock() (*tmservice.GetLatestBlockResponse, error) {
	req := &tmservice.GetLatestBlockRequest{}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/cosmos.base.tendermint.v1beta1.Service/GetLatestBlock", bytes)
	if err != nil {
		return nil, err
	}

	var resp tmservice.GetLatestBlockResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetBlockByHeight(height int64) (*tmservice.GetBlockByHeightResponse, error) {
	req := &tmservice.GetBlockByHeightRequest{
		Height: height,
	}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/cosmos.base.tendermint.v1beta1.Service/GetBlockByHeight", bytes)
	if err != nil {
		return nil, err
	}

	var resp tmservice.GetBlockByHeightResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetLatestValidatorSet() (*tmservice.GetLatestValidatorSetResponse, error) {
	req := &tmservice.GetLatestValidatorSetRequest{}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/cosmos.base.tendermint.v1beta1.Service/GetLatestValidatorSet", bytes)
	if err != nil {
		return nil, err
	}

	var resp tmservice.GetLatestValidatorSetResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetValidatorSetByHeight(height int64) (*tmservice.GetValidatorSetByHeightResponse, error) {
	req := &tmservice.GetValidatorSetByHeightRequest{
		Height: height,
	}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/cosmos.base.tendermint.v1beta1.Service/GetValidatorSetByHeight", bytes)
	if err != nil {
		return nil, err
	}

	var resp tmservice.GetValidatorSetByHeightResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetTx(hash string) (*tx.GetTxResponse, error) {
	req := &tx.GetTxRequest{
		Hash: hash,
	}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/cosmos.tx.v1beta1.Service/GetTx", bytes)
	if err != nil {
		return nil, err
	}
	var resp tx.GetTxResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetTxsEvent(events []string, offset uint64, limit uint64, asc bool) (*tx.GetTxsEventResponse, error) {
	orderBy := tx.OrderBy_ORDER_BY_DESC
	if asc {
		orderBy = tx.OrderBy_ORDER_BY_ASC
	}
	req := &tx.GetTxsEventRequest{
		Events: events,
		Pagination: &query.PageRequest{
			Limit:  limit,
			Offset: offset,
		},
		OrderBy: orderBy,
	}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/cosmos.tx.v1beta1.Service/GetTxsEvent", bytes)
	if err != nil {
		return nil, err
	}
	var resp tx.GetTxsEventResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) BroadcastTx(txBytes []byte, mode string) (*tx.BroadcastTxResponse, error) {
	m := tx.BroadcastMode_BROADCAST_MODE_BLOCK
	switch mode {
	case flags.BroadcastBlock:
		m = tx.BroadcastMode_BROADCAST_MODE_BLOCK
	case flags.BroadcastAsync:
		m = tx.BroadcastMode_BROADCAST_MODE_ASYNC
	case flags.BroadcastSync:
		m = tx.BroadcastMode_BROADCAST_MODE_SYNC
	default:
		return nil, fmt.Errorf("invalid broadcast %s; supported types: sync, async, block", mode)
	}
	req := &tx.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    m,
	}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/cosmos.tx.v1beta1.Service/BroadcastTx", bytes)
	if err != nil {
		return nil, err
	}
	var resp tx.BroadcastTxResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
