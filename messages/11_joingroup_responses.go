package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init11JoinGroupResponse() []schema.Schema {

	return []schema.Schema{
		// Message: JoinGroupResponse, API Key: 11, Version: 0
		schema.NewSchema("JoinGroupResponse:v0",
			&schema.Mfield{Name: FieldJoinGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldJoinGroupResponseGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseProtocolName, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupResponseLeader, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupResponseMemberId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldJoinGroupResponseMembers, Ty: schema.NewSchema("[]JoinGroupResponseMember:v0",
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMemberId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMetadata, Ty: schema.TypeBytes},
			)},
		),

		// Message: JoinGroupResponse, API Key: 11, Version: 1
		schema.NewSchema("JoinGroupResponse:v1",
			&schema.Mfield{Name: FieldJoinGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldJoinGroupResponseGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseProtocolName, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupResponseLeader, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupResponseMemberId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldJoinGroupResponseMembers, Ty: schema.NewSchema("[]JoinGroupResponseMember:v1",
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMemberId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMetadata, Ty: schema.TypeBytes},
			)},
		),

		// Message: JoinGroupResponse, API Key: 11, Version: 2
		schema.NewSchema("JoinGroupResponse:v2",
			&schema.Mfield{Name: FieldJoinGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldJoinGroupResponseGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseProtocolName, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupResponseLeader, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupResponseMemberId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldJoinGroupResponseMembers, Ty: schema.NewSchema("[]JoinGroupResponseMember:v2",
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMemberId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMetadata, Ty: schema.TypeBytes},
			)},
		),

		// Message: JoinGroupResponse, API Key: 11, Version: 3
		schema.NewSchema("JoinGroupResponse:v3",
			&schema.Mfield{Name: FieldJoinGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldJoinGroupResponseGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseProtocolName, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupResponseLeader, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupResponseMemberId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldJoinGroupResponseMembers, Ty: schema.NewSchema("[]JoinGroupResponseMember:v3",
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMemberId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMetadata, Ty: schema.TypeBytes},
			)},
		),

		// Message: JoinGroupResponse, API Key: 11, Version: 4
		schema.NewSchema("JoinGroupResponse:v4",
			&schema.Mfield{Name: FieldJoinGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldJoinGroupResponseGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseProtocolName, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupResponseLeader, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupResponseMemberId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldJoinGroupResponseMembers, Ty: schema.NewSchema("[]JoinGroupResponseMember:v4",
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMemberId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMetadata, Ty: schema.TypeBytes},
			)},
		),

		// Message: JoinGroupResponse, API Key: 11, Version: 5
		schema.NewSchema("JoinGroupResponse:v5",
			&schema.Mfield{Name: FieldJoinGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldJoinGroupResponseGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseProtocolName, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupResponseLeader, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupResponseMemberId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldJoinGroupResponseMembers, Ty: schema.NewSchema("[]JoinGroupResponseMember:v5",
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMemberId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldJoinGroupResponseMembersGroupInstanceId, Ty: schema.TypeStrNullable},
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMetadata, Ty: schema.TypeBytes},
			)},
		),

		// Message: JoinGroupResponse, API Key: 11, Version: 6
		schema.NewSchema("JoinGroupResponse:v6",
			&schema.Mfield{Name: FieldJoinGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldJoinGroupResponseGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseProtocolName, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldJoinGroupResponseLeader, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldJoinGroupResponseMemberId, Ty: schema.TypeStrCompact},
			&schema.ArrayCompact{Name: FieldJoinGroupResponseMembers, Ty: schema.NewSchema("[]JoinGroupResponseMember:v6",
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMemberId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldJoinGroupResponseMembersGroupInstanceId, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMetadata, Ty: schema.TypeBytesCompact},
				&schema.SchemaTaggedFields{Name: FieldJoinGroupResponseMembersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldJoinGroupResponseTags},
		),

		// Message: JoinGroupResponse, API Key: 11, Version: 7
		schema.NewSchema("JoinGroupResponse:v7",
			&schema.Mfield{Name: FieldJoinGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldJoinGroupResponseGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseProtocolType, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldJoinGroupResponseProtocolName, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldJoinGroupResponseLeader, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldJoinGroupResponseMemberId, Ty: schema.TypeStrCompact},
			&schema.ArrayCompact{Name: FieldJoinGroupResponseMembers, Ty: schema.NewSchema("[]JoinGroupResponseMember:v7",
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMemberId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldJoinGroupResponseMembersGroupInstanceId, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMetadata, Ty: schema.TypeBytesCompact},
				&schema.SchemaTaggedFields{Name: FieldJoinGroupResponseMembersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldJoinGroupResponseTags},
		),

		// Message: JoinGroupResponse, API Key: 11, Version: 8
		schema.NewSchema("JoinGroupResponse:v8",
			&schema.Mfield{Name: FieldJoinGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldJoinGroupResponseGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseProtocolType, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldJoinGroupResponseProtocolName, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldJoinGroupResponseLeader, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldJoinGroupResponseMemberId, Ty: schema.TypeStrCompact},
			&schema.ArrayCompact{Name: FieldJoinGroupResponseMembers, Ty: schema.NewSchema("[]JoinGroupResponseMember:v8",
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMemberId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldJoinGroupResponseMembersGroupInstanceId, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMetadata, Ty: schema.TypeBytesCompact},
				&schema.SchemaTaggedFields{Name: FieldJoinGroupResponseMembersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldJoinGroupResponseTags},
		),

		// Message: JoinGroupResponse, API Key: 11, Version: 9
		schema.NewSchema("JoinGroupResponse:v9",
			&schema.Mfield{Name: FieldJoinGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldJoinGroupResponseGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupResponseProtocolType, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldJoinGroupResponseProtocolName, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldJoinGroupResponseLeader, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldJoinGroupResponseSkipAssignment, Ty: schema.TypeBool},
			&schema.Mfield{Name: FieldJoinGroupResponseMemberId, Ty: schema.TypeStrCompact},
			&schema.ArrayCompact{Name: FieldJoinGroupResponseMembers, Ty: schema.NewSchema("[]JoinGroupResponseMember:v9",
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMemberId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldJoinGroupResponseMembersGroupInstanceId, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldJoinGroupResponseMembersMetadata, Ty: schema.TypeBytesCompact},
				&schema.SchemaTaggedFields{Name: FieldJoinGroupResponseMembersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldJoinGroupResponseTags},
		),
	}

}

const (

	// FieldJoinGroupResponseErrorCode is: The error code, or 0 if there was no error.
	FieldJoinGroupResponseErrorCode = "ErrorCode"

	// FieldJoinGroupResponseGenerationId is: The generation ID of the group.
	FieldJoinGroupResponseGenerationId = "GenerationId"

	// FieldJoinGroupResponseLeader is: The leader of the group.
	FieldJoinGroupResponseLeader = "Leader"

	// FieldJoinGroupResponseMemberId is: The member ID assigned by the group coordinator.
	FieldJoinGroupResponseMemberId = "MemberId"

	// FieldJoinGroupResponseMembers is:
	FieldJoinGroupResponseMembers = "Members"

	// FieldJoinGroupResponseMembersGroupInstanceId is: The unique identifier of the consumer instance provided by end user.
	FieldJoinGroupResponseMembersGroupInstanceId = "GroupInstanceId"

	// FieldJoinGroupResponseMembersMemberId is: The group member ID.
	FieldJoinGroupResponseMembersMemberId = "MemberId"

	// FieldJoinGroupResponseMembersMetadata is: The group member metadata.
	FieldJoinGroupResponseMembersMetadata = "Metadata"

	// FieldJoinGroupResponseMembersTags is: The tagged fields.
	FieldJoinGroupResponseMembersTags = "Tags"

	// FieldJoinGroupResponseProtocolName is: The group protocol selected by the coordinator.
	FieldJoinGroupResponseProtocolName = "ProtocolName"

	// FieldJoinGroupResponseProtocolType is: The group protocol name.
	FieldJoinGroupResponseProtocolType = "ProtocolType"

	// FieldJoinGroupResponseSkipAssignment is: True if the leader must skip running the assignment.
	FieldJoinGroupResponseSkipAssignment = "SkipAssignment"

	// FieldJoinGroupResponseTags is: The tagged fields.
	FieldJoinGroupResponseTags = "Tags"

	// FieldJoinGroupResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldJoinGroupResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/JoinGroupResponse.json
const originalJoinGroupResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 11,
  "type": "response",
  "name": "JoinGroupResponse",
  // Version 1 is the same as version 0.
  //
  // Version 2 adds throttle time.
  //
  // Starting in version 3, on quota violation, brokers send out responses before throttling.
  //
  // Starting in version 4, the client needs to issue a second request to join group
  // with assigned id.
  //
  // Version 5 is bumped to apply group.instance.id to identify member across restarts.
  //
  // Version 6 is the first flexible version.
  //
  // Starting from version 7, the broker sends back the Protocol Type to the client (KIP-559).
  //
  // Version 8 is the same as version 7.
  //
  // Version 9 adds the SkipAssignment field.
  "validVersions": "0-9",
  "flexibleVersions": "6+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "2+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The error code, or 0 if there was no error." },
    { "name": "GenerationId", "type": "int32", "versions": "0+", "default": "-1",
      "about": "The generation ID of the group." },
    { "name": "ProtocolType", "type": "string", "versions": "7+",
      "nullableVersions": "7+", "default": "null", "ignorable": true,
      "about": "The group protocol name." },
    { "name": "ProtocolName", "type": "string", "versions": "0+", "nullableVersions": "7+",
      "about": "The group protocol selected by the coordinator." },
    { "name": "Leader", "type": "string", "versions": "0+",
      "about": "The leader of the group." },
    { "name": "SkipAssignment", "type": "bool", "versions": "9+", "default": "false",
      "about": "True if the leader must skip running the assignment." },
    { "name": "MemberId", "type": "string", "versions": "0+",
      "about": "The member ID assigned by the group coordinator." },
    { "name": "Members", "type": "[]JoinGroupResponseMember", "versions": "0+", "fields": [
      { "name": "MemberId", "type": "string", "versions": "0+",
        "about": "The group member ID." },
      { "name": "GroupInstanceId", "type": "string", "versions": "5+",
        "nullableVersions": "5+", "default": "null",
        "about": "The unique identifier of the consumer instance provided by end user." },
      { "name": "Metadata", "type": "bytes", "versions": "0+",
        "about": "The group member metadata." }
    ]}
  ]
}
`
