package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

type CompactBytes struct{}

func (f *CompactBytes) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetCompactBytes()
}

func (f *CompactBytes) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.([]byte)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a compactbytes", value)}
	}
	pe.PutCompactBytes(in)
	return nil
}

func (f *CompactBytes) GetFields() []boundField {
	return nil
}

func (f *CompactBytes) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *CompactBytes) GetName() string {
	return "compactbytes"
}
