package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init6UpdateMetadataRequest() []schema.Schema {

	return []schema.Schema{

		// Message: UpdateMetadataRequest, API Key: 6, Version: 0
		schema.NewSchema("UpdateMetadataRequestv0", 
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldUpdateMetadataRequestUngroupedPartitionStates, Ty: schema.NewSchema("UngroupedPartitionStatesV0",
			)},
			&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("LiveBrokersV0",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
			)},
		),

		// Message: UpdateMetadataRequest, API Key: 6, Version: 1
		schema.NewSchema("UpdateMetadataRequestv1", 
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldUpdateMetadataRequestUngroupedPartitionStates, Ty: schema.NewSchema("UngroupedPartitionStatesV1",
			)},
			&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("LiveBrokersV1",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
				&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokersEndpoints, Ty: schema.NewSchema("EndpointsV1",
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsPort, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: UpdateMetadataRequest, API Key: 6, Version: 2
		schema.NewSchema("UpdateMetadataRequestv2", 
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldUpdateMetadataRequestUngroupedPartitionStates, Ty: schema.NewSchema("UngroupedPartitionStatesV2",
			)},
			&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("LiveBrokersV2",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
				&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokersEndpoints, Ty: schema.NewSchema("EndpointsV2",
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsPort, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol, Ty: schema.TypeInt16},
				)},
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersRack, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: UpdateMetadataRequest, API Key: 6, Version: 3
		schema.NewSchema("UpdateMetadataRequestv3", 
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldUpdateMetadataRequestUngroupedPartitionStates, Ty: schema.NewSchema("UngroupedPartitionStatesV3",
			)},
			&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("LiveBrokersV3",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
				&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokersEndpoints, Ty: schema.NewSchema("EndpointsV3",
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsPort, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsListener, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol, Ty: schema.TypeInt16},
				)},
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersRack, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: UpdateMetadataRequest, API Key: 6, Version: 4
		schema.NewSchema("UpdateMetadataRequestv4", 
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldUpdateMetadataRequestUngroupedPartitionStates, Ty: schema.NewSchema("UngroupedPartitionStatesV4",
			)},
			&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("LiveBrokersV4",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
				&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokersEndpoints, Ty: schema.NewSchema("EndpointsV4",
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsPort, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsListener, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol, Ty: schema.TypeInt16},
				)},
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersRack, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: UpdateMetadataRequest, API Key: 6, Version: 5
		schema.NewSchema("UpdateMetadataRequestv5", 
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.Array{Name: FieldUpdateMetadataRequestTopicStates, Ty: schema.NewSchema("TopicStatesV5",
				&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesTopicName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldUpdateMetadataRequestTopicStatesPartitionStates, Ty: schema.NewSchema("PartitionStatesV5",
				)},
			)},
			&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("LiveBrokersV5",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
				&schema.Array{Name: FieldUpdateMetadataRequestLiveBrokersEndpoints, Ty: schema.NewSchema("EndpointsV5",
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsPort, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsListener, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol, Ty: schema.TypeInt16},
				)},
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersRack, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: UpdateMetadataRequest, API Key: 6, Version: 6
		schema.NewSchema("UpdateMetadataRequestv6", 
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.ArrayCompact{Name: FieldUpdateMetadataRequestTopicStates, Ty: schema.NewSchema("TopicStatesV6",
				&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesTopicName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldUpdateMetadataRequestTopicStatesPartitionStates, Ty: schema.NewSchema("PartitionStatesV6",
				)},
				&schema.SchemaTaggedFields{Name: FieldUpdateMetadataRequestTopicStatesTags},
			)},
			&schema.ArrayCompact{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("LiveBrokersV6",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
				&schema.ArrayCompact{Name: FieldUpdateMetadataRequestLiveBrokersEndpoints, Ty: schema.NewSchema("EndpointsV6",
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
		schema.NewSchema("UpdateMetadataRequestv7", 
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateMetadataRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.ArrayCompact{Name: FieldUpdateMetadataRequestTopicStates, Ty: schema.NewSchema("TopicStatesV7",
				&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesTopicName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldUpdateMetadataRequestTopicStatesTopicId, Ty: schema.TypeUuid},
				&schema.ArrayCompact{Name: FieldUpdateMetadataRequestTopicStatesPartitionStates, Ty: schema.NewSchema("PartitionStatesV7",
				)},
				&schema.SchemaTaggedFields{Name: FieldUpdateMetadataRequestTopicStatesTags},
			)},
			&schema.ArrayCompact{Name: FieldUpdateMetadataRequestLiveBrokers, Ty: schema.NewSchema("LiveBrokersV7",
				&schema.Mfield{Name: FieldUpdateMetadataRequestLiveBrokersId, Ty: schema.TypeInt32},
				&schema.ArrayCompact{Name: FieldUpdateMetadataRequestLiveBrokersEndpoints, Ty: schema.NewSchema("EndpointsV7",
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
	// FieldUpdateMetadataRequestUngroupedPartitionStates is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestUngroupedPartitionStates = "UngroupedPartitionStates"
	// FieldUpdateMetadataRequestLiveBrokersEndpointsHost is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestLiveBrokersEndpointsHost = "Host"
	// FieldUpdateMetadataRequestTopicStatesPartitionStates is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestTopicStatesPartitionStates = "PartitionStates"
	// FieldUpdateMetadataRequestTags is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestTags = "Tags"
	// FieldUpdateMetadataRequestControllerEpoch is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestControllerEpoch = "ControllerEpoch"
	// FieldUpdateMetadataRequestTopicStatesTags is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestTopicStatesTags = "Tags"
	// FieldUpdateMetadataRequestLiveBrokersTags is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestLiveBrokersTags = "Tags"
	// FieldUpdateMetadataRequestTopicStatesTopicId is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestTopicStatesTopicId = "TopicId"
	// FieldUpdateMetadataRequestLiveBrokers is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestLiveBrokers = "LiveBrokers"
	// FieldUpdateMetadataRequestLiveBrokersId is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestLiveBrokersId = "Id"
	// FieldUpdateMetadataRequestLiveBrokersEndpointsPort is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestLiveBrokersEndpointsPort = "Port"
	// FieldUpdateMetadataRequestLiveBrokersEndpointsListener is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestLiveBrokersEndpointsListener = "Listener"
	// FieldUpdateMetadataRequestBrokerEpoch is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestBrokerEpoch = "BrokerEpoch"
	// FieldUpdateMetadataRequestLiveBrokersEndpointsTags is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestLiveBrokersEndpointsTags = "Tags"
	// FieldUpdateMetadataRequestControllerId is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestControllerId = "ControllerId"
	// FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestLiveBrokersEndpointsSecurityProtocol = "SecurityProtocol"
	// FieldUpdateMetadataRequestLiveBrokersRack is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestLiveBrokersRack = "Rack"
	// FieldUpdateMetadataRequestTopicStates is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestTopicStates = "TopicStates"
	// FieldUpdateMetadataRequestTopicStatesTopicName is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestTopicStatesTopicName = "TopicName"
	// FieldUpdateMetadataRequestLiveBrokersEndpoints is a field name that can be used to resolve the correct struct field.
	FieldUpdateMetadataRequestLiveBrokersEndpoints = "Endpoints"
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

