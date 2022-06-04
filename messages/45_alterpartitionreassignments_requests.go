package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init45AlterPartitionReassignmentsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: AlterPartitionReassignmentsRequest, API Key: 45, Version: 0
		schema.NewSchema("AlterPartitionReassignmentsRequestv0",
			&schema.Mfield{Name: FieldAlterPartitionReassignmentsRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldAlterPartitionReassignmentsRequestTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldAlterPartitionReassignmentsRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldAlterPartitionReassignmentsRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV0",
					&schema.Mfield{Name: FieldAlterPartitionReassignmentsRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAlterPartitionReassignmentsRequestTopicsPartitionsReplicas, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldAlterPartitionReassignmentsRequestTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldAlterPartitionReassignmentsRequestTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldAlterPartitionReassignmentsRequestTags},
		),
	}
}

const (
	// FieldAlterPartitionReassignmentsRequestTags is: The tagged fields.
	FieldAlterPartitionReassignmentsRequestTags = "Tags"
	// FieldAlterPartitionReassignmentsRequestTimeoutMs is: The time in ms to wait for the request to complete.
	FieldAlterPartitionReassignmentsRequestTimeoutMs = "TimeoutMs"
	// FieldAlterPartitionReassignmentsRequestTopics is: The topics to reassign.
	FieldAlterPartitionReassignmentsRequestTopics = "Topics"
	// FieldAlterPartitionReassignmentsRequestTopicsName is: The topic name.
	FieldAlterPartitionReassignmentsRequestTopicsName = "Name"
	// FieldAlterPartitionReassignmentsRequestTopicsPartitions is: The partitions to reassign.
	FieldAlterPartitionReassignmentsRequestTopicsPartitions = "Partitions"
	// FieldAlterPartitionReassignmentsRequestTopicsPartitionsPartitionIndex is: The partition index.
	FieldAlterPartitionReassignmentsRequestTopicsPartitionsPartitionIndex = "PartitionIndex"
	// FieldAlterPartitionReassignmentsRequestTopicsPartitionsReplicas is: The replicas to place the partitions on, or null to cancel a pending reassignment for this partition.
	FieldAlterPartitionReassignmentsRequestTopicsPartitionsReplicas = "Replicas"
	// FieldAlterPartitionReassignmentsRequestTopicsPartitionsTags is: The tagged fields.
	FieldAlterPartitionReassignmentsRequestTopicsPartitionsTags = "Tags"
	// FieldAlterPartitionReassignmentsRequestTopicsTags is: The tagged fields.
	FieldAlterPartitionReassignmentsRequestTopicsTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AlterPartitionReassignmentsRequest.json
const originalAlterPartitionReassignmentsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 45,
  "type": "request",
  "listeners": ["broker", "controller", "zkBroker"],
  "name": "AlterPartitionReassignmentsRequest",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "TimeoutMs", "type": "int32", "versions": "0+", "default": "60000",
      "about": "The time in ms to wait for the request to complete." },
    { "name": "Topics", "type": "[]ReassignableTopic", "versions": "0+",
      "about": "The topics to reassign.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "Partitions", "type": "[]ReassignablePartition", "versions": "0+",
        "about": "The partitions to reassign.", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "Replicas", "type": "[]int32", "versions": "0+", "nullableVersions": "0+", "default": "null", "entityType": "brokerId",
          "about": "The replicas to place the partitions on, or null to cancel a pending reassignment for this partition." }
      ]}
    ]}
  ]
}
`
