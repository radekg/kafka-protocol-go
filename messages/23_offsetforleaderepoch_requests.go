package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init23OffsetForLeaderEpochRequest() []schema.Schema {

	return []schema.Schema{
		// Message: OffsetForLeaderEpochRequest, API Key: 23, Version: 0
		schema.NewSchema("OffsetForLeaderEpochRequest:v0",
			&schema.Array{Name: FieldOffsetForLeaderEpochRequestTopics, Ty: schema.NewSchema("Topics:v0",
				&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetForLeaderEpochRequestTopicsPartitions, Ty: schema.NewSchema("Partitions:v0",
					&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
				)},
			)},
		),

		// Message: OffsetForLeaderEpochRequest, API Key: 23, Version: 1
		schema.NewSchema("OffsetForLeaderEpochRequest:v1",
			&schema.Array{Name: FieldOffsetForLeaderEpochRequestTopics, Ty: schema.NewSchema("Topics:v1",
				&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetForLeaderEpochRequestTopicsPartitions, Ty: schema.NewSchema("Partitions:v1",
					&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
				)},
			)},
		),

		// Message: OffsetForLeaderEpochRequest, API Key: 23, Version: 2
		schema.NewSchema("OffsetForLeaderEpochRequest:v2",
			&schema.Array{Name: FieldOffsetForLeaderEpochRequestTopics, Ty: schema.NewSchema("Topics:v2",
				&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetForLeaderEpochRequestTopicsPartitions, Ty: schema.NewSchema("Partitions:v2",
					&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsPartitionsCurrentLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
				)},
			)},
		),

		// Message: OffsetForLeaderEpochRequest, API Key: 23, Version: 3
		schema.NewSchema("OffsetForLeaderEpochRequest:v3",
			&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldOffsetForLeaderEpochRequestTopics, Ty: schema.NewSchema("Topics:v3",
				&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetForLeaderEpochRequestTopicsPartitions, Ty: schema.NewSchema("Partitions:v3",
					&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsPartitionsCurrentLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
				)},
			)},
		),

		// Message: OffsetForLeaderEpochRequest, API Key: 23, Version: 4
		schema.NewSchema("OffsetForLeaderEpochRequest:v4",
			&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestReplicaId, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldOffsetForLeaderEpochRequestTopics, Ty: schema.NewSchema("Topics:v4",
				&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsTopic, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldOffsetForLeaderEpochRequestTopicsPartitions, Ty: schema.NewSchema("Partitions:v4",
					&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsPartitionsCurrentLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochRequestTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.SchemaTaggedFields{Name: FieldOffsetForLeaderEpochRequestTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldOffsetForLeaderEpochRequestTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldOffsetForLeaderEpochRequestTags},
		),
	}

}

const (

	// FieldOffsetForLeaderEpochRequestReplicaId is: The broker ID of the follower, of -1 if this request is from a consumer.
	FieldOffsetForLeaderEpochRequestReplicaId = "ReplicaId"

	// FieldOffsetForLeaderEpochRequestTags is: The tagged fields.
	FieldOffsetForLeaderEpochRequestTags = "Tags"

	// FieldOffsetForLeaderEpochRequestTopics is: Each topic to get offsets for.
	FieldOffsetForLeaderEpochRequestTopics = "Topics"

	// FieldOffsetForLeaderEpochRequestTopicsPartitions is: Each partition to get offsets for.
	FieldOffsetForLeaderEpochRequestTopicsPartitions = "Partitions"

	// FieldOffsetForLeaderEpochRequestTopicsPartitionsCurrentLeaderEpoch is: An epoch used to fence consumers/replicas with old metadata. If the epoch provided by the client is larger than the current epoch known to the broker, then the UNKNOWN_LEADER_EPOCH error code will be returned. If the provided epoch is smaller, then the FENCED_LEADER_EPOCH error code will be returned.
	FieldOffsetForLeaderEpochRequestTopicsPartitionsCurrentLeaderEpoch = "CurrentLeaderEpoch"

	// FieldOffsetForLeaderEpochRequestTopicsPartitionsLeaderEpoch is: The epoch to look up an offset for.
	FieldOffsetForLeaderEpochRequestTopicsPartitionsLeaderEpoch = "LeaderEpoch"

	// FieldOffsetForLeaderEpochRequestTopicsPartitionsPartition is: The partition index.
	FieldOffsetForLeaderEpochRequestTopicsPartitionsPartition = "Partition"

	// FieldOffsetForLeaderEpochRequestTopicsPartitionsTags is: The tagged fields.
	FieldOffsetForLeaderEpochRequestTopicsPartitionsTags = "Tags"

	// FieldOffsetForLeaderEpochRequestTopicsTags is: The tagged fields.
	FieldOffsetForLeaderEpochRequestTopicsTags = "Tags"

	// FieldOffsetForLeaderEpochRequestTopicsTopic is: The topic name.
	FieldOffsetForLeaderEpochRequestTopicsTopic = "Topic"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/OffsetForLeaderEpochRequest.json
const originalOffsetForLeaderEpochRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 23,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "OffsetForLeaderEpochRequest",
  // Version 1 is the same as version 0.
  //
  // Version 2 adds the current leader epoch to support fencing.
  //
  // Version 3 adds ReplicaId (the default is -2 which conventionally represents a
  //    "debug" consumer which is allowed to see offsets beyond the high watermark).
  //    Followers will use this replicaId when using an older version of the protocol.
  //
  // Version 4 enables flexible versions.
  "validVersions": "0-4",
  "flexibleVersions": "4+",
  "fields": [
    { "name": "ReplicaId", "type": "int32", "versions": "3+", "default": -2, "ignorable": true, "entityType": "brokerId",
      "about": "The broker ID of the follower, of -1 if this request is from a consumer." },
    { "name": "Topics", "type": "[]OffsetForLeaderTopic", "versions": "0+",
      "about": "Each topic to get offsets for.", "fields": [
      { "name": "Topic", "type": "string", "versions": "0+", "entityType": "topicName",
        "mapKey": true, "about": "The topic name." },
      { "name": "Partitions", "type": "[]OffsetForLeaderPartition", "versions": "0+",
        "about": "Each partition to get offsets for.", "fields": [
        { "name": "Partition", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "CurrentLeaderEpoch", "type": "int32", "versions": "2+", "default": "-1", "ignorable": true,
          "about": "An epoch used to fence consumers/replicas with old metadata. If the epoch provided by the client is larger than the current epoch known to the broker, then the UNKNOWN_LEADER_EPOCH error code will be returned. If the provided epoch is smaller, then the FENCED_LEADER_EPOCH error code will be returned." },
        { "name": "LeaderEpoch", "type": "int32", "versions": "0+",
          "about": "The epoch to look up an offset for." }
      ]}
    ]}
  ]
}
`
