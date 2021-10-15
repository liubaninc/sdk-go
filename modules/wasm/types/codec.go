package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgDestroy{}, "wasm/Destroy", nil)

	cdc.RegisterConcrete(&MsgInvoke{}, "wasm/Invoke", nil)

	cdc.RegisterConcrete(&MsgUnfreeze{}, "wasm/Unfreeze", nil)

	cdc.RegisterConcrete(&MsgFreeze{}, "wasm/Freeze", nil)

	cdc.RegisterConcrete(&MsgApprove{}, "wasm/Approve", nil)

	cdc.RegisterConcrete(&MsgSubmit{}, "wasm/Submit", nil)

	cdc.RegisterConcrete(&MsgUpgrade{}, "wasm/Upgrade", nil)

	cdc.RegisterConcrete(&MsgDeploy{}, "wasm/Deploy", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDestroy{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInvoke{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnfreeze{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFreeze{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApprove{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpgrade{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeploy{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
