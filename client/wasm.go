package client

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto/tmhash"
)

type TxSubmitBuilder struct {
	SubmitRequest
}

func (t *TxSubmitBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxSubmitBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxSubmitBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxSubmitBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxSubmitBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxSubmitBuilder) SetCode(code []byte) {
	t.Code = code
}

func (t *TxSubmitBuilder) SetRuntime(runtime string) {
	t.Runtime = runtime
}

func (t *TxSubmitBuilder) SetModule(module string) {
	t.Module = module
}

func (t *TxSubmitBuilder) SetName(name string) {
	t.Name = name
}

func (t *TxSubmitBuilder) SetSrc(src string) {
	t.Src = src
}

func (t *TxSubmitBuilder) SetAbi(abi string) {
	t.Abi = abi
}

func (t *TxSubmitBuilder) SetApprovers(approvers []string) {
	t.Approvers = approvers
}

func (t *TxSubmitBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxSubmit"
}

func (t *TxSubmitBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}

type TxApproveBuilder struct {
	ApproveRequest
}

func (t *TxApproveBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxApproveBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxApproveBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxApproveBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxApproveBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxApproveBuilder) SetCode(code []byte) {
	t.CodeHash = fmt.Sprintf("%X", tmhash.Sum(code))
}

func (t *TxApproveBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxApprove"
}

func (t *TxApproveBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}

type TxDeployBuilder struct {
	DeployRequest
}

func (t *TxDeployBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxDeployBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxDeployBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxDeployBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxDeployBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxDeployBuilder) SetName(name string) {
	t.Name = name
}

func (t *TxDeployBuilder) SetCode(code []byte) {
	t.CodeHash = fmt.Sprintf("%X", tmhash.Sum(code))
}

func (t *TxDeployBuilder) SetInitArgs(initArgs string) {
	t.InitArgs = initArgs
}

func (t *TxDeployBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxDeploy"
}

func (t *TxDeployBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}

type TxInvokeBuilder struct {
	InvokeRequest
}

func (t *TxInvokeBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxInvokeBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxInvokeBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxInvokeBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxInvokeBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxInvokeBuilder) SetName(name string) {
	t.Name = name
}

func (t *TxInvokeBuilder) SetMethod(method string) {
	t.Method = method
}

func (t *TxInvokeBuilder) SetArgs(args string) {
	t.Args = args
}

func (t *TxInvokeBuilder) SetAmounts(amounts string) {
	t.Amounts = amounts
}

func (t *TxInvokeBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxInvoke"
}

func (t *TxInvokeBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}

type TxUpgradeBuilder struct {
	UpgradeRequest
}

func (t *TxUpgradeBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxUpgradeBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxUpgradeBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxUpgradeBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxUpgradeBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxUpgradeBuilder) SetName(name string) {
	t.Name = name
}

func (t *TxUpgradeBuilder) SetCode(code []byte) {
	t.CodeHash = fmt.Sprintf("%X", tmhash.Sum(code))
}

func (t *TxUpgradeBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxUpgrade"
}

func (t *TxUpgradeBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}

type TxFreezeBuilder struct {
	FreezeRequest
}

func (t *TxFreezeBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxFreezeBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxFreezeBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxFreezeBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxFreezeBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxFreezeBuilder) SetName(name string) {
	t.Name = name
}

func (t *TxFreezeBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxFreeze"
}

func (t *TxFreezeBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}

type TxUnFreezeBuilder struct {
	UnFreezeRequest
}

func (t *TxUnFreezeBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxUnFreezeBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxUnFreezeBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxUnFreezeBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxUnFreezeBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxUnFreezeBuilder) SetName(name string) {
	t.Name = name
}

func (t *TxUnFreezeBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxUnFreeze"
}

func (t *TxUnFreezeBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}

type TxDestroyBuilder struct {
	DestroyRequest
}

func (t *TxDestroyBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxDestroyBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxDestroyBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxDestroyBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxDestroyBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxDestroyBuilder) SetName(name string) {
	t.Name = name
}

func (t *TxDestroyBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxDestroy"
}

func (t *TxDestroyBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}
