package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init11JoinGroupRequest() []schema.Schema {

	return []schema.Schema{

		// Message: JoinGroupRequest, API Key: 11, Version: 0
		schema.NewSchema("JoinGroupRequestv0",
			&schema.Mfield{Name: FieldJoinGroupRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupRequestSessionTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupRequestProtocolType, Ty: schema.TypeStr},
			&schema.Array{Name: FieldJoinGroupRequestProtocols, Ty: schema.NewSchema("ProtocolsV0",
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsMetadata, Ty: schema.TypeBytes},
			)},
		),

		// Message: JoinGroupRequest, API Key: 11, Version: 1
		schema.NewSchema("JoinGroupRequestv1",
			&schema.Mfield{Name: FieldJoinGroupRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupRequestSessionTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestRebalanceTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupRequestProtocolType, Ty: schema.TypeStr},
			&schema.Array{Name: FieldJoinGroupRequestProtocols, Ty: schema.NewSchema("ProtocolsV1",
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsMetadata, Ty: schema.TypeBytes},
			)},
		),

		// Message: JoinGroupRequest, API Key: 11, Version: 2
		schema.NewSchema("JoinGroupRequestv2",
			&schema.Mfield{Name: FieldJoinGroupRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupRequestSessionTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestRebalanceTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupRequestProtocolType, Ty: schema.TypeStr},
			&schema.Array{Name: FieldJoinGroupRequestProtocols, Ty: schema.NewSchema("ProtocolsV2",
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsMetadata, Ty: schema.TypeBytes},
			)},
		),

		// Message: JoinGroupRequest, API Key: 11, Version: 3
		schema.NewSchema("JoinGroupRequestv3",
			&schema.Mfield{Name: FieldJoinGroupRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupRequestSessionTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestRebalanceTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupRequestProtocolType, Ty: schema.TypeStr},
			&schema.Array{Name: FieldJoinGroupRequestProtocols, Ty: schema.NewSchema("ProtocolsV3",
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsMetadata, Ty: schema.TypeBytes},
			)},
		),

		// Message: JoinGroupRequest, API Key: 11, Version: 4
		schema.NewSchema("JoinGroupRequestv4",
			&schema.Mfield{Name: FieldJoinGroupRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupRequestSessionTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestRebalanceTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupRequestProtocolType, Ty: schema.TypeStr},
			&schema.Array{Name: FieldJoinGroupRequestProtocols, Ty: schema.NewSchema("ProtocolsV4",
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsMetadata, Ty: schema.TypeBytes},
			)},
		),

		// Message: JoinGroupRequest, API Key: 11, Version: 5
		schema.NewSchema("JoinGroupRequestv5",
			&schema.Mfield{Name: FieldJoinGroupRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupRequestSessionTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestRebalanceTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldJoinGroupRequestGroupInstanceId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldJoinGroupRequestProtocolType, Ty: schema.TypeStr},
			&schema.Array{Name: FieldJoinGroupRequestProtocols, Ty: schema.NewSchema("ProtocolsV5",
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsMetadata, Ty: schema.TypeBytes},
			)},
		),

		// Message: JoinGroupRequest, API Key: 11, Version: 6
		schema.NewSchema("JoinGroupRequestv6",
			&schema.Mfield{Name: FieldJoinGroupRequestGroupId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldJoinGroupRequestSessionTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestRebalanceTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestMemberId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldJoinGroupRequestGroupInstanceId, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldJoinGroupRequestProtocolType, Ty: schema.TypeStrCompact},
			&schema.ArrayCompact{Name: FieldJoinGroupRequestProtocols, Ty: schema.NewSchema("ProtocolsV6",
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsMetadata, Ty: schema.TypeBytesCompact},
				&schema.SchemaTaggedFields{Name: FieldJoinGroupRequestProtocolsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldJoinGroupRequestTags},
		),

		// Message: JoinGroupRequest, API Key: 11, Version: 7
		schema.NewSchema("JoinGroupRequestv7",
			&schema.Mfield{Name: FieldJoinGroupRequestGroupId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldJoinGroupRequestSessionTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestRebalanceTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestMemberId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldJoinGroupRequestGroupInstanceId, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldJoinGroupRequestProtocolType, Ty: schema.TypeStrCompact},
			&schema.ArrayCompact{Name: FieldJoinGroupRequestProtocols, Ty: schema.NewSchema("ProtocolsV7",
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsMetadata, Ty: schema.TypeBytesCompact},
				&schema.SchemaTaggedFields{Name: FieldJoinGroupRequestProtocolsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldJoinGroupRequestTags},
		),

		// Message: JoinGroupRequest, API Key: 11, Version: 8
		schema.NewSchema("JoinGroupRequestv8",
			&schema.Mfield{Name: FieldJoinGroupRequestGroupId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldJoinGroupRequestSessionTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestRebalanceTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestMemberId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldJoinGroupRequestGroupInstanceId, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldJoinGroupRequestProtocolType, Ty: schema.TypeStrCompact},
			&schema.ArrayCompact{Name: FieldJoinGroupRequestProtocols, Ty: schema.NewSchema("ProtocolsV8",
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsMetadata, Ty: schema.TypeBytesCompact},
				&schema.SchemaTaggedFields{Name: FieldJoinGroupRequestProtocolsTags},
			)},
			&schema.Mfield{Name: FieldJoinGroupRequestReason, Ty: schema.TypeStrCompactNullable},
			&schema.SchemaTaggedFields{Name: FieldJoinGroupRequestTags},
		),

		// Message: JoinGroupRequest, API Key: 11, Version: 9
		schema.NewSchema("JoinGroupRequestv9",
			&schema.Mfield{Name: FieldJoinGroupRequestGroupId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldJoinGroupRequestSessionTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestRebalanceTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldJoinGroupRequestMemberId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldJoinGroupRequestGroupInstanceId, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldJoinGroupRequestProtocolType, Ty: schema.TypeStrCompact},
			&schema.ArrayCompact{Name: FieldJoinGroupRequestProtocols, Ty: schema.NewSchema("ProtocolsV9",
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldJoinGroupRequestProtocolsMetadata, Ty: schema.TypeBytesCompact},
				&schema.SchemaTaggedFields{Name: FieldJoinGroupRequestProtocolsTags},
			)},
			&schema.Mfield{Name: FieldJoinGroupRequestReason, Ty: schema.TypeStrCompactNullable},
			&schema.SchemaTaggedFields{Name: FieldJoinGroupRequestTags},
		),
	}
}

const (
	// FieldJoinGroupRequestProtocolType is a field name that can be used to resolve the correct struct field.
	FieldJoinGroupRequestProtocolType = "ProtocolType"
	// FieldJoinGroupRequestProtocolsName is a field name that can be used to resolve the correct struct field.
	FieldJoinGroupRequestProtocolsName = "Name"
	// FieldJoinGroupRequestGroupInstanceId is a field name that can be used to resolve the correct struct field.
	FieldJoinGroupRequestGroupInstanceId = "GroupInstanceId"
	// FieldJoinGroupRequestProtocolsTags is a field name that can be used to resolve the correct struct field.
	FieldJoinGroupRequestProtocolsTags = "Tags"
	// FieldJoinGroupRequestReason is a field name that can be used to resolve the correct struct field.
	FieldJoinGroupRequestReason = "Reason"
	// FieldJoinGroupRequestTags is a field name that can be used to resolve the correct struct field.
	FieldJoinGroupRequestTags = "Tags"
	// FieldJoinGroupRequestGroupId is a field name that can be used to resolve the correct struct field.
	FieldJoinGroupRequestGroupId = "GroupId"
	// FieldJoinGroupRequestSessionTimeoutMs is a field name that can be used to resolve the correct struct field.
	FieldJoinGroupRequestSessionTimeoutMs = "SessionTimeoutMs"
	// FieldJoinGroupRequestMemberId is a field name that can be used to resolve the correct struct field.
	FieldJoinGroupRequestMemberId = "MemberId"
	// FieldJoinGroupRequestProtocols is a field name that can be used to resolve the correct struct field.
	FieldJoinGroupRequestProtocols = "Protocols"
	// FieldJoinGroupRequestProtocolsMetadata is a field name that can be used to resolve the correct struct field.
	FieldJoinGroupRequestProtocolsMetadata = "Metadata"
	// FieldJoinGroupRequestRebalanceTimeoutMs is a field name that can be used to resolve the correct struct field.
	FieldJoinGroupRequestRebalanceTimeoutMs = "RebalanceTimeoutMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/JoinGroupRequest.json
const originalJoinGroupRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "JoinGroupRequest",
  // Version 1 adds RebalanceTimeoutMs.
  //
  // Version 2 and 3 are the same as version 1.
  //
  // Starting from version 4, the client needs to issue a second request to join group
  //
  // Starting from version 5, we add a new field called groupInstanceId to indicate member identity across restarts.
  // with assigned id.
  //
  // Version 6 is the first flexible version.
  //
  // Version 7 is the same as version 6.
  //
  // Version 8 adds the Reason field (KIP-800).
  //
  // Version 9 is the same as version 8.
  "validVersions": "0-9",
  "flexibleVersions": "6+",
  "fields": [
    { "name": "GroupId", "type": "string", "versions": "0+", "entityType": "groupId",
      "about": "The group identifier." },
    { "name": "SessionTimeoutMs", "type": "int32", "versions": "0+",
      "about": "The coordinator considers the consumer dead if it receives no heartbeat after this timeout in milliseconds." },
    // Note: if RebalanceTimeoutMs is not present, SessionTimeoutMs should be
    // used instead.  The default of -1 here is just intended as a placeholder.
    { "name": "RebalanceTimeoutMs", "type": "int32", "versions": "1+", "default": "-1", "ignorable": true,
      "about": "The maximum time in milliseconds that the coordinator will wait for each member to rejoin when rebalancing the group." },
    { "name": "MemberId", "type": "string", "versions": "0+",
      "about": "The member id assigned by the group coordinator." },
    { "name": "GroupInstanceId", "type": "string", "versions": "5+", 
      "nullableVersions": "5+", "default": "null",
      "about": "The unique identifier of the consumer instance provided by end user." },
    { "name": "ProtocolType", "type": "string", "versions": "0+",
      "about": "The unique name the for class of protocols implemented by the group we want to join." },
    { "name": "Protocols", "type": "[]JoinGroupRequestProtocol", "versions": "0+",
      "about": "The list of protocols that the member supports.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "mapKey": true,
        "about": "The protocol name." },
      { "name": "Metadata", "type": "bytes", "versions": "0+",
        "about": "The protocol metadata." }
    ]},
    { "name": "Reason", "type": "string", "versions": "8+", "nullableVersions": "8+", "default": "null", "ignorable": true,
      "about": "The reason why the member (re-)joins the group." }
  ]
}
`
