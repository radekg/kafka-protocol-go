package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init15DescribeGroupsResponse() []schema.Schema {

	return []schema.Schema{

		// Message: DescribeGroupsResponse, API Key: 15, Version: 0
		schema.NewSchema("DescribeGroupsResponsev0",
			&schema.Array{Name: FieldDescribeGroupsResponseGroups, Ty: schema.NewSchema("GroupsV0",
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsGroupId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsGroupState, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsProtocolType, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsProtocolData, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDescribeGroupsResponseGroupsMembers, Ty: schema.NewSchema("MembersV0",
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberId, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersClientId, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersClientHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberMetadata, Ty: schema.TypeBytes},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberAssignment, Ty: schema.TypeBytes},
				)},
			)},
		),

		// Message: DescribeGroupsResponse, API Key: 15, Version: 1
		schema.NewSchema("DescribeGroupsResponsev1",
			&schema.Mfield{Name: FieldDescribeGroupsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDescribeGroupsResponseGroups, Ty: schema.NewSchema("GroupsV1",
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsGroupId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsGroupState, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsProtocolType, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsProtocolData, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDescribeGroupsResponseGroupsMembers, Ty: schema.NewSchema("MembersV1",
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberId, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersClientId, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersClientHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberMetadata, Ty: schema.TypeBytes},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberAssignment, Ty: schema.TypeBytes},
				)},
			)},
		),

		// Message: DescribeGroupsResponse, API Key: 15, Version: 2
		schema.NewSchema("DescribeGroupsResponsev2",
			&schema.Mfield{Name: FieldDescribeGroupsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDescribeGroupsResponseGroups, Ty: schema.NewSchema("GroupsV2",
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsGroupId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsGroupState, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsProtocolType, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsProtocolData, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDescribeGroupsResponseGroupsMembers, Ty: schema.NewSchema("MembersV2",
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberId, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersClientId, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersClientHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberMetadata, Ty: schema.TypeBytes},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberAssignment, Ty: schema.TypeBytes},
				)},
			)},
		),

		// Message: DescribeGroupsResponse, API Key: 15, Version: 3
		schema.NewSchema("DescribeGroupsResponsev3",
			&schema.Mfield{Name: FieldDescribeGroupsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDescribeGroupsResponseGroups, Ty: schema.NewSchema("GroupsV3",
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsGroupId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsGroupState, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsProtocolType, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsProtocolData, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDescribeGroupsResponseGroupsMembers, Ty: schema.NewSchema("MembersV3",
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberId, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersClientId, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersClientHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberMetadata, Ty: schema.TypeBytes},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberAssignment, Ty: schema.TypeBytes},
				)},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsAuthorizedOperations, Ty: schema.TypeInt32},
			)},
		),

		// Message: DescribeGroupsResponse, API Key: 15, Version: 4
		schema.NewSchema("DescribeGroupsResponsev4",
			&schema.Mfield{Name: FieldDescribeGroupsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDescribeGroupsResponseGroups, Ty: schema.NewSchema("GroupsV4",
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsGroupId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsGroupState, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsProtocolType, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsProtocolData, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDescribeGroupsResponseGroupsMembers, Ty: schema.NewSchema("MembersV4",
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberId, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersGroupInstanceId, Ty: schema.TypeStrNullable},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersClientId, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersClientHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberMetadata, Ty: schema.TypeBytes},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberAssignment, Ty: schema.TypeBytes},
				)},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsAuthorizedOperations, Ty: schema.TypeInt32},
			)},
		),

		// Message: DescribeGroupsResponse, API Key: 15, Version: 5
		schema.NewSchema("DescribeGroupsResponsev5",
			&schema.Mfield{Name: FieldDescribeGroupsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldDescribeGroupsResponseGroups, Ty: schema.NewSchema("GroupsV5",
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsGroupId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsGroupState, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsProtocolType, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsProtocolData, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldDescribeGroupsResponseGroupsMembers, Ty: schema.NewSchema("MembersV5",
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberId, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersGroupInstanceId, Ty: schema.TypeStrCompactNullable},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersClientId, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersClientHost, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberMetadata, Ty: schema.TypeBytesCompact},
					&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsMembersMemberAssignment, Ty: schema.TypeBytesCompact},
					&schema.SchemaTaggedFields{Name: FieldDescribeGroupsResponseGroupsMembersTags},
				)},
				&schema.Mfield{Name: FieldDescribeGroupsResponseGroupsAuthorizedOperations, Ty: schema.TypeInt32},
				&schema.SchemaTaggedFields{Name: FieldDescribeGroupsResponseGroupsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDescribeGroupsResponseTags},
		),
	}
}

const (
	// FieldDescribeGroupsResponseTags is: The tagged fields.
	FieldDescribeGroupsResponseTags = "Tags"
	// FieldDescribeGroupsResponseGroupsGroupState is: The group state string, or the empty string.
	FieldDescribeGroupsResponseGroupsGroupState = "GroupState"
	// FieldDescribeGroupsResponseGroupsProtocolData is: The group protocol data, or the empty string.
	FieldDescribeGroupsResponseGroupsProtocolData = "ProtocolData"
	// FieldDescribeGroupsResponseGroupsMembersMemberId is: The member ID assigned by the group coordinator.
	FieldDescribeGroupsResponseGroupsMembersMemberId = "MemberId"
	// FieldDescribeGroupsResponseGroupsMembersClientId is: The client ID used in the member's latest join group request.
	FieldDescribeGroupsResponseGroupsMembersClientId = "ClientId"
	// FieldDescribeGroupsResponseGroupsMembersClientHost is: The client host.
	FieldDescribeGroupsResponseGroupsMembersClientHost = "ClientHost"
	// FieldDescribeGroupsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldDescribeGroupsResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldDescribeGroupsResponseGroupsAuthorizedOperations is: 32-bit bitfield to represent authorized operations for this group.
	FieldDescribeGroupsResponseGroupsAuthorizedOperations = "AuthorizedOperations"
	// FieldDescribeGroupsResponseGroups is: Each described group.
	FieldDescribeGroupsResponseGroups = "Groups"
	// FieldDescribeGroupsResponseGroupsErrorCode is: The describe error, or 0 if there was no error.
	FieldDescribeGroupsResponseGroupsErrorCode = "ErrorCode"
	// FieldDescribeGroupsResponseGroupsMembers is: The group members.
	FieldDescribeGroupsResponseGroupsMembers = "Members"
	// FieldDescribeGroupsResponseGroupsMembersMemberMetadata is: The metadata corresponding to the current group protocol in use.
	FieldDescribeGroupsResponseGroupsMembersMemberMetadata = "MemberMetadata"
	// FieldDescribeGroupsResponseGroupsMembersGroupInstanceId is: The unique identifier of the consumer instance provided by end user.
	FieldDescribeGroupsResponseGroupsMembersGroupInstanceId = "GroupInstanceId"
	// FieldDescribeGroupsResponseGroupsTags is: The tagged fields.
	FieldDescribeGroupsResponseGroupsTags = "Tags"
	// FieldDescribeGroupsResponseGroupsGroupId is: The group ID string.
	FieldDescribeGroupsResponseGroupsGroupId = "GroupId"
	// FieldDescribeGroupsResponseGroupsProtocolType is: The group protocol type, or the empty string.
	FieldDescribeGroupsResponseGroupsProtocolType = "ProtocolType"
	// FieldDescribeGroupsResponseGroupsMembersMemberAssignment is: The current assignment provided by the group leader.
	FieldDescribeGroupsResponseGroupsMembersMemberAssignment = "MemberAssignment"
	// FieldDescribeGroupsResponseGroupsMembersTags is: The tagged fields.
	FieldDescribeGroupsResponseGroupsMembersTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeGroupsResponse.json
const originalDescribeGroupsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 15,
  "type": "response",
  "name": "DescribeGroupsResponse",
  // Version 1 added throttle time.
  //
  // Starting in version 2, on quota violation, brokers send out responses before throttling.
  //
  // Starting in version 3, brokers can send authorized operations.
  //
  // Starting in version 4, the response will optionally include group.instance.id info for members.
  //
  // Version 5 is the first flexible version.
  "validVersions": "0-5",
  "flexibleVersions": "5+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "1+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Groups", "type": "[]DescribedGroup", "versions": "0+",
      "about": "Each described group.", "fields": [
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The describe error, or 0 if there was no error." },
      { "name": "GroupId", "type": "string", "versions": "0+", "entityType": "groupId",
        "about": "The group ID string." },
      { "name": "GroupState", "type": "string", "versions": "0+",
        "about": "The group state string, or the empty string." },
      { "name": "ProtocolType", "type": "string", "versions": "0+",
        "about": "The group protocol type, or the empty string." },
      // ProtocolData is currently only filled in if the group state is in the Stable state.
      { "name": "ProtocolData", "type": "string", "versions": "0+",
        "about": "The group protocol data, or the empty string." },
      // N.B. If the group is in the Dead state, the members array will always be empty.
      { "name": "Members", "type": "[]DescribedGroupMember", "versions": "0+",
        "about": "The group members.", "fields": [
        { "name": "MemberId", "type": "string", "versions": "0+",
          "about": "The member ID assigned by the group coordinator." },
        { "name": "GroupInstanceId", "type": "string", "versions": "4+", "ignorable": true,
          "nullableVersions": "4+", "default": "null",
          "about": "The unique identifier of the consumer instance provided by end user." },
        { "name": "ClientId", "type": "string", "versions": "0+",
          "about": "The client ID used in the member's latest join group request." },
        { "name": "ClientHost", "type": "string", "versions": "0+",
          "about": "The client host." },
        // This is currently only provided if the group is in the Stable state.
        { "name": "MemberMetadata", "type": "bytes", "versions": "0+",
          "about": "The metadata corresponding to the current group protocol in use." },
        // This is currently only provided if the group is in the Stable state.
        { "name": "MemberAssignment", "type": "bytes", "versions": "0+",
          "about": "The current assignment provided by the group leader." }
      ]},
      { "name": "AuthorizedOperations", "type": "int32", "versions": "3+",  "default": "-2147483648",
        "about": "32-bit bitfield to represent authorized operations for this group." }
    ]}
  ]
}
`
