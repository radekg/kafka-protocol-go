package encoder

import (
	"encoding/binary"

	"github.com/google/uuid"
)

type realEncoder struct {
	raw []byte
	off int
}

func NewRealEncoder(raw []byte) EDef {
	return &realEncoder{raw: raw}
}

func (re *realEncoder) PutArrayLength(in int) error {
	re.PutInt32(int32(in))
	return nil
}

func (re *realEncoder) PutBool(in bool) {
	if in {
		re.PutInt8(1)
		return
	}
	re.PutInt8(0)
}

// -- Int8

func (re *realEncoder) PutInt8(in int8) {
	re.raw[re.off] = byte(in)
	re.off++
}

func (re *realEncoder) PutInt16(in int16) {
	binary.BigEndian.PutUint16(re.raw[re.off:], uint16(in))
	re.off += 2
}

// -- Int32

func (re *realEncoder) PutInt32(in int32) {
	binary.BigEndian.PutUint32(re.raw[re.off:], uint32(in))
	re.off += 4
}

func (re *realEncoder) PutInt32CompactArray(in []int32) error {
	if in == nil {
		return re.PutArrayLength(-1)
	}
	err := re.PutArrayCompactLength(len(in))
	if err != nil {
		return err
	}
	for _, val := range in {
		re.PutInt32(val)
	}
	return nil
}

func (re *realEncoder) PutInt32Array(in []int32) error {
	if in == nil {
		return re.PutArrayLength(-1)
	}
	err := re.PutArrayLength(len(in))
	if err != nil {
		return err
	}
	for _, val := range in {
		re.PutInt32(val)
	}
	return nil
}

// -- Int64

func (re *realEncoder) PutInt64(in int64) {
	binary.BigEndian.PutUint64(re.raw[re.off:], uint64(in))
	re.off += 8
}

func (re *realEncoder) PutVarint(in int64) {
	re.off += binary.PutUvarint(re.raw[re.off:], uint64(in))
}

func (re *realEncoder) PutInt64Array(in []int64) error {
	if in == nil {
		return re.PutArrayLength(-1)
	}
	err := re.PutArrayLength(len(in))
	if err != nil {
		return err
	}
	for _, val := range in {
		re.PutInt64(val)
	}
	return nil
}

func (re *realEncoder) PutInt64CompactArray(in []int64) error {
	if in == nil {
		return re.PutArrayLength(-1)
	}
	err := re.PutArrayCompactLength(len(in))
	if err != nil {
		return err
	}
	for _, val := range in {
		re.PutInt64(val)
	}
	return nil
}

// collection

func (re *realEncoder) PutRawBytes(in []byte) error {
	copy(re.raw[re.off:], in)
	re.off += len(in)
	return nil
}

func (re *realEncoder) PutBytes(in []byte) error {
	if in == nil {
		re.PutInt32(-1)
		return nil
	}
	re.PutInt32(int32(len(in)))
	return re.PutRawBytes(in)
}

func (pe *realEncoder) PutUUID(in uuid.UUID) error {
	return pe.PutRawBytes(in[:])
}

func (re *realEncoder) PutBytesCompact(in []byte) error {
	re.PutVarint(int64(len(in) + 1))
	copy(re.raw[re.off:], in)
	re.off += len(in)
	return nil
}

func (re *realEncoder) PutBytesNullable(in []byte) error {
	if in == nil {
		re.PutVarint(-1)
		return nil
	}
	re.PutVarint(int64(len(in)))
	return re.PutRawBytes(in)
}

func (re *realEncoder) PutString(in string) error {
	re.PutInt16(int16(len(in)))
	copy(re.raw[re.off:], in)
	re.off += len(in)
	return nil
}

func (re *realEncoder) PutStringNullable(in *string) error {
	if in == nil {
		re.PutInt16(-1)
		return nil
	}
	return re.PutString(*in)
}

func (re *realEncoder) PutStringArray(in []string) error {
	err := re.PutArrayLength(len(in))
	if err != nil {
		return err
	}

	for _, val := range in {
		if err := re.PutString(val); err != nil {
			return err
		}
	}

	return nil
}

func (re *realEncoder) PutStringCompactArray(in []string) error {
	err := re.PutArrayCompactLength(len(in))
	if err != nil {
		return err
	}

	for _, val := range in {
		if err := re.PutStringCompact(val); err != nil {
			return err
		}
	}

	return nil
}

func (re *realEncoder) Offset() int {
	return re.off
}

func (re *realEncoder) PutStringCompact(in string) error {
	re.PutVarint(int64(len(in) + 1))
	copy(re.raw[re.off:], in)
	re.off += len(in)
	return nil
}

func (re *realEncoder) PutStringCompactNullable(in *string) error {
	if in == nil {
		re.PutVarint(0)
		return nil
	}
	return re.PutStringCompact(*in)
}

func (pe *realEncoder) PutArrayCompactLength(in int) error {
	pe.PutVarint(int64(in + 1))
	return nil
}

func (pe *realEncoder) PutArrayCompactNullableLength(in int) error {
	pe.PutVarint(int64(in + 1))
	return nil
}

func (pe *realEncoder) Bytes() []byte {
	return pe.raw
}
