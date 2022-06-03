package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init66ListTransactionsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: ListTransactionsRequest, API Key: 66, Version: 0
		schema.NewSchema("ListTransactionsRequestv0", 
			&schema.Mfield{Name: FieldListTransactionsRequestStateFilters, Ty: schema.TypeStrCompactArray},
			&schema.Mfield{Name: FieldListTransactionsRequestProducerIdFilters, Ty: schema.TypeInt64CompactArray},
			&schema.SchemaTaggedFields{Name: FieldListTransactionsRequestTags},
		),

	}
}

const (
	// FieldListTransactionsRequestStateFilters is a field name that can be used to resolve the correct struct field.
	FieldListTransactionsRequestStateFilters = "StateFilters"
	// FieldListTransactionsRequestProducerIdFilters is a field name that can be used to resolve the correct struct field.
	FieldListTransactionsRequestProducerIdFilters = "ProducerIdFilters"
	// FieldListTransactionsRequestTags is a field name that can be used to resolve the correct struct field.
	FieldListTransactionsRequestTags = "Tags"
)

