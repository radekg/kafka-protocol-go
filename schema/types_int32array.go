package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

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
