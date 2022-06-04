package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init23OffsetForLeaderEpochResponse() []schema.Schema {

	return []schema.Schema{

		// Message: OffsetForLeaderEpochResponse, API Key: 23, Version: 0
		schema.NewSchema("OffsetForLeaderEpochResponsev0",
			&schema.Array{Name: FieldOffsetForLeaderEpochResponseTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetForLeaderEpochResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV0",
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsEndOffset, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: OffsetForLeaderEpochResponse, API Key: 23, Version: 1
		schema.NewSchema("OffsetForLeaderEpochResponsev1",
			&schema.Array{Name: FieldOffsetForLeaderEpochResponseTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetForLeaderEpochResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV1",
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsEndOffset, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: OffsetForLeaderEpochResponse, API Key: 23, Version: 2
		schema.NewSchema("OffsetForLeaderEpochResponsev2",
			&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldOffsetForLeaderEpochResponseTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetForLeaderEpochResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV2",
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsEndOffset, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: OffsetForLeaderEpochResponse, API Key: 23, Version: 3
		schema.NewSchema("OffsetForLeaderEpochResponsev3",
			&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldOffsetForLeaderEpochResponseTopics, Ty: schema.NewSchema("TopicsV3",
				&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetForLeaderEpochResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV3",
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsEndOffset, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: OffsetForLeaderEpochResponse, API Key: 23, Version: 4
		schema.NewSchema("OffsetForLeaderEpochResponsev4",
			&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldOffsetForLeaderEpochResponseTopics, Ty: schema.NewSchema("TopicsV4",
				&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsTopic, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldOffsetForLeaderEpochResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV4",
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsPartition, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsEndOffset, Ty: schema.TypeInt64},
					&schema.SchemaTaggedFields{Name: FieldOffsetForLeaderEpochResponseTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldOffsetForLeaderEpochResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldOffsetForLeaderEpochResponseTags},
		),
	}
}

const (
	// FieldOffsetForLeaderEpochResponseTopicsTopic is: The topic name.
	FieldOffsetForLeaderEpochResponseTopicsTopic = "Topic"
	// FieldOffsetForLeaderEpochResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldOffsetForLeaderEpochResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldOffsetForLeaderEpochResponseTags is: The tagged fields.
	FieldOffsetForLeaderEpochResponseTags = "Tags"
	// FieldOffsetForLeaderEpochResponseTopicsPartitionsPartition is: The partition index.
	FieldOffsetForLeaderEpochResponseTopicsPartitionsPartition = "Partition"
	// FieldOffsetForLeaderEpochResponseTopicsPartitionsEndOffset is: The end offset of the epoch.
	FieldOffsetForLeaderEpochResponseTopicsPartitionsEndOffset = "EndOffset"
	// FieldOffsetForLeaderEpochResponseTopicsPartitionsLeaderEpoch is: The leader epoch of the partition.
	FieldOffsetForLeaderEpochResponseTopicsPartitionsLeaderEpoch = "LeaderEpoch"
	// FieldOffsetForLeaderEpochResponseTopicsPartitionsTags is: The tagged fields.
	FieldOffsetForLeaderEpochResponseTopicsPartitionsTags = "Tags"
	// FieldOffsetForLeaderEpochResponseTopicsTags is: The tagged fields.
	FieldOffsetForLeaderEpochResponseTopicsTags = "Tags"
	// FieldOffsetForLeaderEpochResponseTopics is: Each topic we fetched offsets for.
	FieldOffsetForLeaderEpochResponseTopics = "Topics"
	// FieldOffsetForLeaderEpochResponseTopicsPartitions is: Each partition in the topic we fetched offsets for.
	FieldOffsetForLeaderEpochResponseTopicsPartitions = "Partitions"
	// FieldOffsetForLeaderEpochResponseTopicsPartitionsErrorCode is: The error code 0, or if there was no error.
	FieldOffsetForLeaderEpochResponseTopicsPartitionsErrorCode = "ErrorCode"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/OffsetForLeaderEpochResponse.json
const originalOffsetForLeaderEpochResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "response",
  "name": "OffsetForLeaderEpochResponse",
  // Version 1 added the leader epoch to the response.
  //
  // Version 2 added the throttle time.
  //
  // Version 3 is the same as version 2.
  //
  // Version 4 enables flexible versions.
  "validVersions": "0-4",
  "flexibleVersions": "4+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "2+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Topics", "type": "[]OffsetForLeaderTopicResult", "versions": "0+",
      "about": "Each topic we fetched offsets for.", "fields": [
      { "name": "Topic", "type": "string", "versions": "0+", "entityType": "topicName",
        "mapKey": true, "about": "The topic name." },
      { "name": "Partitions", "type": "[]EpochEndOffset", "versions": "0+",
        "about": "Each partition in the topic we fetched offsets for.", "fields": [
        { "name": "ErrorCode", "type": "int16", "versions": "0+",
          "about": "The error code 0, or if there was no error." },
        { "name": "Partition", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "LeaderEpoch", "type": "int32", "versions": "1+", "default": "-1", "ignorable": true,
          "about": "The leader epoch of the partition." },
        { "name": "EndOffset", "type": "int64", "versions": "0+", "default": "-1",
          "about": "The end offset of the epoch." }
      ]}
    ]}
  ]
}
`
