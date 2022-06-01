package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

type CompactNullableArray struct {
	Name string
	Ty   EncoderDecoder
}

func (f *CompactNullableArray) Decode(pd decoder.DDef) (interface{}, error) {
	n, err := pd.GetCompactNullableArrayLength()
	if err != nil {
		return nil, err
	}
	if n == -1 {
		return nil, nil
	}
	return decodeArrayElements(n, f.Ty.Decode, pd)
}

func (f *CompactNullableArray) Encode(pe encoder.EDef, value interface{}) error {
	if value == nil {
		return pe.PutCompactNullableArrayLength(-1)
	}
	in, ok := value.([]interface{})
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a []interface{}", value)}
	}

	err := pe.PutCompactNullableArrayLength(len(in))
	if err != nil {
		return err
	}
	return encodeArrayElements(in, f.Ty.Encode, pe)
}

func (f *CompactNullableArray) GetName() string {
	return f.Name
}
