package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgIssue{}

func NewMsgIssue(creator string, base *Base) *MsgIssue {
	return &MsgIssue{
		Creator: creator,
		Base:    base,
	}
}

func (msg *MsgIssue) Route() string {
	return RouterKey
}

func (msg *MsgIssue) Type() string {
	return "Issue"
}

func (msg *MsgIssue) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgIssue) GetSignBytes() []byte {
	return nil
}

func (msg *MsgIssue) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
