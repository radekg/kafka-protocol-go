package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

type Bool struct{}

func (f *Bool) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetBool()
}

func (f *Bool) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.(bool)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a bool", value)}
	}
	pe.PutBool(in)
	return nil
}

func (f *Bool) GetFields() []boundField {
	return nil
}

func (f *Bool) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *Bool) GetName() string {
	return "bool"
}
