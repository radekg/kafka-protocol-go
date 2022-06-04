package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init16ListGroupsResponse() []schema.Schema {

	return []schema.Schema{

		// Message: ListGroupsResponse, API Key: 16, Version: 0
		schema.NewSchema("ListGroupsResponsev0",
			&schema.Mfield{Name: FieldListGroupsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldListGroupsResponseGroups, Ty: schema.NewSchema("GroupsV0",
				&schema.Mfield{Name: FieldListGroupsResponseGroupsGroupId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldListGroupsResponseGroupsProtocolType, Ty: schema.TypeStr},
			)},
		),

		// Message: ListGroupsResponse, API Key: 16, Version: 1
		schema.NewSchema("ListGroupsResponsev1",
			&schema.Mfield{Name: FieldListGroupsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldListGroupsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldListGroupsResponseGroups, Ty: schema.NewSchema("GroupsV1",
				&schema.Mfield{Name: FieldListGroupsResponseGroupsGroupId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldListGroupsResponseGroupsProtocolType, Ty: schema.TypeStr},
			)},
		),

		// Message: ListGroupsResponse, API Key: 16, Version: 2
		schema.NewSchema("ListGroupsResponsev2",
			&schema.Mfield{Name: FieldListGroupsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldListGroupsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldListGroupsResponseGroups, Ty: schema.NewSchema("GroupsV2",
				&schema.Mfield{Name: FieldListGroupsResponseGroupsGroupId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldListGroupsResponseGroupsProtocolType, Ty: schema.TypeStr},
			)},
		),

		// Message: ListGroupsResponse, API Key: 16, Version: 3
		schema.NewSchema("ListGroupsResponsev3",
			&schema.Mfield{Name: FieldListGroupsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldListGroupsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldListGroupsResponseGroups, Ty: schema.NewSchema("GroupsV3",
				&schema.Mfield{Name: FieldListGroupsResponseGroupsGroupId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldListGroupsResponseGroupsProtocolType, Ty: schema.TypeStrCompact},
				&schema.SchemaTaggedFields{Name: FieldListGroupsResponseGroupsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldListGroupsResponseTags},
		),

		// Message: ListGroupsResponse, API Key: 16, Version: 4
		schema.NewSchema("ListGroupsResponsev4",
			&schema.Mfield{Name: FieldListGroupsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldListGroupsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldListGroupsResponseGroups, Ty: schema.NewSchema("GroupsV4",
				&schema.Mfield{Name: FieldListGroupsResponseGroupsGroupId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldListGroupsResponseGroupsProtocolType, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldListGroupsResponseGroupsGroupState, Ty: schema.TypeStrCompact},
				&schema.SchemaTaggedFields{Name: FieldListGroupsResponseGroupsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldListGroupsResponseTags},
		),
	}
}

const (
	// FieldListGroupsResponseErrorCode is: The error code, or 0 if there was no error.
	FieldListGroupsResponseErrorCode = "ErrorCode"
	// FieldListGroupsResponseGroups is: Each group in the response.
	FieldListGroupsResponseGroups = "Groups"
	// FieldListGroupsResponseGroupsGroupId is: The group ID.
	FieldListGroupsResponseGroupsGroupId = "GroupId"
	// FieldListGroupsResponseGroupsProtocolType is: The group protocol type.
	FieldListGroupsResponseGroupsProtocolType = "ProtocolType"
	// FieldListGroupsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldListGroupsResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldListGroupsResponseGroupsTags is: The tagged fields.
	FieldListGroupsResponseGroupsTags = "Tags"
	// FieldListGroupsResponseTags is: The tagged fields.
	FieldListGroupsResponseTags = "Tags"
	// FieldListGroupsResponseGroupsGroupState is: The group state name.
	FieldListGroupsResponseGroupsGroupState = "GroupState"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ListGroupsResponse.json
const originalListGroupsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 16,
  "type": "response",
  "name": "ListGroupsResponse",
  // Version 1 adds the throttle time.
  //
  // Starting in version 2, on quota violation, brokers send out responses before throttling.
  //
  // Version 3 is the first flexible version.
  //
  // Version 4 adds the GroupState field (KIP-518).
  "validVersions": "0-4",
  "flexibleVersions": "3+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "1+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The error code, or 0 if there was no error." },
    { "name": "Groups", "type": "[]ListedGroup", "versions": "0+",
      "about": "Each group in the response.", "fields": [
      { "name": "GroupId", "type": "string", "versions": "0+", "entityType": "groupId",
        "about": "The group ID." },
      { "name": "ProtocolType", "type": "string", "versions": "0+",
        "about": "The group protocol type." },
      { "name": "GroupState", "type": "string", "versions": "4+", "ignorable": true,
        "about": "The group state name." }
    ]}
  ]
}
`
