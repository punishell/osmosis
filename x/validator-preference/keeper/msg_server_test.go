package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/osmosis-labs/osmosis/v12/x/validator-preference/types"
)

func (suite *KeeperTestSuite) TestCreateValidatorSetPreference() {
	type param struct {
		owner       string
		preferences []types.ValidatorPreference
	}

	tests := []struct {
		name       string
		param      param
		expectPass bool
	}{
		{
			name: "creation of new validator set",
			param: param{
				owner: "",
				preferences: []types.ValidatorPreference{
					{
						ValOperAddress: "",
						Weight:         sdk.NewDecWithPrec(0, 5),
					},
					{
						ValOperAddress: "",
						Weight:         sdk.NewDecWithPrec(0, 3),
					},
					{
						ValOperAddress: "",
						Weight:         sdk.NewDecWithPrec(0, 2),
					},
				},
			},
			expectPass: true,
		},
		{
			name: "create validator set with unknown validator address",
			param: param{
				owner: "",
				preferences: []types.ValidatorPreference{
					{
						ValOperAddress: "",
						Weight:         sdk.NewDec(1),
					},
				},
			},
			expectPass: false,
		},
		{
			name: "create validator set with weights != 1",
			param: param{
				owner: "",
				preferences: []types.ValidatorPreference{
					{
						ValOperAddress: "",
						Weight:         sdk.NewDec(1),
					},
				},
			},
			expectPass: false,
		},
		{
			name: "creator validator set with invalid Owner",
			param: param{
				owner: "",
				preferences: []types.ValidatorPreference{
					{
						ValOperAddress: "",
						Weight:         sdk.NewDec(1),
					},
				},
			},
			expectPass: false,
		},
	}

	for _, test := range tests {

	}
}

func (suite *KeeperTestSuite) TestStakeToValidatorSet() {

}

func (suite *KeeperTestSuite) TestUnStakeFromValidatorSet() {

}
