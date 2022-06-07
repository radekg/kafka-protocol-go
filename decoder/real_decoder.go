package decoder

import (
	"encoding/binary"
	"math"

	"github.com/google/uuid"
	"github.com/radekg/kafka-protocol-go/errors"
)

var errInvalidArrayLength = errors.PacketDecodingError{"invalid array length"}
var errInvalidStringLength = errors.PacketDecodingError{"invalid string length"}
var errVarintOverflow = errors.PacketDecodingError{"varint overflow"}
var errInvalidBool = errors.PacketDecodingError{"invalid bool"}
var errInvalidByteSliceLength = errors.PacketDecodingError{"invalid byte slice length"}
var errInvalidCompactLength = errors.PacketDecodingError{"invalid compact length"}
var errInvalidCompactNullableLength = errors.PacketDecodingError{"invalid compact nullable length"}

func ErrInvalidArrayLength() error {
	return errInvalidArrayLength
}

func ErrInvalidStringLength() error {
	return errInvalidStringLength
}

func ErrVarintOverflow() error {
	return errVarintOverflow
}

func ErrInvalidBool() error {
	return errInvalidBool
}

func ErrInvalidByteSliceLength() error {
	return errInvalidByteSliceLength
}

func ErrInvalidCompactLength() error {
	return errInvalidCompactLength
}

func ErrInvalidCompactNullableLength() error {
	return errInvalidCompactNullableLength
}

type realDecoder struct {
	raw []byte
	off int
}

func NewRealDecoder(raw []byte) DDef {
	return &realDecoder{raw: raw}
}

// primitives

func (rd *realDecoder) GetInt8() (int8, error) {
	if rd.Remaining() < 1 {
		rd.off = len(rd.raw)
		return -1, errors.ErrInsufficientData
	}
	tmp := int8(rd.raw[rd.off])
	rd.off++
	return tmp, nil
}

func (rd *realDecoder) GetInt16() (int16, error) {
	if rd.Remaining() < 2 {
		rd.off = len(rd.raw)
		return -1, errors.ErrInsufficientData
	}
	tmp := int16(binary.BigEndian.Uint16(rd.raw[rd.off:]))
	rd.off += 2
	return tmp, nil
}

func (rd *realDecoder) GetInt32() (int32, error) {
	if rd.Remaining() < 4 {
		rd.off = len(rd.raw)
		return -1, errors.ErrInsufficientData
	}
	tmp := int32(binary.BigEndian.Uint32(rd.raw[rd.off:]))
	rd.off += 4
	return tmp, nil
}

func (rd *realDecoder) GetInt64() (int64, error) {
	if rd.Remaining() < 8 {
		rd.off = len(rd.raw)
		return -1, errors.ErrInsufficientData
	}
	tmp := int64(binary.BigEndian.Uint64(rd.raw[rd.off:]))
	rd.off += 8
	return tmp, nil
}

func (rd *realDecoder) GetFloat64() (float64, error) {
	if rd.Remaining() < 8 {
		rd.off = len(rd.raw)
		return -1, errors.ErrInsufficientData
	}
	tmp := binary.BigEndian.Uint64(rd.raw[rd.off:])
	rd.off += 8
	return math.Float64frombits(tmp), nil
}

func (rd *realDecoder) GetVarint() (int64, error) {
	tmp, n := binary.Uvarint(rd.raw[rd.off:])
	if n == 0 {
		rd.off = len(rd.raw)
		return -1, errors.ErrInsufficientData
	}
	if n < 0 {
		rd.off -= n
		return -1, errVarintOverflow
	}
	rd.off += n
	return int64(tmp), nil
}

func (rd *realDecoder) GetArrayLength() (int, error) {
	if rd.Remaining() < 4 {
		rd.off = len(rd.raw)
		return -1, errors.ErrInsufficientData
	}
	tmp := int(int32(binary.BigEndian.Uint32(rd.raw[rd.off:])))
	rd.off += 4
	if tmp > rd.Remaining() {
		rd.off = len(rd.raw)
		return -1, errors.ErrInsufficientData
	} else if tmp > 2*math.MaxUint16 {
		return -1, errInvalidArrayLength
	}
	return tmp, nil
}

func (rd *realDecoder) GetBool() (bool, error) {
	b, err := rd.GetInt8()
	if err != nil || b == 0 {
		return false, err
	}
	if b != 1 {
		return false, errInvalidBool
	}
	return true, nil
}

// collections

func (rd *realDecoder) GetBytes() ([]byte, error) {
	tmp, err := rd.GetInt32()
	if err != nil {
		return nil, err
	}
	if tmp == -1 {
		return nil, nil
	}

	return rd.GetRawBytes(int(tmp))
}

func (rd *realDecoder) GetBytesNullable() ([]byte, error) {
	tmp, err := rd.GetVarint()
	if err != nil {
		return nil, err
	}
	if tmp == -1 {
		return nil, nil
	}
	return rd.GetRawBytes(int(tmp))
}

func (rd *realDecoder) GetRawBytes(length int) ([]byte, error) {
	if length < 0 {
		return nil, errInvalidByteSliceLength
	} else if length > rd.Remaining() {
		rd.off = len(rd.raw)
		return nil, errors.ErrInsufficientData
	}

	start := rd.off
	rd.off += length
	return rd.raw[start:rd.off], nil
}

func (rd *realDecoder) GetUUID() (uuid.UUID, error) {
	r, err := rd.GetRawBytes(16)
	if err != nil {
		var u uuid.UUID
		return u, err
	}
	return uuid.FromBytes(r)
}

func (rd *realDecoder) GetBytesCompact() ([]byte, error) {

	n, err := rd.GetCompactLength()
	if err != nil || n == 0 {
		return []byte{}, err
	}
	tmp := rd.raw[rd.off : rd.off+n]
	rd.off += n
	return tmp, nil
}

func (rd *realDecoder) GetStringLength() (int, error) {
	length, err := rd.GetInt16()
	if err != nil {
		return 0, err
	}

	n := int(length)

	switch {
	case n < -1:
		return 0, errInvalidStringLength
	case n > rd.Remaining():
		rd.off = len(rd.raw)
		return 0, errors.ErrInsufficientData
	}

	return n, nil
}

func (rd *realDecoder) GetString() (string, error) {
	n, err := rd.GetStringLength()
	if err != nil || n == -1 {
		return "", err
	}

	tmpStr := string(rd.raw[rd.off : rd.off+n])
	rd.off += n
	return tmpStr, nil
}

func (rd *realDecoder) GetStringNullable() (*string, error) {
	n, err := rd.GetStringLength()
	if err != nil || n == -1 {
		return nil, err
	}

	tmpStr := string(rd.raw[rd.off : rd.off+n])
	rd.off += n
	return &tmpStr, err
}

func (rd *realDecoder) GetInt32Array() ([]int32, error) {
	n, err := rd.GetArrayLength()

	if err != nil {
		return nil, err
	}
	if n == -1 {
		return nil, nil
	}
	if n == 0 {
		return []int32{}, nil
	}
	if n < -1 {
		return nil, errInvalidArrayLength
	}
	if rd.Remaining() < 4*n {
		rd.off = len(rd.raw)
		return nil, errors.ErrInsufficientData
	}
	ret := make([]int32, n)
	for i := range ret {
		ret[i] = int32(binary.BigEndian.Uint32(rd.raw[rd.off:]))
		rd.off += 4
	}
	return ret, nil
}

func (rd *realDecoder) GetInt32CompactArray() ([]int32, error) {
	n, err := rd.GetArrayLength()

	if err != nil {
		return nil, err
	}
	if n == -1 {
		return nil, nil
	}
	if n == 0 {
		return []int32{}, nil
	}
	if n < -1 {
		return nil, errInvalidArrayLength
	}
	if rd.Remaining() < 4*n {
		rd.off = len(rd.raw)
		return nil, errors.ErrInsufficientData
	}
	ret := make([]int32, n)
	for i := range ret {
		ret[i] = int32(binary.BigEndian.Uint32(rd.raw[rd.off:]))
		rd.off += 4
	}
	return ret, nil
}

func (rd *realDecoder) GetInt64Array() ([]int64, error) {
	n, err := rd.GetArrayLength()

	if err != nil {
		return nil, err
	}
	if n == -1 {
		return nil, nil
	}
	if n == 0 {
		return []int64{}, nil
	}
	if n < -1 {
		return nil, errInvalidArrayLength
	}
	if rd.Remaining() < 8*n {
		rd.off = len(rd.raw)
		return nil, errors.ErrInsufficientData
	}
	ret := make([]int64, n)
	for i := range ret {
		ret[i] = int64(binary.BigEndian.Uint64(rd.raw[rd.off:]))
		rd.off += 8
	}
	return ret, nil
}

func (rd *realDecoder) GetInt64CompactArray() ([]int64, error) {
	n, err := rd.GetArrayCompactLength()

	if err != nil {
		return nil, err
	}
	if n == -1 {
		return nil, nil
	}
	if n == 0 {
		return []int64{}, nil
	}
	if n < -1 {
		return nil, errInvalidArrayLength
	}
	if rd.Remaining() < 8*n {
		rd.off = len(rd.raw)
		return nil, errors.ErrInsufficientData
	}
	ret := make([]int64, n)
	for i := range ret {
		ret[i] = int64(binary.BigEndian.Uint64(rd.raw[rd.off:]))
		rd.off += 8
	}
	return ret, nil
}

func (rd *realDecoder) GetStringArray() ([]string, error) {
	n, err := rd.GetArrayLength()

	if err != nil {
		return nil, err
	}

	if n == -1 {
		return nil, nil
	}

	if n < -1 {
		return nil, errInvalidArrayLength
	}

	ret := make([]string, n)
	for i := range ret {
		str, err := rd.GetString()
		if err != nil {
			return nil, err
		}

		ret[i] = str
	}
	return ret, nil
}

func (rd *realDecoder) GetStringCompactArray() ([]string, error) {
	n, err := rd.GetArrayCompactLength()

	if err != nil {
		return nil, err
	}

	if n == -1 {
		return nil, nil
	}

	if n < -1 {
		return nil, errInvalidArrayLength
	}

	ret := make([]string, n)
	for i := range ret {
		str, err := rd.GetStringCompact()
		if err != nil {
			return nil, err
		}

		ret[i] = str
	}
	return ret, nil
}

func (rd *realDecoder) GetCompactLength() (int, error) {
	length, err := rd.GetVarint()
	if err != nil {
		return 0, err
	}
	n := int(length - 1)

	switch {
	case n < 0:
		return 0, errInvalidCompactLength
	case n > rd.Remaining():
		rd.off = len(rd.raw)
		return 0, errors.ErrInsufficientData
	}
	return n, nil
}

func (rd *realDecoder) GetCompactNullableLength() (int, error) {
	length, err := rd.GetVarint()
	if err != nil {
		return 0, err
	}
	n := int(length - 1)

	switch {
	case n < -1:
		return 0, errInvalidCompactNullableLength
	case n > rd.Remaining():
		rd.off = len(rd.raw)
		return 0, errors.ErrInsufficientData
	}
	return n, nil
}

func (rd *realDecoder) GetStringCompact() (string, error) {

	n, err := rd.GetCompactLength()
	if err != nil || n == 0 {
		return "", err
	}
	tmpStr := string(rd.raw[rd.off : rd.off+n])
	rd.off += n
	return tmpStr, nil
}

func (rd *realDecoder) GetStringCompactNullable() (*string, error) {

	n, err := rd.GetCompactNullableLength()
	if err != nil || n < 0 {
		return nil, err
	}
	tmpStr := string(rd.raw[rd.off : rd.off+n])
	rd.off += n
	return &tmpStr, nil
}

func (rd *realDecoder) GetArrayCompactLength() (int, error) {
	return rd.GetCompactLength()
}

func (rd *realDecoder) GetArrayCompactNullableLength() (int, error) {
	return rd.GetCompactNullableLength()
}

// Other utilities:

func (rd *realDecoder) CopyNew() DDef {
	return &realDecoder{raw: rd.raw, off: 0}
}

func (rd *realDecoder) Offset() int {
	return rd.off
}

func (rd *realDecoder) Remaining() int {
	return len(rd.raw) - rd.off
}
