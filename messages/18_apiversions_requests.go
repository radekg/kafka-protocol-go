package messages

import "github.com/radekg/kafka-protocol-go/schema"

const (
	FieldNameApiVersionsClientSoftwareName    = "client_software_name"
	FieldNameApiVersionsClientSoftwareVersion = "client_software_version"
)

func initApiVersionsRequests() []schema.Schema {

	schemas := []schema.Schema{}

	schemas = append(schemas, schema.NewSchema("apiversions_v0"))
	schemas = append(schemas, schema.NewSchema("apiversions_v1"))
	schemas = append(schemas, schema.NewSchema("apiversions_v2"))

	schemas = append(schemas, schema.NewSchema("apiversions_v3",
		&schema.Mfield{Name: FieldNameApiVersionsClientSoftwareName, Ty: schema.TypeCompactStr},
		&schema.Mfield{Name: FieldNameApiVersionsClientSoftwareVersion, Ty: schema.TypeCompactStr},
		&schema.SchemaTaggedFields{Name: FieldNameTagBuffer}))

	return schemas

}
