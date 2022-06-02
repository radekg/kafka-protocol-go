package messages

import "github.com/radekg/kafka-protocol-go/schema"

const (
	FieldNameProduceAcks            = "acks"
	FieldNameProducePartition       = "partition"
	FieldNameProducePartitionData   = "partition_data"
	FieldNameProduceRawRecords      = "raw_records"
	FieldNameProduceTimeoutMs       = "timeout_ms"
	FieldNameProduceTopicData       = "topic_data"
	FieldNameProduceTopicName       = "topic_name"
	FieldNameProduceTransactionalId = "transactional_id"
)

func initProduceRequests() []schema.Schema {

	schemas := []schema.Schema{}

	// TODO: verify: most likely Array instead of CompactArray
	produceRequestSchemaV0 := schema.NewSchema("produce_v0",
		&schema.Mfield{Name: FieldNameProduceAcks, Ty: schema.TypeInt16},
		&schema.Mfield{Name: FieldNameProduceTimeoutMs, Ty: schema.TypeInt32},
		&schema.CompactArray{Name: FieldNameProduceTopicData, Ty: schema.NewSchema("produce_topic_data_v0",
			&schema.Mfield{Name: FieldNameProduceTopicName, Ty: schema.TypeStr},
			&schema.CompactArray{Name: FieldNameProducePartitionData, Ty: schema.NewSchema("produce_partition_data_v0",
				&schema.Mfield{Name: FieldNameProducePartition, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldNameProduceRawRecords, Ty: schema.TypeBytes},
			)},
		)},
	)

	schemas = append(schemas, produceRequestSchemaV0)
	schemas = append(schemas, produceRequestSchemaV0)
	schemas = append(schemas, produceRequestSchemaV0)

	// Version 3 adds transactional_id:
	// TODO: verify: most likely Array instead of CompactArray
	produceRequestSchemaV3 := schema.NewSchema("produce_v3",
		&schema.Mfield{Name: FieldNameProduceTransactionalId, Ty: schema.TypeNullableStr},
		&schema.Mfield{Name: FieldNameProduceAcks, Ty: schema.TypeInt16},
		&schema.Mfield{Name: FieldNameProduceTimeoutMs, Ty: schema.TypeInt32},
		&schema.CompactArray{Name: FieldNameProduceTopicData, Ty: schema.NewSchema("produce_topic_data_v3",
			&schema.Mfield{Name: FieldNameProduceTopicName, Ty: schema.TypeStr},
			&schema.CompactArray{Name: FieldNameProducePartitionData, Ty: schema.NewSchema("produce_partition_data_v3",
				&schema.Mfield{Name: FieldNameProducePartition, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldNameProduceRawRecords, Ty: schema.TypeBytes},
			)},
		)},
	)

	schemas = append(schemas, produceRequestSchemaV3)
	schemas = append(schemas, produceRequestSchemaV3)
	schemas = append(schemas, produceRequestSchemaV3)
	schemas = append(schemas, produceRequestSchemaV3)
	schemas = append(schemas, produceRequestSchemaV3)
	schemas = append(schemas, produceRequestSchemaV3)

	// Version 9 adds tag buffers and changes:
	//  - Str => CompactStr
	//  - NullableStr => CompactNullableStr
	//  - Records => CompactRecords

	produceRequestSchemaV9 := schema.NewSchema("produce_v9",
		&schema.Mfield{Name: FieldNameProduceTransactionalId, Ty: schema.TypeCompactNullableStr},
		&schema.Mfield{Name: FieldNameProduceAcks, Ty: schema.TypeInt16},
		&schema.Mfield{Name: FieldNameProduceTimeoutMs, Ty: schema.TypeInt32},
		&schema.CompactArray{Name: FieldNameProduceTopicData, Ty: schema.NewSchema("produce_topic_data_v9",
			&schema.Mfield{Name: FieldNameProduceTopicName, Ty: schema.TypeCompactStr},
			&schema.CompactArray{Name: FieldNameProducePartitionData, Ty: schema.NewSchema("produce_partition_data_v9",
				&schema.Mfield{Name: FieldNameProducePartition, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldNameProduceRawRecords, Ty: schema.TypeCompactBytes},
				&schema.SchemaTaggedFields{Name: FieldNameTagBuffer},
			)},
			&schema.SchemaTaggedFields{Name: FieldNameTagBuffer},
		)},
		&schema.SchemaTaggedFields{Name: FieldNameTagBuffer},
	)

	schemas = append(schemas, produceRequestSchemaV9)

	return schemas

}
