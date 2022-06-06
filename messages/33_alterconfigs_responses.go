package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init33AlterConfigsResponse() []schema.Schema {

	return []schema.Schema{
		// Message: AlterConfigsResponse, API Key: 33, Version: 0
		schema.NewSchema("AlterConfigsResponse:v0",
			&schema.Mfield{Name: FieldAlterConfigsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldAlterConfigsResponseResponses, Ty: schema.NewSchema("Responses:v0",
				&schema.Mfield{Name: FieldAlterConfigsResponseResponsesErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldAlterConfigsResponseResponsesErrorMessage, Ty: schema.TypeStrNullable},
				&schema.Mfield{Name: FieldAlterConfigsResponseResponsesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldAlterConfigsResponseResponsesResourceName, Ty: schema.TypeStr},
			)},
		),

		// Message: AlterConfigsResponse, API Key: 33, Version: 1
		schema.NewSchema("AlterConfigsResponse:v1",
			&schema.Mfield{Name: FieldAlterConfigsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldAlterConfigsResponseResponses, Ty: schema.NewSchema("Responses:v1",
				&schema.Mfield{Name: FieldAlterConfigsResponseResponsesErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldAlterConfigsResponseResponsesErrorMessage, Ty: schema.TypeStrNullable},
				&schema.Mfield{Name: FieldAlterConfigsResponseResponsesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldAlterConfigsResponseResponsesResourceName, Ty: schema.TypeStr},
			)},
		),

		// Message: AlterConfigsResponse, API Key: 33, Version: 2
		schema.NewSchema("AlterConfigsResponse:v2",
			&schema.Mfield{Name: FieldAlterConfigsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldAlterConfigsResponseResponses, Ty: schema.NewSchema("Responses:v2",
				&schema.Mfield{Name: FieldAlterConfigsResponseResponsesErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldAlterConfigsResponseResponsesErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldAlterConfigsResponseResponsesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldAlterConfigsResponseResponsesResourceName, Ty: schema.TypeStrCompact},
				&schema.SchemaTaggedFields{Name: FieldAlterConfigsResponseResponsesTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldAlterConfigsResponseTags},
		),
	}

}

const (

	// FieldAlterConfigsResponseResponses is: The responses for each resource.
	FieldAlterConfigsResponseResponses = "Responses"

	// FieldAlterConfigsResponseResponsesErrorCode is: The resource error code.
	FieldAlterConfigsResponseResponsesErrorCode = "ErrorCode"

	// FieldAlterConfigsResponseResponsesErrorMessage is: The resource error message, or null if there was no error.
	FieldAlterConfigsResponseResponsesErrorMessage = "ErrorMessage"

	// FieldAlterConfigsResponseResponsesResourceName is: The resource name.
	FieldAlterConfigsResponseResponsesResourceName = "ResourceName"

	// FieldAlterConfigsResponseResponsesResourceType is: The resource type.
	FieldAlterConfigsResponseResponsesResourceType = "ResourceType"

	// FieldAlterConfigsResponseResponsesTags is: The tagged fields.
	FieldAlterConfigsResponseResponsesTags = "Tags"

	// FieldAlterConfigsResponseTags is: The tagged fields.
	FieldAlterConfigsResponseTags = "Tags"

	// FieldAlterConfigsResponseThrottleTimeMs is: Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldAlterConfigsResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AlterConfigsResponse.json
const originalAlterConfigsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 33,
  "type": "response",
  "name": "AlterConfigsResponse",
  // Starting in version 1, on quota violation brokers send out responses before throttling.
  // Version 2 enables flexible versions.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Responses", "type": "[]AlterConfigsResourceResponse", "versions": "0+",
      "about": "The responses for each resource.", "fields": [
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The resource error code." },
      { "name": "ErrorMessage", "type": "string", "nullableVersions": "0+", "versions": "0+",
        "about": "The resource error message, or null if there was no error." },
      { "name": "ResourceType", "type": "int8", "versions": "0+",
        "about": "The resource type." },
      { "name": "ResourceName", "type": "string", "versions": "0+",
        "about": "The resource name." }
    ]}
  ]
}
`
