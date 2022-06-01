package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	proterrors "github.com/radekg/kafka-protocol-go/errors"
)

type Array struct {
	Name string
	Ty   Schema
}

func (f *Array) Decode(pd decoder.DDef) (interface{}, error) {
	n, err := pd.GetArrayLength()
	if err != nil {
		return nil, err
	}
	return decodeArrayElements(n, f.Ty.Decode, pd)
}

func (f *Array) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.([]interface{})
	if !ok {
		return proterrors.SchemaEncodingError{fmt.Sprintf("value %T not a []interface{}", value)}
	}
	err := pe.PutArrayLength(len(in))
	if err != nil {
		return err
	}
	return encodeArrayElements(in, f.Ty.Encode, pe)
}

func (f *Array) GetName() string {
	return f.Name
}

func (f *Array) GetSchema() Schema {
	return f.Ty
}
