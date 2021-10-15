package params

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	accounttypes "github.com/liubaninc/sdk-go/modules/account/types"
	peertypes "github.com/liubaninc/sdk-go/modules/peer/types"
	utxotypes "github.com/liubaninc/sdk-go/modules/utxo/types"
	validatortypes "github.com/liubaninc/sdk-go/modules/validator/types"
	wasmtypes "github.com/liubaninc/sdk-go/modules/wasm/types"
	"github.com/tendermint/spm/cosmoscmd"
)

func MakeEncodingConfig() cosmoscmd.EncodingConfig {
	encodingConfig := cosmoscmd.MakeEncodingConfig(nil)
	authtypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	banktypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	accounttypes.RegisterCodec(encodingConfig.Amino)
	utxotypes.RegisterCodec(encodingConfig.Amino)
	wasmtypes.RegisterCodec(encodingConfig.Amino)
	peertypes.RegisterCodec(encodingConfig.Amino)
	validatortypes.RegisterCodec(encodingConfig.Amino)

	authtypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	banktypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	accounttypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	utxotypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	wasmtypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	peertypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	validatortypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	return encodingConfig
}

func init() {
	cosmoscmd.SetPrefixes("mc")
}
