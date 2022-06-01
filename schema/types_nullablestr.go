package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

type NullableStr struct{}

func (f *NullableStr) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetNullableString()
}

func (f *NullableStr) Encode(pe encoder.EDef, value interface{}) error {
	if value == nil {
		if err := pe.PutNullableString(nil); err != nil {
			return err
		}
	}

	in, ok := value.(*string)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a *string", value)}
	}
	return pe.PutNullableString(in)
}

func (f *NullableStr) GetFields() []boundField {
	return nil
}

func (f *NullableStr) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *NullableStr) GetName() string {
	return "nullablestr"
}
