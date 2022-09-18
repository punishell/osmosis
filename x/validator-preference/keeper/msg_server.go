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
		// it'd be nice if this value was decimal
		amountToStake := val.Weight.Mul(tokenAmt).RoundInt()

		vals, err := sdk.ValAddressFromBech32(val.ValOperAddress)
		if err != nil {
			return nil, fmt.Errorf("validator not formatted")
		}

		validator, found := server.keeper.stakingKeeper.GetValidator(ctx, vals)
		if !found {
			return nil, fmt.Errorf("validator address doesnot exist")
		}

		_, err = server.keeper.stakingKeeper.Delegate(ctx, owner, amountToStake, validator.Status, validator, true)
		if err != nil {
			return nil, err
		}

	}

	return &types.MsgStakeToValidatorSetResponse{}, nil
}

func (server msgServer) UnStakeFromoValidatorSet(goCtx context.Context, msg *types.MsgUnStakeFromValidatorSet) (*types.MsgUnStakeFromValidatorSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// userA stakes {10osmo ValA ValB} -> [{ValA, ValB} -> {0.4, 0.6}] = 10Osmos = {4osmo, 6osmo}
	// userA unstake {3osmo ValA} -> [{valA, valB} -> {0.142, 0.857}] = 7osmo = {1osmo, 6osmo}

	// User can also do unstake all
	// User gives {amount, validators} to unstake from

	// get the existing validator set preference
	_, found := server.keeper.GetValidatorSetPreference(ctx, msg.Owner)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("user %d doesn't have validator set", msg.Owner))
	}

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	tokens := msg.Coins             // the total amount the user wants to unstake
	validatorSet := msg.Preferences // new validator set
	// Calculate new weights based on tokens provided

	total_weight := sdk.NewDec(0)
	// check if the provided validator set is correct
	for _, val := range validatorSet {
		// validation checks making sure the weights add up to 1 and also the validator given is correct
		valAddr, err := sdk.ValAddressFromBech32(val.ValOperAddress)
		if err != nil {
			return nil, fmt.Errorf("validator not formatted")
		}

		_, found := server.keeper.stakingKeeper.GetValidator(ctx, valAddr)
		if !found {
			return nil, fmt.Errorf("validator address kdoesnot exist")
		}

		amountToStake := val.Weight.Mul(tokenAmt).RoundInt()

		_, err := server.keeper.stakingKeeper.Undelegate(ctx, owner, valAddr)
		if err != nil {
			return nil, err
		}

		total_weight = total_weight.Add(val.Weight)
	}

	// loop through the amount and validator(should be in a list)
	// -> check validator they exist in the current set preferece
	// -> check the user has actually staked the amount to the validator
	// -> undelegate the amount from that specific validator
	// -> update the user val-set-preference set

	// check that the user has staked toekns in first place

	// Undelegate()

	return nil, nil
}
