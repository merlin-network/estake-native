package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/merlin-network/estake-native/x/lscosmos/types"
)

// HandleMinDepositAndFeeChangeProposal changes host chain params for desired min-deposit and protocol fee
func HandleMinDepositAndFeeChangeProposal(ctx sdk.Context, k Keeper, content types.MinDepositAndFeeChangeProposal) error {
	if !k.GetModuleState(ctx) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Module not enabled")
	}

	hostChainParams := k.GetHostChainParams(ctx)
	if hostChainParams.IsEmpty() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "host chain not registered")
	}

	// modify oldData with the new proposal content
	hostChainParams.MinDeposit = content.MinDeposit
	hostChainParams.EstakeParams.EstakeDepositFee = content.EstakeDepositFee
	hostChainParams.EstakeParams.EstakeRestakeFee = content.EstakeRestakeFee
	hostChainParams.EstakeParams.EstakeUnstakeFee = content.EstakeUnstakeFee
	hostChainParams.EstakeParams.EstakeRedemptionFee = content.EstakeRedemptionFee

	k.SetHostChainParams(ctx, hostChainParams)

	return nil
}

// HandleEstakeFeeAddressChangeProposal changes fee collector address
func HandleEstakeFeeAddressChangeProposal(ctx sdk.Context, k Keeper, content types.EstakeFeeAddressChangeProposal) error {
	//Do not check ModuleEnabled state or host chain params here because non-critical proposal and will help not hardcode address inside default genesis

	hostChainParams := k.GetHostChainParams(ctx)

	// modify oldData with the new proposal content
	hostChainParams.EstakeParams.EstakeFeeAddress = content.EstakeFeeAddress

	k.SetHostChainParams(ctx, hostChainParams)

	return nil
}

// HandleAllowListedValidatorSetChangeProposal changes the allowList validator set
func HandleAllowListedValidatorSetChangeProposal(ctx sdk.Context, k Keeper, content types.AllowListedValidatorSetChangeProposal) error {
	if !k.GetModuleState(ctx) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Module not enabled")
	}

	hostChainParams := k.GetHostChainParams(ctx)
	if hostChainParams.IsEmpty() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "host chain not registered")
	}

	if !content.AllowListedValidators.Valid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Allow listed validators is invalid")
	}

	k.SetAllowListedValidators(ctx, content.AllowListedValidators)
	return nil
}
