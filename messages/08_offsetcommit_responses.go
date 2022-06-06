package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init8OffsetCommitResponse() []schema.Schema {

	return []schema.Schema{
		// Message: OffsetCommitResponse, API Key: 8, Version: 0
		schema.NewSchema("OffsetCommitResponse:v0",
			&schema.Array{Name: FieldOffsetCommitResponseTopics, Ty: schema.NewSchema("[]OffsetCommitResponseTopic:v0",
				&schema.Mfield{Name: FieldOffsetCommitResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitResponsePartition:v0",
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: OffsetCommitResponse, API Key: 8, Version: 1
		schema.NewSchema("OffsetCommitResponse:v1",
			&schema.Array{Name: FieldOffsetCommitResponseTopics, Ty: schema.NewSchema("[]OffsetCommitResponseTopic:v1",
				&schema.Mfield{Name: FieldOffsetCommitResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitResponsePartition:v1",
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: OffsetCommitResponse, API Key: 8, Version: 2
		schema.NewSchema("OffsetCommitResponse:v2",
			&schema.Array{Name: FieldOffsetCommitResponseTopics, Ty: schema.NewSchema("[]OffsetCommitResponseTopic:v2",
				&schema.Mfield{Name: FieldOffsetCommitResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitResponsePartition:v2",
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: OffsetCommitResponse, API Key: 8, Version: 3
		schema.NewSchema("OffsetCommitResponse:v3",
			&schema.Mfield{Name: FieldOffsetCommitResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldOffsetCommitResponseTopics, Ty: schema.NewSchema("[]OffsetCommitResponseTopic:v3",
				&schema.Mfield{Name: FieldOffsetCommitResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitResponsePartition:v3",
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: OffsetCommitResponse, API Key: 8, Version: 4
		schema.NewSchema("OffsetCommitResponse:v4",
			&schema.Mfield{Name: FieldOffsetCommitResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldOffsetCommitResponseTopics, Ty: schema.NewSchema("[]OffsetCommitResponseTopic:v4",
				&schema.Mfield{Name: FieldOffsetCommitResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitResponsePartition:v4",
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: OffsetCommitResponse, API Key: 8, Version: 5
		schema.NewSchema("OffsetCommitResponse:v5",
			&schema.Mfield{Name: FieldOffsetCommitResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldOffsetCommitResponseTopics, Ty: schema.NewSchema("[]OffsetCommitResponseTopic:v5",
				&schema.Mfield{Name: FieldOffsetCommitResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitResponsePartition:v5",
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: OffsetCommitResponse, API Key: 8, Version: 6
		schema.NewSchema("OffsetCommitResponse:v6",
			&schema.Mfield{Name: FieldOffsetCommitResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldOffsetCommitResponseTopics, Ty: schema.NewSchema("[]OffsetCommitResponseTopic:v6",
				&schema.Mfield{Name: FieldOffsetCommitResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitResponsePartition:v6",
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: OffsetCommitResponse, API Key: 8, Version: 7
		schema.NewSchema("OffsetCommitResponse:v7",
			&schema.Mfield{Name: FieldOffsetCommitResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldOffsetCommitResponseTopics, Ty: schema.NewSchema("[]OffsetCommitResponseTopic:v7",
				&schema.Mfield{Name: FieldOffsetCommitResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitResponsePartition:v7",
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: OffsetCommitResponse, API Key: 8, Version: 8
		schema.NewSchema("OffsetCommitResponse:v8",
			&schema.Mfield{Name: FieldOffsetCommitResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldOffsetCommitResponseTopics, Ty: schema.NewSchema("[]OffsetCommitResponseTopic:v8",
				&schema.Mfield{Name: FieldOffsetCommitResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldOffsetCommitResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitResponsePartition:v8",
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.SchemaTaggedFields{Name: FieldOffsetCommitResponseTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldOffsetCommitResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldOffsetCommitResponseTags},
		),
	}

}

const (

	// FieldOffsetCommitResponseTags is: The tagged fields.
	FieldOffsetCommitResponseTags = "Tags"

	// FieldOffsetCommitResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldOffsetCommitResponseThrottleTimeMs = "ThrottleTimeMs"

	// FieldOffsetCommitResponseTopics is: The responses for each topic.
	FieldOffsetCommitResponseTopics = "Topics"

	// FieldOffsetCommitResponseTopicsName is: The topic name.
	FieldOffsetCommitResponseTopicsName = "Name"

	// FieldOffsetCommitResponseTopicsPartitions is: The responses for each partition in the topic.
	FieldOffsetCommitResponseTopicsPartitions = "Partitions"

	// FieldOffsetCommitResponseTopicsPartitionsErrorCode is: The error code, or 0 if there was no error.
	FieldOffsetCommitResponseTopicsPartitionsErrorCode = "ErrorCode"

	// FieldOffsetCommitResponseTopicsPartitionsPartitionIndex is: The partition index.
	FieldOffsetCommitResponseTopicsPartitionsPartitionIndex = "PartitionIndex"

	// FieldOffsetCommitResponseTopicsPartitionsTags is: The tagged fields.
	FieldOffsetCommitResponseTopicsPartitionsTags = "Tags"

	// FieldOffsetCommitResponseTopicsTags is: The tagged fields.
	FieldOffsetCommitResponseTopicsTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/OffsetCommitResponse.json
const originalOffsetCommitResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 8,
  "type": "response",
  "name": "OffsetCommitResponse",
  // Versions 1 and 2 are the same as version 0.
  //
  // Version 3 adds the throttle time to the response.
  //
  // Starting in version 4, on quota violation, brokers send out responses before throttling.
  //
  // Versions 5 and 6 are the same as version 4.
  //
  // Version 7 offsetCommitRequest supports a new field called groupInstanceId to indicate member identity across restarts.
  //
  // Version 8 is the first flexible version.
  "validVersions": "0-8",
  "flexibleVersions": "8+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "3+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Topics", "type": "[]OffsetCommitResponseTopic", "versions": "0+",
      "about": "The responses for each topic.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "Partitions", "type": "[]OffsetCommitResponsePartition", "versions": "0+",
        "about": "The responses for each partition in the topic.",  "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "ErrorCode", "type": "int16", "versions": "0+",
          "about": "The error code, or 0 if there was no error." }
      ]}
    ]}
  ]
}
`
