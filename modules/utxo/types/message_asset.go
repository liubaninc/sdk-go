package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

var _ sdk.Msg = &MsgAsset{}

func NewMsgAsset(creator string, base *Base, metadata *banktypes.Metadata) *MsgAsset {
	return &MsgAsset{
		Creator:  creator,
		Base:     base,
		Metadata: metadata,
	}
}

func (msg *MsgAsset) Route() string {
	return RouterKey
}

func (msg *MsgAsset) Type() string {
	return "Asset"
}

func (msg *MsgAsset) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAsset) GetSignBytes() []byte {
	return nil
}

func (msg *MsgAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
