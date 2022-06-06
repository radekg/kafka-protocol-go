package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init4LeaderAndIsrRequest() []schema.Schema {

	return []schema.Schema{
		// Message: LeaderAndIsrRequest, API Key: 4, Version: 0
		schema.NewSchema("LeaderAndIsrRequest:v0",
			&schema.Mfield{Name: FieldLeaderAndIsrRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaderAndIsrRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldLeaderAndIsrRequestUngroupedPartitionStates, Ty: schema.NewSchema("[]LeaderAndIsrPartitionState:v0",
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesTopicName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesControllerEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesLeader, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesIsr, Ty: schema.TypeInt32Array},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesZkVersion, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesReplicas, Ty: schema.TypeInt32Array},
			)},
			&schema.Array{Name: FieldLeaderAndIsrRequestLiveLeaders, Ty: schema.NewSchema("[]LeaderAndIsrLiveLeader:v0",
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersBrokerId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersHostName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersPort, Ty: schema.TypeInt32},
			)},
		),

		// Message: LeaderAndIsrRequest, API Key: 4, Version: 1
		schema.NewSchema("LeaderAndIsrRequest:v1",
			&schema.Mfield{Name: FieldLeaderAndIsrRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaderAndIsrRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldLeaderAndIsrRequestUngroupedPartitionStates, Ty: schema.NewSchema("[]LeaderAndIsrPartitionState:v1",
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesTopicName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesControllerEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesLeader, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesIsr, Ty: schema.TypeInt32Array},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesZkVersion, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesReplicas, Ty: schema.TypeInt32Array},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestUngroupedPartitionStatesIsNew, Ty: schema.TypeBool},
			)},
			&schema.Array{Name: FieldLeaderAndIsrRequestLiveLeaders, Ty: schema.NewSchema("[]LeaderAndIsrLiveLeader:v1",
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersBrokerId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersHostName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersPort, Ty: schema.TypeInt32},
			)},
		),

		// Message: LeaderAndIsrRequest, API Key: 4, Version: 2
		schema.NewSchema("LeaderAndIsrRequest:v2",
			&schema.Mfield{Name: FieldLeaderAndIsrRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaderAndIsrRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaderAndIsrRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.Array{Name: FieldLeaderAndIsrRequestTopicStates, Ty: schema.NewSchema("[]LeaderAndIsrTopicState:v2",
				&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesTopicName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStates, Ty: schema.NewSchema("[]LeaderAndIsrPartitionState:v2",
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesControllerEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeader, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesIsr, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesZkVersion, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesReplicas, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesIsNew, Ty: schema.TypeBool},
				)},
			)},
			&schema.Array{Name: FieldLeaderAndIsrRequestLiveLeaders, Ty: schema.NewSchema("[]LeaderAndIsrLiveLeader:v2",
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersBrokerId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersHostName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersPort, Ty: schema.TypeInt32},
			)},
		),

		// Message: LeaderAndIsrRequest, API Key: 4, Version: 3
		schema.NewSchema("LeaderAndIsrRequest:v3",
			&schema.Mfield{Name: FieldLeaderAndIsrRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaderAndIsrRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaderAndIsrRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.Array{Name: FieldLeaderAndIsrRequestTopicStates, Ty: schema.NewSchema("[]LeaderAndIsrTopicState:v3",
				&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesTopicName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStates, Ty: schema.NewSchema("[]LeaderAndIsrPartitionState:v3",
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesControllerEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeader, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesIsr, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesZkVersion, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesReplicas, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesAddingReplicas, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesRemovingReplicas, Ty: schema.TypeInt32Array},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesIsNew, Ty: schema.TypeBool},
				)},
			)},
			&schema.Array{Name: FieldLeaderAndIsrRequestLiveLeaders, Ty: schema.NewSchema("[]LeaderAndIsrLiveLeader:v3",
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersBrokerId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersHostName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersPort, Ty: schema.TypeInt32},
			)},
		),

		// Message: LeaderAndIsrRequest, API Key: 4, Version: 4
		schema.NewSchema("LeaderAndIsrRequest:v4",
			&schema.Mfield{Name: FieldLeaderAndIsrRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaderAndIsrRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaderAndIsrRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.ArrayCompact{Name: FieldLeaderAndIsrRequestTopicStates, Ty: schema.NewSchema("[]LeaderAndIsrTopicState:v4",
				&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesTopicName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStates, Ty: schema.NewSchema("[]LeaderAndIsrPartitionState:v4",
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesControllerEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeader, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesIsr, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesZkVersion, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesAddingReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesRemovingReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesIsNew, Ty: schema.TypeBool},
					&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrRequestTopicStatesTags},
			)},
			&schema.ArrayCompact{Name: FieldLeaderAndIsrRequestLiveLeaders, Ty: schema.NewSchema("[]LeaderAndIsrLiveLeader:v4",
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersBrokerId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersHostName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersPort, Ty: schema.TypeInt32},
				&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrRequestLiveLeadersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrRequestTags},
		),

		// Message: LeaderAndIsrRequest, API Key: 4, Version: 5
		schema.NewSchema("LeaderAndIsrRequest:v5",
			&schema.Mfield{Name: FieldLeaderAndIsrRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaderAndIsrRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaderAndIsrRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldLeaderAndIsrRequestType, Ty: schema.TypeInt8},
			&schema.ArrayCompact{Name: FieldLeaderAndIsrRequestTopicStates, Ty: schema.NewSchema("[]LeaderAndIsrTopicState:v5",
				&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesTopicName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesTopicId, Ty: schema.TypeUuid},
				&schema.ArrayCompact{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStates, Ty: schema.NewSchema("[]LeaderAndIsrPartitionState:v5",
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesControllerEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeader, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesIsr, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesZkVersion, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesAddingReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesRemovingReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesIsNew, Ty: schema.TypeBool},
					&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrRequestTopicStatesTags},
			)},
			&schema.ArrayCompact{Name: FieldLeaderAndIsrRequestLiveLeaders, Ty: schema.NewSchema("[]LeaderAndIsrLiveLeader:v5",
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersBrokerId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersHostName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersPort, Ty: schema.TypeInt32},
				&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrRequestLiveLeadersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrRequestTags},
		),

		// Message: LeaderAndIsrRequest, API Key: 4, Version: 6
		schema.NewSchema("LeaderAndIsrRequest:v6",
			&schema.Mfield{Name: FieldLeaderAndIsrRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaderAndIsrRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaderAndIsrRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldLeaderAndIsrRequestType, Ty: schema.TypeInt8},
			&schema.ArrayCompact{Name: FieldLeaderAndIsrRequestTopicStates, Ty: schema.NewSchema("[]LeaderAndIsrTopicState:v6",
				&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesTopicName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesTopicId, Ty: schema.TypeUuid},
				&schema.ArrayCompact{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStates, Ty: schema.NewSchema("[]LeaderAndIsrPartitionState:v6",
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesControllerEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeader, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesIsr, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesZkVersion, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesAddingReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesRemovingReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesIsNew, Ty: schema.TypeBool},
					&schema.Mfield{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeaderRecoveryState, Ty: schema.TypeInt8},
					&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrRequestTopicStatesPartitionStatesTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrRequestTopicStatesTags},
			)},
			&schema.ArrayCompact{Name: FieldLeaderAndIsrRequestLiveLeaders, Ty: schema.NewSchema("[]LeaderAndIsrLiveLeader:v6",
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersBrokerId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersHostName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldLeaderAndIsrRequestLiveLeadersPort, Ty: schema.TypeInt32},
				&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrRequestLiveLeadersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrRequestTags},
		),
	}

}

const (

	// FieldLeaderAndIsrRequestBrokerEpoch is: The current broker epoch.
	FieldLeaderAndIsrRequestBrokerEpoch = "BrokerEpoch"

	// FieldLeaderAndIsrRequestControllerEpoch is: The current controller epoch.
	FieldLeaderAndIsrRequestControllerEpoch = "ControllerEpoch"

	// FieldLeaderAndIsrRequestControllerId is: The current controller ID.
	FieldLeaderAndIsrRequestControllerId = "ControllerId"

	// FieldLeaderAndIsrRequestLiveLeaders is: The current live leaders.
	FieldLeaderAndIsrRequestLiveLeaders = "LiveLeaders"

	// FieldLeaderAndIsrRequestLiveLeadersBrokerId is: The leader's broker ID.
	FieldLeaderAndIsrRequestLiveLeadersBrokerId = "BrokerId"

	// FieldLeaderAndIsrRequestLiveLeadersHostName is: The leader's hostname.
	FieldLeaderAndIsrRequestLiveLeadersHostName = "HostName"

	// FieldLeaderAndIsrRequestLiveLeadersPort is: The leader's port.
	FieldLeaderAndIsrRequestLiveLeadersPort = "Port"

	// FieldLeaderAndIsrRequestLiveLeadersTags is: The tagged fields.
	FieldLeaderAndIsrRequestLiveLeadersTags = "Tags"

	// FieldLeaderAndIsrRequestTags is: The tagged fields.
	FieldLeaderAndIsrRequestTags = "Tags"

	// FieldLeaderAndIsrRequestTopicStates is: Each topic.
	FieldLeaderAndIsrRequestTopicStates = "TopicStates"

	// FieldLeaderAndIsrRequestTopicStatesPartitionStates is: The state of each partition
	FieldLeaderAndIsrRequestTopicStatesPartitionStates = "PartitionStates"

	// FieldLeaderAndIsrRequestTopicStatesPartitionStatesAddingReplicas is: The replica IDs that we are adding this partition to, or null if no replicas are being added.
	FieldLeaderAndIsrRequestTopicStatesPartitionStatesAddingReplicas = "AddingReplicas"

	// FieldLeaderAndIsrRequestTopicStatesPartitionStatesControllerEpoch is: The controller epoch.
	FieldLeaderAndIsrRequestTopicStatesPartitionStatesControllerEpoch = "ControllerEpoch"

	// FieldLeaderAndIsrRequestTopicStatesPartitionStatesIsNew is: Whether the replica should have existed on the broker or not.
	FieldLeaderAndIsrRequestTopicStatesPartitionStatesIsNew = "IsNew"

	// FieldLeaderAndIsrRequestTopicStatesPartitionStatesIsr is: The in-sync replica IDs.
	FieldLeaderAndIsrRequestTopicStatesPartitionStatesIsr = "Isr"

	// FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeader is: The broker ID of the leader.
	FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeader = "Leader"

	// FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeaderEpoch is: The leader epoch.
	FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeaderEpoch = "LeaderEpoch"

	// FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeaderRecoveryState is: 1 if the partition is recovering from an unclean leader election; 0 otherwise.
	FieldLeaderAndIsrRequestTopicStatesPartitionStatesLeaderRecoveryState = "LeaderRecoveryState"

	// FieldLeaderAndIsrRequestTopicStatesPartitionStatesPartitionIndex is: The partition index.
	FieldLeaderAndIsrRequestTopicStatesPartitionStatesPartitionIndex = "PartitionIndex"

	// FieldLeaderAndIsrRequestTopicStatesPartitionStatesRemovingReplicas is: The replica IDs that we are removing this partition from, or null if no replicas are being removed.
	FieldLeaderAndIsrRequestTopicStatesPartitionStatesRemovingReplicas = "RemovingReplicas"

	// FieldLeaderAndIsrRequestTopicStatesPartitionStatesReplicas is: The replica IDs.
	FieldLeaderAndIsrRequestTopicStatesPartitionStatesReplicas = "Replicas"

	// FieldLeaderAndIsrRequestTopicStatesPartitionStatesTags is: The tagged fields.
	FieldLeaderAndIsrRequestTopicStatesPartitionStatesTags = "Tags"

	// FieldLeaderAndIsrRequestTopicStatesPartitionStatesZkVersion is: The ZooKeeper version.
	FieldLeaderAndIsrRequestTopicStatesPartitionStatesZkVersion = "ZkVersion"

	// FieldLeaderAndIsrRequestTopicStatesTags is: The tagged fields.
	FieldLeaderAndIsrRequestTopicStatesTags = "Tags"

	// FieldLeaderAndIsrRequestTopicStatesTopicId is: The unique topic ID.
	FieldLeaderAndIsrRequestTopicStatesTopicId = "TopicId"

	// FieldLeaderAndIsrRequestTopicStatesTopicName is: The topic name.
	FieldLeaderAndIsrRequestTopicStatesTopicName = "TopicName"

	// FieldLeaderAndIsrRequestType is: The type that indicates whether all topics are included in the request
	FieldLeaderAndIsrRequestType = "Type"

	// FieldLeaderAndIsrRequestUngroupedPartitionStates is: The state of each partition, in a v0 or v1 message.
	FieldLeaderAndIsrRequestUngroupedPartitionStates = "UngroupedPartitionStates"

	// FieldLeaderAndIsrRequestUngroupedPartitionStatesControllerEpoch is: The controller epoch.
	FieldLeaderAndIsrRequestUngroupedPartitionStatesControllerEpoch = "ControllerEpoch"

	// FieldLeaderAndIsrRequestUngroupedPartitionStatesIsNew is: Whether the replica should have existed on the broker or not.
	FieldLeaderAndIsrRequestUngroupedPartitionStatesIsNew = "IsNew"

	// FieldLeaderAndIsrRequestUngroupedPartitionStatesIsr is: The in-sync replica IDs.
	FieldLeaderAndIsrRequestUngroupedPartitionStatesIsr = "Isr"

	// FieldLeaderAndIsrRequestUngroupedPartitionStatesLeader is: The broker ID of the leader.
	FieldLeaderAndIsrRequestUngroupedPartitionStatesLeader = "Leader"

	// FieldLeaderAndIsrRequestUngroupedPartitionStatesLeaderEpoch is: The leader epoch.
	FieldLeaderAndIsrRequestUngroupedPartitionStatesLeaderEpoch = "LeaderEpoch"

	// FieldLeaderAndIsrRequestUngroupedPartitionStatesPartitionIndex is: The partition index.
	FieldLeaderAndIsrRequestUngroupedPartitionStatesPartitionIndex = "PartitionIndex"

	// FieldLeaderAndIsrRequestUngroupedPartitionStatesReplicas is: The replica IDs.
	FieldLeaderAndIsrRequestUngroupedPartitionStatesReplicas = "Replicas"

	// FieldLeaderAndIsrRequestUngroupedPartitionStatesTopicName is: The topic name.  This is only present in v0 or v1.
	FieldLeaderAndIsrRequestUngroupedPartitionStatesTopicName = "TopicName"

	// FieldLeaderAndIsrRequestUngroupedPartitionStatesZkVersion is: The ZooKeeper version.
	FieldLeaderAndIsrRequestUngroupedPartitionStatesZkVersion = "ZkVersion"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/LeaderAndIsrRequest.json
const originalLeaderAndIsrRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 4,
  "type": "request",
  "listeners": ["zkBroker"],
  "name": "LeaderAndIsrRequest",
  // Version 1 adds IsNew.
  //
  // Version 2 adds broker epoch and reorganizes the partitions by topic.
  //
  // Version 3 adds AddingReplicas and RemovingReplicas.
  //
  // Version 4 is the first flexible version.
  //
  // Version 5 adds Topic ID and Type to the TopicStates, as described in KIP-516.
  //
  // Version 6 adds ElectionState as described in KIP-704.
  "validVersions": "0-6",
  "flexibleVersions": "4+",
  "fields": [
    { "name": "ControllerId", "type": "int32", "versions": "0+", "entityType": "brokerId",
      "about": "The current controller ID." },
    { "name": "ControllerEpoch", "type": "int32", "versions": "0+",
      "about": "The current controller epoch." },
    { "name": "BrokerEpoch", "type": "int64", "versions": "2+", "ignorable": true, "default": "-1",
      "about": "The current broker epoch." },
    { "name": "Type", "type": "int8", "versions": "5+",
      "about": "The type that indicates whether all topics are included in the request"},
    { "name": "UngroupedPartitionStates", "type": "[]LeaderAndIsrPartitionState", "versions": "0-1",
      "about": "The state of each partition, in a v0 or v1 message." },
    // In v0 or v1 requests, each partition is listed alongside its topic name.
    // In v2+ requests, partitions are organized by topic, so that each topic name
    // only needs to be listed once.
    { "name": "TopicStates", "type": "[]LeaderAndIsrTopicState", "versions": "2+",
      "about": "Each topic.", "fields": [
      { "name": "TopicName", "type": "string", "versions": "2+", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "TopicId", "type": "uuid", "versions": "5+", "ignorable": true,
        "about": "The unique topic ID." },
      { "name": "PartitionStates", "type": "[]LeaderAndIsrPartitionState", "versions": "2+",
        "about": "The state of each partition" }
    ]},
    { "name": "LiveLeaders", "type": "[]LeaderAndIsrLiveLeader", "versions": "0+",
      "about": "The current live leaders.", "fields": [
      { "name": "BrokerId", "type": "int32", "versions": "0+", "entityType": "brokerId",
        "about": "The leader's broker ID." },
      { "name": "HostName", "type": "string", "versions": "0+",
        "about": "The leader's hostname." },
      { "name": "Port", "type": "int32", "versions": "0+",
        "about": "The leader's port." }
    ]}
  ],
  "commonStructs": [
    { "name": "LeaderAndIsrPartitionState", "versions": "0+", "fields": [
      { "name": "TopicName", "type": "string", "versions": "0-1", "entityType": "topicName", "ignorable": true,
        "about": "The topic name.  This is only present in v0 or v1." },
      { "name": "PartitionIndex", "type": "int32", "versions": "0+",
        "about": "The partition index." },
      { "name": "ControllerEpoch", "type": "int32", "versions": "0+",
        "about": "The controller epoch." },
      { "name": "Leader", "type": "int32", "versions": "0+", "entityType": "brokerId",
        "about": "The broker ID of the leader." },
      { "name": "LeaderEpoch", "type": "int32", "versions": "0+",
        "about": "The leader epoch." },
      { "name": "Isr", "type": "[]int32", "versions": "0+", "entityType": "brokerId",
        "about": "The in-sync replica IDs." },
      { "name": "ZkVersion", "type": "int32", "versions": "0+",
        "about": "The ZooKeeper version." },
      { "name": "Replicas", "type": "[]int32", "versions": "0+", "entityType": "brokerId",
        "about": "The replica IDs." },
      { "name": "AddingReplicas", "type": "[]int32", "versions": "3+", "ignorable": true, "entityType": "brokerId",
        "about": "The replica IDs that we are adding this partition to, or null if no replicas are being added." },
      { "name": "RemovingReplicas", "type": "[]int32", "versions": "3+", "ignorable": true, "entityType": "brokerId",
        "about": "The replica IDs that we are removing this partition from, or null if no replicas are being removed." },
      { "name": "IsNew", "type": "bool", "versions": "1+", "default": "false", "ignorable": true,
        "about": "Whether the replica should have existed on the broker or not." },
      { "name": "LeaderRecoveryState", "type": "int8", "versions": "6+", "default": "0",
        "about": "1 if the partition is recovering from an unclean leader election; 0 otherwise." }
    ]}
  ]
}
`
