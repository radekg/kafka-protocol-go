package sasl

import (
	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
	"github.com/radekg/kafka-protocol-go/schema"
)

type SaslAuthenticateRequestV2 struct {
	SaslAuthBytes []byte
	TaggedFields  schema.TaggedFields
}

func (r *SaslAuthenticateRequestV2) Encode(pe encoder.EDef) error {
	if err := pe.PutCompactBytes(r.SaslAuthBytes); err != nil {
		return err
	}
	return r.TaggedFields.Encode(pe)
}

func (r *SaslAuthenticateRequestV2) Decode(pd decoder.DDef) (err error) {
	if r.SaslAuthBytes, err = pd.GetCompactBytes(); err != nil {
		return err
	}
	return r.TaggedFields.Decode(pd)
}

func (r *SaslAuthenticateRequestV2) Key() int16 {
	return 36
}

func (r *SaslAuthenticateRequestV2) Version() int16 {
	return 2
}

type SaslAuthenticateResponseV2 struct {
	Err               errors.KError
	ErrMsg            *string
	SaslAuthBytes     []byte
	SessionLifetimeMs int64
	TaggedFields      schema.TaggedFields
}

func (r *SaslAuthenticateResponseV2) Encode(pe encoder.EDef) error {
	pe.PutInt16(int16(r.Err))

	if err := pe.PutCompactNullableString(r.ErrMsg); err != nil {
		return err
	}

	if err := pe.PutCompactBytes(r.SaslAuthBytes); err != nil {
		return err
	}
	pe.PutInt64(r.SessionLifetimeMs)

	return r.TaggedFields.Encode(pe)
}

func (r *SaslAuthenticateResponseV2) Decode(pd decoder.DDef) error {
	kerr, err := pd.GetInt16()
	if err != nil {
		return err
	}
	r.Err = errors.KError(kerr)

	if r.ErrMsg, err = pd.GetCompactNullableString(); err != nil {
		return err
	}
	if r.SaslAuthBytes, err = pd.GetCompactBytes(); err != nil {
		return err
	}
	if r.SessionLifetimeMs, err = pd.GetInt64(); err != nil {
		return err
	}
	return r.TaggedFields.Decode(pd)
}
