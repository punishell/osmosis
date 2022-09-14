package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/osmosis-labs/osmosis/v12/x/validator-preference/types"
)

type msgServer struct {
	keeper *Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

var _ types.MsgServer = msgServer{}

func (server msgServer) CreateValidatorSetPreference(goCtx context.Context, msg *types.MsgValidatorSetPreference) (*types.MsgValidatorSetPreferenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: might want to check if a user already have a validator-set created

	total_weight := sdk.NewDec(0)
	for _, val := range msg.Preferences {
		// validation checks making sure the weights add up to 1 and also the validator given is correct
		vals, err := sdk.ValAddressFromBech32(val.ValOperAddress)
		if err != nil {
			return nil, fmt.Errorf("validator not formatted")
		}

		_, found := server.keeper.stakingKeeper.GetValidator(ctx, vals)
		if !found {
			return nil, fmt.Errorf("validator address kdoesnot exist")
		}

		total_weight = total_weight.Add(val.Weight)
	}

	if total_weight != sdk.NewDec(1) {
		return nil, fmt.Errorf("The weights allocated to the validators do not add up to 1")
	}

	server.keeper.SetValidatorSetPreferences(ctx, *msg)

	return &types.MsgValidatorSetPreferenceResponse{}, nil
}

func (server msgServer) StakeToValidatorSet(goCtx context.Context, msg *types.MsgStakeToValidatorSet) (*types.MsgValidatorSetPreferenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// loop through the validatorSetPreference and delegate the proportion of the tokens based on weights
	// user account address
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	//tokens := msg.Coins
	//validatorPreferences := msg.Preferences

	//newShares, err := server.keeper.stakingKeeper.Delegate(ctx, owner)

	return nil, nil

}
