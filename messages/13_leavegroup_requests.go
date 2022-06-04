package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init13LeaveGroupRequest() []schema.Schema {

	return []schema.Schema{

		// Message: LeaveGroupRequest, API Key: 13, Version: 0
		schema.NewSchema("LeaveGroupRequestv0",
			&schema.Mfield{Name: FieldLeaveGroupRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldLeaveGroupRequestMemberId, Ty: schema.TypeStr},
		),

		// Message: LeaveGroupRequest, API Key: 13, Version: 1
		schema.NewSchema("LeaveGroupRequestv1",
			&schema.Mfield{Name: FieldLeaveGroupRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldLeaveGroupRequestMemberId, Ty: schema.TypeStr},
		),

		// Message: LeaveGroupRequest, API Key: 13, Version: 2
		schema.NewSchema("LeaveGroupRequestv2",
			&schema.Mfield{Name: FieldLeaveGroupRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldLeaveGroupRequestMemberId, Ty: schema.TypeStr},
		),

		// Message: LeaveGroupRequest, API Key: 13, Version: 3
		schema.NewSchema("LeaveGroupRequestv3",
			&schema.Mfield{Name: FieldLeaveGroupRequestGroupId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldLeaveGroupRequestMembers, Ty: schema.NewSchema("MembersV3",
				&schema.Mfield{Name: FieldLeaveGroupRequestMembersMemberId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldLeaveGroupRequestMembersGroupInstanceId, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: LeaveGroupRequest, API Key: 13, Version: 4
		schema.NewSchema("LeaveGroupRequestv4",
			&schema.Mfield{Name: FieldLeaveGroupRequestGroupId, Ty: schema.TypeStrCompact},
			&schema.ArrayCompact{Name: FieldLeaveGroupRequestMembers, Ty: schema.NewSchema("MembersV4",
				&schema.Mfield{Name: FieldLeaveGroupRequestMembersMemberId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldLeaveGroupRequestMembersGroupInstanceId, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldLeaveGroupRequestMembersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldLeaveGroupRequestTags},
		),

		// Message: LeaveGroupRequest, API Key: 13, Version: 5
		schema.NewSchema("LeaveGroupRequestv5",
			&schema.Mfield{Name: FieldLeaveGroupRequestGroupId, Ty: schema.TypeStrCompact},
			&schema.ArrayCompact{Name: FieldLeaveGroupRequestMembers, Ty: schema.NewSchema("MembersV5",
				&schema.Mfield{Name: FieldLeaveGroupRequestMembersMemberId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldLeaveGroupRequestMembersGroupInstanceId, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldLeaveGroupRequestMembersReason, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldLeaveGroupRequestMembersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldLeaveGroupRequestTags},
		),
	}
}

const (
	// FieldLeaveGroupRequestGroupId is: The ID of the group to leave.
	FieldLeaveGroupRequestGroupId = "GroupId"
	// FieldLeaveGroupRequestMemberId is: The member ID to remove from the group.
	FieldLeaveGroupRequestMemberId = "MemberId"
	// FieldLeaveGroupRequestMembers is: List of leaving member identities.
	FieldLeaveGroupRequestMembers = "Members"
	// FieldLeaveGroupRequestMembersGroupInstanceId is: The group instance ID to remove from the group.
	FieldLeaveGroupRequestMembersGroupInstanceId = "GroupInstanceId"
	// FieldLeaveGroupRequestMembersMemberId is: The member ID to remove from the group.
	FieldLeaveGroupRequestMembersMemberId = "MemberId"
	// FieldLeaveGroupRequestMembersReason is: The reason why the member left the group.
	FieldLeaveGroupRequestMembersReason = "Reason"
	// FieldLeaveGroupRequestMembersTags is: The tagged fields.
	FieldLeaveGroupRequestMembersTags = "Tags"
	// FieldLeaveGroupRequestTags is: The tagged fields.
	FieldLeaveGroupRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/LeaveGroupRequest.json
const originalLeaveGroupRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 13,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "LeaveGroupRequest",
  // Version 1 and 2 are the same as version 0.
  //
  // Version 3 defines batch processing scheme with group.instance.id + member.id for identity
  //
  // Version 4 is the first flexible version.
  //
  // Version 5 adds the Reason field (KIP-800).
  "validVersions": "0-5",
  "flexibleVersions": "4+",
  "fields": [
    { "name": "GroupId", "type": "string", "versions": "0+", "entityType": "groupId",
      "about": "The ID of the group to leave." },
    { "name": "MemberId", "type": "string", "versions": "0-2",
      "about": "The member ID to remove from the group." },
    { "name": "Members", "type": "[]MemberIdentity", "versions": "3+",
      "about": "List of leaving member identities.", "fields": [
      { "name": "MemberId", "type": "string", "versions": "3+",
        "about": "The member ID to remove from the group." },
      { "name": "GroupInstanceId", "type": "string",
        "versions": "3+", "nullableVersions": "3+", "default": "null",
        "about": "The group instance ID to remove from the group." },
      { "name": "Reason", "type": "string",
        "versions": "5+", "nullableVersions": "5+", "default": "null", "ignorable": true,
        "about": "The reason why the member left the group." }
    ]}
  ]
}
`
