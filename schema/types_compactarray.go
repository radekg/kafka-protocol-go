package schema

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	proterrors "github.com/radekg/kafka-protocol-go/errors"
)

type CompactArray struct {
	Name string
	Ty   Schema
}

func (f *CompactArray) Decode(pd decoder.DDef) (interface{}, error) {
	n, err := pd.GetCompactArrayLength()
	if err != nil {
		return nil, errors.Wrapf(err, "getCompactArrayLength field %s", f.Name)
	}
	result, err := decodeArrayElements(n, f.Ty.Decode, pd)
	if err != nil {
		return nil, errors.Wrapf(err, "decodeArrayElements field %s", f.Name)
	}
	return result, err
}

func (f *CompactArray) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.([]interface{})
	if !ok {
		return proterrors.SchemaEncodingError{fmt.Sprintf("value %T not a []interface{}", value)}
	}
	err := pe.PutCompactArrayLength(len(in))
	if err != nil {
		return err
	}
	return encodeArrayElements(in, f.Ty.Encode, pe)
}

func (f *CompactArray) GetName() string {
	return f.Name
}
func (f *CompactArray) GetSchema() Schema {
	return f.Ty
}
