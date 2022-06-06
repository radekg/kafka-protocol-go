package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init5StopReplicaResponse() []schema.Schema {

	return []schema.Schema{
		// Message: StopReplicaResponse, API Key: 5, Version: 0
		schema.NewSchema("StopReplicaResponse:v0",
			&schema.Mfield{Name: FieldStopReplicaResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldStopReplicaResponsePartitionErrors, Ty: schema.NewSchema("[]StopReplicaPartitionError:v0",
				&schema.Mfield{Name: FieldStopReplicaResponsePartitionErrorsTopicName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldStopReplicaResponsePartitionErrorsPartitionIndex, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldStopReplicaResponsePartitionErrorsErrorCode, Ty: schema.TypeInt16},
			)},
		),

		// Message: StopReplicaResponse, API Key: 5, Version: 1
		schema.NewSchema("StopReplicaResponse:v1",
			&schema.Mfield{Name: FieldStopReplicaResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldStopReplicaResponsePartitionErrors, Ty: schema.NewSchema("[]StopReplicaPartitionError:v1",
				&schema.Mfield{Name: FieldStopReplicaResponsePartitionErrorsTopicName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldStopReplicaResponsePartitionErrorsPartitionIndex, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldStopReplicaResponsePartitionErrorsErrorCode, Ty: schema.TypeInt16},
			)},
		),

		// Message: StopReplicaResponse, API Key: 5, Version: 2
		schema.NewSchema("StopReplicaResponse:v2",
			&schema.Mfield{Name: FieldStopReplicaResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldStopReplicaResponsePartitionErrors, Ty: schema.NewSchema("[]StopReplicaPartitionError:v2",
				&schema.Mfield{Name: FieldStopReplicaResponsePartitionErrorsTopicName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldStopReplicaResponsePartitionErrorsPartitionIndex, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldStopReplicaResponsePartitionErrorsErrorCode, Ty: schema.TypeInt16},
				&schema.SchemaTaggedFields{Name: FieldStopReplicaResponsePartitionErrorsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldStopReplicaResponseTags},
		),

		// Message: StopReplicaResponse, API Key: 5, Version: 3
		schema.NewSchema("StopReplicaResponse:v3",
			&schema.Mfield{Name: FieldStopReplicaResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldStopReplicaResponsePartitionErrors, Ty: schema.NewSchema("[]StopReplicaPartitionError:v3",
				&schema.Mfield{Name: FieldStopReplicaResponsePartitionErrorsTopicName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldStopReplicaResponsePartitionErrorsPartitionIndex, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldStopReplicaResponsePartitionErrorsErrorCode, Ty: schema.TypeInt16},
				&schema.SchemaTaggedFields{Name: FieldStopReplicaResponsePartitionErrorsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldStopReplicaResponseTags},
		),
	}

}

const (

	// FieldStopReplicaResponseErrorCode is: The top-level error code, or 0 if there was no top-level error.
	FieldStopReplicaResponseErrorCode = "ErrorCode"

	// FieldStopReplicaResponsePartitionErrors is: The responses for each partition.
	FieldStopReplicaResponsePartitionErrors = "PartitionErrors"

	// FieldStopReplicaResponsePartitionErrorsErrorCode is: The partition error code, or 0 if there was no partition error.
	FieldStopReplicaResponsePartitionErrorsErrorCode = "ErrorCode"

	// FieldStopReplicaResponsePartitionErrorsPartitionIndex is: The partition index.
	FieldStopReplicaResponsePartitionErrorsPartitionIndex = "PartitionIndex"

	// FieldStopReplicaResponsePartitionErrorsTags is: The tagged fields.
	FieldStopReplicaResponsePartitionErrorsTags = "Tags"

	// FieldStopReplicaResponsePartitionErrorsTopicName is: The topic name.
	FieldStopReplicaResponsePartitionErrorsTopicName = "TopicName"

	// FieldStopReplicaResponseTags is: The tagged fields.
	FieldStopReplicaResponseTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/StopReplicaResponse.json
const originalStopReplicaResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 5,
  "type": "response",
  "name": "StopReplicaResponse",
  // Version 1 is the same as version 0.
  //
  // Version 2 is the first flexible version.
  //
  // Version 3 returns FENCED_LEADER_EPOCH if the epoch of the leader is stale (KIP-570).
  "validVersions": "0-3",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The top-level error code, or 0 if there was no top-level error." },
    { "name": "PartitionErrors", "type": "[]StopReplicaPartitionError", "versions": "0+",
      "about": "The responses for each partition.", "fields": [
      { "name": "TopicName", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "PartitionIndex", "type": "int32", "versions": "0+",
        "about": "The partition index." },
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The partition error code, or 0 if there was no partition error." }
    ]}
  ]
}
`
