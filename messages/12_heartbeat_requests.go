package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init12HeartbeatRequest() []schema.Schema {

	return []schema.Schema{

		// Message: HeartbeatRequest, API Key: 12, Version: 0
		schema.NewSchema("HeartbeatRequestv0",
			&schema.Mfield{Name: FieldHeartbeatRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldHeartbeatRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldHeartbeatRequestMemberId, Ty: schema.TypeStr},
		),

		// Message: HeartbeatRequest, API Key: 12, Version: 1
		schema.NewSchema("HeartbeatRequestv1",
			&schema.Mfield{Name: FieldHeartbeatRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldHeartbeatRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldHeartbeatRequestMemberId, Ty: schema.TypeStr},
		),

		// Message: HeartbeatRequest, API Key: 12, Version: 2
		schema.NewSchema("HeartbeatRequestv2",
			&schema.Mfield{Name: FieldHeartbeatRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldHeartbeatRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldHeartbeatRequestMemberId, Ty: schema.TypeStr},
		),

		// Message: HeartbeatRequest, API Key: 12, Version: 3
		schema.NewSchema("HeartbeatRequestv3",
			&schema.Mfield{Name: FieldHeartbeatRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldHeartbeatRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldHeartbeatRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldHeartbeatRequestGroupInstanceId, Ty: schema.TypeStrNullable},
		),

		// Message: HeartbeatRequest, API Key: 12, Version: 4
		schema.NewSchema("HeartbeatRequestv4",
			&schema.Mfield{Name: FieldHeartbeatRequestGroupId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldHeartbeatRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldHeartbeatRequestMemberId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldHeartbeatRequestGroupInstanceId, Ty: schema.TypeStrCompactNullable},
			&schema.SchemaTaggedFields{Name: FieldHeartbeatRequestTags},
		),
	}
}

const (
	// FieldHeartbeatRequestTags is a field name that can be used to resolve the correct struct field.
	FieldHeartbeatRequestTags = "Tags"
	// FieldHeartbeatRequestGroupId is a field name that can be used to resolve the correct struct field.
	FieldHeartbeatRequestGroupId = "GroupId"
	// FieldHeartbeatRequestGenerationId is a field name that can be used to resolve the correct struct field.
	FieldHeartbeatRequestGenerationId = "GenerationId"
	// FieldHeartbeatRequestMemberId is a field name that can be used to resolve the correct struct field.
	FieldHeartbeatRequestMemberId = "MemberId"
	// FieldHeartbeatRequestGroupInstanceId is a field name that can be used to resolve the correct struct field.
	FieldHeartbeatRequestGroupInstanceId = "GroupInstanceId"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/HeartbeatRequest.json
const originalHeartbeatRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 12,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "HeartbeatRequest",
  // Version 1 and version 2 are the same as version 0.
  //
  // Starting from version 3, we add a new field called groupInstanceId to indicate member identity across restarts.
  //
  // Version 4 is the first flexible version.
  "validVersions": "0-4",
  "flexibleVersions": "4+",
  "fields": [
    { "name": "GroupId", "type": "string", "versions": "0+", "entityType": "groupId",
      "about": "The group id." },
    { "name": "GenerationId", "type": "int32", "versions": "0+",
      "about": "The generation of the group." },
    { "name": "MemberId", "type": "string", "versions": "0+",
      "about": "The member ID." },
    { "name": "GroupInstanceId", "type": "string", "versions": "3+",
      "nullableVersions": "3+", "default": "null",
      "about": "The unique identifier of the consumer instance provided by end user." }
  ]
}
`
