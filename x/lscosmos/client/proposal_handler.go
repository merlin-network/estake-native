package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/merlin-network/estake-native/x/lscosmos/client/cli"
)

var (
	MinDepositAndFeeChangeProposalHandler      = govclient.NewProposalHandler(cli.NewMinDepositAndFeeChangeCmd)
	EstakeFeeAddressChangeProposalHandler      = govclient.NewProposalHandler(cli.NewEstakeFeeAddressChangeCmd)
	AllowListValidatorSetChangeProposalHandler = govclient.NewProposalHandler(cli.NewAllowListedValidatorSetChangeProposalCmd)
)
