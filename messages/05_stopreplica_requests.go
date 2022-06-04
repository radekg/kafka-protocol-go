package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init5StopReplicaRequest() []schema.Schema {

	return []schema.Schema{

		// Message: StopReplicaRequest, API Key: 5, Version: 0
		schema.NewSchema("StopReplicaRequestv0",
			&schema.Mfield{Name: FieldStopReplicaRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldStopReplicaRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldStopReplicaRequestDeletePartitions, Ty: schema.TypeBool},
		),

		// Message: StopReplicaRequest, API Key: 5, Version: 1
		schema.NewSchema("StopReplicaRequestv1",
			&schema.Mfield{Name: FieldStopReplicaRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldStopReplicaRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldStopReplicaRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldStopReplicaRequestDeletePartitions, Ty: schema.TypeBool},
			&schema.Array{Name: FieldStopReplicaRequestTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldStopReplicaRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldStopReplicaRequestTopicsPartitionIndexes, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: StopReplicaRequest, API Key: 5, Version: 2
		schema.NewSchema("StopReplicaRequestv2",
			&schema.Mfield{Name: FieldStopReplicaRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldStopReplicaRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldStopReplicaRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldStopReplicaRequestDeletePartitions, Ty: schema.TypeBool},
			&schema.ArrayCompact{Name: FieldStopReplicaRequestTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldStopReplicaRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldStopReplicaRequestTopicsPartitionIndexes, Ty: schema.TypeInt32CompactArray},
				&schema.SchemaTaggedFields{Name: FieldStopReplicaRequestTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldStopReplicaRequestTags},
		),

		// Message: StopReplicaRequest, API Key: 5, Version: 3
		schema.NewSchema("StopReplicaRequestv3",
			&schema.Mfield{Name: FieldStopReplicaRequestControllerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldStopReplicaRequestControllerEpoch, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldStopReplicaRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.ArrayCompact{Name: FieldStopReplicaRequestTopicStates, Ty: schema.NewSchema("TopicStatesV3",
				&schema.Mfield{Name: FieldStopReplicaRequestTopicStatesTopicName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldStopReplicaRequestTopicStatesPartitionStates, Ty: schema.NewSchema("PartitionStatesV3",
					&schema.Mfield{Name: FieldStopReplicaRequestTopicStatesPartitionStatesPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldStopReplicaRequestTopicStatesPartitionStatesLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldStopReplicaRequestTopicStatesPartitionStatesDeletePartition, Ty: schema.TypeBool},
					&schema.SchemaTaggedFields{Name: FieldStopReplicaRequestTopicStatesPartitionStatesTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldStopReplicaRequestTopicStatesTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldStopReplicaRequestTags},
		),
	}
}

const (
	// FieldStopReplicaRequestControllerId is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestControllerId = "ControllerId"
	// FieldStopReplicaRequestTopicsName is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestTopicsName = "Name"
	// FieldStopReplicaRequestTags is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestTags = "Tags"
	// FieldStopReplicaRequestTopicStatesPartitionStates is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestTopicStatesPartitionStates = "PartitionStates"
	// FieldStopReplicaRequestTopicStatesTags is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestTopicStatesTags = "Tags"
	// FieldStopReplicaRequestControllerEpoch is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestControllerEpoch = "ControllerEpoch"
	// FieldStopReplicaRequestTopicsPartitionIndexes is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestTopicsPartitionIndexes = "PartitionIndexes"
	// FieldStopReplicaRequestTopicsTags is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestTopicsTags = "Tags"
	// FieldStopReplicaRequestTopicStatesTopicName is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestTopicStatesTopicName = "TopicName"
	// FieldStopReplicaRequestTopicStatesPartitionStatesDeletePartition is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestTopicStatesPartitionStatesDeletePartition = "DeletePartition"
	// FieldStopReplicaRequestBrokerEpoch is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestBrokerEpoch = "BrokerEpoch"
	// FieldStopReplicaRequestTopicStates is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestTopicStates = "TopicStates"
	// FieldStopReplicaRequestTopicStatesPartitionStatesPartitionIndex is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestTopicStatesPartitionStatesPartitionIndex = "PartitionIndex"
	// FieldStopReplicaRequestTopicStatesPartitionStatesTags is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestTopicStatesPartitionStatesTags = "Tags"
	// FieldStopReplicaRequestDeletePartitions is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestDeletePartitions = "DeletePartitions"
	// FieldStopReplicaRequestTopics is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestTopics = "Topics"
	// FieldStopReplicaRequestTopicStatesPartitionStatesLeaderEpoch is a field name that can be used to resolve the correct struct field.
	FieldStopReplicaRequestTopicStatesPartitionStatesLeaderEpoch = "LeaderEpoch"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/StopReplicaRequest.json
const originalStopReplicaRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 5,
  "type": "request",
  "listeners": ["zkBroker"],
  "name": "StopReplicaRequest",
  // Version 1 adds the broker epoch and reorganizes the partitions to be stored
  // per topic.
  //
  // Version 2 is the first flexible version.
  //
  // Version 3 adds the leader epoch per partition (KIP-570).
  "validVersions": "0-3",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ControllerId", "type": "int32", "versions": "0+", "entityType": "brokerId",
      "about": "The controller id." },
    { "name": "ControllerEpoch", "type": "int32", "versions": "0+",
      "about": "The controller epoch." },
    { "name": "BrokerEpoch", "type": "int64", "versions": "1+", "default": "-1", "ignorable": true,
      "about": "The broker epoch." },
    { "name": "DeletePartitions", "type": "bool", "versions": "0-2",
      "about": "Whether these partitions should be deleted." },
    { "name": "UngroupedPartitions", "type": "[]StopReplicaPartitionV0", "versions": "0",
      "about": "The partitions to stop.", "fields": [
      { "name": "TopicName", "type": "string", "versions": "0", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "PartitionIndex", "type": "int32", "versions": "0",
        "about": "The partition index." }
    ]},
    { "name": "Topics", "type": "[]StopReplicaTopicV1", "versions": "1-2",
      "about": "The topics to stop.", "fields": [
      { "name": "Name", "type": "string", "versions": "1-2", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "PartitionIndexes", "type": "[]int32", "versions": "1-2",
        "about": "The partition indexes." }
    ]},
    { "name": "TopicStates", "type": "[]StopReplicaTopicState", "versions": "3+",
      "about": "Each topic.", "fields": [
      { "name": "TopicName", "type": "string", "versions": "3+", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "PartitionStates", "type": "[]StopReplicaPartitionState", "versions": "3+",
        "about": "The state of each partition", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "3+",
          "about": "The partition index." },
        { "name": "LeaderEpoch", "type": "int32", "versions": "3+", "default": "-1",
          "about": "The leader epoch." },
        { "name": "DeletePartition", "type": "bool", "versions": "3+",
          "about": "Whether this partition should be deleted." }
      ]}
    ]}
  ]
}
`
