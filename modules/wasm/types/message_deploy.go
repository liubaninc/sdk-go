package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utxotypes "github.com/liubaninc/sdk-go/modules/utxo/types"
	"github.com/liubaninc/sdk-go/modules/wasm/xmodel/contract"
)

var _ sdk.Msg = &MsgDeploy{}

func NewMsgDeploy(creator string, base *utxotypes.Base, baseExt *BaseExt, name string, codeHash string, initArgs string, resourceLimits []*contract.ResourceLimit) *MsgDeploy {
	return &MsgDeploy{
		Creator:        creator,
		Base:           base,
		BaseExt:        baseExt,
		Name:           name,
		CodeHash:       codeHash,
		InitArgs:       initArgs,
		ResourceLimits: resourceLimits,
	}
}

func (msg *MsgDeploy) Route() string {
	return RouterKey
}

func (msg *MsgDeploy) Type() string {
	return "Deploy"
}

func (msg *MsgDeploy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeploy) GetSignBytes() []byte {
	return nil
}

func (msg *MsgDeploy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}

func (msg *MsgDeploy) GetInputsExt() []*contract.InputExt {
	if msg.BaseExt == nil {
		return nil
	}
	return msg.BaseExt.InputExts
}
func (msg *MsgDeploy) GetOutputsExt() []*contract.OutputExt {
	if msg.BaseExt == nil {
		return nil
	}
	return msg.BaseExt.OutputExts
}

func (msg *MsgDeploy) GetRequests() []*InvokeRequest {
	return nil
}
