package keeper

import (
	"context"
	"encoding/binary"

	"github.com/airchains-network/junction/x/zksequencer/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetCollegeCount get the total number of college
func (k Keeper) GetCollegeCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.CollegeCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCollegeCount set the total number of college
func (k Keeper) SetCollegeCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.CollegeCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendCollege appends a college in the store with a new id and update the count
func (k Keeper) AppendCollege(
	ctx context.Context,
	college types.College,
) uint64 {
	// Create the college
	count := k.GetCollegeCount(ctx)

	// Set the ID of the appended value
	college.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CollegeKey))
	appendedValue := k.cdc.MustMarshal(&college)
	store.Set(GetCollegeIDBytes(college.Id), appendedValue)

	// Update college count
	k.SetCollegeCount(ctx, count+1)

	return count
}

// SetCollege set a specific college in the store
func (k Keeper) SetCollege(ctx context.Context, college types.College) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CollegeKey))
	b := k.cdc.MustMarshal(&college)
	store.Set(GetCollegeIDBytes(college.Id), b)
}

// GetCollege returns a college from its id
func (k Keeper) GetCollege(ctx context.Context, id uint64) (val types.College, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CollegeKey))
	b := store.Get(GetCollegeIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCollege removes a college from the store
func (k Keeper) RemoveCollege(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CollegeKey))
	store.Delete(GetCollegeIDBytes(id))
}

// GetAllCollege returns all college
func (k Keeper) GetAllCollege(ctx context.Context) (list []types.College) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CollegeKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.College
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCollegeIDBytes returns the byte representation of the ID
func GetCollegeIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.CollegeKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
