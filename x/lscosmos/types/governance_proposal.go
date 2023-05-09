package types

import (
	"fmt"
	"strings"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeMinDepositAndFeeChange    = "MinDepositAndFeeChange"
	ProposalEstakeFeeAddressChange        = "EstakeFeeAddressChange"
	ProposalAllowListedValidatorSetChange = "AllowListedValidatorSetChange"
)

var (
	_ govtypes.Content = &MinDepositAndFeeChangeProposal{}
	_ govtypes.Content = &EstakeFeeAddressChangeProposal{}
	_ govtypes.Content = &AllowListedValidatorSetChangeProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeMinDepositAndFeeChange)
	govtypes.RegisterProposalType(ProposalEstakeFeeAddressChange)
	govtypes.RegisterProposalType(ProposalAllowListedValidatorSetChange)
}

// NewHostChainParams returns HostChainParams with the input provided
func NewHostChainParams(chainID, connectionID, channel, port, baseDenom, mintDenom, estakefeeAddress string, minDeposit math.Int, estakeDepositFee, estakeRestakeFee, estakeUnstakeFee, estakeRedemptionFee sdktypes.Dec) HostChainParams {
	return HostChainParams{
		ChainID:         chainID,
		ConnectionID:    connectionID,
		TransferChannel: channel,
		TransferPort:    port,
		BaseDenom:       baseDenom,
		MintDenom:       mintDenom,
		MinDeposit:      minDeposit,
		EstakeParams: EstakeParams{
			EstakeDepositFee:    estakeDepositFee,
			EstakeRestakeFee:    estakeRestakeFee,
			EstakeUnstakeFee:    estakeUnstakeFee,
			EstakeRedemptionFee: estakeRedemptionFee,
			EstakeFeeAddress:    estakefeeAddress,
		},
	}
}

// IsEmpty Checks if HostChainParams were initialised
func (c *HostChainParams) IsEmpty() bool {
	if c.TransferChannel == "" ||
		c.TransferPort == "" ||
		c.ConnectionID == "" ||
		c.ChainID == "" ||
		c.BaseDenom == "" ||
		c.MintDenom == "" ||
		c.EstakeParams.EstakeFeeAddress == "" {
		return true
	}
	// can add more, but this should be good enough

	return false
}

// NewMinDepositAndFeeChangeProposal creates a protocol fee and min deposit change proposal.
func NewMinDepositAndFeeChangeProposal(title, description string, minDeposit math.Int, estakeDepositFee,
	estakeRestakeFee, estakeUnstakeFee, estakeRedemptionFee sdktypes.Dec) *MinDepositAndFeeChangeProposal {

	return &MinDepositAndFeeChangeProposal{
		Title:               title,
		Description:         description,
		MinDeposit:          minDeposit,
		EstakeDepositFee:    estakeDepositFee,
		EstakeRestakeFee:    estakeRestakeFee,
		EstakeUnstakeFee:    estakeUnstakeFee,
		EstakeRedemptionFee: estakeRedemptionFee,
	}
}

// GetTitle returns the title of the min-deposit and fee change proposal.
func (m *MinDepositAndFeeChangeProposal) GetTitle() string {
	return m.Title
}

// GetDescription returns the description of the min-deposit and fee change proposal.
func (m *MinDepositAndFeeChangeProposal) GetDescription() string {
	return m.Description
}

// ProposalRoute returns the proposal-route of the min-deposit and fee change proposal.
func (m *MinDepositAndFeeChangeProposal) ProposalRoute() string {
	return RouterKey
}

// ProposalType returns the proposal-type of the min-deposit and fee change proposal.
func (m *MinDepositAndFeeChangeProposal) ProposalType() string {
	return ProposalTypeMinDepositAndFeeChange
}

// ValidateBasic runs basic stateless validity checks
func (m *MinDepositAndFeeChangeProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(m)
	if err != nil {
		return err
	}

	if m.EstakeDepositFee.IsNegative() || m.EstakeDepositFee.GTE(MaxEstakeDepositFee) {
		return errorsmod.Wrapf(ErrInvalidFee, "estake deposit fee must be between %s and %s", sdktypes.ZeroDec(), MaxEstakeDepositFee)
	}

	if m.EstakeRestakeFee.IsNegative() || m.EstakeRestakeFee.GTE(MaxEstakeRestakeFee) {
		return errorsmod.Wrapf(ErrInvalidFee, "estake restake fee must be between %s and %s", sdktypes.ZeroDec(), MaxEstakeRestakeFee)
	}

	if m.EstakeUnstakeFee.IsNegative() || m.EstakeUnstakeFee.GTE(MaxEstakeUnstakeFee) {
		return errorsmod.Wrapf(ErrInvalidFee, "estake unstake fee must be between %s and %s", sdktypes.ZeroDec(), MaxEstakeUnstakeFee)
	}

	if m.EstakeRedemptionFee.IsNegative() || m.EstakeRedemptionFee.GTE(MaxEstakeRedemptionFee) {
		return errorsmod.Wrapf(ErrInvalidFee, "estake redemption fee must be between %s and %s", sdktypes.ZeroDec(), MaxEstakeRedemptionFee)
	}

	if m.MinDeposit.LTE(sdktypes.ZeroInt()) {
		return errorsmod.Wrapf(ErrInvalidDeposit, "min deposit must be positive")
	}

	return nil
}

// String returns the string of proposal details
func (m *MinDepositAndFeeChangeProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`MinDepositAndFeeChange:
Title:                 %s
Description:           %s
MinDeposit:             %s
EstakeDepositFee:	   %s
EstakeRestakeFee: 	   %s
EstakeUnstakeFee: 	   %s
EstakeRedemptionFee:   %s

`,
		m.Title,
		m.Description,
		m.MinDeposit,
		m.EstakeDepositFee,
		m.EstakeRestakeFee,
		m.EstakeUnstakeFee,
		m.EstakeRedemptionFee),
	)
	return b.String()
}

// NewEstakeFeeAddressChangeProposal creates a estake fee  address change proposal.
func NewEstakeFeeAddressChangeProposal(title, description,
	estakeFeeAddress string) *EstakeFeeAddressChangeProposal {
	return &EstakeFeeAddressChangeProposal{
		Title:            title,
		Description:      description,
		EstakeFeeAddress: estakeFeeAddress,
	}
}

// GetTitle returns the title of fee collector estake fee address change proposal.
func (m *EstakeFeeAddressChangeProposal) GetTitle() string {
	return m.Title
}

// GetDescription returns the description of the estake fee address proposal.
func (m *EstakeFeeAddressChangeProposal) GetDescription() string {
	return m.Description
}

// ProposalRoute returns the proposal-route of estake fee address proposal.
func (m *EstakeFeeAddressChangeProposal) ProposalRoute() string {
	return RouterKey
}

// ProposalType returns the proposal-type of estake fee address change proposal.
func (m *EstakeFeeAddressChangeProposal) ProposalType() string {
	return ProposalEstakeFeeAddressChange
}

// ValidateBasic runs basic stateless validity checks
func (m *EstakeFeeAddressChangeProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(m)
	if err != nil {
		return err
	}

	_, err = sdktypes.AccAddressFromBech32(m.EstakeFeeAddress)
	if err != nil {
		return err
	}

	return nil
}

// String returns the string of proposal details
func (m *EstakeFeeAddressChangeProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`EstakeFeeAddressChange:
Title:                 %s
Description:           %s
EstakeFeeAddress: 	   %s

`,
		m.Title,
		m.Description,
		m.EstakeFeeAddress,
	),
	)
	return b.String()
}

// NewAllowListedValidatorSetChangeProposal creates a allowListed validator set change proposal.
func NewAllowListedValidatorSetChangeProposal(title, description string, allowListedValidators AllowListedValidators) *AllowListedValidatorSetChangeProposal {
	return &AllowListedValidatorSetChangeProposal{
		Title:                 title,
		Description:           description,
		AllowListedValidators: allowListedValidators,
	}
}

// GetTitle returns the title of allowListed validator set change proposal.
func (m *AllowListedValidatorSetChangeProposal) GetTitle() string {
	return m.Title
}

// GetDescription returns the description of allowListed validator set change proposal.
func (m *AllowListedValidatorSetChangeProposal) GetDescription() string {
	return m.Description
}

// ProposalRoute returns the proposal-route of allowListed validator set change proposal.
func (m *AllowListedValidatorSetChangeProposal) ProposalRoute() string {
	return RouterKey
}

// ProposalType returns the proposal-type of allowListed validator set change proposal.
func (m *AllowListedValidatorSetChangeProposal) ProposalType() string {
	return ProposalAllowListedValidatorSetChange
}

// ValidateBasic runs basic stateless validity checks
func (m *AllowListedValidatorSetChangeProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(m)
	if err != nil {
		return err
	}

	if !m.AllowListedValidators.Valid() {
		return errorsmod.Wrapf(ErrInValidAllowListedValidators, "allow listed validators is not valid")
	}

	return nil
}

// String returns the string of proposal details
func (m *AllowListedValidatorSetChangeProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`AllowListedValidatorSetChange:
Title:                 %s
Description:           %s
AllowListedValidators: 	   %s

`,
		m.Title,
		m.Description,
		m.AllowListedValidators,
	),
	)
	return b.String()
}
