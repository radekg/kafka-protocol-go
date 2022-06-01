package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

type Varint struct{}

func (f *Varint) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetVarint()
}

func (f *Varint) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.(int64)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a varint", value)}
	}
	pe.PutVarint(in)
	return nil
}

func (f *Varint) GetFields() []boundField {
	return nil
}

func (f *Varint) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *Varint) GetName() string {
	return "varint"
}
