package messages

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/errors"
	"github.com/radekg/kafka-protocol-go/schema"
)

type RequestReader interface {
	ReadHeaderDiscoverVersion() (MessageHeader, MessageHeaderVersion, error)
	ReadHeader(version MessageHeaderVersion) (MessageHeader, error)
}

func NewDefaultRequestReader(decoder decoder.DDef) RequestReader {
	return &defaultRequestReader{decoder: decoder}
}

type defaultRequestReader struct {
	decoder decoder.DDef
}

func (rr *defaultRequestReader) ReadHeaderDiscoverVersion() (MessageHeader, MessageHeaderVersion, error) {
	for i := len(headers) - 1; i >= 0; i = i - 1 {
		tempDecoder := rr.decoder.CopyNew()
		if val, err := headers[i].Decode(tempDecoder); err == nil {
			if st, ok := val.(*schema.Struct); ok {
				// We're good after parsing the header with selected version,
				// let's hold on to the decoder instance.
				// With this, we can continue reading the message out of the underlying decoder
				// after reading the header.
				rr.decoder = tempDecoder
				return &defaultMessageHeader{underlying: st}, MessageHeaderVersion(i), nil
			} else {
				return nil, -1, errors.SchemaDecodingError{Info: "header read error: schema decode should return *Struct"}
			}
		}
	}
	return nil, -1, errors.SchemaDecodingError{Info: "header read error: unrecognised message header"}
}

func (rr *defaultRequestReader) ReadHeader(version MessageHeaderVersion) (MessageHeader, error) {
	if version < MessageHeaderV0 || int(version) > len(headers) {
		return nil, errors.SchemaDecodingError{Info: fmt.Sprintf("header read error: unsupported message header version: %d", version)}
	}
	val, err := headers[version].Decode(rr.decoder)
	if err != nil {
		return nil, err
	}
	if st, ok := val.(*schema.Struct); ok {
		return &defaultMessageHeader{underlying: st}, nil

	}
	return nil, errors.SchemaDecodingError{Info: "header read error: schema decode should return *Struct"}
}
