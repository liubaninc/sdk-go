package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utxotypes "github.com/liubaninc/sdk-go/modules/utxo/types"
)

var _ sdk.Msg = &MsgCreatePeer{}

func NewMsgCreatePeer(creator string, base *utxotypes.Base, index string, addr string, moniker string) *MsgCreatePeer {
	return &MsgCreatePeer{
		Creator: creator,
		Base:    base,
		Index:   index,
		Addr:    addr,
		Moniker: moniker,
	}
}

func (msg *MsgCreatePeer) Route() string {
	return RouterKey
}

func (msg *MsgCreatePeer) Type() string {
	return "CreatePeer"
}

func (msg *MsgCreatePeer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePeer) GetSignBytes() []byte {
	return nil
}

func (msg *MsgCreatePeer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePeer{}

func NewMsgUpdatePeer(creator string, base *utxotypes.Base, index string, addr string) *MsgUpdatePeer {
	return &MsgUpdatePeer{
		Creator: creator,
		Base:    base,
		Index:   index,
		Addr:    addr,
	}
}

func (msg *MsgUpdatePeer) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePeer) Type() string {
	return "UpdatePeer"
}

func (msg *MsgUpdatePeer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePeer) GetSignBytes() []byte {
	return nil
}

func (msg *MsgUpdatePeer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePeer{}

func NewMsgDeletePeer(creator string, base *utxotypes.Base, index string) *MsgDeletePeer {
	return &MsgDeletePeer{
		Creator: creator,
		Base:    base,
		Index:   index,
	}
}
func (msg *MsgDeletePeer) Route() string {
	return RouterKey
}

func (msg *MsgDeletePeer) Type() string {
	return "DeletePeer"
}

func (msg *MsgDeletePeer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePeer) GetSignBytes() []byte {
	return nil
}

func (msg *MsgDeletePeer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
