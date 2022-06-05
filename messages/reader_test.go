package messages

import (
	"testing"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/schema"
	"github.com/stretchr/testify/assert"
)

func TestDiscoversMessageHeaderVersion(t *testing.T) {

	expectedLength := int32(100)
	expectedApiKey := int16(3)
	expectedApiVersion := int16(1)
	expectedCorrelationId := int32(9999)
	expectedClientId := "some-test-client-id"

	expectedTagTag := int64(1000)
	expectedTagData := []byte("tag-data")

	expectedTagBuffer := []schema.RawTaggedField{
		schema.NewRawTaggedField(expectedTagTag, expectedTagData),
	}

	t.Run("it=discovers message header V2", func(tt *testing.T) {
		encoder := encoder.NewRealEncoder(make([]byte, 1000))
		headerSchema := headers[2]
		headerStruct := &schema.Struct{Schema: headerSchema, Values: []interface{}{
			expectedLength,     // length is irrelevant
			expectedApiKey,     // produce
			expectedApiVersion, // version 0
			expectedCorrelationId,
			&expectedClientId,
			expectedTagBuffer,
		}}
		encodeErr := headerStruct.Schema.Encode(encoder, headerStruct)
		assert.Nil(tt, encodeErr)
		encodedBytes := encoder.Bytes()[0:encoder.Offset()]

		headerDecoder := decoder.NewRealDecoder(encodedBytes)
		rr := NewDefaultRequestReader(headerDecoder)
		header, version, err := rr.ReadHeaderDiscoverVersion()
		assert.Nil(tt, err)
		assert.Equal(tt, MessageHeaderV2, version)

		assert.Equal(tt, expectedLength, header.Length())
		assert.Equal(tt, expectedApiKey, header.APIKey())
		assert.Equal(tt, expectedApiVersion, header.APIVersion())
		assert.Equal(tt, expectedCorrelationId, header.CorrelationID())
		assert.Equal(tt, expectedClientId, *header.ClientID())

		headerTags := header.Tags().Values()
		assert.Equal(tt, len(expectedTagBuffer), len(headerTags))
		assert.Equal(tt, expectedTagTag, headerTags[0].Tag())
		assert.Equal(tt, string(expectedTagData), string(headerTags[0].Data()))
	})

	t.Run("it=discovers message header V1", func(tt *testing.T) {
		encoder := encoder.NewRealEncoder(make([]byte, 1000))
		headerSchema := headers[1]
		headerStruct := &schema.Struct{Schema: headerSchema, Values: []interface{}{
			expectedLength,     // length is irrelevant
			expectedApiKey,     // produce
			expectedApiVersion, // version 0
			expectedCorrelationId,
			&expectedClientId,
		}}
		encodeErr := headerStruct.Schema.Encode(encoder, headerStruct)
		assert.Nil(tt, encodeErr)
		encodedBytes := encoder.Bytes()[0:encoder.Offset()]

		headerDecoder := decoder.NewRealDecoder(encodedBytes)
		rr := NewDefaultRequestReader(headerDecoder)
		header, version, err := rr.ReadHeaderDiscoverVersion()
		assert.Nil(tt, err)
		assert.Equal(tt, MessageHeaderV1, version)

		assert.Equal(tt, expectedLength, header.Length())
		assert.Equal(tt, expectedApiKey, header.APIKey())
		assert.Equal(tt, expectedApiVersion, header.APIVersion())
		assert.Equal(tt, expectedCorrelationId, header.CorrelationID())
		assert.Equal(tt, expectedClientId, *header.ClientID())
		assert.Nil(tt, header.Tags())
	})

	t.Run("it=discovers message header V0", func(tt *testing.T) {
		encoder := encoder.NewRealEncoder(make([]byte, 1000))
		headerSchema := headers[0]
		headerStruct := &schema.Struct{Schema: headerSchema, Values: []interface{}{
			expectedLength, // length is irrelevant
			expectedApiKey,
			expectedApiVersion,
			expectedCorrelationId,
		}}
		encodeErr := headerStruct.Schema.Encode(encoder, headerStruct)
		assert.Nil(tt, encodeErr)
		encodedBytes := encoder.Bytes()[0:encoder.Offset()]

		headerDecoder := decoder.NewRealDecoder(encodedBytes)
		rr := NewDefaultRequestReader(headerDecoder)
		header, version, err := rr.ReadHeaderDiscoverVersion()
		assert.Nil(tt, err)
		assert.Equal(tt, MessageHeaderV0, version)

		assert.Equal(tt, expectedLength, header.Length())
		assert.Equal(tt, expectedApiKey, header.APIKey())
		assert.Equal(tt, expectedApiVersion, header.APIVersion())
		assert.Equal(tt, expectedCorrelationId, header.CorrelationID())
		assert.Nil(tt, header.Tags())
	})

	t.Run("it=header V0 does not decode as V1", func(tt *testing.T) {
		encoder := encoder.NewRealEncoder(make([]byte, 1000))
		headerSchema := headers[0]
		headerStruct := &schema.Struct{Schema: headerSchema, Values: []interface{}{
			expectedLength, // length is irrelevant
			expectedApiKey,
			expectedApiVersion,
			expectedCorrelationId,
		}}
		encodeErr := headerStruct.Schema.Encode(encoder, headerStruct)
		assert.Nil(tt, encodeErr)
		encodedBytes := encoder.Bytes()[0:encoder.Offset()]

		headerDecoder := decoder.NewRealDecoder(encodedBytes)
		rr := NewDefaultRequestReader(headerDecoder)
		_, err := rr.ReadHeader(MessageHeaderV1)
		assert.NotNil(tt, err)
	})

	t.Run("it=header V0 does not decode as V2", func(tt *testing.T) {
		encoder := encoder.NewRealEncoder(make([]byte, 1000))
		headerSchema := headers[0]
		headerStruct := &schema.Struct{Schema: headerSchema, Values: []interface{}{
			expectedLength, // length is irrelevant
			expectedApiKey,
			expectedApiVersion,
			expectedCorrelationId,
		}}
		encodeErr := headerStruct.Schema.Encode(encoder, headerStruct)
		assert.Nil(tt, encodeErr)
		encodedBytes := encoder.Bytes()[0:encoder.Offset()]

		headerDecoder := decoder.NewRealDecoder(encodedBytes)
		rr := NewDefaultRequestReader(headerDecoder)
		_, err := rr.ReadHeader(MessageHeaderV2)
		assert.NotNil(tt, err)
	})

	t.Run("it=header V1 does not decode as V2", func(tt *testing.T) {
		encoder := encoder.NewRealEncoder(make([]byte, 1000))
		headerSchema := headers[1]
		headerStruct := &schema.Struct{Schema: headerSchema, Values: []interface{}{
			expectedLength, // length is irrelevant
			expectedApiKey,
			expectedApiVersion,
			expectedCorrelationId,
			&expectedClientId,
		}}
		encodeErr := headerStruct.Schema.Encode(encoder, headerStruct)
		assert.Nil(tt, encodeErr)
		encodedBytes := encoder.Bytes()[0:encoder.Offset()]

		headerDecoder := decoder.NewRealDecoder(encodedBytes)
		rr := NewDefaultRequestReader(headerDecoder)
		_, err := rr.ReadHeader(MessageHeaderV2)
		assert.NotNil(tt, err)
	})

}
