package messages

import "github.com/radekg/kafka-protocol-go/schema"

const (
	FieldNameFetchCurrentLeaderEpoch  = "current_leader_epoch"
	FieldNameFetchForgottenTopicsData = "forgotten_topics_data"
	FieldNameFetchIsolationLevel      = "isolation_level"
	FieldNameFetchLogStartOffset      = "log_start_offset"
	FieldNameFetchMaxBytes            = "max_bytes"
	FieldNameFetchMaxWaitMs           = "max_wait_ms"
	FieldNameFetchMinBytes            = "min_bytes"
	FieldNameFetchOffset              = "fetch_offset"
	FieldNameFetchPartitionData       = "partition_data"
	FieldNameFetchPartition           = "partition"
	FieldNameFetchPartitionMaxBytes   = "partition_max_bytes"
	FieldNameFetchRackId              = "rack_id"
	FieldNameFetchReplicaId           = "replica_id"
	FieldNameFetchSessionId           = "session_id"
	FieldNameFetchSessionEpoch        = "session_epoch"
	FieldNameFetchTopicData           = "topic_data"
	FieldNameFetchTopicName           = "topic"
)

func initFetchRequests() []schema.Schema {

	schemas := []schema.Schema{}

	// TODO: verify: most likely Array instead of CompactArray
	fetchDataSchemaV0 := schema.NewSchema("fetch_v0",
		&schema.Mfield{Name: FieldNameFetchReplicaId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMaxWaitMs, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMinBytes, Ty: schema.TypeInt32},
		&schema.CompactArray{Name: FieldNameFetchTopicData, Ty: schema.NewSchema("fetch_topic_data_v0",
			&schema.Mfield{Name: FieldNameFetchTopicName, Ty: schema.TypeStr},
			&schema.CompactArray{Name: FieldNameFetchPartitionData, Ty: schema.NewSchema("fetch_partition_data_v0",
				&schema.Mfield{Name: FieldNameFetchPartition, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldNameFetchOffset, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldNameFetchPartitionMaxBytes, Ty: schema.TypeInt32},
			)},
		)},
	)

	schemas = append(schemas, fetchDataSchemaV0)
	schemas = append(schemas, fetchDataSchemaV0)
	schemas = append(schemas, fetchDataSchemaV0)

	// Version 3 adds max_bytes
	// TODO: verify: most likely Array instead of CompactArray
	fetchDataSchemaV3 := schema.NewSchema("fetch_v3",
		&schema.Mfield{Name: FieldNameFetchReplicaId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMaxWaitMs, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMinBytes, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMaxBytes, Ty: schema.TypeInt32},
		&schema.CompactArray{Name: FieldNameFetchTopicData, Ty: schema.NewSchema("fetch_topic_data_v3",
			&schema.Mfield{Name: FieldNameFetchTopicName, Ty: schema.TypeStr},
			&schema.CompactArray{Name: FieldNameFetchPartitionData, Ty: schema.NewSchema("fetch_partition_data_v3",
				&schema.Mfield{Name: FieldNameFetchPartition, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldNameFetchOffset, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldNameFetchPartitionMaxBytes, Ty: schema.TypeInt32},
			)},
		)},
	)

	schemas = append(schemas, fetchDataSchemaV3)

	// Version 4 adds isolation_level
	// TODO: verify: most likely Array instead of CompactArray
	fetchDataSchemaV4 := schema.NewSchema("fetch_v4",
		&schema.Mfield{Name: FieldNameFetchReplicaId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMaxWaitMs, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMinBytes, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMaxBytes, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchIsolationLevel, Ty: schema.TypeInt8},
		&schema.CompactArray{Name: FieldNameFetchTopicData, Ty: schema.NewSchema("fetch_topic_data_v4",
			&schema.Mfield{Name: FieldNameFetchTopicName, Ty: schema.TypeStr},
			&schema.CompactArray{Name: FieldNameFetchPartitionData, Ty: schema.NewSchema("fetch_partition_data_v4",
				&schema.Mfield{Name: FieldNameFetchPartition, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldNameFetchOffset, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldNameFetchPartitionMaxBytes, Ty: schema.TypeInt32},
			)},
		)},
	)

	schemas = append(schemas, fetchDataSchemaV4)

	// Version 5 adds topic.partition.log_start_offset
	// TODO: verify: most likely Array instead of CompactArray
	fetchDataSchemaV5 := schema.NewSchema("fetch_v5",
		&schema.Mfield{Name: FieldNameFetchReplicaId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMaxWaitMs, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMinBytes, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMaxBytes, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchIsolationLevel, Ty: schema.TypeInt8},
		&schema.CompactArray{Name: FieldNameFetchTopicData, Ty: schema.NewSchema("fetch_topic_data_v5",
			&schema.Mfield{Name: FieldNameFetchTopicName, Ty: schema.TypeStr},
			&schema.CompactArray{Name: FieldNameFetchPartitionData, Ty: schema.NewSchema("fetch_partition_data_v5",
				&schema.Mfield{Name: FieldNameFetchPartition, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldNameFetchOffset, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldNameFetchLogStartOffset, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldNameFetchPartitionMaxBytes, Ty: schema.TypeInt32},
			)},
		)},
	)

	schemas = append(schemas, fetchDataSchemaV5)
	schemas = append(schemas, fetchDataSchemaV5)

	// Version 7 adds:
	//  - session_id, session_epoch
	//  - forgotten_topics_data
	// TODO: verify: most likely Array instead of CompactArray
	fetchDataSchemaV7 := schema.NewSchema("fetch_v7",
		&schema.Mfield{Name: FieldNameFetchReplicaId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMaxWaitMs, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMinBytes, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMaxBytes, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchIsolationLevel, Ty: schema.TypeInt8},
		&schema.Mfield{Name: FieldNameFetchSessionId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchSessionEpoch, Ty: schema.TypeInt32},
		&schema.CompactArray{Name: FieldNameFetchTopicData, Ty: schema.NewSchema("fetch_topic_data_v7",
			&schema.Mfield{Name: FieldNameFetchTopicName, Ty: schema.TypeStr},
			&schema.CompactArray{Name: FieldNameFetchPartitionData, Ty: schema.NewSchema("fetch_partition_data_v7",
				&schema.Mfield{Name: FieldNameFetchPartition, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldNameFetchOffset, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldNameFetchLogStartOffset, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldNameFetchPartitionMaxBytes, Ty: schema.TypeInt32},
			)},
		)},
		&schema.CompactArray{Name: FieldNameFetchForgottenTopicsData, Ty: schema.NewSchema("fetch_forgotten_topics_data_v7",
			&schema.Mfield{Name: FieldNameFetchTopicName, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldNameFetchPartition, Ty: schema.TypeInt32Array},
		)},
	)

	schemas = append(schemas, fetchDataSchemaV7)
	schemas = append(schemas, fetchDataSchemaV7)

	// Version 9 adds topic.partition.current_leader_epoch
	// TODO: verify: most likely Array instead of CompactArray
	fetchDataSchemaV9 := schema.NewSchema("fetch_v9",
		&schema.Mfield{Name: FieldNameFetchReplicaId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMaxWaitMs, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMinBytes, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMaxBytes, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchIsolationLevel, Ty: schema.TypeInt8},
		&schema.Mfield{Name: FieldNameFetchSessionId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchSessionEpoch, Ty: schema.TypeInt32},
		&schema.CompactArray{Name: FieldNameFetchTopicData, Ty: schema.NewSchema("fetch_topic_data_v9",
			&schema.Mfield{Name: FieldNameFetchTopicName, Ty: schema.TypeStr},
			&schema.CompactArray{Name: FieldNameFetchPartitionData, Ty: schema.NewSchema("fetch_partition_data_v9",
				&schema.Mfield{Name: FieldNameFetchPartition, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldNameFetchCurrentLeaderEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldNameFetchOffset, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldNameFetchLogStartOffset, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldNameFetchPartitionMaxBytes, Ty: schema.TypeInt32},
			)},
		)},
		&schema.CompactArray{Name: FieldNameFetchForgottenTopicsData, Ty: schema.NewSchema("fetch_forgotten_topics_data_v9",
			&schema.Mfield{Name: FieldNameFetchTopicName, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldNameFetchPartition, Ty: schema.TypeInt32Array},
		)},
	)

	schemas = append(schemas, fetchDataSchemaV9)
	schemas = append(schemas, fetchDataSchemaV9)

	// Version 11 adds rack_id
	// TODO: verify: most likely Array instead of CompactArray
	fetchDataSchemaV11 := schema.NewSchema("fetch_v11",
		&schema.Mfield{Name: FieldNameFetchReplicaId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMaxWaitMs, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMinBytes, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMaxBytes, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchIsolationLevel, Ty: schema.TypeInt8},
		&schema.Mfield{Name: FieldNameFetchSessionId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchSessionEpoch, Ty: schema.TypeInt32},
		&schema.CompactArray{Name: FieldNameFetchTopicData, Ty: schema.NewSchema("fetch_topic_data_v11",
			&schema.Mfield{Name: FieldNameFetchTopicName, Ty: schema.TypeStr},
			&schema.CompactArray{Name: FieldNameFetchPartitionData, Ty: schema.NewSchema("fetch_partition_data_v11",
				&schema.Mfield{Name: FieldNameFetchPartition, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldNameFetchCurrentLeaderEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldNameFetchOffset, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldNameFetchLogStartOffset, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldNameFetchPartitionMaxBytes, Ty: schema.TypeInt32},
			)},
		)},
		&schema.CompactArray{Name: FieldNameFetchForgottenTopicsData, Ty: schema.NewSchema("fetch_forgotten_topics_data_v11",
			&schema.Mfield{Name: FieldNameFetchTopicName, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldNameFetchPartition, Ty: schema.TypeInt32Array},
		)},
		&schema.Mfield{Name: FieldNameFetchRackId, Ty: schema.TypeStr},
	)

	schemas = append(schemas, fetchDataSchemaV11)

	// Version 12 adds tag buffers and changes:
	//  - Str => CompactStr

	fetchDataSchemaV12 := schema.NewSchema("fetch_v12",
		&schema.Mfield{Name: FieldNameFetchReplicaId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMaxWaitMs, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMinBytes, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchMaxBytes, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchIsolationLevel, Ty: schema.TypeInt8},
		&schema.Mfield{Name: FieldNameFetchSessionId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameFetchSessionEpoch, Ty: schema.TypeInt32},
		&schema.CompactArray{Name: FieldNameFetchTopicData, Ty: schema.NewSchema("fetch_topic_data_v12",
			&schema.Mfield{Name: FieldNameFetchTopicName, Ty: schema.TypeCompactStr},
			&schema.CompactArray{Name: FieldNameFetchPartitionData, Ty: schema.NewSchema("fetch_partition_data_v12",
				&schema.Mfield{Name: FieldNameFetchPartition, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldNameFetchCurrentLeaderEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldNameFetchOffset, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldNameFetchLogStartOffset, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldNameFetchPartitionMaxBytes, Ty: schema.TypeInt32},
				&schema.SchemaTaggedFields{Name: FieldNameTagBuffer},
			)},
			&schema.SchemaTaggedFields{Name: FieldNameTagBuffer},
		)},
		&schema.CompactArray{Name: FieldNameFetchForgottenTopicsData, Ty: schema.NewSchema("fetch_forgotten_topics_data_v12",
			&schema.Mfield{Name: FieldNameFetchTopicName, Ty: schema.TypeCompactStr},
			&schema.Mfield{Name: FieldNameFetchPartition, Ty: schema.TypeInt32Array}, // TODO: is this a compact array?
			&schema.SchemaTaggedFields{Name: FieldNameTagBuffer},
		)},
		&schema.Mfield{Name: FieldNameFetchRackId, Ty: schema.TypeCompactStr},
		&schema.SchemaTaggedFields{Name: FieldNameTagBuffer},
	)

	schemas = append(schemas, fetchDataSchemaV12)

	return schemas

}
