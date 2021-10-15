package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utxotypes "github.com/liubaninc/sdk-go/modules/utxo/types"
)

var _ sdk.Msg = &MsgCreateValidator{}

func NewMsgCreateValidator(creator string, base *utxotypes.Base, pubKey string, description *Description) *MsgCreateValidator {
	return &MsgCreateValidator{
		Creator:     creator,
		Base:        base,
		PubKey:      pubKey,
		Description: description,
	}
}

func (msg *MsgCreateValidator) Route() string {
	return RouterKey
}

func (msg *MsgCreateValidator) Type() string {
	return "CreateValidator"
}

func (msg *MsgCreateValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateValidator) GetSignBytes() []byte {
	return nil
}

func (msg *MsgCreateValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
