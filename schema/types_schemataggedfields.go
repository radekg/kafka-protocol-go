package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

type SchemaTaggedFields struct {
	Name string
}

func (f SchemaTaggedFields) Decode(pd decoder.DDef) (interface{}, error) {
	numTaggedFields, err := pd.GetVarint()
	if err != nil {
		return nil, err
	}
	if numTaggedFields == 0 {
		result := make([]RawTaggedField, 0)
		return result, nil
	}
	if numTaggedFields < 0 {
		return nil, fmt.Errorf("Negative number of tagged fields %d", numTaggedFields)
	}
	result := make([]RawTaggedField, numTaggedFields)
	for i := 0; i < int(numTaggedFields); i++ {
		result[i].tag, err = pd.GetVarint()
		if err != nil {
			return nil, err
		}
		result[i].data, err = pd.GetVarintBytes()
		if err != nil {
			return nil, err
		}

	}
	return result, nil
}

func (f SchemaTaggedFields) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.([]RawTaggedField)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a []rawTaggedField", value)}
	}
	pe.PutVarint(int64(len(in)))
	for _, rawTaggedField := range in {
		pe.PutVarint(rawTaggedField.tag)
		err := pe.PutVarintBytes(rawTaggedField.data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f SchemaTaggedFields) GetName() string {
	return f.Name
}

func (f SchemaTaggedFields) GetSchema() Schema {
	return nil
}
