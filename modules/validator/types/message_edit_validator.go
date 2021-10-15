package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utxotypes "github.com/liubaninc/sdk-go/modules/utxo/types"
)

var _ sdk.Msg = &MsgEditValidator{}

func NewMsgEditValidator(creator string, base *utxotypes.Base, description *Description) *MsgEditValidator {
	return &MsgEditValidator{
		Creator:     creator,
		Base:        base,
		Description: description,
	}
}

func (msg *MsgEditValidator) Route() string {
	return RouterKey
}

func (msg *MsgEditValidator) Type() string {
	return "EditValidator"
}

func (msg *MsgEditValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEditValidator) GetSignBytes() []byte {
	return nil
}

func (msg *MsgEditValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
