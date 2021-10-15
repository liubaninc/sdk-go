package client

import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type TxAssetBuilder struct {
	AssetRequest
}

func (t *TxAssetBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxAssetBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxAssetBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxAssetBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxAssetBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxAssetBuilder) SetAsset(base string, display string, exponent uint32, description string) {
	t.Asset = &banktypes.Metadata{
		Description: description,
		Base:        base,
		Display:     display,
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    base,
				Exponent: 0,
			},
			{
				Denom:    display,
				Exponent: exponent,
			},
		},
	}
}

func (t *TxAssetBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxAsset"
}

func (t *TxAssetBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}

type TxIssueBuilder struct {
	IssueRequest
}

func (t *TxIssueBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxIssueBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxIssueBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxIssueBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxIssueBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxIssueBuilder) AppendReceiver(address string, amounts string) {
	t.Receivers = append(t.Receivers, &Receiver{Address: address, Amounts: amounts})
}

func (t *TxIssueBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxIssue"
}

func (t *TxIssueBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}

type TxTransferBuilder struct {
	TransferRequest
}

func (t *TxTransferBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxTransferBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxTransferBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxTransferBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxTransferBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxTransferBuilder) AppendReceiver(address string, amounts string) {
	t.Receivers = append(t.Receivers, &Receiver{Address: address, Amounts: amounts})
}

func (t *TxTransferBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxTransfer"
}

func (t *TxTransferBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}

type TxBurnBuilder struct {
	BurnRequest
}

func (t *TxBurnBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxBurnBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxBurnBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxBurnBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxBurnBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxBurnBuilder) SetAmounts(amounts string) {
	t.Amounts = amounts
}

func (t *TxBurnBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxBurn"
}

func (t *TxBurnBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}
