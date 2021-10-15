package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	utxotypes "github.com/liubaninc/sdk-go/modules/utxo/types"
)

var _ sdk.Msg = &MsgFreeze{}

func NewMsgFreeze(creator string, base *utxotypes.Base, name string) *MsgFreeze {
	return &MsgFreeze{
		Creator: creator,
		Base:    base,
		Name:    name,
	}
}

func (msg *MsgFreeze) Route() string {
	return RouterKey
}

func (msg *MsgFreeze) Type() string {
	return "Freeze"
}

func (msg *MsgFreeze) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFreeze) GetSignBytes() []byte {
	return nil
}

func (msg *MsgFreeze) ValidateBasic() error {
	return nil
}
