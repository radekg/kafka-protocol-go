package sasl

import (
	"fmt"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

type SaslHandshakeRequestV0orV1 struct {
	Version   int16 // not encoded / decoded
	Mechanism string
}

func (r *SaslHandshakeRequestV0orV1) Encode(pe encoder.EDef) error {
	if r.Version != 0 && r.Version != 1 {
		return fmt.Errorf("SaslHandshakeRequestV0orV1 expects version 0 or 1")
	}

	if err := pe.PutString(r.Mechanism); err != nil {
		return err
	}
	return nil
}
func (r *SaslHandshakeRequestV0orV1) Decode(pd decoder.DDef) (err error) {
	if r.Version != 0 && r.Version != 1 {
		return fmt.Errorf("SaslHandshakeRequestV0orV1 expects version 0 or 1")
	}
	if r.Mechanism, err = pd.GetString(); err != nil {
		return err
	}

	return nil
}

func (r *SaslHandshakeRequestV0orV1) key() int16 {
	return 17
}

func (r *SaslHandshakeRequestV0orV1) version() int16 {
	return r.Version
}

type SaslHandshakeResponseV0orV1 struct {
	Err               errors.KError
	EnabledMechanisms []string
}

func (r *SaslHandshakeResponseV0orV1) encode(pe encoder.EDef) error {
	pe.PutInt16(int16(r.Err))
	return pe.PutStringArray(r.EnabledMechanisms)
}

func (r *SaslHandshakeResponseV0orV1) decode(pd decoder.DDef) error {
	kerr, err := pd.GetInt16()
	if err != nil {
		return err
	}
	r.Err = errors.KError(kerr)
	if r.EnabledMechanisms, err = pd.GetStringArray(); err != nil {
		return err
	}
	return nil
}
