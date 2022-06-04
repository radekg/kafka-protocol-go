package encoder

import "github.com/google/uuid"

// EDef is the interface providing helpers for writing with Kafka's encoding rules.
// Types implementing Encoder only need to worry about calling methods like PutString,
// not about how a string is represented in Kafka.
type EDef interface {

	// Array lengths
	PutArrayLength(in int) error
	PutArrayCompactLength(in int) error
	PutArrayCompactNullableLength(in int) error

	// Booleans
	PutBool(in bool)

	// Bytes:
	PutBytes(in []byte) error
	PutBytesCompact(in []byte) error
	PutBytesNullable(in []byte) error

	// Floats
	PutFloat64(in float64)

	// Integers and their arrays
	PutInt8(in int8)
	PutInt16(in int16)
	PutInt32(in int32)
	PutInt32Array(in []int32) error
	PutInt32CompactArray(in []int32) error
	PutInt64(in int64)
	PutInt64Array(in []int64) error
	PutInt64CompactArray(in []int64) error
	PutVarint(in int64)

	// Strings and their arrays
	PutString(in string) error
	PutStringArray(in []string) error
	PutStringCompact(in string) error
	PutStringCompactArray(in []string) error
	PutStringCompactNullable(in *string) error
	PutStringNullable(in *string) error

	// UUIDs
	PutUUID(in uuid.UUID) error

	// Other utilities:
	Bytes() []byte
	// Provide the current offset to record the batch size metric
	Offset() int
}

type PEDef interface {
	EDef
	Length() int
}
