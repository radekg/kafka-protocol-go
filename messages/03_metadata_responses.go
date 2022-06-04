package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init3MetadataResponse() []schema.Schema {

	return []schema.Schema{

		// Message: MetadataResponse, API Key: 3, Version: 0
		schema.NewSchema("MetadataResponsev0",
			&schema.Array{Name: FieldMetadataResponseBrokers, Ty: schema.NewSchema("BrokersV0",
				&schema.Mfield{Name: FieldMetadataResponseBrokersNodeId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersHost, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseBrokersPort, Ty: schema.TypeInt32},
			)},
			&schema.Array{Name: FieldMetadataResponseTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldMetadataResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldMetadataResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldMetadataResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV0",
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsReplicaNodes, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsIsrNodes, Ty: schema.TypeInt32Array},
				)},
			)},
		),

		// Message: MetadataResponse, API Key: 3, Version: 1
		schema.NewSchema("MetadataResponsev1",
			&schema.Array{Name: FieldMetadataResponseBrokers, Ty: schema.NewSchema("BrokersV1",
				&schema.Mfield{Name: FieldMetadataResponseBrokersNodeId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersHost, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseBrokersPort, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersRack, Ty: schema.TypeStrNullable},
			)},
			&schema.Mfield{Name: FieldMetadataResponseControllerId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldMetadataResponseTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldMetadataResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldMetadataResponseTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseTopicsIsInternal, Ty: schema.TypeBool},
				&schema.Array{Name: FieldMetadataResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV1",
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsReplicaNodes, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsIsrNodes, Ty: schema.TypeInt32Array},
				)},
			)},
		),

		// Message: MetadataResponse, API Key: 3, Version: 2
		schema.NewSchema("MetadataResponsev2",
			&schema.Array{Name: FieldMetadataResponseBrokers, Ty: schema.NewSchema("BrokersV2",
				&schema.Mfield{Name: FieldMetadataResponseBrokersNodeId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersHost, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseBrokersPort, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersRack, Ty: schema.TypeStrNullable},
			)},
			&schema.Mfield{Name: FieldMetadataResponseClusterId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldMetadataResponseControllerId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldMetadataResponseTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldMetadataResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldMetadataResponseTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseTopicsIsInternal, Ty: schema.TypeBool},
				&schema.Array{Name: FieldMetadataResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV2",
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsReplicaNodes, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsIsrNodes, Ty: schema.TypeInt32Array},
				)},
			)},
		),

		// Message: MetadataResponse, API Key: 3, Version: 3
		schema.NewSchema("MetadataResponsev3",
			&schema.Mfield{Name: FieldMetadataResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldMetadataResponseBrokers, Ty: schema.NewSchema("BrokersV3",
				&schema.Mfield{Name: FieldMetadataResponseBrokersNodeId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersHost, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseBrokersPort, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersRack, Ty: schema.TypeStrNullable},
			)},
			&schema.Mfield{Name: FieldMetadataResponseClusterId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldMetadataResponseControllerId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldMetadataResponseTopics, Ty: schema.NewSchema("TopicsV3",
				&schema.Mfield{Name: FieldMetadataResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldMetadataResponseTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseTopicsIsInternal, Ty: schema.TypeBool},
				&schema.Array{Name: FieldMetadataResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV3",
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsReplicaNodes, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsIsrNodes, Ty: schema.TypeInt32Array},
				)},
			)},
		),

		// Message: MetadataResponse, API Key: 3, Version: 4
		schema.NewSchema("MetadataResponsev4",
			&schema.Mfield{Name: FieldMetadataResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldMetadataResponseBrokers, Ty: schema.NewSchema("BrokersV4",
				&schema.Mfield{Name: FieldMetadataResponseBrokersNodeId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersHost, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseBrokersPort, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersRack, Ty: schema.TypeStrNullable},
			)},
			&schema.Mfield{Name: FieldMetadataResponseClusterId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldMetadataResponseControllerId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldMetadataResponseTopics, Ty: schema.NewSchema("TopicsV4",
				&schema.Mfield{Name: FieldMetadataResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldMetadataResponseTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseTopicsIsInternal, Ty: schema.TypeBool},
				&schema.Array{Name: FieldMetadataResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV4",
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsReplicaNodes, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsIsrNodes, Ty: schema.TypeInt32Array},
				)},
			)},
		),

		// Message: MetadataResponse, API Key: 3, Version: 5
		schema.NewSchema("MetadataResponsev5",
			&schema.Mfield{Name: FieldMetadataResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldMetadataResponseBrokers, Ty: schema.NewSchema("BrokersV5",
				&schema.Mfield{Name: FieldMetadataResponseBrokersNodeId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersHost, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseBrokersPort, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersRack, Ty: schema.TypeStrNullable},
			)},
			&schema.Mfield{Name: FieldMetadataResponseClusterId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldMetadataResponseControllerId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldMetadataResponseTopics, Ty: schema.NewSchema("TopicsV5",
				&schema.Mfield{Name: FieldMetadataResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldMetadataResponseTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseTopicsIsInternal, Ty: schema.TypeBool},
				&schema.Array{Name: FieldMetadataResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV5",
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsReplicaNodes, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsIsrNodes, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsOfflineReplicas, Ty: schema.TypeInt32Array},
				)},
			)},
		),

		// Message: MetadataResponse, API Key: 3, Version: 6
		schema.NewSchema("MetadataResponsev6",
			&schema.Mfield{Name: FieldMetadataResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldMetadataResponseBrokers, Ty: schema.NewSchema("BrokersV6",
				&schema.Mfield{Name: FieldMetadataResponseBrokersNodeId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersHost, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseBrokersPort, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersRack, Ty: schema.TypeStrNullable},
			)},
			&schema.Mfield{Name: FieldMetadataResponseClusterId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldMetadataResponseControllerId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldMetadataResponseTopics, Ty: schema.NewSchema("TopicsV6",
				&schema.Mfield{Name: FieldMetadataResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldMetadataResponseTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseTopicsIsInternal, Ty: schema.TypeBool},
				&schema.Array{Name: FieldMetadataResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV6",
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsReplicaNodes, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsIsrNodes, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsOfflineReplicas, Ty: schema.TypeInt32Array},
				)},
			)},
		),

		// Message: MetadataResponse, API Key: 3, Version: 7
		schema.NewSchema("MetadataResponsev7",
			&schema.Mfield{Name: FieldMetadataResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldMetadataResponseBrokers, Ty: schema.NewSchema("BrokersV7",
				&schema.Mfield{Name: FieldMetadataResponseBrokersNodeId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersHost, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseBrokersPort, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersRack, Ty: schema.TypeStrNullable},
			)},
			&schema.Mfield{Name: FieldMetadataResponseClusterId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldMetadataResponseControllerId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldMetadataResponseTopics, Ty: schema.NewSchema("TopicsV7",
				&schema.Mfield{Name: FieldMetadataResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldMetadataResponseTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseTopicsIsInternal, Ty: schema.TypeBool},
				&schema.Array{Name: FieldMetadataResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV7",
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsReplicaNodes, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsIsrNodes, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsOfflineReplicas, Ty: schema.TypeInt32Array},
				)},
			)},
		),

		// Message: MetadataResponse, API Key: 3, Version: 8
		schema.NewSchema("MetadataResponsev8",
			&schema.Mfield{Name: FieldMetadataResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldMetadataResponseBrokers, Ty: schema.NewSchema("BrokersV8",
				&schema.Mfield{Name: FieldMetadataResponseBrokersNodeId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersHost, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseBrokersPort, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersRack, Ty: schema.TypeStrNullable},
			)},
			&schema.Mfield{Name: FieldMetadataResponseClusterId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldMetadataResponseControllerId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldMetadataResponseTopics, Ty: schema.NewSchema("TopicsV8",
				&schema.Mfield{Name: FieldMetadataResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldMetadataResponseTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldMetadataResponseTopicsIsInternal, Ty: schema.TypeBool},
				&schema.Array{Name: FieldMetadataResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV8",
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsReplicaNodes, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsIsrNodes, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsOfflineReplicas, Ty: schema.TypeInt32Array},
				)},
				&schema.Mfield{Name: FieldMetadataResponseTopicsTopicAuthorizedOperations, Ty: schema.TypeInt32},
			)},
			&schema.Mfield{Name: FieldMetadataResponseClusterAuthorizedOperations, Ty: schema.TypeInt32},
		),

		// Message: MetadataResponse, API Key: 3, Version: 9
		schema.NewSchema("MetadataResponsev9",
			&schema.Mfield{Name: FieldMetadataResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldMetadataResponseBrokers, Ty: schema.NewSchema("BrokersV9",
				&schema.Mfield{Name: FieldMetadataResponseBrokersNodeId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersHost, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldMetadataResponseBrokersPort, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersRack, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldMetadataResponseBrokersTags},
			)},
			&schema.Mfield{Name: FieldMetadataResponseClusterId, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldMetadataResponseControllerId, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldMetadataResponseTopics, Ty: schema.NewSchema("TopicsV9",
				&schema.Mfield{Name: FieldMetadataResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldMetadataResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldMetadataResponseTopicsIsInternal, Ty: schema.TypeBool},
				&schema.ArrayCompact{Name: FieldMetadataResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV9",
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsReplicaNodes, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsIsrNodes, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsOfflineReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldMetadataResponseTopicsPartitionsTags},
				)},
				&schema.Mfield{Name: FieldMetadataResponseTopicsTopicAuthorizedOperations, Ty: schema.TypeInt32},
				&schema.SchemaTaggedFields{Name: FieldMetadataResponseTopicsTags},
			)},
			&schema.Mfield{Name: FieldMetadataResponseClusterAuthorizedOperations, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldMetadataResponseTags},
		),

		// Message: MetadataResponse, API Key: 3, Version: 10
		schema.NewSchema("MetadataResponsev10",
			&schema.Mfield{Name: FieldMetadataResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldMetadataResponseBrokers, Ty: schema.NewSchema("BrokersV10",
				&schema.Mfield{Name: FieldMetadataResponseBrokersNodeId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersHost, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldMetadataResponseBrokersPort, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersRack, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldMetadataResponseBrokersTags},
			)},
			&schema.Mfield{Name: FieldMetadataResponseClusterId, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldMetadataResponseControllerId, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldMetadataResponseTopics, Ty: schema.NewSchema("TopicsV10",
				&schema.Mfield{Name: FieldMetadataResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldMetadataResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldMetadataResponseTopicsTopicId, Ty: schema.TypeUuid},
				&schema.Mfield{Name: FieldMetadataResponseTopicsIsInternal, Ty: schema.TypeBool},
				&schema.ArrayCompact{Name: FieldMetadataResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV10",
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsReplicaNodes, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsIsrNodes, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsOfflineReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldMetadataResponseTopicsPartitionsTags},
				)},
				&schema.Mfield{Name: FieldMetadataResponseTopicsTopicAuthorizedOperations, Ty: schema.TypeInt32},
				&schema.SchemaTaggedFields{Name: FieldMetadataResponseTopicsTags},
			)},
			&schema.Mfield{Name: FieldMetadataResponseClusterAuthorizedOperations, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldMetadataResponseTags},
		),

		// Message: MetadataResponse, API Key: 3, Version: 11
		schema.NewSchema("MetadataResponsev11",
			&schema.Mfield{Name: FieldMetadataResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldMetadataResponseBrokers, Ty: schema.NewSchema("BrokersV11",
				&schema.Mfield{Name: FieldMetadataResponseBrokersNodeId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersHost, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldMetadataResponseBrokersPort, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersRack, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldMetadataResponseBrokersTags},
			)},
			&schema.Mfield{Name: FieldMetadataResponseClusterId, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldMetadataResponseControllerId, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldMetadataResponseTopics, Ty: schema.NewSchema("TopicsV11",
				&schema.Mfield{Name: FieldMetadataResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldMetadataResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldMetadataResponseTopicsTopicId, Ty: schema.TypeUuid},
				&schema.Mfield{Name: FieldMetadataResponseTopicsIsInternal, Ty: schema.TypeBool},
				&schema.ArrayCompact{Name: FieldMetadataResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV11",
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsReplicaNodes, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsIsrNodes, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsOfflineReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldMetadataResponseTopicsPartitionsTags},
				)},
				&schema.Mfield{Name: FieldMetadataResponseTopicsTopicAuthorizedOperations, Ty: schema.TypeInt32},
				&schema.SchemaTaggedFields{Name: FieldMetadataResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldMetadataResponseTags},
		),

		// Message: MetadataResponse, API Key: 3, Version: 12
		schema.NewSchema("MetadataResponsev12",
			&schema.Mfield{Name: FieldMetadataResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldMetadataResponseBrokers, Ty: schema.NewSchema("BrokersV12",
				&schema.Mfield{Name: FieldMetadataResponseBrokersNodeId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersHost, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldMetadataResponseBrokersPort, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldMetadataResponseBrokersRack, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldMetadataResponseBrokersTags},
			)},
			&schema.Mfield{Name: FieldMetadataResponseClusterId, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldMetadataResponseControllerId, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldMetadataResponseTopics, Ty: schema.NewSchema("TopicsV12",
				&schema.Mfield{Name: FieldMetadataResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldMetadataResponseTopicsName, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldMetadataResponseTopicsTopicId, Ty: schema.TypeUuid},
				&schema.Mfield{Name: FieldMetadataResponseTopicsIsInternal, Ty: schema.TypeBool},
				&schema.ArrayCompact{Name: FieldMetadataResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV12",
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsReplicaNodes, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsIsrNodes, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldMetadataResponseTopicsPartitionsOfflineReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldMetadataResponseTopicsPartitionsTags},
				)},
				&schema.Mfield{Name: FieldMetadataResponseTopicsTopicAuthorizedOperations, Ty: schema.TypeInt32},
				&schema.SchemaTaggedFields{Name: FieldMetadataResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldMetadataResponseTags},
		),
	}
}

const (
	// FieldMetadataResponseTopicsPartitionsOfflineReplicas is: The set of offline replicas of this partition.
	FieldMetadataResponseTopicsPartitionsOfflineReplicas = "OfflineReplicas"
	// FieldMetadataResponseTopicsPartitionsLeaderEpoch is: The leader epoch of this partition.
	FieldMetadataResponseTopicsPartitionsLeaderEpoch = "LeaderEpoch"
	// FieldMetadataResponseClusterAuthorizedOperations is: 32-bit bitfield to represent authorized operations for this cluster.
	FieldMetadataResponseClusterAuthorizedOperations = "ClusterAuthorizedOperations"
	// FieldMetadataResponseBrokersNodeId is: The broker ID.
	FieldMetadataResponseBrokersNodeId = "NodeId"
	// FieldMetadataResponseTopics is: Each topic in the response.
	FieldMetadataResponseTopics = "Topics"
	// FieldMetadataResponseTopicsPartitionsErrorCode is: The partition error, or 0 if there was no error.
	FieldMetadataResponseTopicsPartitionsErrorCode = "ErrorCode"
	// FieldMetadataResponseTopicsPartitionsIsrNodes is: The set of nodes that are in sync with the leader for this partition.
	FieldMetadataResponseTopicsPartitionsIsrNodes = "IsrNodes"
	// FieldMetadataResponseTopicsPartitionsTags is: The tagged fields.
	FieldMetadataResponseTopicsPartitionsTags = "Tags"
	// FieldMetadataResponseTopicsErrorCode is: The topic error, or 0 if there was no error.
	FieldMetadataResponseTopicsErrorCode = "ErrorCode"
	// FieldMetadataResponseTopicsName is: The topic name.
	FieldMetadataResponseTopicsName = "Name"
	// FieldMetadataResponseTopicsIsInternal is: True if the topic is internal.
	FieldMetadataResponseTopicsIsInternal = "IsInternal"
	// FieldMetadataResponseClusterId is: The cluster ID that responding broker belongs to.
	FieldMetadataResponseClusterId = "ClusterId"
	// FieldMetadataResponseTopicsTopicAuthorizedOperations is: 32-bit bitfield to represent authorized operations for this topic.
	FieldMetadataResponseTopicsTopicAuthorizedOperations = "TopicAuthorizedOperations"
	// FieldMetadataResponseTags is: The tagged fields.
	FieldMetadataResponseTags = "Tags"
	// FieldMetadataResponseBrokersPort is: The broker port.
	FieldMetadataResponseBrokersPort = "Port"
	// FieldMetadataResponseTopicsPartitionsLeaderId is: The ID of the leader broker.
	FieldMetadataResponseTopicsPartitionsLeaderId = "LeaderId"
	// FieldMetadataResponseBrokersRack is: The rack of the broker, or null if it has not been assigned to a rack.
	FieldMetadataResponseBrokersRack = "Rack"
	// FieldMetadataResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldMetadataResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldMetadataResponseBrokersHost is: The broker hostname.
	FieldMetadataResponseBrokersHost = "Host"
	// FieldMetadataResponseControllerId is: The ID of the controller broker.
	FieldMetadataResponseControllerId = "ControllerId"
	// FieldMetadataResponseTopicsTopicId is: The topic id.
	FieldMetadataResponseTopicsTopicId = "TopicId"
	// FieldMetadataResponseTopicsTags is: The tagged fields.
	FieldMetadataResponseTopicsTags = "Tags"
	// FieldMetadataResponseBrokers is: Each broker in the response.
	FieldMetadataResponseBrokers = "Brokers"
	// FieldMetadataResponseTopicsPartitionsPartitionIndex is: The partition index.
	FieldMetadataResponseTopicsPartitionsPartitionIndex = "PartitionIndex"
	// FieldMetadataResponseTopicsPartitions is: Each partition in the topic.
	FieldMetadataResponseTopicsPartitions = "Partitions"
	// FieldMetadataResponseTopicsPartitionsReplicaNodes is: The set of all nodes that host this partition.
	FieldMetadataResponseTopicsPartitionsReplicaNodes = "ReplicaNodes"
	// FieldMetadataResponseBrokersTags is: The tagged fields.
	FieldMetadataResponseBrokersTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/MetadataResponse.json
const originalMetadataResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 3,
  "type": "response",
  "name": "MetadataResponse",
  // Version 1 adds fields for the rack of each broker, the controller id, and
  // whether or not the topic is internal.
  //
  // Version 2 adds the cluster ID field.
  //
  // Version 3 adds the throttle time.
  //
  // Version 4 is the same as version 3.
  //
  // Version 5 adds a per-partition offline_replicas field. This field specifies
  // the list of replicas that are offline.
  //
  // Starting in version 6, on quota violation, brokers send out responses before throttling.
  //
  // Version 7 adds the leader epoch to the partition metadata.
  //
  // Starting in version 8, brokers can send authorized operations for topic and cluster.
  //
  // Version 9 is the first flexible version.
  //
  // Version 10 adds topicId.
  //
  // Version 11 deprecates ClusterAuthorizedOperations. This is now exposed
  // by the DescribeCluster API (KIP-700).
  // Version 12 supports topicId.
  "validVersions": "0-12",
  "flexibleVersions": "9+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "3+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Brokers", "type": "[]MetadataResponseBroker", "versions": "0+",
      "about": "Each broker in the response.", "fields": [
      { "name": "NodeId", "type": "int32", "versions": "0+", "mapKey": true, "entityType": "brokerId",
        "about": "The broker ID." },
      { "name": "Host", "type": "string", "versions": "0+",
        "about": "The broker hostname." },
      { "name": "Port", "type": "int32", "versions": "0+",
        "about": "The broker port." },
      { "name": "Rack", "type": "string", "versions": "1+", "nullableVersions": "1+", "ignorable": true, "default": "null",
        "about": "The rack of the broker, or null if it has not been assigned to a rack." }
    ]},
    { "name": "ClusterId", "type": "string", "nullableVersions": "2+", "versions": "2+", "ignorable": true, "default": "null",
      "about": "The cluster ID that responding broker belongs to." },
    { "name": "ControllerId", "type": "int32", "versions": "1+", "default": "-1", "ignorable": true, "entityType": "brokerId",
      "about": "The ID of the controller broker." },
    { "name": "Topics", "type": "[]MetadataResponseTopic", "versions": "0+",
      "about": "Each topic in the response.", "fields": [
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The topic error, or 0 if there was no error." },
      { "name": "Name", "type": "string", "versions": "0+", "mapKey": true, "entityType": "topicName", "nullableVersions": "12+",
        "about": "The topic name." },
      { "name": "TopicId", "type": "uuid", "versions": "10+", "ignorable": true, "about": "The topic id." },
      { "name": "IsInternal", "type": "bool", "versions": "1+", "default": "false", "ignorable": true,
        "about": "True if the topic is internal." },
      { "name": "Partitions", "type": "[]MetadataResponsePartition", "versions": "0+",
        "about": "Each partition in the topic.", "fields": [
        { "name": "ErrorCode", "type": "int16", "versions": "0+",
          "about": "The partition error, or 0 if there was no error." },
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "LeaderId", "type": "int32", "versions": "0+", "entityType": "brokerId",
          "about": "The ID of the leader broker." },
        { "name": "LeaderEpoch", "type": "int32", "versions": "7+", "default": "-1", "ignorable": true,
          "about": "The leader epoch of this partition." },
        { "name": "ReplicaNodes", "type": "[]int32", "versions": "0+", "entityType": "brokerId",
          "about": "The set of all nodes that host this partition." },
        { "name": "IsrNodes", "type": "[]int32", "versions": "0+", "entityType": "brokerId",
          "about": "The set of nodes that are in sync with the leader for this partition." },
        { "name": "OfflineReplicas", "type": "[]int32", "versions": "5+", "ignorable": true, "entityType": "brokerId",
          "about": "The set of offline replicas of this partition." }
      ]},
      { "name": "TopicAuthorizedOperations", "type": "int32", "versions": "8+", "default": "-2147483648",
        "about": "32-bit bitfield to represent authorized operations for this topic." }
    ]},
    { "name": "ClusterAuthorizedOperations", "type": "int32", "versions": "8-10", "default": "-2147483648",
      "about": "32-bit bitfield to represent authorized operations for this cluster." }
  ]
}
`
