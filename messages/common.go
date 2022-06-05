package messages

import "github.com/radekg/kafka-protocol-go/schema"

func readTagBuffer(st *schema.Struct, fieldName string) *schema.TaggedFields {
	vals := st.Get(fieldName)
	if vals == nil {
		return nil
	}
	return schema.NewTaggedFields(vals.([]schema.RawTaggedField))
}
