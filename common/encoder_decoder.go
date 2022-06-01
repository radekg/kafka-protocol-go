package common

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/defs"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

// Encoder is the interface that wraps the basic Encode method.
// Anything implementing Encoder can be turned into bytes using Kafka's encoding rules.
type Encoder interface {
	Encode(pe encoder.EDef) error
}

// Encode takes an Encoder and turns it into bytes while potentially recording metrics.
func Encode(e Encoder) ([]byte, error) {
	if e == nil {
		return nil, nil
	}

	prepEnc := encoder.NewPrepEncoder()

	err := e.Encode(prepEnc)
	if err != nil {
		return nil, err
	}

	if prepEnc.Length() < 0 || prepEnc.Length() > int(defs.MaxRequestSize) {
		return nil, errors.PacketEncodingError{fmt.Sprintf("invalid request size (%d)", prepEnc.Length())}
	}

	realEnc := encoder.NewRealEncoder(make([]byte, prepEnc.Length()))

	err = e.Encode(realEnc)
	if err != nil {
		return nil, err
	}

	return realEnc.Bytes(), nil
}

// Decoder is the interface that wraps the basic Decode method.
// Anything implementing Decoder can be extracted from bytes using Kafka's encoding rules.
type Decoder interface {
	Decode(pd decoder.DDef) error
}

// Decode takes bytes and a Decoder and fills the fields of the decoder from the bytes,
// interpreted using Kafka's encoding rules.
func Decode(buf []byte, in Decoder) error {
	if buf == nil {
		return nil
	}

	helper := decoder.NewRealDecoder(buf)
	err := in.Decode(helper)
	if err != nil {
		return err
	}

	if helper.Offset() != len(buf) {
		return errors.PacketDecodingError{"invalid length"}
	}

	return nil
}
