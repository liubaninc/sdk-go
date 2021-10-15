package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgTransfer{}

func NewMsgTransfer(creator string, base *Base) *MsgTransfer {
	return &MsgTransfer{
		Creator: creator,
		Base:    base,
	}
}

func (msg *MsgTransfer) Route() string {
	return RouterKey
}

func (msg *MsgTransfer) Type() string {
	return "Transfer"
}

func (msg *MsgTransfer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransfer) GetSignBytes() []byte {
	return nil
}

func (msg *MsgTransfer) ValidateBasic() error {
	return nil
}
