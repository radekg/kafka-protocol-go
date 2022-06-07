package decoder

import "github.com/google/uuid"

// DDef is the interface providing helpers for reading with Kafka's encoding rules.
// Types implementing Decoder only need to worry about calling methods like GetString,
// not about how a string is represented in Kafka.
type DDef interface {

	// Array lengths
	GetArrayLength() (int, error)
	GetArrayCompactLength() (int, error)
	GetArrayCompactNullableLength() (int, error)

	// Booleans
	GetBool() (bool, error)

	// Bytes
	GetBytes() ([]byte, error)
	GetBytesCompact() ([]byte, error)
	GetBytesNullable() ([]byte, error)

	// Floats
	GetFloat64() (float64, error)

	// Integers and their arrays
	GetInt8() (int8, error)
	GetInt16() (int16, error)
	GetInt32() (int32, error)
	GetInt32Array() ([]int32, error)
	GetInt32CompactArray() ([]int32, error)
	GetInt64() (int64, error)
	GetInt64Array() ([]int64, error)
	GetInt64CompactArray() ([]int64, error)
	GetVarint() (int64, error)

	// Strings and their arrays
	GetString() (string, error)
	GetStringArray() ([]string, error)
	GetStringCompact() (string, error)
	GetStringCompactArray() ([]string, error)
	GetStringCompactNullable() (*string, error)
	GetStringNullable() (*string, error)

	// UUIDs
	GetUUID() (uuid.UUID, error)

	// Other utility
	CopyNew() DDef
	Offset() int
	Remaining() int
}
