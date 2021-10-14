package client

import sdktypes "github.com/liubaninc/m1/sdk"

type TxPeerCreateBuilder struct {
	sdktypes.CreatePeerRequest
}

func (t *TxPeerCreateBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxPeerCreateBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxPeerCreateBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxPeerCreateBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxPeerCreateBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxPeerCreateBuilder) SetIndex(index string) {
	t.Index = index
}

func (t *TxPeerCreateBuilder) SetAddr(addr string) {
	t.Addr = addr
}

func (t *TxPeerCreateBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxPeerCreate"
}

func (t *TxPeerCreateBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}

type TxPeerUpdateBuilder struct {
	sdktypes.UpdatePeerRequest
}

func (t *TxPeerUpdateBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxPeerUpdateBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxPeerUpdateBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxPeerUpdateBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxPeerUpdateBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxPeerUpdateBuilder) SetIndex(index string) {
	t.Index = index
}

func (t *TxPeerUpdateBuilder) SetAddr(addr string) {
	t.Addr = addr
}

func (t *TxPeerUpdateBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxPeerUpdate"
}

func (t *TxPeerUpdateBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}

type TxPeerDeleteBuilder struct {
	sdktypes.DeletePeerRequest
}

func (t *TxPeerDeleteBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxPeerDeleteBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxPeerDeleteBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxPeerDeleteBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxPeerDeleteBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxPeerDeleteBuilder) SetIndex(index string) {
	t.Index = index
}

func (t *TxPeerDeleteBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxPeerDelete"
}

func (t *TxPeerDeleteBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}
