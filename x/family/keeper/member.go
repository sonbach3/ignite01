package keeper

import (
	"encoding/binary"

	"capy/x/family/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetMemberCount get the total number of member
func (k Keeper) GetMemberCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MemberCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMemberCount set the total number of member
func (k Keeper) SetMemberCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MemberCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMember appends a member in the store with a new id and update the count
func (k Keeper) AppendMember(
	ctx sdk.Context,
	member types.Member,
) uint64 {
	// Create the member
	count := k.GetMemberCount(ctx)

	// Set the ID of the appended value
	member.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MemberKey))
	appendedValue := k.cdc.MustMarshal(&member)
	store.Set(GetMemberIDBytes(member.Id), appendedValue)

	// Update member count
	k.SetMemberCount(ctx, count+1)

	return count
}

// SetMember set a specific member in the store
func (k Keeper) SetMember(ctx sdk.Context, member types.Member) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MemberKey))
	b := k.cdc.MustMarshal(&member)
	store.Set(GetMemberIDBytes(member.Id), b)
}

// GetMember returns a member from its id
func (k Keeper) GetMember(ctx sdk.Context, id uint64) (val types.Member, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MemberKey))
	b := store.Get(GetMemberIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMember removes a member from the store
func (k Keeper) RemoveMember(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MemberKey))
	store.Delete(GetMemberIDBytes(id))
}

// GetAllMember returns all member
func (k Keeper) GetAllMember(ctx sdk.Context) (list []types.Member) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MemberKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Member
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMemberIDBytes returns the byte representation of the ID
func GetMemberIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMemberIDFromBytes returns ID in uint64 format from a byte array
func GetMemberIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
