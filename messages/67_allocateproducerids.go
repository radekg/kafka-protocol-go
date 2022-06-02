package messages

import "github.com/radekg/kafka-protocol-go/schema"

const (
	FieldNameAllocateProducerIdsBrokerEpoch = "broker_epoch"
	FieldNameAllocateProducerIdsBrokerId    = "broker_id"
)

func initAllocateProducerIdRequests() []schema.Schema {

	schemas := []schema.Schema{}

	schemas = append(schemas, schema.NewSchema("allocate_producer_ids_v0",
		&schema.Mfield{Name: FieldNameAllocateProducerIdsBrokerId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameAllocateProducerIdsBrokerEpoch, Ty: schema.TypeInt64},
		&schema.SchemaTaggedFields{Name: FieldNameTagBuffer}))

	return schemas

}
