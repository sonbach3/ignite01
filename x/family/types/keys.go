package types

const (
	// ModuleName defines the module name
	ModuleName = "family"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_family"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	MemberKey      = "Member/value/"
	MemberCountKey = "Member/count/"
)
