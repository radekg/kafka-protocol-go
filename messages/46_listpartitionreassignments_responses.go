package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init46ListPartitionReassignmentsResponse() []schema.Schema {

	return []schema.Schema{

		// Message: ListPartitionReassignmentsResponse, API Key: 46, Version: 0
		schema.NewSchema("ListPartitionReassignmentsResponsev0",
			&schema.Mfield{Name: FieldListPartitionReassignmentsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldListPartitionReassignmentsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldListPartitionReassignmentsResponseErrorMessage, Ty: schema.TypeStrCompactNullable},
			&schema.ArrayCompact{Name: FieldListPartitionReassignmentsResponseTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldListPartitionReassignmentsResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldListPartitionReassignmentsResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV0",
					&schema.Mfield{Name: FieldListPartitionReassignmentsResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldListPartitionReassignmentsResponseTopicsPartitionsReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldListPartitionReassignmentsResponseTopicsPartitionsAddingReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldListPartitionReassignmentsResponseTopicsPartitionsRemovingReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldListPartitionReassignmentsResponseTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldListPartitionReassignmentsResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldListPartitionReassignmentsResponseTags},
		),
	}
}

const (
	// FieldListPartitionReassignmentsResponseErrorCode is: The top-level error code, or 0 if there was no error
	FieldListPartitionReassignmentsResponseErrorCode = "ErrorCode"
	// FieldListPartitionReassignmentsResponseErrorMessage is: The top-level error message, or null if there was no error.
	FieldListPartitionReassignmentsResponseErrorMessage = "ErrorMessage"
	// FieldListPartitionReassignmentsResponseTags is: The tagged fields.
	FieldListPartitionReassignmentsResponseTags = "Tags"
	// FieldListPartitionReassignmentsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldListPartitionReassignmentsResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldListPartitionReassignmentsResponseTopics is: The ongoing reassignments for each topic.
	FieldListPartitionReassignmentsResponseTopics = "Topics"
	// FieldListPartitionReassignmentsResponseTopicsName is: The topic name.
	FieldListPartitionReassignmentsResponseTopicsName = "Name"
	// FieldListPartitionReassignmentsResponseTopicsPartitions is: The ongoing reassignments for each partition.
	FieldListPartitionReassignmentsResponseTopicsPartitions = "Partitions"
	// FieldListPartitionReassignmentsResponseTopicsPartitionsAddingReplicas is: The set of replicas we are currently adding.
	FieldListPartitionReassignmentsResponseTopicsPartitionsAddingReplicas = "AddingReplicas"
	// FieldListPartitionReassignmentsResponseTopicsPartitionsPartitionIndex is: The index of the partition.
	FieldListPartitionReassignmentsResponseTopicsPartitionsPartitionIndex = "PartitionIndex"
	// FieldListPartitionReassignmentsResponseTopicsPartitionsRemovingReplicas is: The set of replicas we are currently removing.
	FieldListPartitionReassignmentsResponseTopicsPartitionsRemovingReplicas = "RemovingReplicas"
	// FieldListPartitionReassignmentsResponseTopicsPartitionsReplicas is: The current replica set.
	FieldListPartitionReassignmentsResponseTopicsPartitionsReplicas = "Replicas"
	// FieldListPartitionReassignmentsResponseTopicsPartitionsTags is: The tagged fields.
	FieldListPartitionReassignmentsResponseTopicsPartitionsTags = "Tags"
	// FieldListPartitionReassignmentsResponseTopicsTags is: The tagged fields.
	FieldListPartitionReassignmentsResponseTopicsTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ListPartitionReassignmentsResponse.json
const originalListPartitionReassignmentsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "response",
  "name": "ListPartitionReassignmentsResponse",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The top-level error code, or 0 if there was no error" },
    { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
      "about": "The top-level error message, or null if there was no error." },
    { "name": "Topics", "type": "[]OngoingTopicReassignment", "versions": "0+",
      "about": "The ongoing reassignments for each topic.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "Partitions", "type": "[]OngoingPartitionReassignment", "versions": "0+",
        "about": "The ongoing reassignments for each partition.", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The index of the partition." },
        { "name": "Replicas", "type": "[]int32", "versions": "0+", "entityType": "brokerId",
          "about": "The current replica set." },
        { "name": "AddingReplicas", "type": "[]int32", "versions": "0+", "entityType": "brokerId",
          "about": "The set of replicas we are currently adding." },
        { "name": "RemovingReplicas", "type": "[]int32", "versions": "0+", "entityType": "brokerId",
          "about": "The set of replicas we are currently removing." }
      ]}
    ]}
  ]
}
`
