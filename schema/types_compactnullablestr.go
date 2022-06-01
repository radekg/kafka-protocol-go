package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

type CompactNullableStr struct{}

func (f *CompactNullableStr) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetCompactNullableString()
}

func (f *CompactNullableStr) Encode(pe encoder.EDef, value interface{}) error {
	if value == nil {
		if err := pe.PutCompactNullableString(nil); err != nil {
			return err
		}
	}

	in, ok := value.(*string)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a *string", value)}
	}
	return pe.PutCompactNullableString(in)
}

func (f *CompactNullableStr) GetFields() []boundField {
	return nil
}

func (f *CompactNullableStr) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *CompactNullableStr) GetName() string {
	return "compactnullablestr"
}
