package keeper

import (
	"encoding/binary"

	"capy/x/capy/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAnimalsCount get the total number of animals
func (k Keeper) GetAnimalsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AnimalsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAnimalsCount set the total number of animals
func (k Keeper) SetAnimalsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AnimalsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAnimals appends a animals in the store with a new id and update the count
func (k Keeper) AppendAnimals(
	ctx sdk.Context,
	animals types.Animals,
) uint64 {
	// Create the animals
	count := k.GetAnimalsCount(ctx)

	// Set the ID of the appended value
	animals.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnimalsKey))
	appendedValue := k.cdc.MustMarshal(&animals)
	store.Set(GetAnimalsIDBytes(animals.Id), appendedValue)

	// Update animals count
	k.SetAnimalsCount(ctx, count+1)

	return count
}

// SetAnimals set a specific animals in the store
func (k Keeper) SetAnimals(ctx sdk.Context, animals types.Animals) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnimalsKey))
	b := k.cdc.MustMarshal(&animals)
	store.Set(GetAnimalsIDBytes(animals.Id), b)
}

// GetAnimals returns a animals from its id
func (k Keeper) GetAnimals(ctx sdk.Context, id uint64) (val types.Animals, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnimalsKey))
	b := store.Get(GetAnimalsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAnimals removes a animals from the store
func (k Keeper) RemoveAnimals(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnimalsKey))
	store.Delete(GetAnimalsIDBytes(id))
}

// GetAllAnimals returns all animals
func (k Keeper) GetAllAnimals(ctx sdk.Context) (list []types.Animals) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnimalsKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Animals
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAnimalsIDBytes returns the byte representation of the ID
func GetAnimalsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAnimalsIDFromBytes returns ID in uint64 format from a byte array
func GetAnimalsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
