package decoder

import "github.com/google/uuid"

// DDef is the interface providing helpers for reading with Kafka's encoding rules.
// Types implementing Decoder only need to worry about calling methods like GetString,
// not about how a string is represented in Kafka.
type DDef interface {
	// Primitives
	GetInt8() (int8, error)
	GetInt16() (int16, error)
	GetInt32() (int32, error)
	GetInt64() (int64, error)
	GetVarint() (int64, error)
	GetArrayLength() (int, error)
	GetBool() (bool, error)

	GetBytes() ([]byte, error)
	GetUUID() (uuid.UUID, error)
	GetString() (string, error)
	GetNullableString() (*string, error)
	GetInt32Array() ([]int32, error)
	GetInt64Array() ([]int64, error)
	GetStringArray() ([]string, error)

	GetVarintBytes() ([]byte, error)

	GetCompactBytes() ([]byte, error)
	GetCompactString() (string, error)
	GetCompactNullableString() (*string, error)
	GetCompactArrayLength() (int, error)
	GetCompactNullableArrayLength() (int, error)

	Offset() int

	// Subsets
	Remaining() int
}
