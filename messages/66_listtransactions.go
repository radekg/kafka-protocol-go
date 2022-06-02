package messages

import "github.com/radekg/kafka-protocol-go/schema"

const (
	FieldNameListTransactionsStateFilters      = "state_filters"
	FieldNameListTransactionsProducerIdFilters = "producer_id_filters"
)

func initListTransactionsRequests() []schema.Schema {

	schemas := []schema.Schema{}

	schemas = append(schemas, schema.NewSchema("list_transactions_v0",
		&schema.Mfield{Name: FieldNameListTransactionsStateFilters, Ty: schema.TypeCompactStr},
		&schema.Mfield{Name: FieldNameListTransactionsProducerIdFilters, Ty: schema.TypeInt64},
		&schema.SchemaTaggedFields{Name: FieldNameTagBuffer}))

	return schemas

}
