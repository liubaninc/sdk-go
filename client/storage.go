package client

type TxDataCreateBuilder struct {
	CreateDataRequest
}

func (t *TxDataCreateBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxDataCreateBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxDataCreateBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxDataCreateBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxDataCreateBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxDataCreateBuilder) SetContent(content string) {
	t.Content = content
}

func (t *TxDataCreateBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxDataCreate"
}

func (t *TxDataCreateBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}

type TxDataUpdateBuilder struct {
	UpdateDataRequest
}

func (t *TxDataUpdateBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxDataUpdateBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxDataUpdateBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxDataUpdateBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxDataUpdateBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxDataUpdateBuilder) SetID(id uint64) {
	t.Id = id
}

func (t *TxDataUpdateBuilder) SetContent(content string) {
	t.Content = content
}

func (t *TxDataUpdateBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxDataUpdate"
}

func (t *TxDataUpdateBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}

type TxDataDeleteBuilder struct {
	DeleteDataRequest
}

func (t *TxDataDeleteBuilder) SetCreator(creator string) {
	t.Base.Creator = creator
}

func (t *TxDataDeleteBuilder) SetFees(fees string) {
	t.Base.Fees = fees
}

func (t *TxDataDeleteBuilder) SetMemo(memo string) {
	t.Base.Memo = memo
}

func (t *TxDataDeleteBuilder) SetLock(lock uint64) {
	t.Base.Lock = lock
}

func (t *TxDataDeleteBuilder) SetTimeoutHeight(timeoutHeight uint64) {
	t.Base.TimeoutHeight = timeoutHeight
}

func (t *TxDataDeleteBuilder) SetID(id uint64) {
	t.Id = id
}

func (t *TxDataDeleteBuilder) Path() string {
	return "/m1.sdk.v1beta1.Service/BuildUnsignedTxDataDelete"
}

func (t *TxDataDeleteBuilder) Bytes() []byte {
	bts, err := t.Marshal()
	if err != nil {
		panic(err)
	}
	return bts
}
