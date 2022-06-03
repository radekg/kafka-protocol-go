package messages

import "github.com/radekg/kafka-protocol-go/schema"

const (
	FieldNameHeaderAcks          = "acks"
	FieldNameHeaderApiKey        = "api_key"
	FieldNameHeaderApiVersion    = "api_version"
	FieldNameHeaderClientId      = "client_id"
	FieldNameHeaderCorrelationId = "correlation_id"
	FieldNameHeaderMessageLength = "message_length"
	FieldNameTagBuffer           = "tag_buffer"
)

func initHeaders() []schema.Schema {

	// Length is not part of the specification
	// but keeping it in the schema makes it easier to parse
	// in one go.

	schemas := []schema.Schema{}

	schemas = append(schemas, schema.NewSchema("header_v0",
		&schema.Mfield{Name: FieldNameHeaderMessageLength, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameHeaderApiKey, Ty: schema.TypeInt16},
		&schema.Mfield{Name: FieldNameHeaderApiVersion, Ty: schema.TypeInt16},
		&schema.Mfield{Name: FieldNameHeaderCorrelationId, Ty: schema.TypeInt32}))

	schemas = append(schemas, schema.NewSchema("header_v1",
		&schema.Mfield{Name: FieldNameHeaderMessageLength, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameHeaderApiKey, Ty: schema.TypeInt16},
		&schema.Mfield{Name: FieldNameHeaderApiVersion, Ty: schema.TypeInt16},
		&schema.Mfield{Name: FieldNameHeaderCorrelationId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameHeaderClientId, Ty: schema.TypeStrNullable}))

	schemas = append(schemas, schema.NewSchema("header_v2",
		&schema.Mfield{Name: FieldNameHeaderMessageLength, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameHeaderApiKey, Ty: schema.TypeInt16},
		&schema.Mfield{Name: FieldNameHeaderApiVersion, Ty: schema.TypeInt16},
		&schema.Mfield{Name: FieldNameHeaderCorrelationId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameHeaderClientId, Ty: schema.TypeStrNullable},
		&schema.SchemaTaggedFields{Name: FieldNameTagBuffer}))

	return schemas

}
