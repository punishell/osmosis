package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
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

func (k Keeper) SetValidatorSetPreferences(ctx sdk.Context, validators types.MsgValidatorSetPreference) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixValidatorSet)
	bz, err := proto.Marshal(&validators)
	if err != nil {
		panic("Error during marshalling")
	}
	prefixStore.Set(types.KeyPrefixValidatorSet, bz)
}

func (k Keeper) GetValidatorSetPreference(ctx sdk.Context) []types.ValidatorPreference {
	validatorSet := []types.ValidatorPreference{}

	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixValidatorSet)
	bz := prefixStore.Get(types.KeyPrefixValidatorSet)
	err := proto.Unmarshal(bz, &validatorSet)
	if err != nil {
		panic("Error during unmarshalling")
	}

	return validatorSet
}
