package client

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/golang/protobuf/proto"
	wasmtypes "github.com/liubaninc/sdk-go/modules/wasm/types"
)

func (c Client) GetContract(name string) (*wasmtypes.QueryGetContractResponse, error) {
	req := &wasmtypes.QueryGetContractRequest{
		Name: name,
	}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/liubaninc.m1.wasm.Query/Contract", bytes)
	if err != nil {
		return nil, err
	}

	var resp wasmtypes.QueryGetContractResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetContractAll(key []byte, offset uint64, limit uint64, countTotal bool) (*wasmtypes.QueryAllContractResponse, error) {
	req := &wasmtypes.QueryAllContractRequest{
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
	value, err := c.query("/liubaninc.m1.wasm.Query/ContractAll", bytes)
	if err != nil {
		return nil, err
	}

	var resp wasmtypes.QueryAllContractResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetContractCode(hash string) (*wasmtypes.QueryGetContractCodeResponse, error) {
	req := &wasmtypes.QueryGetContractCodeRequest{
		Hash: hash,
	}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/liubaninc.m1.wasm.Query/ContractCode", bytes)
	if err != nil {
		return nil, err
	}

	var resp wasmtypes.QueryGetContractCodeResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetContractCodeAll(key []byte, offset uint64, limit uint64, countTotal bool) (*wasmtypes.QueryAllContractCodeResponse, error) {
	req := &wasmtypes.QueryAllContractCodeRequest{
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
	value, err := c.query("/liubaninc.m1.wasm.Query/ContractCodeAll", bytes)
	if err != nil {
		return nil, err
	}

	var resp wasmtypes.QueryAllContractCodeResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
