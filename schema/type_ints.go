package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

// -- Int8

type Int8 struct{}

func (f *Int8) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetInt8()
}
func (f *Int8) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.(int8)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a int8", value)}
	}
	pe.PutInt8(in)
	return nil
}

func (f *Int8) GetFields() []boundField {
	return nil
}

func (f *Int8) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *Int8) GetName() string {
	return "int8"
}

// -- Int16

type Int16 struct{}

func (f *Int16) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetInt16()
}
func (f *Int16) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.(int16)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a int16", value)}
	}
	pe.PutInt16(in)
	return nil
}

func (f *Int16) GetFields() []boundField {
	return nil
}

func (f *Int16) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *Int16) GetName() string {
	return "int16"
}

// -- Int32

type Int32 struct{}

func (f *Int32) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetInt32()
}

func (f *Int32) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.(int32)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a int32", value)}
	}
	pe.PutInt32(in)
	return nil
}

func (f *Int32) GetFields() []boundField {
	return nil
}

func (f *Int32) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *Int32) GetName() string {
	return "int32"
}

// -- Int64

type Int64 struct{}

func (f *Int64) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetInt64()
}

func (f *Int64) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.(int64)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a int64", value)}
	}
	pe.PutInt64(in)
	return nil
}

func (f *Int64) GetFields() []boundField {
	return nil
}

func (f *Int64) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *Int64) GetName() string {
	return "int64"
}
