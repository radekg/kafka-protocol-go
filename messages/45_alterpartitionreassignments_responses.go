package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init45AlterPartitionReassignmentsResponse() []schema.Schema {

	return []schema.Schema{
		// Message: AlterPartitionReassignmentsResponse, API Key: 45, Version: 0
		schema.NewSchema("AlterPartitionReassignmentsResponse:v0",
			&schema.Mfield{Name: FieldAlterPartitionReassignmentsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldAlterPartitionReassignmentsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldAlterPartitionReassignmentsResponseErrorMessage, Ty: schema.TypeStrCompactNullable},
			&schema.ArrayCompact{Name: FieldAlterPartitionReassignmentsResponseResponses, Ty: schema.NewSchema("Responses:v0",
				&schema.Mfield{Name: FieldAlterPartitionReassignmentsResponseResponsesName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldAlterPartitionReassignmentsResponseResponsesPartitions, Ty: schema.NewSchema("Partitions:v0",
					&schema.Mfield{Name: FieldAlterPartitionReassignmentsResponseResponsesPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAlterPartitionReassignmentsResponseResponsesPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldAlterPartitionReassignmentsResponseResponsesPartitionsErrorMessage, Ty: schema.TypeStrCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldAlterPartitionReassignmentsResponseResponsesPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldAlterPartitionReassignmentsResponseResponsesTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldAlterPartitionReassignmentsResponseTags},
		),
	}

}

const (

	// FieldAlterPartitionReassignmentsResponseErrorCode is: The top-level error code, or 0 if there was no error.
	FieldAlterPartitionReassignmentsResponseErrorCode = "ErrorCode"

	// FieldAlterPartitionReassignmentsResponseErrorMessage is: The top-level error message, or null if there was no error.
	FieldAlterPartitionReassignmentsResponseErrorMessage = "ErrorMessage"

	// FieldAlterPartitionReassignmentsResponseResponses is: The responses to topics to reassign.
	FieldAlterPartitionReassignmentsResponseResponses = "Responses"

	// FieldAlterPartitionReassignmentsResponseResponsesName is: The topic name
	FieldAlterPartitionReassignmentsResponseResponsesName = "Name"

	// FieldAlterPartitionReassignmentsResponseResponsesPartitions is: The responses to partitions to reassign
	FieldAlterPartitionReassignmentsResponseResponsesPartitions = "Partitions"

	// FieldAlterPartitionReassignmentsResponseResponsesPartitionsErrorCode is: The error code for this partition, or 0 if there was no error.
	FieldAlterPartitionReassignmentsResponseResponsesPartitionsErrorCode = "ErrorCode"

	// FieldAlterPartitionReassignmentsResponseResponsesPartitionsErrorMessage is: The error message for this partition, or null if there was no error.
	FieldAlterPartitionReassignmentsResponseResponsesPartitionsErrorMessage = "ErrorMessage"

	// FieldAlterPartitionReassignmentsResponseResponsesPartitionsPartitionIndex is: The partition index.
	FieldAlterPartitionReassignmentsResponseResponsesPartitionsPartitionIndex = "PartitionIndex"

	// FieldAlterPartitionReassignmentsResponseResponsesPartitionsTags is: The tagged fields.
	FieldAlterPartitionReassignmentsResponseResponsesPartitionsTags = "Tags"

	// FieldAlterPartitionReassignmentsResponseResponsesTags is: The tagged fields.
	FieldAlterPartitionReassignmentsResponseResponsesTags = "Tags"

	// FieldAlterPartitionReassignmentsResponseTags is: The tagged fields.
	FieldAlterPartitionReassignmentsResponseTags = "Tags"

	// FieldAlterPartitionReassignmentsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldAlterPartitionReassignmentsResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AlterPartitionReassignmentsResponse.json
const originalAlterPartitionReassignmentsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "response",
  "name": "AlterPartitionReassignmentsResponse",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The top-level error code, or 0 if there was no error." },
    { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
      "about": "The top-level error message, or null if there was no error." },
    { "name": "Responses", "type": "[]ReassignableTopicResponse", "versions": "0+",
      "about": "The responses to topics to reassign.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name" },
      { "name": "Partitions", "type": "[]ReassignablePartitionResponse", "versions": "0+",
        "about": "The responses to partitions to reassign", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "ErrorCode", "type": "int16", "versions": "0+",
          "about": "The error code for this partition, or 0 if there was no error." },
        { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
          "about": "The error message for this partition, or null if there was no error." }
      ]}
    ]}
  ]
}
`
