package main

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
	"github.com/radekg/kafka-protocol-go/schema"
)

type ResponseHeaderV1 struct {
	Length          int32
	CorrelationID   int32
	RawTaggedFields []schema.RawTaggedField
}

func (r *ResponseHeaderV1) Decode(pd decoder.DDef) (err error) {
	r.Length, err = pd.GetInt32()
	if err != nil {
		return err
	}
	if r.Length <= 4 {
		return errors.PacketDecodingError{fmt.Sprintf("message of length %d too small", r.Length)}
	}
	r.CorrelationID, err = pd.GetInt32()

	tf := &schema.SchemaTaggedFields{}
	taggedFields, err := tf.Decode(pd)
	if err != nil {
		return err
	}
	if rawTaggedFields, ok := taggedFields.([]schema.RawTaggedField); ok {
		r.RawTaggedFields = rawTaggedFields
	} else {
		return errors.PacketDecodingError{fmt.Sprintf("taggedFields type %v", taggedFields)}
	}
	return err
}

func (r *ResponseHeaderV1) Encode(pe encoder.EDef) (err error) {
	pe.PutInt32(r.Length)
	pe.PutInt32(r.CorrelationID)

	tf := &schema.SchemaTaggedFields{}
	err = tf.Encode(pe, r.RawTaggedFields)
	if err != nil {
		return err
	}
	return nil
}
