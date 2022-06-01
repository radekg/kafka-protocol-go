package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

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
