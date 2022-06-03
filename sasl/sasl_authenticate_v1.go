package sasl

import (
	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

type SaslAuthenticateRequestV1 struct {
	SaslAuthBytes []byte
}

func (r *SaslAuthenticateRequestV1) Encode(pe encoder.EDef) error {
	if err := pe.PutBytes(r.SaslAuthBytes); err != nil {
		return err
	}
	return nil
}
func (r *SaslAuthenticateRequestV1) Decode(pd decoder.DDef) (err error) {
	if r.SaslAuthBytes, err = pd.GetBytes(); err != nil {
		return err
	}

	return nil
}

func (r *SaslAuthenticateRequestV1) Key() int16 {
	return 36
}

func (r *SaslAuthenticateRequestV1) Version() int16 {
	return 1
}

type SaslAuthenticateResponseV1 struct {
	Err               errors.KError
	ErrMsg            *string
	SaslAuthBytes     []byte
	SessionLifetimeMs int64
}

func (r *SaslAuthenticateResponseV1) Encode(pe encoder.EDef) error {
	pe.PutInt16(int16(r.Err))

	if err := pe.PutStringNullable(r.ErrMsg); err != nil {
		return err
	}

	if err := pe.PutBytes(r.SaslAuthBytes); err != nil {
		return err
	}
	pe.PutInt64(r.SessionLifetimeMs)

	return nil
}

func (r *SaslAuthenticateResponseV1) Decode(pd decoder.DDef) error {
	kerr, err := pd.GetInt16()
	if err != nil {
		return err
	}
	r.Err = errors.KError(kerr)

	if r.ErrMsg, err = pd.GetStringNullable(); err != nil {
		return err
	}
	if r.SaslAuthBytes, err = pd.GetBytes(); err != nil {
		return err
	}
	if r.SessionLifetimeMs, err = pd.GetInt64(); err != nil {
		return err
	}
	return nil
}
