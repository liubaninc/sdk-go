package client

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/golang/protobuf/proto"
	acctypes "github.com/liubaninc/m1/x/account/types"
)

func (c Client) GetAccountACL(address string) (*acctypes.QueryGetAccountResponse, error) {
	req := &acctypes.QueryGetAccountRequest{
		Address: address,
	}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/liubaninc.m1.account.Query/Account", bytes)
	if err != nil {
		return nil, err
	}

	var resp acctypes.QueryGetAccountResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetAccountACLAll(key []byte, offset uint64, limit uint64, countTotal bool) (*acctypes.QueryAllAccountResponse, error) {
	req := &acctypes.QueryAllAccountRequest{
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
	value, err := c.query("/liubaninc.m1.account.Query/AccountAll", bytes)
	if err != nil {
		return nil, err
	}

	var resp acctypes.QueryAllAccountResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
