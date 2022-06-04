package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init61DescribeProducersResponse() []schema.Schema {

	return []schema.Schema{

		// Message: DescribeProducersResponse, API Key: 61, Version: 0
		schema.NewSchema("DescribeProducersResponsev0",
			&schema.Mfield{Name: FieldDescribeProducersResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldDescribeProducersResponseTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldDescribeProducersResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldDescribeProducersResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV0",
					&schema.Mfield{Name: FieldDescribeProducersResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldDescribeProducersResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldDescribeProducersResponseTopicsPartitionsErrorMessage, Ty: schema.TypeStrCompactNullable},
					&schema.ArrayCompact{Name: FieldDescribeProducersResponseTopicsPartitionsActiveProducers, Ty: schema.NewSchema("ActiveProducersV0",
						&schema.Mfield{Name: FieldDescribeProducersResponseTopicsPartitionsActiveProducersProducerId, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldDescribeProducersResponseTopicsPartitionsActiveProducersProducerEpoch, Ty: schema.TypeInt32},
						&schema.Mfield{Name: FieldDescribeProducersResponseTopicsPartitionsActiveProducersLastSequence, Ty: schema.TypeInt32},
						&schema.Mfield{Name: FieldDescribeProducersResponseTopicsPartitionsActiveProducersLastTimestamp, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldDescribeProducersResponseTopicsPartitionsActiveProducersCoordinatorEpoch, Ty: schema.TypeInt32},
						&schema.Mfield{Name: FieldDescribeProducersResponseTopicsPartitionsActiveProducersCurrentTxnStartOffset, Ty: schema.TypeInt64},
						&schema.SchemaTaggedFields{Name: FieldDescribeProducersResponseTopicsPartitionsActiveProducersTags},
					)},
					&schema.SchemaTaggedFields{Name: FieldDescribeProducersResponseTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldDescribeProducersResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDescribeProducersResponseTags},
		),
	}
}

const (
	// FieldDescribeProducersResponseTags is: The tagged fields.
	FieldDescribeProducersResponseTags = "Tags"
	// FieldDescribeProducersResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldDescribeProducersResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldDescribeProducersResponseTopics is: Each topic in the response.
	FieldDescribeProducersResponseTopics = "Topics"
	// FieldDescribeProducersResponseTopicsName is: The topic name
	FieldDescribeProducersResponseTopicsName = "Name"
	// FieldDescribeProducersResponseTopicsPartitions is: Each partition in the response.
	FieldDescribeProducersResponseTopicsPartitions = "Partitions"
	// FieldDescribeProducersResponseTopicsPartitionsActiveProducers is:
	FieldDescribeProducersResponseTopicsPartitionsActiveProducers = "ActiveProducers"
	// FieldDescribeProducersResponseTopicsPartitionsActiveProducersCoordinatorEpoch is:
	FieldDescribeProducersResponseTopicsPartitionsActiveProducersCoordinatorEpoch = "CoordinatorEpoch"
	// FieldDescribeProducersResponseTopicsPartitionsActiveProducersCurrentTxnStartOffset is:
	FieldDescribeProducersResponseTopicsPartitionsActiveProducersCurrentTxnStartOffset = "CurrentTxnStartOffset"
	// FieldDescribeProducersResponseTopicsPartitionsActiveProducersLastSequence is:
	FieldDescribeProducersResponseTopicsPartitionsActiveProducersLastSequence = "LastSequence"
	// FieldDescribeProducersResponseTopicsPartitionsActiveProducersLastTimestamp is:
	FieldDescribeProducersResponseTopicsPartitionsActiveProducersLastTimestamp = "LastTimestamp"
	// FieldDescribeProducersResponseTopicsPartitionsActiveProducersProducerEpoch is:
	FieldDescribeProducersResponseTopicsPartitionsActiveProducersProducerEpoch = "ProducerEpoch"
	// FieldDescribeProducersResponseTopicsPartitionsActiveProducersProducerId is:
	FieldDescribeProducersResponseTopicsPartitionsActiveProducersProducerId = "ProducerId"
	// FieldDescribeProducersResponseTopicsPartitionsActiveProducersTags is: The tagged fields.
	FieldDescribeProducersResponseTopicsPartitionsActiveProducersTags = "Tags"
	// FieldDescribeProducersResponseTopicsPartitionsErrorCode is: The partition error code, or 0 if there was no error.
	FieldDescribeProducersResponseTopicsPartitionsErrorCode = "ErrorCode"
	// FieldDescribeProducersResponseTopicsPartitionsErrorMessage is: The partition error message, which may be null if no additional details are available
	FieldDescribeProducersResponseTopicsPartitionsErrorMessage = "ErrorMessage"
	// FieldDescribeProducersResponseTopicsPartitionsPartitionIndex is: The partition index.
	FieldDescribeProducersResponseTopicsPartitionsPartitionIndex = "PartitionIndex"
	// FieldDescribeProducersResponseTopicsPartitionsTags is: The tagged fields.
	FieldDescribeProducersResponseTopicsPartitionsTags = "Tags"
	// FieldDescribeProducersResponseTopicsTags is: The tagged fields.
	FieldDescribeProducersResponseTopicsTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeProducersResponse.json
const originalDescribeProducersResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 61,
  "type": "response",
  "name": "DescribeProducersResponse",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Topics", "type": "[]TopicResponse", "versions": "0+",
      "about": "Each topic in the response.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name" },
      { "name": "Partitions", "type": "[]PartitionResponse", "versions": "0+",
        "about": "Each partition in the response.", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "ErrorCode", "type": "int16", "versions": "0+",
          "about": "The partition error code, or 0 if there was no error." },
        { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+", "default": "null",
          "about": "The partition error message, which may be null if no additional details are available" },
        { "name": "ActiveProducers", "type": "[]ProducerState", "versions": "0+", "fields": [
          { "name": "ProducerId", "type": "int64", "versions": "0+", "entityType": "producerId" },
          { "name": "ProducerEpoch", "type": "int32", "versions": "0+" },
          { "name": "LastSequence", "type": "int32", "versions": "0+", "default":  "-1" },
          { "name": "LastTimestamp", "type": "int64", "versions": "0+", "default": "-1" },
          { "name": "CoordinatorEpoch", "type":  "int32", "versions":  "0+" },
          { "name": "CurrentTxnStartOffset", "type": "int64", "versions": "0+", "default": "-1" }
        ]}
      ]}
    ]}
  ]
}
`
