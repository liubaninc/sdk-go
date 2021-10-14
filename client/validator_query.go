package client

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/golang/protobuf/proto"
	validatortypes "github.com/liubaninc/m1/x/validator/types"
)

func (c Client) GetValidator(address string) (*validatortypes.QueryGetValidatorResponse, error) {
	req := &validatortypes.QueryGetValidatorRequest{
		Address: address,
	}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/liubaninc.m1.validator.Query/Validator", bytes)
	if err != nil {
		return nil, err
	}

	var resp validatortypes.QueryGetValidatorResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetValidatorAll(key []byte, offset uint64, limit uint64, countTotal bool) (*validatortypes.QueryAllValidatorResponse, error) {
	req := &validatortypes.QueryAllValidatorRequest{
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
	value, err := c.query("/liubaninc.m1.validator.Query/ValidatorAll", bytes)
	if err != nil {
		return nil, err
	}

	var resp validatortypes.QueryAllValidatorResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
