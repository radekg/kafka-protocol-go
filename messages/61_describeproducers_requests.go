package messages

import "github.com/radekg/kafka-protocol-go/schema"

const (
	FieldNameDescribeTopicsTopicData        = "topics"
	FieldNameDescribeTopicsTopicName        = "topic"
	FieldNameDescribeTopicsPartitionIndexes = "partition_indexes"
)

func initDescribeProducersRequests() []schema.Schema {

	schemas := []schema.Schema{}

	schemas = append(schemas, schema.NewSchema("describe_producers_v0",
		&schema.CompactArray{Name: FieldNameDescribeTopicsTopicData, Ty: schema.NewSchema("describe_producers_topics_v0",
			&schema.Mfield{Name: FieldNameDescribeTopicsTopicName, Ty: schema.TypeCompactStr},
			&schema.Mfield{Name: FieldNameDescribeTopicsPartitionIndexes, Ty: schema.TypeInt32Array}, // TODO: is this a compact array?
			&schema.SchemaTaggedFields{Name: FieldNameTagBuffer},
		)},
		&schema.SchemaTaggedFields{Name: FieldNameTagBuffer}))

	return schemas

}
