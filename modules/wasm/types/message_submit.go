package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utxotypes "github.com/liubaninc/sdk-go/modules/utxo/types"
	"github.com/liubaninc/sdk-go/modules/wasm/xmodel/contract"
)

var _ sdk.Msg = &MsgSubmit{}

func NewMsgSubmit(creator string, base *utxotypes.Base, name string, code []byte, abi string, src string, desc *contract.WasmCodeDesc, approvers []string) *MsgSubmit {
	return &MsgSubmit{
		Creator:   creator,
		Base:      base,
		Name:      name,
		Code:      code,
		Abi:       abi,
		Src:       src,
		Desc:      desc,
		Approvers: approvers,
	}
}

func (msg *MsgSubmit) Route() string {
	return RouterKey
}

func (msg *MsgSubmit) Type() string {
	return "Submit"
}

func (msg *MsgSubmit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmit) GetSignBytes() []byte {
	return nil
}

func (msg *MsgSubmit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
