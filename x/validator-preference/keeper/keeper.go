package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/osmosis-labs/osmosis/v12/osmoutils"
	"github.com/osmosis-labs/osmosis/v12/x/validator-preference/types"
)

type Keeper struct {
	storeKey      sdk.StoreKey
	paramSpace    paramtypes.Subspace
	cdc           codec.BinaryCodec
	stakingKeeper types.StakingInterface
}

func NewKeeper(storeKey sdk.StoreKey, paramSpace paramtypes.Subspace, stakingKeeper types.StakingInterface) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:      storeKey,
		paramSpace:    paramSpace,
		stakingKeeper: stakingKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetValidatorSetPreferences(ctx sdk.Context, validators types.ValidatorSetPreferences) error {
	store := ctx.KVStore(k.storeKey)
	osmoutils.MustSet(store, []byte(validators.Owner), &validators)
	return nil
}

func (k Keeper) GetValidatorSetPreference(ctx sdk.Context, owner string) types.ValidatorSetPreferences {
	validatorSet := types.ValidatorSetPreferences{}

	store := ctx.KVStore(k.storeKey)
	osmoutils.MustGet(store, []byte(owner), &validatorSet)

	return validatorSet
}
