package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

// -- Str Array

type StrArray struct{}

func (f *StrArray) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetStringArray()
}

func (f *StrArray) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.([]string)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a strarray", value)}
	}
	pe.PutStringArray(in)
	return nil
}

func (f *StrArray) GetFields() []boundField {
	return nil
}

func (f *StrArray) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *StrArray) GetName() string {
	return "strarray"
}

// -- Str Array Compact

type StrCompactArray struct{}

func (f *StrCompactArray) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetStringCompactArray()
}

func (f *StrCompactArray) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.([]string)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a strcompactarray", value)}
	}
	pe.PutStringCompactArray(in)
	return nil
}

func (f *StrCompactArray) GetFields() []boundField {
	return nil
}

func (f *StrCompactArray) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *StrCompactArray) GetName() string {
	return "strcompactarray"
}
