package params

// Default simulation operation weights for messages and gov proposals
const (
	DefaultWeightMsgSend                        int = 100
	DefaultWeightMsgMultiSend                   int = 10
	DefaultWeightMsgSetWithdrawAddress          int = 50
	DefaultWeightMsgWithdrawDelegationReward    int = 50
	DefaultWeightMsgWithdrawValidatorCommission int = 50
	DefaultWeightMsgFundCommunityPool           int = 50
	DefaultWeightMsgDeposit                     int = 100
	DefaultWeightMsgVote                        int = 67
	DefaultWeightMsgUnjail                      int = 100
	DefaultWeightMsgCreateValidator             int = 100
	DefaultWeightMsgEditValidator               int = 5
	DefaultWeightMsgDelegate                    int = 100
	DefaultWeightMsgUndelegate                  int = 100
	DefaultWeightMsgBeginRedelegate             int = 100

	DefaultWeightMsgDummy int = 100

	DefaultWeightMsgLiquidStake   int = 80
	DefaultWeightMsgLiquidUnstake int = 30

	DefaultWeightAddWhitelistValidatorsProposal    int = 50
	DefaultWeightUpdateWhitelistValidatorsProposal int = 5
	DefaultWeightDeleteWhitelistValidatorsProposal int = 5
	DefaultWeightCompleteRedelegationUnbonding     int = 30
	DefaultWeightTallyWithLiquidStaking            int = 30

	DefaultWeightCommunitySpendProposal int = 5
	DefaultWeightTextProposal           int = 5
	DefaultWeightParamChangeProposal    int = 5
)
