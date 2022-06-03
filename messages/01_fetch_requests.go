package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init1FetchRequest() []schema.Schema {

	return []schema.Schema{

		// Message: FetchRequest, API Key: 1, Version: 0
		schema.NewSchema("FetchRequestv0", 
			&schema.Mfield{Name: FieldFetchRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxWaitMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMinBytes, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchRequestTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldFetchRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV0",
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsFetchOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartitionMaxBytes, Ty: schema.TypeInt32},
				)},
			)},
		),

		// Message: FetchRequest, API Key: 1, Version: 1
		schema.NewSchema("FetchRequestv1", 
			&schema.Mfield{Name: FieldFetchRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxWaitMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMinBytes, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchRequestTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldFetchRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV1",
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsFetchOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartitionMaxBytes, Ty: schema.TypeInt32},
				)},
			)},
		),

		// Message: FetchRequest, API Key: 1, Version: 2
		schema.NewSchema("FetchRequestv2", 
			&schema.Mfield{Name: FieldFetchRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxWaitMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMinBytes, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchRequestTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldFetchRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV2",
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsFetchOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartitionMaxBytes, Ty: schema.TypeInt32},
				)},
			)},
		),

		// Message: FetchRequest, API Key: 1, Version: 3
		schema.NewSchema("FetchRequestv3", 
			&schema.Mfield{Name: FieldFetchRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxWaitMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMinBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxBytes, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchRequestTopics, Ty: schema.NewSchema("TopicsV3",
				&schema.Mfield{Name: FieldFetchRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV3",
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsFetchOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartitionMaxBytes, Ty: schema.TypeInt32},
				)},
			)},
		),

		// Message: FetchRequest, API Key: 1, Version: 4
		schema.NewSchema("FetchRequestv4", 
			&schema.Mfield{Name: FieldFetchRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxWaitMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMinBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.Array{Name: FieldFetchRequestTopics, Ty: schema.NewSchema("TopicsV4",
				&schema.Mfield{Name: FieldFetchRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV4",
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsFetchOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartitionMaxBytes, Ty: schema.TypeInt32},
				)},
			)},
		),

		// Message: FetchRequest, API Key: 1, Version: 5
		schema.NewSchema("FetchRequestv5", 
			&schema.Mfield{Name: FieldFetchRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxWaitMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMinBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.Array{Name: FieldFetchRequestTopics, Ty: schema.NewSchema("TopicsV5",
				&schema.Mfield{Name: FieldFetchRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV5",
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsFetchOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartitionMaxBytes, Ty: schema.TypeInt32},
				)},
			)},
		),

		// Message: FetchRequest, API Key: 1, Version: 6
		schema.NewSchema("FetchRequestv6", 
			&schema.Mfield{Name: FieldFetchRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxWaitMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMinBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.Array{Name: FieldFetchRequestTopics, Ty: schema.NewSchema("TopicsV6",
				&schema.Mfield{Name: FieldFetchRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV6",
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsFetchOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartitionMaxBytes, Ty: schema.TypeInt32},
				)},
			)},
		),

		// Message: FetchRequest, API Key: 1, Version: 7
		schema.NewSchema("FetchRequestv7", 
			&schema.Mfield{Name: FieldFetchRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxWaitMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMinBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldFetchRequestSessionId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestSessionEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchRequestTopics, Ty: schema.NewSchema("TopicsV7",
				&schema.Mfield{Name: FieldFetchRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV7",
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsFetchOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartitionMaxBytes, Ty: schema.TypeInt32},
				)},
			)},
			&schema.Array{Name: FieldFetchRequestForgottenTopicsData, Ty: schema.NewSchema("ForgottenTopicsDataV7",
				&schema.Mfield{Name: FieldFetchRequestForgottenTopicsDataTopic, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldFetchRequestForgottenTopicsDataPartitions, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: FetchRequest, API Key: 1, Version: 8
		schema.NewSchema("FetchRequestv8", 
			&schema.Mfield{Name: FieldFetchRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxWaitMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMinBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldFetchRequestSessionId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestSessionEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchRequestTopics, Ty: schema.NewSchema("TopicsV8",
				&schema.Mfield{Name: FieldFetchRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV8",
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsFetchOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartitionMaxBytes, Ty: schema.TypeInt32},
				)},
			)},
			&schema.Array{Name: FieldFetchRequestForgottenTopicsData, Ty: schema.NewSchema("ForgottenTopicsDataV8",
				&schema.Mfield{Name: FieldFetchRequestForgottenTopicsDataTopic, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldFetchRequestForgottenTopicsDataPartitions, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: FetchRequest, API Key: 1, Version: 9
		schema.NewSchema("FetchRequestv9", 
			&schema.Mfield{Name: FieldFetchRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxWaitMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMinBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldFetchRequestSessionId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestSessionEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchRequestTopics, Ty: schema.NewSchema("TopicsV9",
				&schema.Mfield{Name: FieldFetchRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV9",
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsCurrentLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsFetchOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartitionMaxBytes, Ty: schema.TypeInt32},
				)},
			)},
			&schema.Array{Name: FieldFetchRequestForgottenTopicsData, Ty: schema.NewSchema("ForgottenTopicsDataV9",
				&schema.Mfield{Name: FieldFetchRequestForgottenTopicsDataTopic, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldFetchRequestForgottenTopicsDataPartitions, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: FetchRequest, API Key: 1, Version: 10
		schema.NewSchema("FetchRequestv10", 
			&schema.Mfield{Name: FieldFetchRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxWaitMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMinBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldFetchRequestSessionId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestSessionEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchRequestTopics, Ty: schema.NewSchema("TopicsV10",
				&schema.Mfield{Name: FieldFetchRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV10",
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsCurrentLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsFetchOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartitionMaxBytes, Ty: schema.TypeInt32},
				)},
			)},
			&schema.Array{Name: FieldFetchRequestForgottenTopicsData, Ty: schema.NewSchema("ForgottenTopicsDataV10",
				&schema.Mfield{Name: FieldFetchRequestForgottenTopicsDataTopic, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldFetchRequestForgottenTopicsDataPartitions, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: FetchRequest, API Key: 1, Version: 11
		schema.NewSchema("FetchRequestv11", 
			&schema.Mfield{Name: FieldFetchRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxWaitMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMinBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldFetchRequestSessionId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestSessionEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchRequestTopics, Ty: schema.NewSchema("TopicsV11",
				&schema.Mfield{Name: FieldFetchRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV11",
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsCurrentLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsFetchOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartitionMaxBytes, Ty: schema.TypeInt32},
				)},
			)},
			&schema.Array{Name: FieldFetchRequestForgottenTopicsData, Ty: schema.NewSchema("ForgottenTopicsDataV11",
				&schema.Mfield{Name: FieldFetchRequestForgottenTopicsDataTopic, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldFetchRequestForgottenTopicsDataPartitions, Ty: schema.TypeInt32Array},
			)},
			&schema.Mfield{Name: FieldFetchRequestRackId, Ty: schema.TypeStr},
		),

		// Message: FetchRequest, API Key: 1, Version: 12
		schema.NewSchema("FetchRequestv12", 
			&schema.Mfield{Name: FieldFetchRequestClusterId, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldFetchRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxWaitMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMinBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldFetchRequestSessionId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestSessionEpoch, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldFetchRequestTopics, Ty: schema.NewSchema("TopicsV12",
				&schema.Mfield{Name: FieldFetchRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.ArrayCompact{Name: FieldFetchRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV12",
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsCurrentLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsFetchOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsLastFetchedEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartitionMaxBytes, Ty: schema.TypeInt32},
					&schema.SchemaTaggedFields{Name: FieldFetchRequestTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldFetchRequestTopicsTags},
			)},
			&schema.ArrayCompact{Name: FieldFetchRequestForgottenTopicsData, Ty: schema.NewSchema("ForgottenTopicsDataV12",
				&schema.Mfield{Name: FieldFetchRequestForgottenTopicsDataTopic, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldFetchRequestForgottenTopicsDataPartitions, Ty: schema.TypeInt32CompactArray},
				&schema.SchemaTaggedFields{Name: FieldFetchRequestForgottenTopicsDataTags},
			)},
			&schema.Mfield{Name: FieldFetchRequestRackId, Ty: schema.TypeStr},
			&schema.SchemaTaggedFields{Name: FieldFetchRequestTags},
		),

		// Message: FetchRequest, API Key: 1, Version: 13
		schema.NewSchema("FetchRequestv13", 
			&schema.Mfield{Name: FieldFetchRequestClusterId, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldFetchRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxWaitMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMinBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestMaxBytes, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldFetchRequestSessionId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchRequestSessionEpoch, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldFetchRequestTopics, Ty: schema.NewSchema("TopicsV13",
				&schema.Mfield{Name: FieldFetchRequestTopicsTopicId, Ty: schema.TypeUuid},
				&schema.ArrayCompact{Name: FieldFetchRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV13",
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsCurrentLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsFetchOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsLastFetchedEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchRequestTopicsPartitionsPartitionMaxBytes, Ty: schema.TypeInt32},
					&schema.SchemaTaggedFields{Name: FieldFetchRequestTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldFetchRequestTopicsTags},
			)},
			&schema.ArrayCompact{Name: FieldFetchRequestForgottenTopicsData, Ty: schema.NewSchema("ForgottenTopicsDataV13",
				&schema.Mfield{Name: FieldFetchRequestForgottenTopicsDataTopicId, Ty: schema.TypeUuid},
				&schema.Mfield{Name: FieldFetchRequestForgottenTopicsDataPartitions, Ty: schema.TypeInt32CompactArray},
				&schema.SchemaTaggedFields{Name: FieldFetchRequestForgottenTopicsDataTags},
			)},
			&schema.Mfield{Name: FieldFetchRequestRackId, Ty: schema.TypeStr},
			&schema.SchemaTaggedFields{Name: FieldFetchRequestTags},
		),

	}
}

const (
	// FieldFetchRequestTopicsTopic is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestTopicsTopic = "Topic"
	// FieldFetchRequestMaxWaitMs is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestMaxWaitMs = "MaxWaitMs"
	// FieldFetchRequestTopicsPartitionsCurrentLeaderEpoch is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestTopicsPartitionsCurrentLeaderEpoch = "CurrentLeaderEpoch"
	// FieldFetchRequestTopicsTags is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestTopicsTags = "Tags"
	// FieldFetchRequestForgottenTopicsDataTopicId is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestForgottenTopicsDataTopicId = "TopicId"
	// FieldFetchRequestReplicaId is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestReplicaId = "ReplicaId"
	// FieldFetchRequestTopicsPartitionsPartition is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestTopicsPartitionsPartition = "Partition"
	// FieldFetchRequestTopicsPartitionsPartitionMaxBytes is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestTopicsPartitionsPartitionMaxBytes = "PartitionMaxBytes"
	// FieldFetchRequestForgottenTopicsDataPartitions is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestForgottenTopicsDataPartitions = "Partitions"
	// FieldFetchRequestRackId is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestRackId = "RackId"
	// FieldFetchRequestTopics is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestTopics = "Topics"
	// FieldFetchRequestForgottenTopicsDataTopic is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestForgottenTopicsDataTopic = "Topic"
	// FieldFetchRequestClusterId is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestClusterId = "ClusterId"
	// FieldFetchRequestTopicsPartitionsTags is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestTopicsPartitionsTags = "Tags"
	// FieldFetchRequestForgottenTopicsDataTags is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestForgottenTopicsDataTags = "Tags"
	// FieldFetchRequestTopicsTopicId is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestTopicsTopicId = "TopicId"
	// FieldFetchRequestMinBytes is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestMinBytes = "MinBytes"
	// FieldFetchRequestTopicsPartitionsLogStartOffset is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestTopicsPartitionsLogStartOffset = "LogStartOffset"
	// FieldFetchRequestTopicsPartitions is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestTopicsPartitions = "Partitions"
	// FieldFetchRequestSessionEpoch is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestSessionEpoch = "SessionEpoch"
	// FieldFetchRequestTopicsPartitionsLastFetchedEpoch is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestTopicsPartitionsLastFetchedEpoch = "LastFetchedEpoch"
	// FieldFetchRequestMaxBytes is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestMaxBytes = "MaxBytes"
	// FieldFetchRequestSessionId is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestSessionId = "SessionId"
	// FieldFetchRequestIsolationLevel is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestIsolationLevel = "IsolationLevel"
	// FieldFetchRequestForgottenTopicsData is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestForgottenTopicsData = "ForgottenTopicsData"
	// FieldFetchRequestTags is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestTags = "Tags"
	// FieldFetchRequestTopicsPartitionsFetchOffset is a field name that can be used to resolve the correct struct field.
	FieldFetchRequestTopicsPartitionsFetchOffset = "FetchOffset"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/FetchRequest.json
const originalFetchRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

{
  "apiKey": 1,
  "type": "request",
  "listeners": ["zkBroker", "broker", "controller"],
  "name": "FetchRequest",
  //
  // Version 1 is the same as version 0.
  //
  // Starting in Version 2, the requestor must be able to handle Kafka Log
  // Message format version 1.
  //
  // Version 3 adds MaxBytes.  Starting in version 3, the partition ordering in
  // the request is now relevant.  Partitions will be processed in the order
  // they appear in the request.
  //
  // Version 4 adds IsolationLevel.  Starting in version 4, the reqestor must be
  // able to handle Kafka log message format version 2.
  //
  // Version 5 adds LogStartOffset to indicate the earliest available offset of
  // partition data that can be consumed.
  //
  // Version 6 is the same as version 5.
  //
  // Version 7 adds incremental fetch request support.
  //
  // Version 8 is the same as version 7.
  //
  // Version 9 adds CurrentLeaderEpoch, as described in KIP-320.
  //
  // Version 10 indicates that we can use the ZStd compression algorithm, as
  // described in KIP-110.
  // Version 12 adds flexible versions support as well as epoch validation through
  // the 'LastFetchedEpoch' field
  //
  // Version 13 replaces topic names with topic IDs (KIP-516). May return UNKNOWN_TOPIC_ID error code.
  "validVersions": "0-13",
  "flexibleVersions": "12+",
  "fields": [
    { "name": "ClusterId", "type": "string", "versions": "12+", "nullableVersions": "12+", "default": "null",
      "taggedVersions": "12+", "tag": 0, "ignorable": true,
      "about": "The clusterId if known. This is used to validate metadata fetches prior to broker registration." },
    { "name": "ReplicaId", "type": "int32", "versions": "0+", "entityType": "brokerId",
      "about": "The broker ID of the follower, of -1 if this request is from a consumer." },
    { "name": "MaxWaitMs", "type": "int32", "versions": "0+",
      "about": "The maximum time in milliseconds to wait for the response." },
    { "name": "MinBytes", "type": "int32", "versions": "0+",
      "about": "The minimum bytes to accumulate in the response." },
    { "name": "MaxBytes", "type": "int32", "versions": "3+", "default": "0x7fffffff", "ignorable": true,
      "about": "The maximum bytes to fetch.  See KIP-74 for cases where this limit may not be honored." },
    { "name": "IsolationLevel", "type": "int8", "versions": "4+", "default": "0", "ignorable": true,
      "about": "This setting controls the visibility of transactional records. Using READ_UNCOMMITTED (isolation_level = 0) makes all records visible. With READ_COMMITTED (isolation_level = 1), non-transactional and COMMITTED transactional records are visible. To be more concrete, READ_COMMITTED returns all data from offsets smaller than the current LSO (last stable offset), and enables the inclusion of the list of aborted transactions in the result, which allows consumers to discard ABORTED transactional records" },
    { "name": "SessionId", "type": "int32", "versions": "7+", "default": "0", "ignorable": true,
      "about": "The fetch session ID." },
    { "name": "SessionEpoch", "type": "int32", "versions": "7+", "default": "-1", "ignorable": true,
      "about": "The fetch session epoch, which is used for ordering requests in a session." },
    { "name": "Topics", "type": "[]FetchTopic", "versions": "0+",
      "about": "The topics to fetch.", "fields": [
      { "name": "Topic", "type": "string", "versions": "0-12", "entityType": "topicName", "ignorable": true,
        "about": "The name of the topic to fetch." },
      { "name": "TopicId", "type": "uuid", "versions": "13+", "ignorable": true, "about": "The unique topic ID"},
      { "name": "Partitions", "type": "[]FetchPartition", "versions": "0+",
        "about": "The partitions to fetch.", "fields": [
        { "name": "Partition", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "CurrentLeaderEpoch", "type": "int32", "versions": "9+", "default": "-1", "ignorable": true,
          "about": "The current leader epoch of the partition." },
        { "name": "FetchOffset", "type": "int64", "versions": "0+",
          "about": "The message offset." },
        { "name": "LastFetchedEpoch", "type": "int32", "versions": "12+", "default": "-1", "ignorable": false,
          "about": "The epoch of the last fetched record or -1 if there is none"},
        { "name": "LogStartOffset", "type": "int64", "versions": "5+", "default": "-1", "ignorable": true,
          "about": "The earliest available offset of the follower replica.  The field is only used when the request is sent by the follower."},
        { "name": "PartitionMaxBytes", "type": "int32", "versions": "0+",
          "about": "The maximum bytes to fetch from this partition.  See KIP-74 for cases where this limit may not be honored." }
      ]}
    ]},
    { "name": "ForgottenTopicsData", "type": "[]ForgottenTopic", "versions": "7+", "ignorable": false,
      "about": "In an incremental fetch request, the partitions to remove.", "fields": [
      { "name": "Topic", "type": "string", "versions": "7-12", "entityType": "topicName", "ignorable": true,
        "about": "The partition name." },
      { "name": "TopicId", "type": "uuid", "versions": "13+", "ignorable": true, "about": "The unique topic ID"},
      { "name": "Partitions", "type": "[]int32", "versions": "7+",
        "about": "The partitions indexes to forget." }
    ]},
    { "name": "RackId", "type":  "string", "versions": "11+", "default": "", "ignorable": true,
      "about": "Rack ID of the consumer making this request"}
  ]
}
`

