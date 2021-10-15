package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type ACLDecorator struct {
	keeper ACLKeeper
}

func NewACLDecorator(keeper ACLKeeper) ACLDecorator {
	return ACLDecorator{keeper: keeper}
}

func (vad ACLDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if vad.keeper != nil {
		for _, msg := range tx.GetMsgs() {
			route := msg.Route()
			for _, acc := range msg.GetSigners() {
				if has, _ := vad.keeper.HasPerm(ctx, acc.String(), route); !has {
					return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s had no perm %s", acc.String(), route)
				}
			}
		}
	}
	return next(ctx, tx, simulate)
}
