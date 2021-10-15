package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utxotypes "github.com/liubaninc/sdk-go/modules/utxo/types"
)

var _ sdk.Msg = &MsgDestroy{}

func NewMsgDestroy(creator string, base *utxotypes.Base, name string) *MsgDestroy {
	return &MsgDestroy{
		Creator: creator,
		Base:    base,
		Name:    name,
	}
}

func (msg *MsgDestroy) Route() string {
	return RouterKey
}

func (msg *MsgDestroy) Type() string {
	return "Destroy"
}

func (msg *MsgDestroy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDestroy) GetSignBytes() []byte {
	return nil
}

func (msg *MsgDestroy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
