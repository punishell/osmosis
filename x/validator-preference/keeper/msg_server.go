package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	// check if a user already have a validator-set created
	_, found := server.keeper.GetValidatorSetPreference(ctx, msg.Owner)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("user %d already has a validator set", msg.Owner))
	}

	total_weight := sdk.NewDec(0)
	for _, val := range msg.Preferences {
		// validation checks making sure the weights add up to 1 and also the validator given is correct
		_, err := server.keeper.ValidateValidator(ctx, val.ValOperAddress)
		if err != nil {
			return nil, err
		}

		total_weight = total_weight.Add(val.Weight)
	}

	if total_weight != sdk.NewDec(1) {
		return nil, fmt.Errorf("The weights allocated to the validators do not add up to 1")
	}

	server.keeper.SetValidatorSetPreferences(ctx, *msg)

	return &types.MsgValidatorSetPreferenceResponse{}, nil
}

func (server msgServer) StakeToValidatorSet(goCtx context.Context, msg *types.MsgStakeToValidatorSet) (*types.MsgStakeToValidatorSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// loop through the validatorSetPreference and delegate the proportion of the tokens based on weights
	// user account address
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	tokenAmt := sdk.NewDec(msg.Coins[0].Amount.Int64())

	for _, val := range msg.Preferences {
		validator, err := server.keeper.ValidateValidator(ctx, val.ValOperAddress)
		if err != nil {
			return nil, err
		}

		// NOTE: it'd be nice if this value was decimal
		amountToStake := val.Weight.Mul(tokenAmt).RoundInt()

		_, err = server.keeper.stakingKeeper.Delegate(ctx, owner, amountToStake, validator.Status, validator, true)
		if err != nil {
			return nil, err
		}
	}

	return &types.MsgStakeToValidatorSetResponse{}, nil
}

// userA stakes {10osmo ValA ValB} -> [{ValA, ValB} -> {0.4, 0.6}] = 10osmo = {4osmo, 6osmo}
// userA unstake {3osmo ValA} -> [{valA, valB} -> {0.142, 0.857}] = 7osmo = {1osmo, 6osmo}
// 	User can also do unstake all
// 	User gives {amount, validators} to unstake from (partial undelegation)
func (server msgServer) UnStakeFromValidatorSet(goCtx context.Context, msg *types.MsgUnStakeFromValidatorSet) (*types.MsgUnStakeFromValidatorSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get the existing validator set preference
	existingSet, found := server.keeper.GetValidatorSetPreference(ctx, msg.Owner)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("user %d doesn't have validator set", msg.Owner))
	}

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	// the total amount the user wants to unstakes
	tokenAmt := sdk.NewDec(msg.Coins[0].Amount.Int64())

	total_weight := sdk.NewDec(0)
	// check if the provided validator set is correct
	for i, val := range msg.Preferences {
		// validation checks making sure the weights add up to 1 and also the validator given is correct
		_, err := server.keeper.ValidateValidator(ctx, val.ValOperAddress)
		if err != nil {
			return nil, err
		}

		// Calculate the amount to unstake based on the weight provided
		amountToUnStake := val.Weight.Mul(tokenAmt)

		// ValidateValidator gurantees that this exist
		valAddr, err := sdk.ValAddressFromBech32(val.ValOperAddress)
		if err != nil {
			return nil, err
		}

		_, err = server.keeper.stakingKeeper.Undelegate(ctx, owner, valAddr, amountToUnStake)
		if err != nil {
			return nil, err
		}

		total_weight = total_weight.Add(val.Weight)

		// updating the existing validator set with new weights
		existingSet[i].Weight = val.Weight
		existingSet[i].ValOperAddress = val.ValOperAddress
	}

	if total_weight != sdk.NewDec(1) {
		return nil, fmt.Errorf("The weights allocated to the validators do not add up to 1")
	}

	server.keeper.SetValidatorSetPreferences(ctx, existingSet)

	return &types.MsgUnStakeFromValidatorSetResponse{}, nil
}
