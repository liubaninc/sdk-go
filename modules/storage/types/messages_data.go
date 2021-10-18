package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utxotypes "github.com/liubaninc/sdk-go/modules/utxo/types"
)

var _ sdk.Msg = &MsgCreateData{}

func NewMsgCreateData(creator string, base *utxotypes.Base, content string) *MsgCreateData {
	return &MsgCreateData{
		Creator: creator,
		Base:    base,
		Content: content,
	}
}

func (msg *MsgCreateData) Route() string {
	return RouterKey
}

func (msg *MsgCreateData) Type() string {
	return "CreateData"
}

func (msg *MsgCreateData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}

var _ sdk.Msg = &MsgUpdateData{}

func NewMsgUpdateData(creator string, base *utxotypes.Base, id uint64, content string) *MsgUpdateData {
	return &MsgUpdateData{
		Id:      id,
		Base:    base,
		Creator: creator,
		Content: content,
	}
}

func (msg *MsgUpdateData) Route() string {
	return RouterKey
}

func (msg *MsgUpdateData) Type() string {
	return "UpdateData"
}

func (msg *MsgUpdateData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteData{}

func NewMsgDeleteData(creator string, base *utxotypes.Base, id uint64) *MsgDeleteData {
	return &MsgDeleteData{
		Id:      id,
		Base:    base,
		Creator: creator,
	}
}
func (msg *MsgDeleteData) Route() string {
	return RouterKey
}

func (msg *MsgDeleteData) Type() string {
	return "DeleteData"
}

func (msg *MsgDeleteData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
