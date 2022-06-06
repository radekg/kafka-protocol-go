package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init21DeleteRecordsRequest() []schema.Schema {

	return []schema.Schema{
		// Message: DeleteRecordsRequest, API Key: 21, Version: 0
		schema.NewSchema("DeleteRecordsRequest:v0",
			&schema.Array{Name: FieldDeleteRecordsRequestTopics, Ty: schema.NewSchema("[]DeleteRecordsTopic:v0",
				&schema.Mfield{Name: FieldDeleteRecordsRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDeleteRecordsRequestTopicsPartitions, Ty: schema.NewSchema("[]DeleteRecordsPartition:v0",
					&schema.Mfield{Name: FieldDeleteRecordsRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldDeleteRecordsRequestTopicsPartitionsOffset, Ty: schema.TypeInt64},
				)},
			)},
			&schema.Mfield{Name: FieldDeleteRecordsRequestTimeoutMs, Ty: schema.TypeInt32},
		),

		// Message: DeleteRecordsRequest, API Key: 21, Version: 1
		schema.NewSchema("DeleteRecordsRequest:v1",
			&schema.Array{Name: FieldDeleteRecordsRequestTopics, Ty: schema.NewSchema("[]DeleteRecordsTopic:v1",
				&schema.Mfield{Name: FieldDeleteRecordsRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDeleteRecordsRequestTopicsPartitions, Ty: schema.NewSchema("[]DeleteRecordsPartition:v1",
					&schema.Mfield{Name: FieldDeleteRecordsRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldDeleteRecordsRequestTopicsPartitionsOffset, Ty: schema.TypeInt64},
				)},
			)},
			&schema.Mfield{Name: FieldDeleteRecordsRequestTimeoutMs, Ty: schema.TypeInt32},
		),

		// Message: DeleteRecordsRequest, API Key: 21, Version: 2
		schema.NewSchema("DeleteRecordsRequest:v2",
			&schema.ArrayCompact{Name: FieldDeleteRecordsRequestTopics, Ty: schema.NewSchema("[]DeleteRecordsTopic:v2",
				&schema.Mfield{Name: FieldDeleteRecordsRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldDeleteRecordsRequestTopicsPartitions, Ty: schema.NewSchema("[]DeleteRecordsPartition:v2",
					&schema.Mfield{Name: FieldDeleteRecordsRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldDeleteRecordsRequestTopicsPartitionsOffset, Ty: schema.TypeInt64},
					&schema.SchemaTaggedFields{Name: FieldDeleteRecordsRequestTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldDeleteRecordsRequestTopicsTags},
			)},
			&schema.Mfield{Name: FieldDeleteRecordsRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldDeleteRecordsRequestTags},
		),
	}

}

const (

	// FieldDeleteRecordsRequestTags is: The tagged fields.
	FieldDeleteRecordsRequestTags = "Tags"

	// FieldDeleteRecordsRequestTimeoutMs is: How long to wait for the deletion to complete, in milliseconds.
	FieldDeleteRecordsRequestTimeoutMs = "TimeoutMs"

	// FieldDeleteRecordsRequestTopics is: Each topic that we want to delete records from.
	FieldDeleteRecordsRequestTopics = "Topics"

	// FieldDeleteRecordsRequestTopicsName is: The topic name.
	FieldDeleteRecordsRequestTopicsName = "Name"

	// FieldDeleteRecordsRequestTopicsPartitions is: Each partition that we want to delete records from.
	FieldDeleteRecordsRequestTopicsPartitions = "Partitions"

	// FieldDeleteRecordsRequestTopicsPartitionsOffset is: The deletion offset.
	FieldDeleteRecordsRequestTopicsPartitionsOffset = "Offset"

	// FieldDeleteRecordsRequestTopicsPartitionsPartitionIndex is: The partition index.
	FieldDeleteRecordsRequestTopicsPartitionsPartitionIndex = "PartitionIndex"

	// FieldDeleteRecordsRequestTopicsPartitionsTags is: The tagged fields.
	FieldDeleteRecordsRequestTopicsPartitionsTags = "Tags"

	// FieldDeleteRecordsRequestTopicsTags is: The tagged fields.
	FieldDeleteRecordsRequestTopicsTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DeleteRecordsRequest.json
const originalDeleteRecordsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "DeleteRecordsRequest",
  // Version 1 is the same as version 0.

  // Version 2 is the first flexible version.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "Topics", "type": "[]DeleteRecordsTopic", "versions": "0+",
      "about": "Each topic that we want to delete records from.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "Partitions", "type": "[]DeleteRecordsPartition", "versions": "0+",
        "about": "Each partition that we want to delete records from.", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "Offset", "type": "int64", "versions": "0+",
          "about": "The deletion offset." }
      ]}
    ]},
    { "name": "TimeoutMs", "type": "int32", "versions": "0+",
      "about": "How long to wait for the deletion to complete, in milliseconds." }
  ]
}
`
