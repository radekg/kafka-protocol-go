package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init2ListOffsetsResponse() []schema.Schema {

	return []schema.Schema{
		// Message: ListOffsetsResponse, API Key: 2, Version: 0
		schema.NewSchema("ListOffsetsResponse:v0",
			&schema.Array{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("[]ListOffsetsTopicResponse:v0",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("[]ListOffsetsPartitionResponse:v0",
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: ListOffsetsResponse, API Key: 2, Version: 1
		schema.NewSchema("ListOffsetsResponse:v1",
			&schema.Array{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("[]ListOffsetsTopicResponse:v1",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("[]ListOffsetsPartitionResponse:v1",
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsOffset, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: ListOffsetsResponse, API Key: 2, Version: 2
		schema.NewSchema("ListOffsetsResponse:v2",
			&schema.Mfield{Name: FieldListOffsetsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("[]ListOffsetsTopicResponse:v2",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("[]ListOffsetsPartitionResponse:v2",
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsOffset, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: ListOffsetsResponse, API Key: 2, Version: 3
		schema.NewSchema("ListOffsetsResponse:v3",
			&schema.Mfield{Name: FieldListOffsetsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("[]ListOffsetsTopicResponse:v3",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("[]ListOffsetsPartitionResponse:v3",
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsOffset, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: ListOffsetsResponse, API Key: 2, Version: 4
		schema.NewSchema("ListOffsetsResponse:v4",
			&schema.Mfield{Name: FieldListOffsetsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("[]ListOffsetsTopicResponse:v4",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("[]ListOffsetsPartitionResponse:v4",
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
				)},
			)},
		),

		// Message: ListOffsetsResponse, API Key: 2, Version: 5
		schema.NewSchema("ListOffsetsResponse:v5",
			&schema.Mfield{Name: FieldListOffsetsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("[]ListOffsetsTopicResponse:v5",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("[]ListOffsetsPartitionResponse:v5",
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsTimestamp, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldListOffsetsResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
				)},
			)},
		),

		// Message: ListOffsetsResponse, API Key: 2, Version: 6
		schema.NewSchema("ListOffsetsResponse:v6",
			&schema.Mfield{Name: FieldListOffsetsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("[]ListOffsetsTopicResponse:v6",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("[]ListOffsetsPartitionResponse:v6",
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
		schema.NewSchema("ListOffsetsResponse:v7",
			&schema.Mfield{Name: FieldListOffsetsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldListOffsetsResponseTopics, Ty: schema.NewSchema("[]ListOffsetsTopicResponse:v7",
				&schema.Mfield{Name: FieldListOffsetsResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldListOffsetsResponseTopicsPartitions, Ty: schema.NewSchema("[]ListOffsetsPartitionResponse:v7",
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

	// FieldListOffsetsResponseTags is: The tagged fields.
	FieldListOffsetsResponseTags = "Tags"

	// FieldListOffsetsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldListOffsetsResponseThrottleTimeMs = "ThrottleTimeMs"

	// FieldListOffsetsResponseTopics is: Each topic in the response.
	FieldListOffsetsResponseTopics = "Topics"

	// FieldListOffsetsResponseTopicsName is: The topic name
	FieldListOffsetsResponseTopicsName = "Name"

	// FieldListOffsetsResponseTopicsPartitions is: Each partition in the response.
	FieldListOffsetsResponseTopicsPartitions = "Partitions"

	// FieldListOffsetsResponseTopicsPartitionsErrorCode is: The partition error code, or 0 if there was no error.
	FieldListOffsetsResponseTopicsPartitionsErrorCode = "ErrorCode"

	// FieldListOffsetsResponseTopicsPartitionsLeaderEpoch is:
	FieldListOffsetsResponseTopicsPartitionsLeaderEpoch = "LeaderEpoch"

	// FieldListOffsetsResponseTopicsPartitionsOffset is: The returned offset.
	FieldListOffsetsResponseTopicsPartitionsOffset = "Offset"

	// FieldListOffsetsResponseTopicsPartitionsPartitionIndex is: The partition index.
	FieldListOffsetsResponseTopicsPartitionsPartitionIndex = "PartitionIndex"

	// FieldListOffsetsResponseTopicsPartitionsTags is: The tagged fields.
	FieldListOffsetsResponseTopicsPartitionsTags = "Tags"

	// FieldListOffsetsResponseTopicsPartitionsTimestamp is: The timestamp associated with the returned offset.
	FieldListOffsetsResponseTopicsPartitionsTimestamp = "Timestamp"

	// FieldListOffsetsResponseTopicsTags is: The tagged fields.
	FieldListOffsetsResponseTopicsTags = "Tags"
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
