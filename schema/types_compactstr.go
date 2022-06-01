package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

type CompactStr struct {
}

func (f *CompactStr) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetCompactString()
}

func (f *CompactStr) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.(string)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a string", value)}
	}
	return pe.PutCompactString(in)
}

func (f *CompactStr) GetFields() []boundField {
	return nil
}

func (f *CompactStr) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *CompactStr) GetName() string {
	return "compactstr"
}
