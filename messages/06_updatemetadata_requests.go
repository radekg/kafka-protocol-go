package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init6UpdateMetadataRequest() []schema.Schema {

	return []schema.Schema{
		// Message: UpdateMetadataRequest, API Key: 6, Version: 0
		schema.NewSchema("UpdateMetadataRequest:v0",
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldUpdateMetadataRequestUngroupedPartitionStates, Ty: schema.NewSchema("[]UpdateMetadataPartitionState:v0",
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesTopicName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesControllerEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesLeader, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesIsr, Ty: schema.TypeInt32Array},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesZkVersion, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesReplicas, Ty: schema.TypeInt32Array},
			)},
			&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("[]UpdateMetadataBroker:v0",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
			)},
		),

		// Message: UpdateMetadataRequest, API Key: 6, Version: 1
		schema.NewSchema("UpdateMetadataRequest:v1",
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldUpdateMetadataRequestUngroupedPartitionStates, Ty: schema.NewSchema("[]UpdateMetadataPartitionState:v1",
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesTopicName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesControllerEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesLeader, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesIsr, Ty: schema.TypeInt32Array},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesZkVersion, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesReplicas, Ty: schema.TypeInt32Array},
			)},
			&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("[]UpdateMetadataBroker:v1",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
				&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokersEndpoints, Ty: schema.NewSchema("[]UpdateMetadataEndpoint:v1",
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsPort, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: UpdateMetadataRequest, API Key: 6, Version: 2
		schema.NewSchema("UpdateMetadataRequest:v2",
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldUpdateMetadataRequestUngroupedPartitionStates, Ty: schema.NewSchema("[]UpdateMetadataPartitionState:v2",
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesTopicName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesControllerEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesLeader, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesIsr, Ty: schema.TypeInt32Array},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesZkVersion, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesReplicas, Ty: schema.TypeInt32Array},
			)},
			&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("[]UpdateMetadataBroker:v2",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
				&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokersEndpoints, Ty: schema.NewSchema("[]UpdateMetadataEndpoint:v2",
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsPort, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol, Ty: schema.TypeInt16},
				)},
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersRack, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: UpdateMetadataRequest, API Key: 6, Version: 3
		schema.NewSchema("UpdateMetadataRequest:v3",
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldUpdateMetadataRequestUngroupedPartitionStates, Ty: schema.NewSchema("[]UpdateMetadataPartitionState:v3",
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesTopicName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesControllerEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesLeader, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesIsr, Ty: schema.TypeInt32Array},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesZkVersion, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesReplicas, Ty: schema.TypeInt32Array},
			)},
			&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("[]UpdateMetadataBroker:v3",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
				&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokersEndpoints, Ty: schema.NewSchema("[]UpdateMetadataEndpoint:v3",
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsPort, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsListener, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol, Ty: schema.TypeInt16},
				)},
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersRack, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: UpdateMetadataRequest, API Key: 6, Version: 4
		schema.NewSchema("UpdateMetadataRequest:v4",
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldUpdateMetadataRequestUngroupedPartitionStates, Ty: schema.NewSchema("[]UpdateMetadataPartitionState:v4",
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesTopicName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesControllerEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesLeader, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesIsr, Ty: schema.TypeInt32Array},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesZkVersion, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesReplicas, Ty: schema.TypeInt32Array},
				&schema.Mfield{Name: FieldUpdateMetadataRequestUngroupedPartitionStatesOfflineReplicas, Ty: schema.TypeInt32Array},
			)},
			&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("[]UpdateMetadataBroker:v4",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
				&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokersEndpoints, Ty: schema.NewSchema("[]UpdateMetadataEndpoint:v4",
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsPort, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsListener, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol, Ty: schema.TypeInt16},
				)},
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersRack, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: UpdateMetadataRequest, API Key: 6, Version: 5
		schema.NewSchema("UpdateMetadataRequest:v5",
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.Array{Name: FieldUpdateMetadataRequestTopicStates, Ty: schema.NewSchema("[]UpdateMetadataTopicState:v5",
				&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesTopicName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldUpdateMetadataRequestTopicStatesPartitionStates, Ty: schema.NewSchema("[]UpdateMetadataPartitionState:v5",
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesControllerEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesLeader, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesIsr, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesZkVersion, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesReplicas, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesOfflineReplicas, Ty: schema.TypeInt32Array},
				)},
			)},
			&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("[]UpdateMetadataBroker:v5",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
				&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokersEndpoints, Ty: schema.NewSchema("[]UpdateMetadataEndpoint:v5",
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsPort, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsListener, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol, Ty: schema.TypeInt16},
				)},
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersRack, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: UpdateMetadataRequest, API Key: 6, Version: 6
		schema.NewSchema("UpdateMetadataRequest:v6",
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.ArrayCompact{Name: FieldUpdateMetadataRequestTopicStates, Ty: schema.NewSchema("[]UpdateMetadataTopicState:v6",
				&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesTopicName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldUpdateMetadataRequestTopicStatesPartitionStates, Ty: schema.NewSchema("[]UpdateMetadataPartitionState:v6",
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesControllerEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesLeader, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesIsr, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesZkVersion, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesOfflineReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldUpdateMetadataRequestTopicStatesTags},
			)},
			&schema.ArrayCompact{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("[]UpdateMetadataBroker:v6",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
				&schema.ArrayCompact{Name: FieldUpdateMetadataRequestLiveBrokersEndpoints, Ty: schema.NewSchema("[]UpdateMetadataEndpoint:v6",
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsPort, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsHost, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsListener, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol, Ty: schema.TypeInt16},
					&schema.SchemaTaggedFields{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsTags},
				)},
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersRack, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldUpdateMetadataRequestLiveBrokersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldUpdateMetadataRequestTags},
		),

		// Message: UpdateMetadataRequest, API Key: 6, Version: 7
		schema.NewSchema("UpdateMetadataRequest:v7",
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.ArrayCompact{Name: FieldUpdateMetadataRequestTopicStates, Ty: schema.NewSchema("[]UpdateMetadataTopicState:v7",
				&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesTopicName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesTopicId, Ty: schema.TypeUuid},
				&schema.ArrayCompact{Name: FieldUpdateMetadataRequestTopicStatesPartitionStates, Ty: schema.NewSchema("[]UpdateMetadataPartitionState:v7",
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesControllerEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesLeader, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesIsr, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesZkVersion, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesOfflineReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldUpdateMetadataRequestTopicStatesPartitionStatesTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldUpdateMetadataRequestTopicStatesTags},
			)},
			&schema.ArrayCompact{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("[]UpdateMetadataBroker:v7",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
				&schema.ArrayCompact{Name: FieldUpdateMetadataRequestLiveBrokersEndpoints, Ty: schema.NewSchema("[]UpdateMetadataEndpoint:v7",
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsPort, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsHost, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsListener, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol, Ty: schema.TypeInt16},
					&schema.SchemaTaggedFields{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsTags},
				)},
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersRack, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldUpdateMetadataRequestLiveBrokersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldUpdateMetadataRequestTags},
		),
	}

}

const (

	// FieldUpdateMetadataRequestBrokerEpoch is: The broker epoch.
	FieldUpdateMetadataRequestBrokerEpoch = "BrokerEpoch"

	// FieldUpdateMetadataRequestControllerEpoch is: The controller epoch.
	FieldUpdateMetadataRequestControllerEpoch = "ControllerEpoch"

	// FieldUpdateMetadataRequestControllerId is: The controller id.
	FieldUpdateMetadataRequestControllerId = "ControllerId"

	// FieldUpdateMetadataRequestLiveBrokers is:
	FieldUpdateMetadataRequestLiveBrokers = "LiveBrokers"

	// FieldUpdateMetadataRequestLiveBrokersEndpoints is: The broker endpoints.
	FieldUpdateMetadataRequestLiveBrokersEndpoints = "Endpoints"

	// FieldUpdateMetadataRequestLiveBrokersEndpointsHost is: The hostname of this endpoint
	FieldUpdateMetadataRequestLiveBrokersEndpointsHost = "Host"

	// FieldUpdateMetadataRequestLiveBrokersEndpointsListener is: The listener name.
	FieldUpdateMetadataRequestLiveBrokersEndpointsListener = "Listener"

	// FieldUpdateMetadataRequestLiveBrokersEndpointsPort is: The port of this endpoint
	FieldUpdateMetadataRequestLiveBrokersEndpointsPort = "Port"

	// FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol is: The security protocol type.
	FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol = "SecurityProtocol"

	// FieldUpdateMetadataRequestLiveBrokersEndpointsTags is: The tagged fields.
	FieldUpdateMetadataRequestLiveBrokersEndpointsTags = "Tags"

	// FieldUpdateMetadataRequestLiveBrokersId is: The broker id.
	FieldUpdateMetadataRequestLiveBrokersId = "Id"

	// FieldUpdateMetadataRequestLiveBrokersRack is: The rack which this broker belongs to.
	FieldUpdateMetadataRequestLiveBrokersRack = "Rack"

	// FieldUpdateMetadataRequestLiveBrokersTags is: The tagged fields.
	FieldUpdateMetadataRequestLiveBrokersTags = "Tags"

	// FieldUpdateMetadataRequestTags is: The tagged fields.
	FieldUpdateMetadataRequestTags = "Tags"

	// FieldUpdateMetadataRequestTopicStates is: In newer versions of this RPC, each topic that we would like to update.
	FieldUpdateMetadataRequestTopicStates = "TopicStates"

	// FieldUpdateMetadataRequestTopicStatesPartitionStates is: The partition that we would like to update.
	FieldUpdateMetadataRequestTopicStatesPartitionStates = "PartitionStates"

	// FieldUpdateMetadataRequestTopicStatesPartitionStatesControllerEpoch is: The controller epoch.
	FieldUpdateMetadataRequestTopicStatesPartitionStatesControllerEpoch = "ControllerEpoch"

	// FieldUpdateMetadataRequestTopicStatesPartitionStatesIsr is: The brokers which are in the ISR for this partition.
	FieldUpdateMetadataRequestTopicStatesPartitionStatesIsr = "Isr"

	// FieldUpdateMetadataRequestTopicStatesPartitionStatesLeader is: The ID of the broker which is the current partition leader.
	FieldUpdateMetadataRequestTopicStatesPartitionStatesLeader = "Leader"

	// FieldUpdateMetadataRequestTopicStatesPartitionStatesLeaderEpoch is: The leader epoch of this partition.
	FieldUpdateMetadataRequestTopicStatesPartitionStatesLeaderEpoch = "LeaderEpoch"

	// FieldUpdateMetadataRequestTopicStatesPartitionStatesOfflineReplicas is: The replicas of this partition which are offline.
	FieldUpdateMetadataRequestTopicStatesPartitionStatesOfflineReplicas = "OfflineReplicas"

	// FieldUpdateMetadataRequestTopicStatesPartitionStatesPartitionIndex is: The partition index.
	FieldUpdateMetadataRequestTopicStatesPartitionStatesPartitionIndex = "PartitionIndex"

	// FieldUpdateMetadataRequestTopicStatesPartitionStatesReplicas is: All the replicas of this partition.
	FieldUpdateMetadataRequestTopicStatesPartitionStatesReplicas = "Replicas"

	// FieldUpdateMetadataRequestTopicStatesPartitionStatesTags is: The tagged fields.
	FieldUpdateMetadataRequestTopicStatesPartitionStatesTags = "Tags"

	// FieldUpdateMetadataRequestTopicStatesPartitionStatesZkVersion is: The Zookeeper version.
	FieldUpdateMetadataRequestTopicStatesPartitionStatesZkVersion = "ZkVersion"

	// FieldUpdateMetadataRequestTopicStatesTags is: The tagged fields.
	FieldUpdateMetadataRequestTopicStatesTags = "Tags"

	// FieldUpdateMetadataRequestTopicStatesTopicId is: The topic id.
	FieldUpdateMetadataRequestTopicStatesTopicId = "TopicId"

	// FieldUpdateMetadataRequestTopicStatesTopicName is: The topic name.
	FieldUpdateMetadataRequestTopicStatesTopicName = "TopicName"

	// FieldUpdateMetadataRequestUngroupedPartitionStates is: In older versions of this RPC, each partition that we would like to update.
	FieldUpdateMetadataRequestUngroupedPartitionStates = "UngroupedPartitionStates"

	// FieldUpdateMetadataRequestUngroupedPartitionStatesControllerEpoch is: The controller epoch.
	FieldUpdateMetadataRequestUngroupedPartitionStatesControllerEpoch = "ControllerEpoch"

	// FieldUpdateMetadataRequestUngroupedPartitionStatesIsr is: The brokers which are in the ISR for this partition.
	FieldUpdateMetadataRequestUngroupedPartitionStatesIsr = "Isr"

	// FieldUpdateMetadataRequestUngroupedPartitionStatesLeader is: The ID of the broker which is the current partition leader.
	FieldUpdateMetadataRequestUngroupedPartitionStatesLeader = "Leader"

	// FieldUpdateMetadataRequestUngroupedPartitionStatesLeaderEpoch is: The leader epoch of this partition.
	FieldUpdateMetadataRequestUngroupedPartitionStatesLeaderEpoch = "LeaderEpoch"

	// FieldUpdateMetadataRequestUngroupedPartitionStatesOfflineReplicas is: The replicas of this partition which are offline.
	FieldUpdateMetadataRequestUngroupedPartitionStatesOfflineReplicas = "OfflineReplicas"

	// FieldUpdateMetadataRequestUngroupedPartitionStatesPartitionIndex is: The partition index.
	FieldUpdateMetadataRequestUngroupedPartitionStatesPartitionIndex = "PartitionIndex"

	// FieldUpdateMetadataRequestUngroupedPartitionStatesReplicas is: All the replicas of this partition.
	FieldUpdateMetadataRequestUngroupedPartitionStatesReplicas = "Replicas"

	// FieldUpdateMetadataRequestUngroupedPartitionStatesTopicName is: In older versions of this RPC, the topic name.
	FieldUpdateMetadataRequestUngroupedPartitionStatesTopicName = "TopicName"

	// FieldUpdateMetadataRequestUngroupedPartitionStatesZkVersion is: The Zookeeper version.
	FieldUpdateMetadataRequestUngroupedPartitionStatesZkVersion = "ZkVersion"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/UpdateMetadataRequest.json
const originalUpdateMetadataRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 6,
  "type": "request",
  "listeners": ["zkBroker"],
  "name": "UpdateMetadataRequest",
  // Version 1 allows specifying multiple endpoints for each broker.
  //
  // Version 2 adds the rack.
  //
  // Version 3 adds the listener name.
  //
  // Version 4 adds the offline replica list.
  //
  // Version 5 adds the broker epoch field and normalizes partitions by topic.
  // Version 7 adds topicId
  "validVersions": "0-7",
  "flexibleVersions": "6+",
  "fields": [
    { "name": "ControllerId", "type": "int32", "versions": "0+", "entityType": "brokerId",
      "about": "The controller id." },
    { "name": "ControllerEpoch", "type": "int32", "versions": "0+",
      "about": "The controller epoch." },
    { "name": "BrokerEpoch", "type": "int64", "versions": "5+", "ignorable": true, "default": "-1",
      "about": "The broker epoch." },
    { "name": "UngroupedPartitionStates", "type": "[]UpdateMetadataPartitionState", "versions": "0-4",
      "about": "In older versions of this RPC, each partition that we would like to update." },
    { "name": "TopicStates", "type": "[]UpdateMetadataTopicState", "versions": "5+",
      "about": "In newer versions of this RPC, each topic that we would like to update.", "fields": [
      { "name": "TopicName", "type": "string", "versions": "5+", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "TopicId", "type": "uuid", "versions": "7+", "ignorable": true, "about": "The topic id."},
      { "name": "PartitionStates", "type": "[]UpdateMetadataPartitionState", "versions": "5+",
        "about": "The partition that we would like to update." }
    ]},
    { "name": "LiveBrokers", "type": "[]UpdateMetadataBroker", "versions": "0+", "fields": [
        { "name": "Id", "type": "int32", "versions": "0+", "entityType": "brokerId",
          "about": "The broker id." },
        // Version 0 of the protocol only allowed specifying a single host and
        // port per broker, rather than an array of endpoints.
        { "name": "V0Host", "type": "string", "versions": "0", "ignorable": true,
          "about": "The broker hostname." },
        { "name": "V0Port", "type": "int32", "versions": "0", "ignorable": true,
          "about": "The broker port." },
        { "name": "Endpoints", "type": "[]UpdateMetadataEndpoint", "versions": "1+", "ignorable": true,
          "about": "The broker endpoints.", "fields": [
          { "name": "Port", "type": "int32", "versions": "1+",
            "about": "The port of this endpoint" },
          { "name": "Host", "type": "string", "versions": "1+",
            "about": "The hostname of this endpoint" },
          { "name": "Listener", "type": "string", "versions": "3+", "ignorable": true,
            "about": "The listener name." },
          { "name": "SecurityProtocol", "type": "int16", "versions": "1+",
            "about": "The security protocol type." }
        ]},
        { "name": "Rack", "type": "string", "versions": "2+", "nullableVersions": "0+", "ignorable": true,
          "about": "The rack which this broker belongs to." }
    ]}
  ],
  "commonStructs": [
    { "name": "UpdateMetadataPartitionState", "versions": "0+", "fields": [
      { "name": "TopicName", "type": "string", "versions": "0-4", "entityType": "topicName", "ignorable": true,
        "about": "In older versions of this RPC, the topic name." },
      { "name": "PartitionIndex", "type": "int32", "versions": "0+",
        "about": "The partition index." },
      { "name": "ControllerEpoch", "type": "int32", "versions": "0+",
        "about": "The controller epoch." },
      { "name": "Leader", "type": "int32", "versions": "0+", "entityType": "brokerId",
        "about": "The ID of the broker which is the current partition leader." },
      { "name": "LeaderEpoch", "type": "int32", "versions": "0+",
        "about": "The leader epoch of this partition." },
      { "name": "Isr", "type": "[]int32", "versions": "0+", "entityType": "brokerId",
        "about": "The brokers which are in the ISR for this partition." },
      { "name": "ZkVersion", "type": "int32", "versions": "0+",
        "about": "The Zookeeper version." },
      { "name": "Replicas", "type": "[]int32", "versions": "0+", "entityType": "brokerId",
        "about": "All the replicas of this partition." },
      { "name": "OfflineReplicas", "type": "[]int32", "versions": "4+", "entityType": "brokerId", "ignorable": true,
        "about": "The replicas of this partition which are offline." }
    ]}
  ]
}
`
