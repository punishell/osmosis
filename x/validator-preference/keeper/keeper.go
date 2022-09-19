package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/osmosis-labs/osmosis/v12/x/validator-preference/types"
	"google.golang.org/protobuf/proto"
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

func (k Keeper) SetValidatorSetPreferences(ctx sdk.Context, validators types.MsgStakeToValidatorSet) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := proto.Marshal(&validators.Preferences)
	if err != nil {
		return err
	}
	store.Set([]byte(validators.Owner), bz)
	return nil
}

func (k Keeper) GetValidatorSetPreference(ctx sdk.Context, owner string) ([]types.ValidatorPreference, bool) {
	validatorSet := []types.ValidatorPreference{}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(owner))
	if bz == nil {
		return validatorSet, false
	}
	err := proto.Unmarshal(bz, &validatorSet)
	if err != nil {
		return nil, false
	}

	return validatorSet, true
}
