package keeper

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/log"

	"github.com/archway-network/archway/internal/collcompat"
	"github.com/archway-network/archway/x/cwica/types"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		Codec               codec.BinaryCodec
		storeKey            storetypes.StoreKey
		channelKeeper       types.ChannelKeeper
		connectionKeeper    types.ConnectionKeeper
		errorsKeeper        types.ErrorsKeeper
		icaControllerKeeper types.ICAControllerKeeper
		sudoKeeper          types.WasmKeeper
		authority           string
		logger              log.Logger
		Schema              collections.Schema

		// Params key: ParamsKeyPrefix | value: Params
		Params collections.Item[types.Params]
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	channelKeeper types.ChannelKeeper,
	connectionKeeper types.ConnectionKeeper,
	errorsKeeper types.ErrorsKeeper,
	icaControllerKeeper types.ICAControllerKeeper,
	sudoKeeper types.WasmKeeper,
	authority string,
	logger log.Logger,
) Keeper {
	sb := collections.NewSchemaBuilder(collcompat.NewKVStoreService(storeKey))

	k := Keeper{
		Codec:               cdc,
		storeKey:            storeKey,
		channelKeeper:       channelKeeper,
		connectionKeeper:    connectionKeeper,
		errorsKeeper:        errorsKeeper,
		icaControllerKeeper: icaControllerKeeper,
		sudoKeeper:          sudoKeeper,
		authority:           authority,
		logger:              logger.With("module", "x/"+types.ModuleName),
		Params: collections.NewItem(
			sb,
			types.ParamsKeyPrefix,
			"params",
			collcompat.ProtoValue[types.Params](cdc),
		),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema
	return k
}

// GetAuthority returns the authority of the keeper. Should be the governance module address.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return k.logger
}

// SetWasmKeeper sets the given wasm keeper.
// NOTE: Only for testing purposes
func (k *Keeper) SetWasmKeeper(wk types.WasmKeeper) {
	k.sudoKeeper = wk
}

// SetICAControllerKeeper sets the given ica controller keeper.
// NOTE: Only for testing purposes
func (k *Keeper) SetICAControllerKeeper(icak types.ICAControllerKeeper) {
	k.icaControllerKeeper = icak
}

// SetChannelKeeper sets the given channel keeper.
// NOTE: Only for testing purposes
func (k *Keeper) SetChannelKeeper(ck types.ChannelKeeper) {
	k.channelKeeper = ck
}

// SetConnectionKeeper sets the given connection keeper.
// NOTE: Only for testing purposes
func (k *Keeper) SetConnectionKeeper(ck types.ConnectionKeeper) {
	k.connectionKeeper = ck
}

// SetErrorsKeeper sets the given errors keeper.
// NOTE: Only for testing purposes
func (k *Keeper) SetErrorsKeeper(ek types.ErrorsKeeper) {
	k.errorsKeeper = ek
}
