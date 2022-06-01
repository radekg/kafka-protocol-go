package schema

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/encoder"
	"github.com/radekg/kafka-protocol-go/errors"
)

type Uuid struct {
}

func (f *Uuid) Decode(pd decoder.DDef) (interface{}, error) {
	return pd.GetUUID()
}

func (f *Uuid) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.(uuid.UUID)
	if !ok {
		return errors.SchemaEncodingError{fmt.Sprintf("value %T not a uuid", value)}
	}
	return pe.PutUUID(in)
}

func (f *Uuid) GetFields() []boundField {
	return nil
}

func (f *Uuid) GetFieldsByName() map[string]*boundField {
	return nil
}

func (f *Uuid) GetName() string {
	return "uuid"
}
