package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utxotypes "github.com/liubaninc/sdk-go/modules/utxo/types"
	"github.com/liubaninc/sdk-go/modules/wasm/xmodel/contract"
)

var _ sdk.Msg = &MsgUpgrade{}

func NewMsgUpgrade(creator string, base *utxotypes.Base, baseExt *BaseExt, name string, codeHash string, resourceLimits []*contract.ResourceLimit) *MsgUpgrade {
	return &MsgUpgrade{
		Creator:        creator,
		Base:           base,
		BaseExt:        baseExt,
		Name:           name,
		CodeHash:       codeHash,
		ResourceLimits: resourceLimits,
	}
}

func (msg *MsgUpgrade) Route() string {
	return RouterKey
}

func (msg *MsgUpgrade) Type() string {
	return "Upgrade"
}

func (msg *MsgUpgrade) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpgrade) GetSignBytes() []byte {
	return nil
}

func (msg *MsgUpgrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}

func (msg *MsgUpgrade) GetInputsExt() []*contract.InputExt {
	if msg.BaseExt == nil {
		return nil
	}
	return msg.BaseExt.InputExts
}
func (msg *MsgUpgrade) GetOutputsExt() []*contract.OutputExt {
	if msg.BaseExt == nil {
		return nil
	}
	return msg.BaseExt.OutputExts
}

func (msg *MsgUpgrade) GetRequests() []*InvokeRequest {
	args := map[string]string{
		"contract_name": msg.Name,
		"contract_code": msg.CodeHash,
	}
	bts, _ := json.Marshal(args)
	return []*InvokeRequest{
		{
			ModuleName:     "xkernel",
			ContractName:   "$contract",
			MethodName:     "upgradeContract",
			Args:           string(bts),
			ResourceLimits: msg.ResourceLimits,
		},
	}
}
