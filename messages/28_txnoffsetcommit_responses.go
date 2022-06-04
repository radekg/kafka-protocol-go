package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init28TxnOffsetCommitResponse() []schema.Schema {

	return []schema.Schema{

		// Message: TxnOffsetCommitResponse, API Key: 28, Version: 0
		schema.NewSchema("TxnOffsetCommitResponsev0",
			&schema.Mfield{Name: FieldTxnOffsetCommitResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldTxnOffsetCommitResponseTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldTxnOffsetCommitResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldTxnOffsetCommitResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV0",
					&schema.Mfield{Name: FieldTxnOffsetCommitResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldTxnOffsetCommitResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: TxnOffsetCommitResponse, API Key: 28, Version: 1
		schema.NewSchema("TxnOffsetCommitResponsev1",
			&schema.Mfield{Name: FieldTxnOffsetCommitResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldTxnOffsetCommitResponseTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldTxnOffsetCommitResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldTxnOffsetCommitResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV1",
					&schema.Mfield{Name: FieldTxnOffsetCommitResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldTxnOffsetCommitResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: TxnOffsetCommitResponse, API Key: 28, Version: 2
		schema.NewSchema("TxnOffsetCommitResponsev2",
			&schema.Mfield{Name: FieldTxnOffsetCommitResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldTxnOffsetCommitResponseTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldTxnOffsetCommitResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldTxnOffsetCommitResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV2",
					&schema.Mfield{Name: FieldTxnOffsetCommitResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldTxnOffsetCommitResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: TxnOffsetCommitResponse, API Key: 28, Version: 3
		schema.NewSchema("TxnOffsetCommitResponsev3",
			&schema.Mfield{Name: FieldTxnOffsetCommitResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldTxnOffsetCommitResponseTopics, Ty: schema.NewSchema("TopicsV3",
				&schema.Mfield{Name: FieldTxnOffsetCommitResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldTxnOffsetCommitResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV3",
					&schema.Mfield{Name: FieldTxnOffsetCommitResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldTxnOffsetCommitResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.SchemaTaggedFields{Name: FieldTxnOffsetCommitResponseTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldTxnOffsetCommitResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldTxnOffsetCommitResponseTags},
		),
	}
}

const (
	// FieldTxnOffsetCommitResponseTags is: The tagged fields.
	FieldTxnOffsetCommitResponseTags = "Tags"
	// FieldTxnOffsetCommitResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldTxnOffsetCommitResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldTxnOffsetCommitResponseTopics is: The responses for each topic.
	FieldTxnOffsetCommitResponseTopics = "Topics"
	// FieldTxnOffsetCommitResponseTopicsName is: The topic name.
	FieldTxnOffsetCommitResponseTopicsName = "Name"
	// FieldTxnOffsetCommitResponseTopicsPartitions is: The responses for each partition in the topic.
	FieldTxnOffsetCommitResponseTopicsPartitions = "Partitions"
	// FieldTxnOffsetCommitResponseTopicsPartitionsErrorCode is: The error code, or 0 if there was no error.
	FieldTxnOffsetCommitResponseTopicsPartitionsErrorCode = "ErrorCode"
	// FieldTxnOffsetCommitResponseTopicsPartitionsPartitionIndex is: The partition index.
	FieldTxnOffsetCommitResponseTopicsPartitionsPartitionIndex = "PartitionIndex"
	// FieldTxnOffsetCommitResponseTopicsPartitionsTags is: The tagged fields.
	FieldTxnOffsetCommitResponseTopicsPartitionsTags = "Tags"
	// FieldTxnOffsetCommitResponseTopicsTags is: The tagged fields.
	FieldTxnOffsetCommitResponseTopicsTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/TxnOffsetCommitResponse.json
const originalTxnOffsetCommitResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 28,
  "type": "response",
  "name": "TxnOffsetCommitResponse",
  // Starting in version 1, on quota violation, brokers send out responses before throttling.
  //
  // Version 2 is the same as version 1.
  //
  // Version 3 adds illegal generation, fenced instance id, and unknown member id errors.
  "validVersions": "0-3",
  "flexibleVersions": "3+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Topics", "type": "[]TxnOffsetCommitResponseTopic", "versions": "0+",
      "about": "The responses for each topic.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "Partitions", "type": "[]TxnOffsetCommitResponsePartition", "versions": "0+",
        "about": "The responses for each partition in the topic.", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "ErrorCode", "type": "int16", "versions": "0+",
          "about": "The error code, or 0 if there was no error." }
      ]}
    ]}
  ]
}
`
