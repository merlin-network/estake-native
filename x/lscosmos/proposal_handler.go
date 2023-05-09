package lscosmos

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/merlin-network/estake-native/x/lscosmos/keeper"
	"github.com/merlin-network/estake-native/x/lscosmos/types"
)

// NewLSCosmosProposalHandler creates a new governance Handler for lscosmos module
func NewLSCosmosProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.MinDepositAndFeeChangeProposal:
			return keeper.HandleMinDepositAndFeeChangeProposal(ctx, k, *c)
		case *types.EstakeFeeAddressChangeProposal:
			return keeper.HandleEstakeFeeAddressChangeProposal(ctx, k, *c)
		case *types.AllowListedValidatorSetChangeProposal:
			return keeper.HandleAllowListedValidatorSetChangeProposal(ctx, k, *c)

		default:
			return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized proposal content type: %T", c)
		}
	}
}
