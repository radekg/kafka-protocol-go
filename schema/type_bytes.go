package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

// -- Bytes

type Bytes struct{}

func (f *Bytes) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetBytes()
}

func (f *Bytes) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.([]byte)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a bytes", value)}
	}
	pe.PutBytes(in)
	return nil
}

func (f *Bytes) GetFields() []boundField {
	return nil
}

func (f *Bytes) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *Bytes) GetName() string {
	return "bytes"
}

// -- Bytes Compact

type BytesCompact struct{}

func (f *BytesCompact) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetBytesCompact()
}

func (f *BytesCompact) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.([]byte)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a bytescompact", value)}
	}
	pe.PutBytesCompact(in)
	return nil
}

func (f *BytesCompact) GetFields() []boundField {
	return nil
}

func (f *BytesCompact) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *BytesCompact) GetName() string {
	return "bytescompact"
}
