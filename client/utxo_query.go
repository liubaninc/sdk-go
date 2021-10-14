package client

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/golang/protobuf/proto"
	utxotypes "github.com/liubaninc/m1/x/utxo/types"
)

func (c Client) GetToken(base string) (*utxotypes.QueryGetTokenResponse, error) {
	req := &utxotypes.QueryGetTokenRequest{
		Base: base,
	}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/liubaninc.m1.utxo.Query/Token", bytes)
	if err != nil {
		return nil, err
	}

	var resp utxotypes.QueryGetTokenResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetTokenAll(key []byte, offset uint64, limit uint64, countTotal bool) (*utxotypes.QueryAllTokenResponse, error) {
	req := &utxotypes.QueryAllTokenRequest{
		Pagination: &query.PageRequest{
			Key:        key,
			Offset:     offset,
			Limit:      limit,
			CountTotal: countTotal,
		},
	}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/liubaninc.m1.utxo.Query/TokenAll", bytes)
	if err != nil {
		return nil, err
	}

	var resp utxotypes.QueryAllTokenResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
