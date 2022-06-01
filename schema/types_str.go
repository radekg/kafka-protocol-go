package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

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
