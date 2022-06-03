package schema

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	proterrors "github.com/radekg/kafka-protocol-go/errors"
)

// -- Array

type Array struct {
	Name string
	Ty   Schema
}

func (f *Array) Decode(pd decoder.DDef) (interface{}, error) {
	n, err := pd.GetArrayLength()
	if err != nil {
		return nil, err
	}
	return decodeArrayElements(n, f.Ty.Decode, pd)
}

func (f *Array) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.([]interface{})
	if !ok {
		return proterrors.SchemaEncodingError{fmt.Sprintf("value %T not a []interface{}", value)}
	}
	err := pe.PutArrayLength(len(in))
	if err != nil {
		return err
	}
	return encodeArrayElements(in, f.Ty.Encode, pe)
}

func (f *Array) GetName() string {
	return f.Name
}

func (f *Array) GetSchema() Schema {
	return f.Ty
}

// -- Array Compact

type ArrayCompact struct {
	Name string
	Ty   Schema
}

func (f *ArrayCompact) Decode(pd decoder.DDef) (interface{}, error) {
	n, err := pd.GetArrayCompactLength()
	if err != nil {
		return nil, errors.Wrapf(err, "getCompactArrayLength field %s", f.Name)
	}
	result, err := decodeArrayElements(n, f.Ty.Decode, pd)
	if err != nil {
		return nil, errors.Wrapf(err, "decodeArrayElements field %s", f.Name)
	}
	return result, err
}

func (f *ArrayCompact) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.([]interface{})
	if !ok {
		return proterrors.SchemaEncodingError{fmt.Sprintf("value %T not a []interface{}", value)}
	}
	err := pe.PutArrayCompactLength(len(in))
	if err != nil {
		return err
	}
	return encodeArrayElements(in, f.Ty.Encode, pe)
}

func (f *ArrayCompact) GetName() string {
	return f.Name
}
func (f *ArrayCompact) GetSchema() Schema {
	return f.Ty
}

// -- Array Compact Nullable

type ArrayCompactNullable struct {
	Name string
	Ty   EncoderDecoder
}

func (f *ArrayCompactNullable) Decode(pd decoder.DDef) (interface{}, error) {
	n, err := pd.GetArrayCompactNullableLength() // TODO: revisit implementation, compact nullable arrays are 0-length
	if err != nil {
		return nil, err
	}
	if n == -1 {
		return nil, nil
	}
	return decodeArrayElements(n, f.Ty.Decode, pd)
}

func (f *ArrayCompactNullable) Encode(pe encoder.EDef, value interface{}) error {
	if value == nil {
		return pe.PutArrayCompactNullableLength(0) // compact null array is 0-length
	}
	in, ok := value.([]interface{})
	if !ok {
		return proterrors.SchemaEncodingError{fmt.Sprintf("value %T not a []interface{}", value)}
	}

	err := pe.PutArrayCompactNullableLength(len(in))
	if err != nil {
		return err
	}
	return encodeArrayElements(in, f.Ty.Encode, pe)
}

func (f *ArrayCompactNullable) GetName() string {
	return f.Name
}
