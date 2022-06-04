package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init13LeaveGroupResponse() []schema.Schema {

	return []schema.Schema{

		// Message: LeaveGroupResponse, API Key: 13, Version: 0
		schema.NewSchema("LeaveGroupResponsev0",
			&schema.Mfield{Name: FieldLeaveGroupResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: LeaveGroupResponse, API Key: 13, Version: 1
		schema.NewSchema("LeaveGroupResponsev1",
			&schema.Mfield{Name: FieldLeaveGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaveGroupResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: LeaveGroupResponse, API Key: 13, Version: 2
		schema.NewSchema("LeaveGroupResponsev2",
			&schema.Mfield{Name: FieldLeaveGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaveGroupResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: LeaveGroupResponse, API Key: 13, Version: 3
		schema.NewSchema("LeaveGroupResponsev3",
			&schema.Mfield{Name: FieldLeaveGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaveGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldLeaveGroupResponseMembers, Ty: schema.NewSchema("MembersV3",
				&schema.Mfield{Name: FieldLeaveGroupResponseMembersMemberId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldLeaveGroupResponseMembersGroupInstanceId, Ty: schema.TypeStrNullable},
				&schema.Mfield{Name: FieldLeaveGroupResponseMembersErrorCode, Ty: schema.TypeInt16},
			)},
		),

		// Message: LeaveGroupResponse, API Key: 13, Version: 4
		schema.NewSchema("LeaveGroupResponsev4",
			&schema.Mfield{Name: FieldLeaveGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaveGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldLeaveGroupResponseMembers, Ty: schema.NewSchema("MembersV4",
				&schema.Mfield{Name: FieldLeaveGroupResponseMembersMemberId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldLeaveGroupResponseMembersGroupInstanceId, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldLeaveGroupResponseMembersErrorCode, Ty: schema.TypeInt16},
				&schema.SchemaTaggedFields{Name: FieldLeaveGroupResponseMembersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldLeaveGroupResponseTags},
		),

		// Message: LeaveGroupResponse, API Key: 13, Version: 5
		schema.NewSchema("LeaveGroupResponsev5",
			&schema.Mfield{Name: FieldLeaveGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldLeaveGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldLeaveGroupResponseMembers, Ty: schema.NewSchema("MembersV5",
				&schema.Mfield{Name: FieldLeaveGroupResponseMembersMemberId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldLeaveGroupResponseMembersGroupInstanceId, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldLeaveGroupResponseMembersErrorCode, Ty: schema.TypeInt16},
				&schema.SchemaTaggedFields{Name: FieldLeaveGroupResponseMembersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldLeaveGroupResponseTags},
		),
	}
}

const (
	// FieldLeaveGroupResponseMembersTags is: The tagged fields.
	FieldLeaveGroupResponseMembersTags = "Tags"
	// FieldLeaveGroupResponseTags is: The tagged fields.
	FieldLeaveGroupResponseTags = "Tags"
	// FieldLeaveGroupResponseErrorCode is: The error code, or 0 if there was no error.
	FieldLeaveGroupResponseErrorCode = "ErrorCode"
	// FieldLeaveGroupResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldLeaveGroupResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldLeaveGroupResponseMembers is: List of leaving member responses.
	FieldLeaveGroupResponseMembers = "Members"
	// FieldLeaveGroupResponseMembersMemberId is: The member ID to remove from the group.
	FieldLeaveGroupResponseMembersMemberId = "MemberId"
	// FieldLeaveGroupResponseMembersGroupInstanceId is: The group instance ID to remove from the group.
	FieldLeaveGroupResponseMembersGroupInstanceId = "GroupInstanceId"
	// FieldLeaveGroupResponseMembersErrorCode is: The error code, or 0 if there was no error.
	FieldLeaveGroupResponseMembersErrorCode = "ErrorCode"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/LeaveGroupResponse.json
const originalLeaveGroupResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "response",
  "name": "LeaveGroupResponse",
  // Version 1 adds the throttle time.
  //
  // Starting in version 2, on quota violation, brokers send out responses before throttling.
  //
  // Starting in version 3, we will make leave group request into batch mode and add group.instance.id.
  //
  // Version 4 is the first flexible version.
  //
  // Version 5 is the same as version 4.
  "validVersions": "0-5",
  "flexibleVersions": "4+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "1+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The error code, or 0 if there was no error." },

    { "name": "Members", "type": "[]MemberResponse", "versions": "3+",
      "about": "List of leaving member responses.", "fields": [
      { "name": "MemberId", "type": "string", "versions": "3+",
        "about": "The member ID to remove from the group." },
      { "name": "GroupInstanceId", "type": "string", "versions": "3+", "nullableVersions": "3+",
        "about": "The group instance ID to remove from the group." },
      { "name": "ErrorCode", "type": "int16", "versions": "3+",
        "about": "The error code, or 0 if there was no error." }
    ]}
  ]
}
`
