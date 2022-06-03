package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

// -- Str

type Str struct{}

func (f *Str) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetString()
}

func (f *Str) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.(string)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a string", value)}
	}
	return pe.PutString(in)
}

func (f *Str) GetFields() []boundField {
	return nil
}

func (f *Str) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *Str) GetName() string {
	return "str"
}

// -- Str Compact

type StrCompact struct {
}

func (f *StrCompact) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetStringCompact()
}

func (f *StrCompact) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.(string)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a string", value)}
	}
	return pe.PutStringCompact(in)
}

func (f *StrCompact) GetFields() []boundField {
	return nil
}

func (f *StrCompact) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *StrCompact) GetName() string {
	return "compactstr"
}

// -- Str Compact Nullable

type StrCompactNullable struct{}

func (f *StrCompactNullable) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetStringCompactNullable()
}

func (f *StrCompactNullable) Encode(pe encoder.EDef, value interface{}) error {
	if value == nil {
		if err := pe.PutStringCompactNullable(nil); err != nil {
			return err
		}
	}

	in, ok := value.(*string)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a *string", value)}
	}
	return pe.PutStringCompactNullable(in)
}

func (f *StrCompactNullable) GetFields() []boundField {
	return nil
}

func (f *StrCompactNullable) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *StrCompactNullable) GetName() string {
	return "compactnullablestr"
}

// -- Str Nullable

type StrNullable struct{}

func (f *StrNullable) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetStringNullable()
}

func (f *StrNullable) Encode(pe encoder.EDef, value interface{}) error {
	if value == nil {
		if err := pe.PutStringNullable(nil); err != nil {
			return err
		}
	}

	in, ok := value.(*string)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a *string", value)}
	}
	return pe.PutStringNullable(in)
}

func (f *StrNullable) GetFields() []boundField {
	return nil
}

func (f *StrNullable) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *StrNullable) GetName() string {
	return "nullablestr"
}
