package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init14SyncGroupRequest() []schema.Schema {

	return []schema.Schema{
		// Message: SyncGroupRequest, API Key: 14, Version: 0
		schema.NewSchema("SyncGroupRequest:v0",
			&schema.Mfield{Name: FieldSyncGroupRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldSyncGroupRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldSyncGroupRequestMemberId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldSyncGroupRequestAssignments, Ty: schema.NewSchema("Assignments:v0",
				&schema.Mfield{Name: FieldSyncGroupRequestAssignmentsMemberId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldSyncGroupRequestAssignmentsAssignment, Ty: schema.TypeBytes},
			)},
		),

		// Message: SyncGroupRequest, API Key: 14, Version: 1
		schema.NewSchema("SyncGroupRequest:v1",
			&schema.Mfield{Name: FieldSyncGroupRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldSyncGroupRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldSyncGroupRequestMemberId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldSyncGroupRequestAssignments, Ty: schema.NewSchema("Assignments:v1",
				&schema.Mfield{Name: FieldSyncGroupRequestAssignmentsMemberId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldSyncGroupRequestAssignmentsAssignment, Ty: schema.TypeBytes},
			)},
		),

		// Message: SyncGroupRequest, API Key: 14, Version: 2
		schema.NewSchema("SyncGroupRequest:v2",
			&schema.Mfield{Name: FieldSyncGroupRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldSyncGroupRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldSyncGroupRequestMemberId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldSyncGroupRequestAssignments, Ty: schema.NewSchema("Assignments:v2",
				&schema.Mfield{Name: FieldSyncGroupRequestAssignmentsMemberId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldSyncGroupRequestAssignmentsAssignment, Ty: schema.TypeBytes},
			)},
		),

		// Message: SyncGroupRequest, API Key: 14, Version: 3
		schema.NewSchema("SyncGroupRequest:v3",
			&schema.Mfield{Name: FieldSyncGroupRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldSyncGroupRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldSyncGroupRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldSyncGroupRequestGroupInstanceId, Ty: schema.TypeStrNullable},
			&schema.Array{Name: FieldSyncGroupRequestAssignments, Ty: schema.NewSchema("Assignments:v3",
				&schema.Mfield{Name: FieldSyncGroupRequestAssignmentsMemberId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldSyncGroupRequestAssignmentsAssignment, Ty: schema.TypeBytes},
			)},
		),

		// Message: SyncGroupRequest, API Key: 14, Version: 4
		schema.NewSchema("SyncGroupRequest:v4",
			&schema.Mfield{Name: FieldSyncGroupRequestGroupId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldSyncGroupRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldSyncGroupRequestMemberId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldSyncGroupRequestGroupInstanceId, Ty: schema.TypeStrCompactNullable},
			&schema.ArrayCompact{Name: FieldSyncGroupRequestAssignments, Ty: schema.NewSchema("Assignments:v4",
				&schema.Mfield{Name: FieldSyncGroupRequestAssignmentsMemberId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldSyncGroupRequestAssignmentsAssignment, Ty: schema.TypeBytesCompact},
				&schema.SchemaTaggedFields{Name: FieldSyncGroupRequestAssignmentsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldSyncGroupRequestTags},
		),

		// Message: SyncGroupRequest, API Key: 14, Version: 5
		schema.NewSchema("SyncGroupRequest:v5",
			&schema.Mfield{Name: FieldSyncGroupRequestGroupId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldSyncGroupRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldSyncGroupRequestMemberId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldSyncGroupRequestGroupInstanceId, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldSyncGroupRequestProtocolType, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldSyncGroupRequestProtocolName, Ty: schema.TypeStrCompactNullable},
			&schema.ArrayCompact{Name: FieldSyncGroupRequestAssignments, Ty: schema.NewSchema("Assignments:v5",
				&schema.Mfield{Name: FieldSyncGroupRequestAssignmentsMemberId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldSyncGroupRequestAssignmentsAssignment, Ty: schema.TypeBytesCompact},
				&schema.SchemaTaggedFields{Name: FieldSyncGroupRequestAssignmentsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldSyncGroupRequestTags},
		),
	}

}

const (

	// FieldSyncGroupRequestAssignments is: Each assignment.
	FieldSyncGroupRequestAssignments = "Assignments"

	// FieldSyncGroupRequestAssignmentsAssignment is: The member assignment.
	FieldSyncGroupRequestAssignmentsAssignment = "Assignment"

	// FieldSyncGroupRequestAssignmentsMemberId is: The ID of the member to assign.
	FieldSyncGroupRequestAssignmentsMemberId = "MemberId"

	// FieldSyncGroupRequestAssignmentsTags is: The tagged fields.
	FieldSyncGroupRequestAssignmentsTags = "Tags"

	// FieldSyncGroupRequestGenerationId is: The generation of the group.
	FieldSyncGroupRequestGenerationId = "GenerationId"

	// FieldSyncGroupRequestGroupId is: The unique group identifier.
	FieldSyncGroupRequestGroupId = "GroupId"

	// FieldSyncGroupRequestGroupInstanceId is: The unique identifier of the consumer instance provided by end user.
	FieldSyncGroupRequestGroupInstanceId = "GroupInstanceId"

	// FieldSyncGroupRequestMemberId is: The member ID assigned by the group.
	FieldSyncGroupRequestMemberId = "MemberId"

	// FieldSyncGroupRequestProtocolName is: The group protocol name.
	FieldSyncGroupRequestProtocolName = "ProtocolName"

	// FieldSyncGroupRequestProtocolType is: The group protocol type.
	FieldSyncGroupRequestProtocolType = "ProtocolType"

	// FieldSyncGroupRequestTags is: The tagged fields.
	FieldSyncGroupRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/SyncGroupRequest.json
const originalSyncGroupRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 14,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "SyncGroupRequest",
  // Versions 1 and 2 are the same as version 0.
  //
  // Starting from version 3, we add a new field called groupInstanceId to indicate member identity across restarts.
  //
  // Version 4 is the first flexible version.
  //
  // Starting from version 5, the client sends the Protocol Type and the Protocol Name
  // to the broker (KIP-559). The broker will reject the request if they are inconsistent
  // with the Type and Name known by the broker.
  "validVersions": "0-5",
  "flexibleVersions": "4+",
  "fields": [
    { "name": "GroupId", "type": "string", "versions": "0+", "entityType": "groupId",
      "about": "The unique group identifier." },
    { "name": "GenerationId", "type": "int32", "versions": "0+",
      "about": "The generation of the group." },
    { "name": "MemberId", "type": "string", "versions": "0+",
      "about": "The member ID assigned by the group." },
    { "name": "GroupInstanceId", "type": "string", "versions": "3+", 
      "nullableVersions": "3+", "default": "null",
      "about": "The unique identifier of the consumer instance provided by end user." },
    { "name": "ProtocolType", "type": "string", "versions": "5+",
      "nullableVersions": "5+", "default": "null", "ignorable": true,
      "about": "The group protocol type." },
    { "name": "ProtocolName", "type": "string", "versions": "5+",
      "nullableVersions": "5+", "default": "null", "ignorable": true,
      "about": "The group protocol name." },
    { "name": "Assignments", "type": "[]SyncGroupRequestAssignment", "versions": "0+",
      "about": "Each assignment.", "fields": [
      { "name": "MemberId", "type": "string", "versions": "0+",
        "about": "The ID of the member to assign." },
      { "name": "Assignment", "type": "bytes", "versions": "0+",
        "about": "The member assignment." }
    ]}
  ]
}
`
