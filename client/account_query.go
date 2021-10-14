package client

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/golang/protobuf/proto"
)

func (c Client) GetAccounts(key []byte, offset uint64, limit uint64, countTotal bool) (*authtypes.QueryAccountsResponse, error) {
	req := &authtypes.QueryAccountsRequest{
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
	value, err := c.query("/cosmos.auth.v1beta1.Query/Accounts", bytes)
	if err != nil {
		return nil, err
	}

	var resp authtypes.QueryAccountsResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetAccount(address string) (*authtypes.QueryAccountResponse, error) {
	req := &authtypes.QueryAccountRequest{
		Address: address,
	}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/cosmos.auth.v1beta1.Query/Account", bytes)
	if err != nil {
		return nil, err
	}

	var resp authtypes.QueryAccountResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetAccountAllBalances(address string) (*banktypes.QueryAllBalancesResponse, error) {
	req := &banktypes.QueryAllBalancesRequest{
		Address: address,
	}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/cosmos.bank.v1beta1.Query/AllBalances", bytes)
	if err != nil {
		return nil, err
	}

	var resp banktypes.QueryAllBalancesResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
