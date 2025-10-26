package permissions

const (
	MASK_TYPE             uint32 = 0b11110111001000000000000000000000
	MASK_MODE             uint32 = 0b00001000110100000000000000000000
	MASK_UNKNOWN          uint32 = 0b00000000000011111111111000000000
	MASK_PERM             uint32 = 0b00000000000000000000000111111111
	MASK_PERM_OWNER       uint32 = 0b00000000000000000000000111000000
	MASK_PERM_GROUP       uint32 = 0b00000000000000000000000000111000
	MASK_PERM_WORLD       uint32 = 0b00000000000000000000000000000111
	MASK_TYPE_SHIFT       int    = 21
	MASK_MODE_SHIFT       int    = 20
	MASK_UNKNOWN_SHIFT    int    = 9
	MASK_PERM_SHIFT       int    = 0
	MASK_PERM_OWNER_SHIFT int    = 6
	MASK_PERM_GROUP_SHIFT int    = 3
	MASK_PERM_WORLD_SHIFT int    = 0
	MODE_STICKY           uint32 = 0b00000001
	MODE_GID              uint32 = 0b00000100
	MODE_UID              uint32 = 0b00001000
	MODE_LINK             uint32 = 0b10000000
	TYPE_CHAR             uint32 = 0b00000000001
	TYPE_SOCKET           uint32 = 0b00000001000
	TYPE_FIFO             uint32 = 0b00000010000
	TYPE_BLOCK_DEVICE     uint32 = 0b00000100000
	TYPE_CHAR_DEVICE      uint32 = 0b00000100001
	TYPE_DIR              uint32 = 0b10000000000
	type_unknown_8        uint32 = 0b00010000000 // unknown purpose
	type_unknown_9        uint32 = 0b00100000000 // unknown purpose
	type_unknown_10       uint32 = 0b01000000000 // unknown purpose
	PERM_READ             uint32 = 0b100
	PERM_WRITE            uint32 = 0b010
	PERM_EXEC             uint32 = 0b001
)
