package client

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"
	"github.com/liubaninc/sdk-go/params"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

var (
	encodingConfig = params.MakeEncodingConfig()
)

type Client struct {
	conn rpcclient.Client
	opts rpcclient.ABCIQueryOptions
}

func New(rpcURL string) (*Client, error) {
	rpc, err := rpchttp.New(rpcURL, "/websocket")
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: rpc,
		opts: rpcclient.DefaultABCIQueryOptions,
	}, nil
}

func MustNew(rpcURL string) *Client {
	client, err := New(rpcURL)
	if err != nil {
		panic(err)
	}
	return client
}

func (c Client) query(path string, bytes []byte) ([]byte, error) {
	queryResp, err := c.conn.ABCIQueryWithOptions(context.Background(), path, bytes, c.opts)
	if err != nil {
		return nil, err
	}
	if queryResp.Response.Code != 0 {
		return nil, fmt.Errorf(queryResp.Response.Log)
	}
	return queryResp.Response.Value, nil
}

func (c Client) WithHeight(height int64) *Client {
	return &Client{
		conn: c.conn,
		opts: rpcclient.ABCIQueryOptions{
			Height: height,
			Prove:  false,
		},
	}
}

func (c Client) GenerateAndBroadcastTx(fromKey cryptotypes.PrivKey, builder TxBuilder, broadcast string) (*sdk.TxResponse, error) {
	var from sdk.AccAddress
	from = fromKey.PubKey().Address().Bytes()
	pubKey := sdk.MustBech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, fromKey.PubKey())

	builder.SetCreator(from.String())

	// 构建未签名交易
	value, err := c.query(builder.Path(), builder.Bytes())
	if err != nil {
		return nil, err
	}
	var unsignedTxResp BuildUnsignedTxResponse
	if err := proto.Unmarshal(value, &unsignedTxResp); err != nil {
		return nil, err
	}

	// 获取签名内容
	signBytesReq := GetSignBytesRequest{
		Tx:     unsignedTxResp.Tx,
		PubKey: pubKey,
	}
	bytes, _ := signBytesReq.Marshal()
	value, err = c.query("/m1.sdk.v1beta1.Service/GetSignBytes", bytes)
	if err != nil {
		return nil, err
	}
	var signBytesResp GetSignBytesResponse
	if err := proto.Unmarshal(value, &signBytesResp); err != nil {
		return nil, err
	}
	signBytes := signBytesResp.SignBytes

	// 本地签名
	signedBytes, err := fromKey.Sign(signBytes)
	if err != nil {
		return nil, err
	}
	// 广播交易
	buildTxReq := BuildTxRequest{
		Tx:       unsignedTxResp.Tx,
		PubKey:   pubKey,
		Sequence: signBytesResp.Sequence,
		SigBytes: [][]byte{signedBytes},
	}
	bytes, _ = buildTxReq.Marshal()
	value, err = c.query("/m1.sdk.v1beta1.Service/BuildTx", bytes)
	if err != nil {
		return nil, err
	}
	var buildTxResp BuildTxResponse
	if err := proto.Unmarshal(value, &buildTxResp); err != nil {
		return nil, err
	}

	//resp, err := c.BroadcastTx(buildTxResp.Raw, broadcast)
	//if err != nil {
	//	return nil, err
	//}
	//return resp.TxResponse, nil

	return c.broadcast(buildTxResp.Raw, broadcast)
}

func (c Client) broadcast(tx []byte, mode string) (*sdk.TxResponse, error) {
	switch mode {
	case flags.BroadcastBlock:
		res, err := c.conn.BroadcastTxCommit(context.Background(), tx)
		return sdk.NewResponseFormatBroadcastTxCommit(res), err
	case flags.BroadcastAsync:
		res, err := c.conn.BroadcastTxAsync(context.Background(), tx)
		return sdk.NewResponseFormatBroadcastTx(res), err
	case flags.BroadcastSync:
		res, err := c.conn.BroadcastTxSync(context.Background(), tx)
		return sdk.NewResponseFormatBroadcastTx(res), err
	default:
		return nil, fmt.Errorf("invalid broadcast %s; supported types: sync, async, block", mode)
	}
}

type TxBuilder interface {
	Path() string
	Bytes() []byte
	SetCreator(string)
}
