package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init2ListOffsetsResponse() []schema.Schema {

	return []schema.Schema{

		// Message: ListOffsetsResponse, API Key: 2, Version: 0
		schema.NewSchema("ListOffsetsResponsev0",
			&schema.Array{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV0",
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: ListOffsetsResponse, API Key: 2, Version: 1
		schema.NewSchema("ListOffsetsResponsev1",
			&schema.Array{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV1",
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsOffset, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: ListOffsetsResponse, API Key: 2, Version: 2
		schema.NewSchema("ListOffsetsResponsev2",
			&schema.Mfield{Name: FieldListOffsetsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV2",
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsOffset, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: ListOffsetsResponse, API Key: 2, Version: 3
		schema.NewSchema("ListOffsetsResponsev3",
			&schema.Mfield{Name: FieldListOffsetsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("TopicsV3",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV3",
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsOffset, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: ListOffsetsResponse, API Key: 2, Version: 4
		schema.NewSchema("ListOffsetsResponsev4",
			&schema.Mfield{Name: FieldListOffsetsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("TopicsV4",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV4",
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
				)},
			)},
		),

		// Message: ListOffsetsResponse, API Key: 2, Version: 5
		schema.NewSchema("ListOffsetsResponsev5",
			&schema.Mfield{Name: FieldListOffsetsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("TopicsV5",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV5",
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
				)},
			)},
		),

		// Message: ListOffsetsResponse, API Key: 2, Version: 6
		schema.NewSchema("ListOffsetsResponsev6",
			&schema.Mfield{Name: FieldListOffsetsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("TopicsV6",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV6",
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.SchemaTaggedFields{Name: FieldListOffsetsResponseTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldListOffsetsResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldListOffsetsResponseTags},
		),

		// Message: ListOffsetsResponse, API Key: 2, Version: 7
		schema.NewSchema("ListOffsetsResponsev7",
			&schema.Mfield{Name: FieldListOffsetsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("TopicsV7",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV7",
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.SchemaTaggedFields{Name: FieldListOffsetsResponseTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldListOffsetsResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldListOffsetsResponseTags},
		),
	}
}

const (
	// FieldListOffsetsResponseTopicsPartitions is: Each partition in the response.
	FieldListOffsetsResponseTopicsPartitions = "Partitions"
	// FieldListOffsetsResponseTopicsPartitionsTimestamp is: The timestamp associated with the returned offset.
	FieldListOffsetsResponseTopicsPartitionsTimestamp = "Timestamp"
	// FieldListOffsetsResponseTopicsPartitionsTags is: The tagged fields.
	FieldListOffsetsResponseTopicsPartitionsTags = "Tags"
	// FieldListOffsetsResponseTopicsPartitionsLeaderEpoch is:
	FieldListOffsetsResponseTopicsPartitionsLeaderEpoch = "LeaderEpoch"
	// FieldListOffsetsResponseTopicsTags is: The tagged fields.
	FieldListOffsetsResponseTopicsTags = "Tags"
	// FieldListOffsetsResponseTopics is: Each topic in the response.
	FieldListOffsetsResponseTopics = "Topics"
	// FieldListOffsetsResponseTopicsName is: The topic name
	FieldListOffsetsResponseTopicsName = "Name"
	// FieldListOffsetsResponseTopicsPartitionsPartitionIndex is: The partition index.
	FieldListOffsetsResponseTopicsPartitionsPartitionIndex = "PartitionIndex"
	// FieldListOffsetsResponseTopicsPartitionsErrorCode is: The partition error code, or 0 if there was no error.
	FieldListOffsetsResponseTopicsPartitionsErrorCode = "ErrorCode"
	// FieldListOffsetsResponseTopicsPartitionsOffset is: The returned offset.
	FieldListOffsetsResponseTopicsPartitionsOffset = "Offset"
	// FieldListOffsetsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldListOffsetsResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldListOffsetsResponseTags is: The tagged fields.
	FieldListOffsetsResponseTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ListOffsetsResponse.json
const originalListOffsetsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "response",
  "name": "ListOffsetsResponse",
  // Version 1 removes the offsets array in favor of returning a single offset.
  // Version 1 also adds the timestamp associated with the returned offset.
  //
  // Version 2 adds the throttle time.
  //
  // Starting in version 3, on quota violation, brokers send out responses before throttling.
  //
  // Version 4 adds the leader epoch, which is used for fencing.
  //
  // Version 5 adds a new error code, OFFSET_NOT_AVAILABLE.
  //
  // Version 6 enables flexible versions.
  //
  // Version 7 is the same as version 6 (KIP-734).
  "validVersions": "0-7",
  "flexibleVersions": "6+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "2+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Topics", "type": "[]ListOffsetsTopicResponse", "versions": "0+",
      "about": "Each topic in the response.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name" },
      { "name": "Partitions", "type": "[]ListOffsetsPartitionResponse", "versions": "0+",
        "about": "Each partition in the response.", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "ErrorCode", "type": "int16", "versions": "0+",
          "about": "The partition error code, or 0 if there was no error." },
        { "name": "OldStyleOffsets", "type": "[]int64", "versions": "0", "ignorable": false,
          "about": "The result offsets." },
        { "name": "Timestamp", "type": "int64", "versions": "1+", "default": "-1", "ignorable": false,
          "about": "The timestamp associated with the returned offset." },
        { "name": "Offset", "type": "int64", "versions": "1+", "default": "-1", "ignorable": false,
          "about": "The returned offset." },
        { "name": "LeaderEpoch", "type": "int32", "versions": "4+", "default": "-1" }
      ]}
    ]}
  ]
}
`
