package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init46ListPartitionReassignmentsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: ListPartitionReassignmentsRequest, API Key: 46, Version: 0
		schema.NewSchema("ListPartitionReassignmentsRequestv0",
			&schema.Mfield{Name: FieldListPartitionReassignmentsRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldListPartitionReassignmentsRequestTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldListPartitionReassignmentsRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldListPartitionReassignmentsRequestTopicsPartitionIndexes, Ty: schema.TypeInt32CompactArray},
				&schema.SchemaTaggedFields{Name: FieldListPartitionReassignmentsRequestTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldListPartitionReassignmentsRequestTags},
		),
	}
}

const (
	// FieldListPartitionReassignmentsRequestTimeoutMs is: The time in ms to wait for the request to complete.
	FieldListPartitionReassignmentsRequestTimeoutMs = "TimeoutMs"
	// FieldListPartitionReassignmentsRequestTopics is: The topics to list partition reassignments for, or null to list everything.
	FieldListPartitionReassignmentsRequestTopics = "Topics"
	// FieldListPartitionReassignmentsRequestTopicsName is: The topic name
	FieldListPartitionReassignmentsRequestTopicsName = "Name"
	// FieldListPartitionReassignmentsRequestTopicsPartitionIndexes is: The partitions to list partition reassignments for.
	FieldListPartitionReassignmentsRequestTopicsPartitionIndexes = "PartitionIndexes"
	// FieldListPartitionReassignmentsRequestTopicsTags is: The tagged fields.
	FieldListPartitionReassignmentsRequestTopicsTags = "Tags"
	// FieldListPartitionReassignmentsRequestTags is: The tagged fields.
	FieldListPartitionReassignmentsRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ListPartitionReassignmentsRequest.json
const originalListPartitionReassignmentsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 46,
  "type": "request",
  "listeners": ["broker", "controller", "zkBroker"],
  "name": "ListPartitionReassignmentsRequest",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "TimeoutMs", "type": "int32", "versions": "0+", "default": "60000",
      "about": "The time in ms to wait for the request to complete." },
    { "name": "Topics", "type": "[]ListPartitionReassignmentsTopics", "versions": "0+", "nullableVersions": "0+", "default": "null",
      "about": "The topics to list partition reassignments for, or null to list everything.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name" },
      { "name": "PartitionIndexes", "type": "[]int32", "versions": "0+",
        "about": "The partitions to list partition reassignments for." }
    ]}
  ]
}
`
