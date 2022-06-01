package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

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
