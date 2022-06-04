package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init43ElectLeadersResponse() []schema.Schema {

	return []schema.Schema{

		// Message: ElectLeadersResponse, API Key: 43, Version: 0
		schema.NewSchema("ElectLeadersResponsev0",
			&schema.Mfield{Name: FieldElectLeadersResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldElectLeadersResponseReplicaElectionResults, Ty: schema.NewSchema("ReplicaElectionResultsV0",
				&schema.Mfield{Name: FieldElectLeadersResponseReplicaElectionResultsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldElectLeadersResponseReplicaElectionResultsPartitionResult, Ty: schema.NewSchema("PartitionResultV0",
					&schema.Mfield{Name: FieldElectLeadersResponseReplicaElectionResultsPartitionResultPartitionId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldElectLeadersResponseReplicaElectionResultsPartitionResultErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldElectLeadersResponseReplicaElectionResultsPartitionResultErrorMessage, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: ElectLeadersResponse, API Key: 43, Version: 1
		schema.NewSchema("ElectLeadersResponsev1",
			&schema.Mfield{Name: FieldElectLeadersResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldElectLeadersResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldElectLeadersResponseReplicaElectionResults, Ty: schema.NewSchema("ReplicaElectionResultsV1",
				&schema.Mfield{Name: FieldElectLeadersResponseReplicaElectionResultsTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldElectLeadersResponseReplicaElectionResultsPartitionResult, Ty: schema.NewSchema("PartitionResultV1",
					&schema.Mfield{Name: FieldElectLeadersResponseReplicaElectionResultsPartitionResultPartitionId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldElectLeadersResponseReplicaElectionResultsPartitionResultErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldElectLeadersResponseReplicaElectionResultsPartitionResultErrorMessage, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: ElectLeadersResponse, API Key: 43, Version: 2
		schema.NewSchema("ElectLeadersResponsev2",
			&schema.Mfield{Name: FieldElectLeadersResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldElectLeadersResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldElectLeadersResponseReplicaElectionResults, Ty: schema.NewSchema("ReplicaElectionResultsV2",
				&schema.Mfield{Name: FieldElectLeadersResponseReplicaElectionResultsTopic, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldElectLeadersResponseReplicaElectionResultsPartitionResult, Ty: schema.NewSchema("PartitionResultV2",
					&schema.Mfield{Name: FieldElectLeadersResponseReplicaElectionResultsPartitionResultPartitionId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldElectLeadersResponseReplicaElectionResultsPartitionResultErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldElectLeadersResponseReplicaElectionResultsPartitionResultErrorMessage, Ty: schema.TypeStrCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldElectLeadersResponseReplicaElectionResultsPartitionResultTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldElectLeadersResponseReplicaElectionResultsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldElectLeadersResponseTags},
		),
	}
}

const (
	// FieldElectLeadersResponseReplicaElectionResultsPartitionResultErrorCode is: The result error, or zero if there was no error.
	FieldElectLeadersResponseReplicaElectionResultsPartitionResultErrorCode = "ErrorCode"
	// FieldElectLeadersResponseErrorCode is: The top level response error code.
	FieldElectLeadersResponseErrorCode = "ErrorCode"
	// FieldElectLeadersResponseTags is: The tagged fields.
	FieldElectLeadersResponseTags = "Tags"
	// FieldElectLeadersResponseReplicaElectionResults is: The election results, or an empty array if the requester did not have permission and the request asks for all partitions.
	FieldElectLeadersResponseReplicaElectionResults = "ReplicaElectionResults"
	// FieldElectLeadersResponseReplicaElectionResultsPartitionResult is: The results for each partition
	FieldElectLeadersResponseReplicaElectionResultsPartitionResult = "PartitionResult"
	// FieldElectLeadersResponseReplicaElectionResultsPartitionResultPartitionId is: The partition id
	FieldElectLeadersResponseReplicaElectionResultsPartitionResultPartitionId = "PartitionId"
	// FieldElectLeadersResponseReplicaElectionResultsPartitionResultErrorMessage is: The result message, or null if there was no error.
	FieldElectLeadersResponseReplicaElectionResultsPartitionResultErrorMessage = "ErrorMessage"
	// FieldElectLeadersResponseReplicaElectionResultsPartitionResultTags is: The tagged fields.
	FieldElectLeadersResponseReplicaElectionResultsPartitionResultTags = "Tags"
	// FieldElectLeadersResponseReplicaElectionResultsTags is: The tagged fields.
	FieldElectLeadersResponseReplicaElectionResultsTags = "Tags"
	// FieldElectLeadersResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldElectLeadersResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldElectLeadersResponseReplicaElectionResultsTopic is: The topic name
	FieldElectLeadersResponseReplicaElectionResultsTopic = "Topic"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ElectLeadersResponse.json
const originalElectLeadersResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 43,
  "type": "response",
  "name": "ElectLeadersResponse",
  // Version 1 adds a top-level error code.
  //
  // Version 2 is the first flexible version.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "1+", "ignorable": false,
      "about": "The top level response error code." },
    { "name": "ReplicaElectionResults", "type": "[]ReplicaElectionResult", "versions": "0+",
      "about": "The election results, or an empty array if the requester did not have permission and the request asks for all partitions.", "fields": [
      { "name": "Topic", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name" },
      { "name": "PartitionResult", "type": "[]PartitionResult", "versions": "0+",
        "about": "The results for each partition", "fields": [
        { "name": "PartitionId", "type": "int32", "versions": "0+",
          "about": "The partition id" },
        { "name": "ErrorCode", "type": "int16", "versions": "0+",
          "about": "The result error, or zero if there was no error."},
        { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
          "about": "The result message, or null if there was no error."}
      ]}
    ]}
  ]
}
`
