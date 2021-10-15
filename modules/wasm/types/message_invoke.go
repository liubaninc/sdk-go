package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	utxotypes "github.com/liubaninc/sdk-go/modules/utxo/types"
	"github.com/liubaninc/sdk-go/modules/wasm/xmodel/contract"
)

var _ sdk.Msg = &MsgInvoke{}

func NewMsgInvoke(creator string, base *utxotypes.Base, baseExt *BaseExt, request []*InvokeRequest) *MsgInvoke {
	return &MsgInvoke{
		Creator:  creator,
		Base:     base,
		BaseExt:  baseExt,
		Requests: request,
	}
}

func (msg *MsgInvoke) Route() string {
	return RouterKey
}

func (msg *MsgInvoke) Type() string {
	return "Invoke"
}

func (msg *MsgInvoke) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInvoke) GetSignBytes() []byte {
	return nil
}

func (msg *MsgInvoke) ValidateBasic() error {
	return nil
}

func (msg *MsgInvoke) GetInputsExt() []*contract.InputExt {
	if msg.BaseExt == nil {
		return nil
	}
	return msg.BaseExt.InputExts
}
func (msg *MsgInvoke) GetOutputsExt() []*contract.OutputExt {
	if msg.BaseExt == nil {
		return nil
	}
	return msg.BaseExt.OutputExts
}
