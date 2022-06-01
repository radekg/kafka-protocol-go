package encoder

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/google/uuid"
	"github.com/radekg/kafka-protocol-go/errors"
)

type prepEncoder struct {
	length int
}

func NewPrepEncoder() PEDef {
	return &prepEncoder{}
}

// primitives

func (pe *prepEncoder) PutInt8(in int8) {
	pe.length++
}

func (pe *prepEncoder) PutInt16(in int16) {
	pe.length += 2
}

func (pe *prepEncoder) PutInt32(in int32) {
	pe.length += 4
}

func (pe *prepEncoder) PutInt64(in int64) {
	pe.length += 8
}

func (pe *prepEncoder) PutVarint(in int64) {
	var buf [binary.MaxVarintLen64]byte
	pe.length += binary.PutUvarint(buf[:], uint64(in))
}

func (pe *prepEncoder) PutArrayLength(in int) error {
	if in > math.MaxInt32 {
		return errors.PacketEncodingError{fmt.Sprintf("array too long (%d)", in)}
	}
	pe.length += 4
	return nil
}

func (pe *prepEncoder) PutBool(in bool) {
	pe.length++
}

// arrays
func (pe *prepEncoder) PutBytes(in []byte) error {
	pe.length += 4
	if in == nil {
		return nil
	}
	return pe.PutRawBytes(in)
}

func (pe *prepEncoder) PutUUID(in uuid.UUID) error {
	pe.length += 16
	return nil
}

func (pe *prepEncoder) PutVarintBytes(in []byte) error {
	if in == nil {
		pe.PutVarint(-1)
		return nil
	}
	pe.PutVarint(int64(len(in)))
	return pe.PutRawBytes(in)
}

func (pe *prepEncoder) PutRawBytes(in []byte) error {
	if len(in) > math.MaxInt32 {
		return errors.PacketEncodingError{fmt.Sprintf("byteslice too long (%d)", len(in))}
	}
	pe.length += len(in)
	return nil
}

func (pe *prepEncoder) PutCompactBytes(in []byte) error {
	pe.PutVarint(int64(len(in) + 1))
	if len(in) > math.MaxInt16 {
		return errors.PacketEncodingError{fmt.Sprintf("compact bytes too long (%d)", len(in))}
	}
	pe.length += len(in)
	return nil
}

func (pe *prepEncoder) PutNullableString(in *string) error {
	if in == nil {
		pe.length += 2
		return nil
	}
	return pe.PutString(*in)
}

func (pe *prepEncoder) PutString(in string) error {
	pe.length += 2
	if len(in) > math.MaxInt16 {
		return errors.PacketEncodingError{fmt.Sprintf("string too long (%d)", len(in))}
	}
	pe.length += len(in)
	return nil
}

func (pe *prepEncoder) PutStringArray(in []string) error {
	err := pe.PutArrayLength(len(in))
	if err != nil {
		return err
	}

	for _, str := range in {
		if err := pe.PutString(str); err != nil {
			return err
		}
	}

	return nil
}

func (pe *prepEncoder) PutInt32Array(in []int32) error {
	if in == nil {
		return pe.PutArrayLength(-1)
	}
	err := pe.PutArrayLength(len(in))
	if err != nil {
		return err
	}
	pe.length += 4 * len(in)
	return nil
}

func (pe *prepEncoder) PutInt64Array(in []int64) error {
	if in == nil {
		return pe.PutArrayLength(-1)
	}
	err := pe.PutArrayLength(len(in))
	if err != nil {
		return err
	}
	pe.length += 8 * len(in)
	return nil
}

func (pe *prepEncoder) Offset() int {
	return pe.length
}

func (pe *prepEncoder) PutCompactString(in string) error {
	pe.PutVarint(int64(len(in) + 1))
	if len(in) > math.MaxInt16 {
		return errors.PacketEncodingError{fmt.Sprintf("string too long (%d)", len(in))}
	}
	pe.length += len(in)
	return nil
}

func (pe *prepEncoder) PutCompactNullableString(in *string) error {
	if in == nil {
		// A null string is represented with a length of 0.
		pe.length += 1 // pe.putVarint(0) is always 1
		return nil
	}
	return pe.PutCompactString(*in)
}

func (pe *prepEncoder) PutCompactArrayLength(in int) error {
	switch {
	case in > math.MaxInt16:
		return errors.PacketEncodingError{fmt.Sprintf("comact array too long (%d)", in)}
	case in == -1:
		return errors.PacketEncodingError{fmt.Sprintf("comact array is null (%d)", in)}
	case in < -1:
		return errors.PacketEncodingError{fmt.Sprintf("comact array invalid length (%d)", in)}
	}
	pe.PutVarint(int64(in + 1))
	return nil
}

func (pe *prepEncoder) PutCompactNullableArrayLength(in int) error {
	switch {
	case in > math.MaxInt16:
		return errors.PacketEncodingError{fmt.Sprintf("comact array too long (%d)", in)}
	case in < -1:
		return errors.PacketEncodingError{fmt.Sprintf("comact array invalid length (%d)", in)}
	}
	pe.PutVarint(int64(in + 1))
	return nil
}

func (pe *prepEncoder) Bytes() []byte {
	return nil
}

func (pe *prepEncoder) Length() int {
	return pe.length
}
