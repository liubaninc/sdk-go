package client

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/golang/protobuf/proto"
)

func (c Client) GetNodeInfo() (*GetNodeInfoResponse, error) {
	req := &GetNodeInfoRequest{}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/cosmos.base.tendermint.v1beta1.Service/GetNodeInfo", bytes)
	if err != nil {
		return nil, err
	}

	var resp GetNodeInfoResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetSyncing() (*GetSyncingResponse, error) {
	req := &GetSyncingRequest{}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/cosmos.base.tendermint.v1beta1.Service/GetSyncing", bytes)
	if err != nil {
		return nil, err
	}

	var resp GetSyncingResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetLatestBlock() (*GetLatestBlockResponse, error) {
	req := &GetLatestBlockRequest{}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/cosmos.base.tendermint.v1beta1.Service/GetLatestBlock", bytes)
	if err != nil {
		return nil, err
	}

	var resp GetLatestBlockResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetBlockByHeight(height int64) (*GetBlockByHeightResponse, error) {
	req := &GetBlockByHeightRequest{
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

	var resp GetBlockByHeightResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetLatestValidatorSet() (*GetLatestValidatorSetResponse, error) {
	req := &GetLatestValidatorSetRequest{}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/cosmos.base.tendermint.v1beta1.Service/GetLatestValidatorSet", bytes)
	if err != nil {
		return nil, err
	}

	var resp GetLatestValidatorSetResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetValidatorSetByHeight(height int64) (*GetValidatorSetByHeightResponse, error) {
	req := &GetValidatorSetByHeightRequest{
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

	var resp GetValidatorSetByHeightResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetTx(hash string) (*GetTxResponse, error) {
	req := &GetTxRequest{
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
	var resp GetTxResponse
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
