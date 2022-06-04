package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init2ListOffsetsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: ListOffsetsRequest, API Key: 2, Version: 0
		schema.NewSchema("ListOffsetsRequestv0",
			&schema.Mfield{Name: FieldListOffsetsRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldListOffsetsRequestTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldListOffsetsRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV0",
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: ListOffsetsRequest, API Key: 2, Version: 1
		schema.NewSchema("ListOffsetsRequestv1",
			&schema.Mfield{Name: FieldListOffsetsRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldListOffsetsRequestTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldListOffsetsRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV1",
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: ListOffsetsRequest, API Key: 2, Version: 2
		schema.NewSchema("ListOffsetsRequestv2",
			&schema.Mfield{Name: FieldListOffsetsRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldListOffsetsRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.Array{Name: FieldListOffsetsRequestTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldListOffsetsRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV2",
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: ListOffsetsRequest, API Key: 2, Version: 3
		schema.NewSchema("ListOffsetsRequestv3",
			&schema.Mfield{Name: FieldListOffsetsRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldListOffsetsRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.Array{Name: FieldListOffsetsRequestTopics, Ty: schema.NewSchema("TopicsV3",
				&schema.Mfield{Name: FieldListOffsetsRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV3",
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: ListOffsetsRequest, API Key: 2, Version: 4
		schema.NewSchema("ListOffsetsRequestv4",
			&schema.Mfield{Name: FieldListOffsetsRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldListOffsetsRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.Array{Name: FieldListOffsetsRequestTopics, Ty: schema.NewSchema("TopicsV4",
				&schema.Mfield{Name: FieldListOffsetsRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV4",
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsCurrentLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: ListOffsetsRequest, API Key: 2, Version: 5
		schema.NewSchema("ListOffsetsRequestv5",
			&schema.Mfield{Name: FieldListOffsetsRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldListOffsetsRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.Array{Name: FieldListOffsetsRequestTopics, Ty: schema.NewSchema("TopicsV5",
				&schema.Mfield{Name: FieldListOffsetsRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV5",
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsCurrentLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: ListOffsetsRequest, API Key: 2, Version: 6
		schema.NewSchema("ListOffsetsRequestv6",
			&schema.Mfield{Name: FieldListOffsetsRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldListOffsetsRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.ArrayCompact{Name: FieldListOffsetsRequestTopics, Ty: schema.NewSchema("TopicsV6",
				&schema.Mfield{Name: FieldListOffsetsRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldListOffsetsRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV6",
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsCurrentLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
					&schema.SchemaTaggedFields{Name: FieldListOffsetsRequestTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldListOffsetsRequestTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldListOffsetsRequestTags},
		),

		// Message: ListOffsetsRequest, API Key: 2, Version: 7
		schema.NewSchema("ListOffsetsRequestv7",
			&schema.Mfield{Name: FieldListOffsetsRequestReplicaId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldListOffsetsRequestIsolationLevel, Ty: schema.TypeInt8},
			&schema.ArrayCompact{Name: FieldListOffsetsRequestTopics, Ty: schema.NewSchema("TopicsV7",
				&schema.Mfield{Name: FieldListOffsetsRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldListOffsetsRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV7",
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsCurrentLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsRequestTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
					&schema.SchemaTaggedFields{Name: FieldListOffsetsRequestTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldListOffsetsRequestTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldListOffsetsRequestTags},
		),
	}
}

const (
	// FieldListOffsetsRequestReplicaId is: The broker ID of the requestor, or -1 if this request is being made by a normal consumer.
	FieldListOffsetsRequestReplicaId = "ReplicaId"
	// FieldListOffsetsRequestTopicsName is: The topic name.
	FieldListOffsetsRequestTopicsName = "Name"
	// FieldListOffsetsRequestTopicsPartitionsTags is: The tagged fields.
	FieldListOffsetsRequestTopicsPartitionsTags = "Tags"
	// FieldListOffsetsRequestTopicsTags is: The tagged fields.
	FieldListOffsetsRequestTopicsTags = "Tags"
	// FieldListOffsetsRequestTopics is: Each topic in the request.
	FieldListOffsetsRequestTopics = "Topics"
	// FieldListOffsetsRequestTopicsPartitions is: Each partition in the request.
	FieldListOffsetsRequestTopicsPartitions = "Partitions"
	// FieldListOffsetsRequestTopicsPartitionsPartitionIndex is: The partition index.
	FieldListOffsetsRequestTopicsPartitionsPartitionIndex = "PartitionIndex"
	// FieldListOffsetsRequestTopicsPartitionsTimestamp is: The current timestamp.
	FieldListOffsetsRequestTopicsPartitionsTimestamp = "Timestamp"
	// FieldListOffsetsRequestIsolationLevel is: This setting controls the visibility of transactional records. Using READ_UNCOMMITTED (isolation_level = 0) makes all records visible. With READ_COMMITTED (isolation_level = 1), non-transactional and COMMITTED transactional records are visible. To be more concrete, READ_COMMITTED returns all data from offsets smaller than the current LSO (last stable offset), and enables the inclusion of the list of aborted transactions in the result, which allows consumers to discard ABORTED transactional records
	FieldListOffsetsRequestIsolationLevel = "IsolationLevel"
	// FieldListOffsetsRequestTopicsPartitionsCurrentLeaderEpoch is: The current leader epoch.
	FieldListOffsetsRequestTopicsPartitionsCurrentLeaderEpoch = "CurrentLeaderEpoch"
	// FieldListOffsetsRequestTags is: The tagged fields.
	FieldListOffsetsRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ListOffsetsRequest.json
const originalListOffsetsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 2,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "ListOffsetsRequest",
  // Version 1 removes MaxNumOffsets.  From this version forward, only a single
  // offset can be returned.
  //
  // Version 2 adds the isolation level, which is used for transactional reads.
  //
  // Version 3 is the same as version 2.
  //
  // Version 4 adds the current leader epoch, which is used for fencing.
  //
  // Version 5 is the same as version 4.
  //
  // Version 6 enables flexible versions.
  //
  // Version 7 enables listing offsets by max timestamp (KIP-734).
  "validVersions": "0-7",
  "flexibleVersions": "6+",
  "fields": [
    { "name": "ReplicaId", "type": "int32", "versions": "0+", "entityType": "brokerId",
      "about": "The broker ID of the requestor, or -1 if this request is being made by a normal consumer." },
    { "name": "IsolationLevel", "type": "int8", "versions": "2+",
      "about": "This setting controls the visibility of transactional records. Using READ_UNCOMMITTED (isolation_level = 0) makes all records visible. With READ_COMMITTED (isolation_level = 1), non-transactional and COMMITTED transactional records are visible. To be more concrete, READ_COMMITTED returns all data from offsets smaller than the current LSO (last stable offset), and enables the inclusion of the list of aborted transactions in the result, which allows consumers to discard ABORTED transactional records" },
    { "name": "Topics", "type": "[]ListOffsetsTopic", "versions": "0+",
      "about": "Each topic in the request.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "Partitions", "type": "[]ListOffsetsPartition", "versions": "0+",
        "about": "Each partition in the request.", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "CurrentLeaderEpoch", "type": "int32", "versions": "4+", "default": "-1", "ignorable": true,
          "about": "The current leader epoch." },
        { "name": "Timestamp", "type": "int64", "versions": "0+",
          "about": "The current timestamp." },
        { "name": "MaxNumOffsets", "type": "int32", "versions": "0", "default": "1",
          "about": "The maximum number of offsets to report." }
      ]}
    ]}
  ]
}
`
