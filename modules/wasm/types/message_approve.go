package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utxotypes "github.com/liubaninc/sdk-go/modules/utxo/types"
)

var _ sdk.Msg = &MsgApprove{}

func NewMsgApprove(creator string, base *utxotypes.Base, codeHash string) *MsgApprove {
	return &MsgApprove{
		Creator:  creator,
		Base:     base,
		CodeHash: codeHash,
	}
}

func (msg *MsgApprove) Route() string {
	return RouterKey
}

func (msg *MsgApprove) Type() string {
	return "Approve"
}

func (msg *MsgApprove) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApprove) GetSignBytes() []byte {
	return nil
}

func (msg *MsgApprove) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
