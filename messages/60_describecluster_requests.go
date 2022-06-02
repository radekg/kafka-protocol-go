package messages

import "github.com/radekg/kafka-protocol-go/schema"

const (
	FieldNameDescribeClusterIncludeClusterAuthorizedOperations = "include_cluster_authorized_operations"
)

func initDescribeClusterRequests() []schema.Schema {

	schemas := []schema.Schema{}

	schemas = append(schemas, schema.NewSchema("describe_cluster_v0",
		&schema.Mfield{Name: FieldNameDescribeClusterIncludeClusterAuthorizedOperations, Ty: schema.TypeBool},
		&schema.SchemaTaggedFields{Name: FieldNameTagBuffer}))

	return schemas

}
