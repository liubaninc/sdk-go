package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utxotypes "github.com/liubaninc/sdk-go/modules/utxo/types"
)

var _ sdk.Msg = &MsgLeaveValidator{}

func NewMsgLeaveValidator(creator string, base *utxotypes.Base) *MsgLeaveValidator {
	return &MsgLeaveValidator{
		Creator: creator,
		Base:    base,
	}
}

func (msg *MsgLeaveValidator) Route() string {
	return RouterKey
}

func (msg *MsgLeaveValidator) Type() string {
	return "LeaveValidator"
}

func (msg *MsgLeaveValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLeaveValidator) GetSignBytes() []byte {
	return nil
}

func (msg *MsgLeaveValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
