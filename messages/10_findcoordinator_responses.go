package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init10FindCoordinatorResponse() []schema.Schema {

	return []schema.Schema{

		// Message: FindCoordinatorResponse, API Key: 10, Version: 0
		schema.NewSchema("FindCoordinatorResponsev0",
			&schema.Mfield{Name: FieldFindCoordinatorResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldFindCoordinatorResponseNodeId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFindCoordinatorResponseHost, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldFindCoordinatorResponsePort, Ty: schema.TypeInt32},
		),

		// Message: FindCoordinatorResponse, API Key: 10, Version: 1
		schema.NewSchema("FindCoordinatorResponsev1",
			&schema.Mfield{Name: FieldFindCoordinatorResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFindCoordinatorResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldFindCoordinatorResponseErrorMessage, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldFindCoordinatorResponseNodeId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFindCoordinatorResponseHost, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldFindCoordinatorResponsePort, Ty: schema.TypeInt32},
		),

		// Message: FindCoordinatorResponse, API Key: 10, Version: 2
		schema.NewSchema("FindCoordinatorResponsev2",
			&schema.Mfield{Name: FieldFindCoordinatorResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFindCoordinatorResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldFindCoordinatorResponseErrorMessage, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldFindCoordinatorResponseNodeId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFindCoordinatorResponseHost, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldFindCoordinatorResponsePort, Ty: schema.TypeInt32},
		),

		// Message: FindCoordinatorResponse, API Key: 10, Version: 3
		schema.NewSchema("FindCoordinatorResponsev3",
			&schema.Mfield{Name: FieldFindCoordinatorResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFindCoordinatorResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldFindCoordinatorResponseErrorMessage, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldFindCoordinatorResponseNodeId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFindCoordinatorResponseHost, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldFindCoordinatorResponsePort, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldFindCoordinatorResponseTags},
		),

		// Message: FindCoordinatorResponse, API Key: 10, Version: 4
		schema.NewSchema("FindCoordinatorResponsev4",
			&schema.Mfield{Name: FieldFindCoordinatorResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldFindCoordinatorResponseCoordinators, Ty: schema.NewSchema("CoordinatorsV4",
				&schema.Mfield{Name: FieldFindCoordinatorResponseCoordinatorsKey, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldFindCoordinatorResponseCoordinatorsNodeId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldFindCoordinatorResponseCoordinatorsHost, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldFindCoordinatorResponseCoordinatorsPort, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldFindCoordinatorResponseCoordinatorsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldFindCoordinatorResponseCoordinatorsErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldFindCoordinatorResponseCoordinatorsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldFindCoordinatorResponseTags},
		),
	}
}

const (
	// FieldFindCoordinatorResponseCoordinatorsTags is: The tagged fields.
	FieldFindCoordinatorResponseCoordinatorsTags = "Tags"
	// FieldFindCoordinatorResponseErrorMessage is: The error message, or null if there was no error.
	FieldFindCoordinatorResponseErrorMessage = "ErrorMessage"
	// FieldFindCoordinatorResponseCoordinatorsKey is: The coordinator key.
	FieldFindCoordinatorResponseCoordinatorsKey = "Key"
	// FieldFindCoordinatorResponseCoordinatorsErrorMessage is: The error message, or null if there was no error.
	FieldFindCoordinatorResponseCoordinatorsErrorMessage = "ErrorMessage"
	// FieldFindCoordinatorResponsePort is: The port.
	FieldFindCoordinatorResponsePort = "Port"
	// FieldFindCoordinatorResponseCoordinatorsNodeId is: The node id.
	FieldFindCoordinatorResponseCoordinatorsNodeId = "NodeId"
	// FieldFindCoordinatorResponseCoordinatorsPort is: The port.
	FieldFindCoordinatorResponseCoordinatorsPort = "Port"
	// FieldFindCoordinatorResponseCoordinators is: Each coordinator result in the response
	FieldFindCoordinatorResponseCoordinators = "Coordinators"
	// FieldFindCoordinatorResponseCoordinatorsHost is: The host name.
	FieldFindCoordinatorResponseCoordinatorsHost = "Host"
	// FieldFindCoordinatorResponseErrorCode is: The error code, or 0 if there was no error.
	FieldFindCoordinatorResponseErrorCode = "ErrorCode"
	// FieldFindCoordinatorResponseHost is: The host name.
	FieldFindCoordinatorResponseHost = "Host"
	// FieldFindCoordinatorResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldFindCoordinatorResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldFindCoordinatorResponseNodeId is: The node id.
	FieldFindCoordinatorResponseNodeId = "NodeId"
	// FieldFindCoordinatorResponseTags is: The tagged fields.
	FieldFindCoordinatorResponseTags = "Tags"
	// FieldFindCoordinatorResponseCoordinatorsErrorCode is: The error code, or 0 if there was no error.
	FieldFindCoordinatorResponseCoordinatorsErrorCode = "ErrorCode"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/FindCoordinatorResponse.json
const originalFindCoordinatorResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 10,
  "type": "response",
  "name": "FindCoordinatorResponse",
  // Version 1 adds throttle time and error messages.
  //
  // Starting in version 2, on quota violation, brokers send out responses before throttling.
  //
  // Version 3 is the first flexible version.
  //
  // Version 4 adds support for batching via Coordinators (KIP-699)
  "validVersions": "0-4",
  "flexibleVersions": "3+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "1+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0-3",
      "about": "The error code, or 0 if there was no error." },
    { "name": "ErrorMessage", "type": "string", "versions": "1-3", "nullableVersions": "1-3", "ignorable": true,
      "about": "The error message, or null if there was no error." },
    { "name": "NodeId", "type": "int32", "versions": "0-3", "entityType": "brokerId",
      "about": "The node id." },
    { "name": "Host", "type": "string", "versions": "0-3",
      "about": "The host name." },
    { "name": "Port", "type": "int32", "versions": "0-3",
      "about": "The port." },
    { "name": "Coordinators", "type": "[]Coordinator", "versions": "4+", "about": "Each coordinator result in the response", "fields": [
      { "name": "Key", "type": "string", "versions": "4+", "about": "The coordinator key." },
      { "name": "NodeId", "type": "int32", "versions": "4+", "entityType": "brokerId",
        "about": "The node id." },
      { "name": "Host", "type": "string", "versions": "4+", "about": "The host name." },
      { "name": "Port", "type": "int32", "versions": "4+", "about": "The port." },
      { "name": "ErrorCode", "type": "int16", "versions": "4+",
        "about": "The error code, or 0 if there was no error." },
      { "name": "ErrorMessage", "type": "string", "versions": "4+", "nullableVersions": "4+", "ignorable": true,
        "about": "The error message, or null if there was no error." }
    ]}
  ]
}
`
