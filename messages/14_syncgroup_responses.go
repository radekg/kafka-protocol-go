package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init14SyncGroupResponse() []schema.Schema {

	return []schema.Schema{

		// Message: SyncGroupResponse, API Key: 14, Version: 0
		schema.NewSchema("SyncGroupResponsev0",
			&schema.Mfield{Name: FieldSyncGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldSyncGroupResponseAssignment, Ty: schema.TypeBytes},
		),

		// Message: SyncGroupResponse, API Key: 14, Version: 1
		schema.NewSchema("SyncGroupResponsev1",
			&schema.Mfield{Name: FieldSyncGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldSyncGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldSyncGroupResponseAssignment, Ty: schema.TypeBytes},
		),

		// Message: SyncGroupResponse, API Key: 14, Version: 2
		schema.NewSchema("SyncGroupResponsev2",
			&schema.Mfield{Name: FieldSyncGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldSyncGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldSyncGroupResponseAssignment, Ty: schema.TypeBytes},
		),

		// Message: SyncGroupResponse, API Key: 14, Version: 3
		schema.NewSchema("SyncGroupResponsev3",
			&schema.Mfield{Name: FieldSyncGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldSyncGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldSyncGroupResponseAssignment, Ty: schema.TypeBytes},
		),

		// Message: SyncGroupResponse, API Key: 14, Version: 4
		schema.NewSchema("SyncGroupResponsev4",
			&schema.Mfield{Name: FieldSyncGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldSyncGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldSyncGroupResponseAssignment, Ty: schema.TypeBytesCompact},
			&schema.SchemaTaggedFields{Name: FieldSyncGroupResponseTags},
		),

		// Message: SyncGroupResponse, API Key: 14, Version: 5
		schema.NewSchema("SyncGroupResponsev5",
			&schema.Mfield{Name: FieldSyncGroupResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldSyncGroupResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldSyncGroupResponseProtocolType, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldSyncGroupResponseProtocolName, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldSyncGroupResponseAssignment, Ty: schema.TypeBytesCompact},
			&schema.SchemaTaggedFields{Name: FieldSyncGroupResponseTags},
		),
	}
}

const (
	// FieldSyncGroupResponseErrorCode is: The error code, or 0 if there was no error.
	FieldSyncGroupResponseErrorCode = "ErrorCode"
	// FieldSyncGroupResponseAssignment is: The member assignment.
	FieldSyncGroupResponseAssignment = "Assignment"
	// FieldSyncGroupResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldSyncGroupResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldSyncGroupResponseTags is: The tagged fields.
	FieldSyncGroupResponseTags = "Tags"
	// FieldSyncGroupResponseProtocolType is: The group protocol type.
	FieldSyncGroupResponseProtocolType = "ProtocolType"
	// FieldSyncGroupResponseProtocolName is: The group protocol name.
	FieldSyncGroupResponseProtocolName = "ProtocolName"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/SyncGroupResponse.json
const originalSyncGroupResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "response",
  "name": "SyncGroupResponse",
  // Version 1 adds throttle time.
  //
  // Starting in version 2, on quota violation, brokers send out responses before throttling.
  //
  // Starting from version 3, syncGroupRequest supports a new field called groupInstanceId to indicate member identity across restarts.
  //
  // Version 4 is the first flexible version.
  //
  // Starting from version 5, the broker sends back the Protocol Type and the Protocol Name
  // to the client (KIP-559).
  "validVersions": "0-5",
  "flexibleVersions": "4+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "1+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The error code, or 0 if there was no error." },
    { "name": "ProtocolType", "type": "string", "versions": "5+",
      "nullableVersions": "5+", "default": "null", "ignorable": true,
      "about": "The group protocol type." },
    { "name": "ProtocolName", "type": "string", "versions": "5+",
      "nullableVersions": "5+", "default": "null", "ignorable": true,
      "about": "The group protocol name." },
    { "name": "Assignment", "type": "bytes", "versions": "0+",
      "about": "The member assignment." }
  ]
}
`
