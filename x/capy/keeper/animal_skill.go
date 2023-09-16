package keeper

import (
	"encoding/binary"

	"capy/x/capy/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAnimalSkillCount get the total number of animalSkill
func (k Keeper) GetAnimalSkillCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AnimalSkillCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAnimalSkillCount set the total number of animalSkill
func (k Keeper) SetAnimalSkillCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AnimalSkillCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAnimalSkill appends a animalSkill in the store with a new id and update the count
func (k Keeper) AppendAnimalSkill(
	ctx sdk.Context,
	animalSkill types.AnimalSkill,
) uint64 {
	// Create the animalSkill
	count := k.GetAnimalSkillCount(ctx)

	// Set the ID of the appended value
	animalSkill.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnimalSkillKey))
	appendedValue := k.cdc.MustMarshal(&animalSkill)
	store.Set(GetAnimalSkillIDBytes(animalSkill.Id), appendedValue)

	// Update animalSkill count
	k.SetAnimalSkillCount(ctx, count+1)

	return count
}

// SetAnimalSkill set a specific animalSkill in the store
func (k Keeper) SetAnimalSkill(ctx sdk.Context, animalSkill types.AnimalSkill) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnimalSkillKey))
	b := k.cdc.MustMarshal(&animalSkill)
	store.Set(GetAnimalSkillIDBytes(animalSkill.Id), b)
}

// GetAnimalSkill returns a animalSkill from its id
func (k Keeper) GetAnimalSkill(ctx sdk.Context, id uint64) (val types.AnimalSkill, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnimalSkillKey))
	b := store.Get(GetAnimalSkillIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAnimalSkill removes a animalSkill from the store
func (k Keeper) RemoveAnimalSkill(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnimalSkillKey))
	store.Delete(GetAnimalSkillIDBytes(id))
}

// GetAllAnimalSkill returns all animalSkill
func (k Keeper) GetAllAnimalSkill(ctx sdk.Context) (list []types.AnimalSkill) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnimalSkillKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AnimalSkill
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAnimalSkillIDBytes returns the byte representation of the ID
func GetAnimalSkillIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAnimalSkillIDFromBytes returns ID in uint64 format from a byte array
func GetAnimalSkillIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
