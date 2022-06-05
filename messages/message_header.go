package messages

import "github.com/radekg/kafka-protocol-go/schema"

// MessageHeaderVersion is a message header version.
type MessageHeaderVersion int

const (
	// MessageHeaderV0 identifies a header version 0.
	MessageHeaderV0 MessageHeaderVersion = iota
	// MessageHeaderV1 identifies a header version 1.
	MessageHeaderV1
	// MessageHeaderV2 identifies a header version 2.
	MessageHeaderV2
)

// MessageHeader represents the message header.
type MessageHeader interface {
	APIKey() int16
	APIVersion() int16
	CorrelationID() int32
	ClientID() *string
	Length() int32
	Tags() *schema.TaggedFields

	Underlying() *schema.Struct
}

type defaultMessageHeader struct {
	underlying *schema.Struct
}

func (header *defaultMessageHeader) APIKey() int16 {
	return header.underlying.Get(FieldNameHeaderApiKey).(int16)
}
func (header *defaultMessageHeader) APIVersion() int16 {
	return header.underlying.Get(FieldNameHeaderApiVersion).(int16)
}
func (header *defaultMessageHeader) CorrelationID() int32 {
	return header.underlying.Get(FieldNameHeaderCorrelationId).(int32)
}
func (header *defaultMessageHeader) ClientID() *string {
	val := header.underlying.Get(FieldNameHeaderClientId)
	if val != nil {
		return val.(*string)
	}
	return nil
}
func (header *defaultMessageHeader) Length() int32 {
	return header.underlying.Get(FieldNameHeaderMessageLength).(int32)
}
func (header *defaultMessageHeader) Tags() *schema.TaggedFields {
	return readTagBuffer(header.underlying, FieldNameTagBuffer)
}

func (header *defaultMessageHeader) Underlying() *schema.Struct {
	return header.underlying
}
