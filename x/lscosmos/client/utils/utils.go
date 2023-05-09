package utils

import (
	"os"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/merlin-network/estake-native/v2/x/lscosmos/types"
)

// EstakeParams defines the fees and address for register host chain proposal's EstakeParams
type EstakeParams struct {
	EstakeDepositFee    string `json:"estake_deposit_fee" yaml:"estake_deposit_fee"`
	EstakeRestakeFee    string `json:"estake_restake_fee" yaml:"estake_restake_fee"`
	EstakeUnstakeFee    string `json:"estake_unstake_fee" yaml:"estake_unstake_fee"`
	EstakeRedemptionFee string `json:"estake_redemption_fee" yaml:"estake_redemption_fee"`
	EstakeFeeAddress    string `json:"estake_fee_address" yaml:"estake_fee_address"`
}

// MinDepositAndFeeChangeProposalJSON defines a MinDepositAndFeeChangeProposal JSON input to be parsed
// from a JSON file. Deposit is used by gov module to change status of proposal.
type MinDepositAndFeeChangeProposalJSON struct {
	Title               string `json:"title" yaml:"title"`
	Description         string `json:"description" yaml:"description"`
	MinDeposit          string `json:"min_deposit" yaml:"min_deposit"`
	EstakeDepositFee    string `json:"estake_deposit_fee" yaml:"estake_deposit_fee"`
	EstakeRestakeFee    string `json:"estake_restake_fee" yaml:"estake_restake_fee"`
	EstakeUnstakeFee    string `json:"estake_unstake_fee" yaml:"estake_unstake_fee"`
	EstakeRedemptionFee string `json:"estake_redemption_fee" yaml:"estake_redemption_fee"`
	Deposit             string `json:"deposit" yaml:"deposit"`
}

// NewMinDepositAndFeeChangeJSON returns MinDepositAndFeeChangeProposalJSON struct with input values
func NewMinDepositAndFeeChangeJSON(title, description, minDeposit, estakeDepositFee, estakeRestakeFee,
	estakeUnstakeFee, estakeRedemptionFee, deposit string) MinDepositAndFeeChangeProposalJSON {
	return MinDepositAndFeeChangeProposalJSON{
		Title:               title,
		Description:         description,
		MinDeposit:          minDeposit,
		EstakeDepositFee:    estakeDepositFee,
		EstakeRestakeFee:    estakeRestakeFee,
		EstakeUnstakeFee:    estakeUnstakeFee,
		EstakeRedemptionFee: estakeRedemptionFee,
		Deposit:             deposit,
	}

}

// ParseMinDepositAndFeeChangeProposalJSON reads and parses a MinDepositAndFeeChangeProposalJSON from
// file.
func ParseMinDepositAndFeeChangeProposalJSON(cdc *codec.LegacyAmino, proposalFile string) (MinDepositAndFeeChangeProposalJSON, error) {
	proposal := MinDepositAndFeeChangeProposalJSON{}

	contents, err := os.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}
	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

// EstakeFeeAddressChangeProposalJSON defines a EstakeFeeAddressChangeProposal JSON input to be parsed
// from a JSON file. Deposit is used by gov module to change status of proposal.
type EstakeFeeAddressChangeProposalJSON struct {
	Title            string `json:"title" yaml:"title"`
	Description      string `json:"description" yaml:"description"`
	EstakeFeeAddress string `json:"estake_fee_address" yaml:"estake_fee_address"`
	Deposit          string `json:"deposit" yaml:"deposit"`
}

// NewEstakeFeeAddressChangeProposalJSON returns EstakeFeeAddressChangeProposalJSON struct with input values
func NewEstakeFeeAddressChangeProposalJSON(title, description, estakeFeeAddress, deposit string) EstakeFeeAddressChangeProposalJSON {
	return EstakeFeeAddressChangeProposalJSON{
		Title:            title,
		Description:      description,
		EstakeFeeAddress: estakeFeeAddress,
		Deposit:          deposit,
	}

}

// ParseEstakeFeeAddressChangeProposalJSON reads and parses a EstakeFeeAddressChangeProposal  from
// file.
func ParseEstakeFeeAddressChangeProposalJSON(cdc *codec.LegacyAmino, proposalFile string) (EstakeFeeAddressChangeProposalJSON, error) {
	proposal := EstakeFeeAddressChangeProposalJSON{}

	contents, err := os.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}
	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

// JumpstartTxnJSON defines a JumpStart JSON input to be parsed
// from a JSON file.
type AllowListedValidatorSetChangeProposalJSON struct {
	Title                 string                      `json:"title" yaml:"title"`
	Description           string                      `json:"description" yaml:"description"`
	AllowListedValidators types.AllowListedValidators `json:"allow_listed_validators" yaml:"allow_listed_validators"`
	Deposit               string                      `json:"deposit" yaml:"deposit"`
}

// NewAllowListedValidatorSetChangeProposalJSON returns AllowListedValidatorSetChangeProposalJSON struct with input values
func NewAllowListedValidatorSetChangeProposalJSON(title, description, deposit string, allowListedValidators types.AllowListedValidators) AllowListedValidatorSetChangeProposalJSON {
	return AllowListedValidatorSetChangeProposalJSON{
		Title:                 title,
		Description:           description,
		AllowListedValidators: allowListedValidators,
		Deposit:               deposit,
	}

}

// ParseAllowListedValidatorSetChangeProposalJSON  reads and parses a AllowListedValidatorSetChangeProposalJSON  from
// file.
func ParseAllowListedValidatorSetChangeProposalJSON(cdc *codec.LegacyAmino, proposalFile string) (AllowListedValidatorSetChangeProposalJSON, error) {
	proposal := AllowListedValidatorSetChangeProposalJSON{}

	contents, err := os.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}
	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

// JumpstartTxnJSON defines a Jump start JSON input to be parsed
// from a JSON file.
type JumpstartTxnJSON struct {
	ChainID               string                      `json:"chain_id" yaml:"chain_id"`
	ConnectionID          string                      `json:"connection_id" yaml:"connection_id"`
	TransferChannel       string                      `json:"transfer_channel" yaml:"transfer_channel"`
	TransferPort          string                      `json:"transfer_port" yaml:"transfer_port"`
	BaseDenom             string                      `json:"base_denom" yaml:"base_denom"`
	MintDenom             string                      `json:"mint_denom" yaml:"mint_denom"`
	MinDeposit            string                      `json:"min_deposit" yaml:"min_deposit"`
	AllowListedValidators types.AllowListedValidators `json:"allow_listed_validators" yaml:"allow_listed_validators"`
	EstakeParams          EstakeParams                `json:"estake_params" yaml:"estake_params"`
	HostAccounts          types.HostAccounts          `json:"host_accounts" yaml:"host_accounts"`
}

// ParseJumpstartTxnJSON  reads and parses a JumpstartTxnJSON  from
// file.
func ParseJumpstartTxnJSON(cdc *codec.LegacyAmino, file string) (JumpstartTxnJSON, error) {
	jsonTxn := JumpstartTxnJSON{}

	contents, err := os.ReadFile(file)
	if err != nil {
		return jsonTxn, err
	}
	if err := cdc.UnmarshalJSON(contents, &jsonTxn); err != nil {
		return jsonTxn, err
	}

	return jsonTxn, nil
}
