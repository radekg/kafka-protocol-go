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

func NewRawTaggedField(tag int64, data []byte) RawTaggedField {
	return RawTaggedField{tag: tag, data: data}
}

func (rt *RawTaggedField) Tag() int64 {
	return rt.tag
}

func (rt *RawTaggedField) Data() []byte {
	return rt.data
}

type TaggedFields struct {
	values []RawTaggedField
}

func NewTaggedFields(values []RawTaggedField) *TaggedFields {
	return &TaggedFields{values: values}
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

func (r *TaggedFields) Values() []RawTaggedField {
	return r.values
}
