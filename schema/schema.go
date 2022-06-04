package schema

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/radekg/kafka-protocol-go/decoder"
	"github.com/radekg/kafka-protocol-go/defs"
	"github.com/radekg/kafka-protocol-go/encoder"
	proterrors "github.com/radekg/kafka-protocol-go/errors"
)

var (
	TypeBool = &Bool{}

	// Bytes:
	TypeBytes                = &Bytes{}
	TypeBytesCompact         = &BytesCompact{}
	TypeBytesCompactNullable = &BytesCompact{} // Compact nullable bytes
	TypeBytesNullable        = &Bytes{}        // TODO: proper implementation

	// Int types:
	TypeInt8              = &Int8{}
	TypeInt16             = &Int16{}
	TypeInt32             = &Int32{}
	TypeInt32Array        = &Int32Array{}
	TypeInt32CompactArray = &Int32CompactArray{}
	TypeInt64             = &Int64{}
	TypeInt64Array        = &Int64Array{}
	TypeInt64CompactArray = &Int64CompactArray{}
	TypeVarint            = &Varint{}

	// Uint types:
	// TODO: add uint16
	// TODO: add uint32

	// Float types:
	// TODO: add float64

	// String types:
	TypeStr                = &Str{}
	TypeStrArray           = &StrArray{}
	TypeStrCompact         = &StrCompact{}
	TypeStrCompactArray    = &StrCompactArray{}
	TypeStrCompactNullable = &StrCompactNullable{}
	TypeStrNullable        = &StrNullable{}

	// UUID types:
	TypeUuid = &Uuid{}

	// TODO: TypeInt32CompactArray
)

type EncoderDecoder interface {
	Decode(pd decoder.DDef) (interface{}, error)     // reads / gets
	Encode(pe encoder.EDef, value interface{}) error // writes / puts
}

type Field interface {
	EncoderDecoder
	GetName() string
	GetSchema() Schema
}

type Schema interface {
	EncoderDecoder
	GetFields() []boundField
	GetFieldsByName() map[string]*boundField
	GetName() string
}

type schema struct {
	name         string
	fields       []boundField
	fieldsByName map[string]*boundField
}

//
type Mfield struct {
	Name string
	Ty   Schema
}

func (f *Mfield) Decode(pd decoder.DDef) (interface{}, error) {
	return f.Ty.Decode(pd)
}

func (f *Mfield) Encode(pe encoder.EDef, value interface{}) error {
	return f.Ty.Encode(pe, value)
}
func (f *Mfield) GetName() string {
	return f.Name
}
func (f *Mfield) GetSchema() Schema {
	return f.Ty
}

// Arrays helper

func encodeArrayElements(in []interface{}, elementEncode func(pe encoder.EDef, value interface{}) error, pe encoder.EDef) (err error) {
	for _, elem := range in {
		err = elementEncode(pe, elem)
		if err != nil {
			return err
		}
	}
	return nil
}

func decodeArrayElements(n int, elementDecode func(pd decoder.DDef) (interface{}, error), pd decoder.DDef) (interface{}, error) {
	// We could allocate the capacity at once, but in case of malformed payload we could allocate too much memory.
	result := make([]interface{}, 0)

	for i := 0; i < n; i++ {
		elem, err := elementDecode(pd)
		if err != nil {
			return nil, err
		}
		result = append(result, elem)
	}
	return result, nil
}

// Structures

type Struct struct {
	Schema Schema
	Values []interface{}
}

func (s Struct) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(s.GetSchema().GetName() + "{")
	for i, field := range s.GetSchema().GetFields() {
		if i != 0 {
			buffer.WriteString(",")
		}
		name := field.def.GetName()
		buffer.WriteString(fmt.Sprintf("%s:%v", name, s.Get(name)))
	}
	buffer.WriteString("}")
	return buffer.String()
}

func (s Struct) Get(name string) interface{} {
	bf := s.GetSchema().GetFieldsByName()[name]
	if bf == nil || bf.index >= len(s.Values) {
		return nil
	}
	return s.Values[bf.index]
}

func (s *Struct) Replace(name string, value interface{}) error {
	if value == nil {
		return fmt.Errorf("field %s value in struct %s : new value must not be nil", name, s.GetSchema().GetName())
	}
	bf := s.GetSchema().GetFieldsByName()[name]
	if bf == nil {
		return fmt.Errorf("field %s value in struct %s : name not found", name, s.GetSchema().GetName())
	}
	if bf.index >= len(s.Values) {
		return fmt.Errorf("field %s value in struct %s : index %d gte %d", name, s.GetSchema().GetName(), bf.index, len(s.Values))
	}
	v := s.Values[bf.index]
	if v == nil {
		return fmt.Errorf("field %s value in struct %s : old value not found", name, s.GetSchema().GetName())
	}
	oldKind := reflect.TypeOf(v).Kind()
	newKind := reflect.TypeOf(value).Kind()
	if oldKind != newKind {
		return fmt.Errorf("field %s value in struct %s : kinds differ %v to %v", name, s.GetSchema().GetName(), oldKind, newKind)
	}
	s.Values[bf.index] = value
	return nil
}

func (s *Struct) GetSchema() Schema {
	return s.Schema
}

// NewSchema creates new schema. It panics when a duplicate field is provided
func NewSchema(name string, fs ...Field) Schema {

	s := &schema{name: name, fields: make([]boundField, 0), fieldsByName: make(map[string]*boundField)}

	for i, f := range fs {
		if _, ok := s.fieldsByName[f.GetName()]; ok {
			panic(fmt.Sprintf("Schema contains a duplicate field: %s", f.GetName()))
		}
		bf := boundField{
			def:    f,
			index:  i,
			schema: s,
		}
		s.fields = append(s.fields, bf)
		s.fieldsByName[f.GetName()] = &bf
	}
	return s
}

func (s *schema) Encode(pe encoder.EDef, value interface{}) error {
	in, ok := value.(*Struct)
	if !ok {
		return proterrors.SchemaEncodingError{fmt.Sprintf("value %T not a *Struct", value)}
	}
	if len(in.Values) != len(s.GetFields()) {
		return proterrors.SchemaEncodingError{fmt.Sprintf("length difference: values %d, struct fields %d", len(in.Values), len(s.GetFields()))}
	}
	for i, value := range in.Values {
		if err := s.GetFields()[i].def.Encode(pe, value); err != nil {
			return err
		}
	}
	return nil
}

func (s *schema) Decode(pd decoder.DDef) (interface{}, error) {
	values := make([]interface{}, 0)

	for _, field := range s.GetFields() {
		value, err := field.def.Decode(pd)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return &Struct{Schema: s, Values: values}, nil
}

func (s *schema) GetFields() []boundField {
	return s.fields
}
func (s *schema) GetFieldsByName() map[string]*boundField {
	return s.fieldsByName
}
func (s *schema) GetName() string {
	return s.name
}

func DecodeSchema(buf []byte, schema Schema) (*Struct, error) {
	if buf == nil {
		return nil, nil
	}
	helper := decoder.NewRealDecoder(buf)
	v, err := schema.Decode(helper)
	if err != nil {
		return nil, err
	}
	if helper.Offset() != len(buf) {
		return nil, proterrors.SchemaDecodingError{"invalid length"}
	}

	st, ok := v.(*Struct)
	if !ok {
		return nil, proterrors.SchemaDecodingError{"internal error: schema decode should return *Struct"}
	}
	//log.Printf("Decoded Schema %v", v)

	return st, nil
}

func EncodeSchema(s *Struct, schema Schema) ([]byte, error) {
	if s == nil {
		return nil, nil
	}

	prepEnc := encoder.NewPrepEncoder()

	err := schema.Encode(prepEnc, s)
	if err != nil {
		return nil, err
	}

	if prepEnc.Length() < 0 || prepEnc.Length() > int(defs.MaxRequestSize) {
		return nil, proterrors.SchemaEncodingError{fmt.Sprintf("invalid request size (%d)", prepEnc.Length())}
	}

	realEnc := encoder.NewRealEncoder(make([]byte, prepEnc.Length()))

	err = schema.Encode(realEnc, s)
	if err != nil {
		return nil, err
	}

	return realEnc.Bytes(), nil
}
