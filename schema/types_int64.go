package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

type Int64 struct{}

func (f *Int64) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetInt64()
}

func (f *Int64) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.(int64)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a int64", value)}
	}
	pe.PutInt64(in)
	return nil
}

func (f *Int64) GetFields() []boundField {
	return nil
}

func (f *Int64) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *Int64) GetName() string {
	return "int64"
}
