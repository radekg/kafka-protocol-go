package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

// -- Float64

type Float64 struct{}

func (f *Float64) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetFloat64()
}
func (f *Float64) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.(float64)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a float64", value)}
	}
	pe.PutFloat64(in)
	return nil
}

func (f *Float64) GetFields() []boundField {
	return nil
}

func (f *Float64) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *Float64) GetName() string {
	return "float64"
}
