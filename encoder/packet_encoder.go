package encoder

import "github.com/google/uuid"

// EDef is the interface providing helpers for writing with Kafka's encoding rules.
// Types implementing Encoder only need to worry about calling methods like PutString,
// not about how a string is represented in Kafka.
type EDef interface {
	// Primitives
	PutInt8(in int8)
	PutInt16(in int16)
	PutInt32(in int32)
	PutInt64(in int64)
	PutVarint(in int64)
	PutArrayLength(in int) error
	PutBool(in bool)

	PutBytes(in []byte) error
	PutUUID(in uuid.UUID) error
	PutString(in string) error
	PutNullableString(in *string) error
	PutStringArray(in []string) error
	PutInt32Array(in []int32) error
	PutInt64Array(in []int64) error

	PutVarintBytes(in []byte) error

	PutCompactBytes(in []byte) error
	PutCompactString(in string) error
	PutCompactNullableString(in *string) error
	PutCompactArrayLength(in int) error
	PutCompactNullableArrayLength(in int) error

	Bytes() []byte
	// Provide the current offset to record the batch size metric
	Offset() int
}

type PEDef interface {
	EDef
	Length() int
}
