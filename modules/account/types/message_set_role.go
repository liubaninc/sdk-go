package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	utxotypes "github.com/liubaninc/sdk-go/modules/utxo/types"
)

var _ sdk.Msg = &MsgSetRole{}

func NewMsgSetRole(creator string, base *utxotypes.Base, address string, roles []string) *MsgSetRole {
	return &MsgSetRole{
		Creator: creator,
		Base:    base,
		Address: address,
		Roles:   roles,
	}
}

func (msg *MsgSetRole) Route() string {
	return RouterKey
}

func (msg *MsgSetRole) Type() string {
	return "SetRole"
}

func (msg *MsgSetRole) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetRole) GetSignBytes() []byte {
	return nil
}

func (msg *MsgSetRole) ValidateBasic() error {
	return nil
}
