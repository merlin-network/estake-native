package lscosmos_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/stretchr/testify/suite"

	"github.com/merlin-network/estake-native/app"
	"github.com/merlin-network/estake-native/app/helpers"
	"github.com/merlin-network/estake-native/x/lscosmos"
)

type HandlerTestSuite struct {
	suite.Suite

	app        *app.EstakeApp
	ctx        sdk.Context
	govHandler govtypes.Handler
}

func (suite *HandlerTestSuite) SetupTest() {
	_, estakeApp, ctx := helpers.CreateTestApp(suite.T())
	suite.app = &estakeApp
	suite.ctx = ctx
	suite.govHandler = lscosmos.NewLSCosmosProposalHandler(suite.app.LSCosmosKeeper)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
