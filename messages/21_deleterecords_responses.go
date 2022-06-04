package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init21DeleteRecordsResponse() []schema.Schema {

	return []schema.Schema{

		// Message: DeleteRecordsResponse, API Key: 21, Version: 0
		schema.NewSchema("DeleteRecordsResponsev0",
			&schema.Mfield{Name: FieldDeleteRecordsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDeleteRecordsResponseTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldDeleteRecordsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDeleteRecordsResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV0",
					&schema.Mfield{Name: FieldDeleteRecordsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldDeleteRecordsResponseTopicsPartitionsLowWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldDeleteRecordsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: DeleteRecordsResponse, API Key: 21, Version: 1
		schema.NewSchema("DeleteRecordsResponsev1",
			&schema.Mfield{Name: FieldDeleteRecordsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDeleteRecordsResponseTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldDeleteRecordsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDeleteRecordsResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV1",
					&schema.Mfield{Name: FieldDeleteRecordsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldDeleteRecordsResponseTopicsPartitionsLowWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldDeleteRecordsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: DeleteRecordsResponse, API Key: 21, Version: 2
		schema.NewSchema("DeleteRecordsResponsev2",
			&schema.Mfield{Name: FieldDeleteRecordsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldDeleteRecordsResponseTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldDeleteRecordsResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldDeleteRecordsResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV2",
					&schema.Mfield{Name: FieldDeleteRecordsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldDeleteRecordsResponseTopicsPartitionsLowWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldDeleteRecordsResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.SchemaTaggedFields{Name: FieldDeleteRecordsResponseTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldDeleteRecordsResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDeleteRecordsResponseTags},
		),
	}
}

const (
	// FieldDeleteRecordsResponseTopicsPartitionsPartitionIndex is: The partition index.
	FieldDeleteRecordsResponseTopicsPartitionsPartitionIndex = "PartitionIndex"
	// FieldDeleteRecordsResponseTopicsPartitionsErrorCode is: The deletion error code, or 0 if the deletion succeeded.
	FieldDeleteRecordsResponseTopicsPartitionsErrorCode = "ErrorCode"
	// FieldDeleteRecordsResponseTopicsPartitionsTags is: The tagged fields.
	FieldDeleteRecordsResponseTopicsPartitionsTags = "Tags"
	// FieldDeleteRecordsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldDeleteRecordsResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldDeleteRecordsResponseTopicsName is: The topic name.
	FieldDeleteRecordsResponseTopicsName = "Name"
	// FieldDeleteRecordsResponseTopicsPartitions is: Each partition that we wanted to delete records from.
	FieldDeleteRecordsResponseTopicsPartitions = "Partitions"
	// FieldDeleteRecordsResponseTopicsPartitionsLowWatermark is: The partition low water mark.
	FieldDeleteRecordsResponseTopicsPartitionsLowWatermark = "LowWatermark"
	// FieldDeleteRecordsResponseTopicsTags is: The tagged fields.
	FieldDeleteRecordsResponseTopicsTags = "Tags"
	// FieldDeleteRecordsResponseTags is: The tagged fields.
	FieldDeleteRecordsResponseTags = "Tags"
	// FieldDeleteRecordsResponseTopics is: Each topic that we wanted to delete records from.
	FieldDeleteRecordsResponseTopics = "Topics"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DeleteRecordsResponse.json
const originalDeleteRecordsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 21,
  "type": "response",
  "name": "DeleteRecordsResponse",
  // Starting in version 1, on quota violation, brokers send out responses before throttling.

  // Version 2 is the first flexible version.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Topics", "type": "[]DeleteRecordsTopicResult", "versions": "0+",
      "about": "Each topic that we wanted to delete records from.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "mapKey": true, "entityType": "topicName",
        "about": "The topic name." },
      { "name": "Partitions", "type": "[]DeleteRecordsPartitionResult", "versions": "0+",
        "about": "Each partition that we wanted to delete records from.", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+", "mapKey": true,
          "about": "The partition index." },
        { "name": "LowWatermark", "type": "int64", "versions": "0+",
          "about": "The partition low water mark." },
        { "name": "ErrorCode", "type": "int16", "versions": "0+",
          "about": "The deletion error code, or 0 if the deletion succeeded." }
      ]}
    ]}
  ]
}
`
