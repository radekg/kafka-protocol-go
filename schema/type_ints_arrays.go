package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

// -- Int32

type Int32Array struct{}

func (f *Int32Array) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetInt32Array()
}

func (f *Int32Array) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.([]int32)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a int32array", value)}
	}
	pe.PutInt32Array(in)
	return nil
}

func (f *Int32Array) GetFields() []boundField {
	return nil
}

func (f *Int32Array) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *Int32Array) GetName() string {
	return "int32array"
}

// -- Int32 Compact

type Int32CompactArray struct{}

func (f *Int32CompactArray) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetInt32CompactArray()
}

func (f *Int32CompactArray) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.([]int32)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a int32compactarray", value)}
	}
	pe.PutInt32CompactArray(in)
	return nil
}

func (f *Int32CompactArray) GetFields() []boundField {
	return nil
}

func (f *Int32CompactArray) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *Int32CompactArray) GetName() string {
	return "int32compactarray"
}

// -- Int64

type Int64Array struct{}

func (f *Int64Array) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetInt64Array()
}

func (f *Int64Array) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.([]int64)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a int64array", value)}
	}
	pe.PutInt64Array(in)
	return nil
}

func (f *Int64Array) GetFields() []boundField {
	return nil
}

func (f *Int64Array) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *Int64Array) GetName() string {
	return "int64array"
}

// -- Int64 Compact

type Int64CompactArray struct{}

func (f *Int64CompactArray) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetInt64CompactArray()
}

func (f *Int64CompactArray) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.([]int64)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a int64compactarray", value)}
	}
	pe.PutInt64CompactArray(in)
	return nil
}

func (f *Int64CompactArray) GetFields() []boundField {
	return nil
}

func (f *Int64CompactArray) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *Int64CompactArray) GetName() string {
	return "int64compactarray"
}
