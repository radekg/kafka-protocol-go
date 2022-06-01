package schema

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

type RawTaggedField struct {
	tag  int64
	data []byte
}

type TaggedFields struct {
	values []RawTaggedField
}

func (r *TaggedFields) Decode(pd decoder.DDef) (err error) {
	tf := &SchemaTaggedFields{}
	taggedFields, err := tf.Decode(pd)
	if err != nil {
		return err
	}
	if rawTaggedFields, ok := taggedFields.([]RawTaggedField); ok {
		r.values = rawTaggedFields
	} else {
		return errors.PacketDecodingError{fmt.Sprintf("taggedFields type %v", taggedFields)}
	}
	return err
}

func (r *TaggedFields) Encode(pe encoder.EDef) (err error) {
	tf := &SchemaTaggedFields{}
	err = tf.Encode(pe, r.values)
	if err != nil {
		return err
	}
	return nil
}
