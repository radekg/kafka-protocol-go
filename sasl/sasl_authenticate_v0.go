package sasl

import (
	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

type SaslAuthenticateRequestV0 struct {
	SaslAuthBytes []byte
}

func (r *SaslAuthenticateRequestV0) Encode(pe encoder.EDef) error {
	if err := pe.PutBytes(r.SaslAuthBytes); err != nil {
		return err
	}
	return nil
}
func (r *SaslAuthenticateRequestV0) Decode(pd decoder.DDef) (err error) {
	if r.SaslAuthBytes, err = pd.GetBytes(); err != nil {
		return err
	}

	return nil
}

func (r *SaslAuthenticateRequestV0) Key() int16 {
	return 36
}

func (r *SaslAuthenticateRequestV0) Version() int16 {
	return 0
}

type SaslAuthenticateResponseV0 struct {
	Err           errors.KError
	ErrMsg        *string
	SaslAuthBytes []byte
}

func (r *SaslAuthenticateResponseV0) Encode(pe encoder.EDef) error {
	pe.PutInt16(int16(r.Err))

	if err := pe.PutStringNullable(r.ErrMsg); err != nil {
		return err
	}
	return pe.PutBytes(r.SaslAuthBytes)
}

func (r *SaslAuthenticateResponseV0) Decode(pd decoder.DDef) error {
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
	return nil
}
