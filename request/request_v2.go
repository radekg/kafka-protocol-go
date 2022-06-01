package request

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
	"github.com/radekg/kafka-protocol-go/schema"
)

type RequestV2 struct {
	CorrelationID int32
	ClientID      string
	TaggedFields  schema.TaggedFields
	Body          ProtocolBody
}

func (r *RequestV2) Encode(pe encoder.EDef) (err error) {
	pe.PutInt16(r.Body.Key())
	pe.PutInt16(r.Body.Version())
	pe.PutInt32(r.CorrelationID)
	err = pe.PutString(r.ClientID)
	if err != nil {
		return err
	}
	err = r.TaggedFields.Encode(pe)
	if err != nil {
		return err
	}
	err = r.Body.Encode(pe)
	if err != nil {
		return err
	}
	return err
}

func (r *RequestV2) Decode(pd decoder.DDef) (err error) {
	if r.Body == nil {
		return errors.PacketDecodingError{"unknown body decoder"}
	}
	var key int16
	if key, err = pd.GetInt16(); err != nil {
		return err
	}
	var version int16
	if version, err = pd.GetInt16(); err != nil {
		return err
	}
	if r.Body.Key() != key || r.Body.Version() != version {
		return errors.PacketDecodingError{fmt.Sprintf("expected request key,version %d,%d but got %d,%d", r.Body.Key(), r.Body.Version(), key, version)}
	}
	if r.CorrelationID, err = pd.GetInt32(); err != nil {
		return err
	}
	if r.ClientID, err = pd.GetString(); err != nil {
		return err
	}
	err = r.TaggedFields.Decode(pd)
	if err != nil {
		return err
	}
	return r.Body.Decode(pd)
}
