package client

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/golang/protobuf/proto"
	peertypes "github.com/liubaninc/sdk-go/modules/peer/types"
)

func (c Client) GetPeer(index string) (*peertypes.QueryGetPeerResponse, error) {
	req := &peertypes.QueryGetPeerRequest{
		Index: index,
	}
	bytes, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	value, err := c.query("/liubaninc.m1.peer.Query/Peer", bytes)
	if err != nil {
		return nil, err
	}

	var resp peertypes.QueryGetPeerResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c Client) GetPeerAll(key []byte, offset uint64, limit uint64, countTotal bool) (*peertypes.QueryAllPeerResponse, error) {
	req := &peertypes.QueryAllPeerRequest{
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
	value, err := c.query("/liubaninc.m1.peer.Query/PeerAll", bytes)
	if err != nil {
		return nil, err
	}

	var resp peertypes.QueryAllPeerResponse
	if err := proto.Unmarshal(value, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
