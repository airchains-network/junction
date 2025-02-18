package keeper

import (
	"context"
	"encoding/binary"

	"github.com/airchains-network/junction/x/vrf/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetStudentsCount get the total number of students
func (k Keeper) GetStudentsCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.StudentsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetStudentsCount set the total number of students
func (k Keeper) SetStudentsCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.StudentsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendStudents appends a students in the store with a new id and update the count
func (k Keeper) AppendStudents(
	ctx context.Context,
	students types.Students,
) uint64 {
	// Create the students
	count := k.GetStudentsCount(ctx)

	// Set the ID of the appended value
	students.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StudentsKey))
	appendedValue := k.cdc.MustMarshal(&students)
	store.Set(GetStudentsIDBytes(students.Id), appendedValue)

	// Update students count
	k.SetStudentsCount(ctx, count+1)

	return count
}

// SetStudents set a specific students in the store
func (k Keeper) SetStudents(ctx context.Context, students types.Students) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StudentsKey))
	b := k.cdc.MustMarshal(&students)
	store.Set(GetStudentsIDBytes(students.Id), b)
}

// GetStudents returns a students from its id
func (k Keeper) GetStudents(ctx context.Context, id uint64) (val types.Students, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StudentsKey))
	b := store.Get(GetStudentsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStudents removes a students from the store
func (k Keeper) RemoveStudents(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StudentsKey))
	store.Delete(GetStudentsIDBytes(id))
}

// GetAllStudents returns all students
func (k Keeper) GetAllStudents(ctx context.Context) (list []types.Students) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StudentsKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Students
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetStudentsIDBytes returns the byte representation of the ID
func GetStudentsIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.StudentsKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
