package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

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
