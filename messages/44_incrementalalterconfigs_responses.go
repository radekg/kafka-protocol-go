package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init44IncrementalAlterConfigsResponse() []schema.Schema {

	return []schema.Schema{
		// Message: IncrementalAlterConfigsResponse, API Key: 44, Version: 0
		schema.NewSchema("IncrementalAlterConfigsResponse:v0",
			&schema.Mfield{Name: FieldIncrementalAlterConfigsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldIncrementalAlterConfigsResponseResponses, Ty: schema.NewSchema("[]AlterConfigsResourceResponse:v0",
				&schema.Mfield{Name: FieldIncrementalAlterConfigsResponseResponsesErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldIncrementalAlterConfigsResponseResponsesErrorMessage, Ty: schema.TypeStrNullable},
				&schema.Mfield{Name: FieldIncrementalAlterConfigsResponseResponsesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldIncrementalAlterConfigsResponseResponsesResourceName, Ty: schema.TypeStr},
			)},
		),

		// Message: IncrementalAlterConfigsResponse, API Key: 44, Version: 1
		schema.NewSchema("IncrementalAlterConfigsResponse:v1",
			&schema.Mfield{Name: FieldIncrementalAlterConfigsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldIncrementalAlterConfigsResponseResponses, Ty: schema.NewSchema("[]AlterConfigsResourceResponse:v1",
				&schema.Mfield{Name: FieldIncrementalAlterConfigsResponseResponsesErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldIncrementalAlterConfigsResponseResponsesErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldIncrementalAlterConfigsResponseResponsesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldIncrementalAlterConfigsResponseResponsesResourceName, Ty: schema.TypeStrCompact},
				&schema.SchemaTaggedFields{Name: FieldIncrementalAlterConfigsResponseResponsesTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldIncrementalAlterConfigsResponseTags},
		),
	}

}

const (

	// FieldIncrementalAlterConfigsResponseResponses is: The responses for each resource.
	FieldIncrementalAlterConfigsResponseResponses = "Responses"

	// FieldIncrementalAlterConfigsResponseResponsesErrorCode is: The resource error code.
	FieldIncrementalAlterConfigsResponseResponsesErrorCode = "ErrorCode"

	// FieldIncrementalAlterConfigsResponseResponsesErrorMessage is: The resource error message, or null if there was no error.
	FieldIncrementalAlterConfigsResponseResponsesErrorMessage = "ErrorMessage"

	// FieldIncrementalAlterConfigsResponseResponsesResourceName is: The resource name.
	FieldIncrementalAlterConfigsResponseResponsesResourceName = "ResourceName"

	// FieldIncrementalAlterConfigsResponseResponsesResourceType is: The resource type.
	FieldIncrementalAlterConfigsResponseResponsesResourceType = "ResourceType"

	// FieldIncrementalAlterConfigsResponseResponsesTags is: The tagged fields.
	FieldIncrementalAlterConfigsResponseResponsesTags = "Tags"

	// FieldIncrementalAlterConfigsResponseTags is: The tagged fields.
	FieldIncrementalAlterConfigsResponseTags = "Tags"

	// FieldIncrementalAlterConfigsResponseThrottleTimeMs is: Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldIncrementalAlterConfigsResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/IncrementalAlterConfigsResponse.json
const originalIncrementalAlterConfigsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 44,
  "type": "response",
  "name": "IncrementalAlterConfigsResponse",
  // Version 1 is the first flexible version.
  "validVersions": "0-1",
  "flexibleVersions": "1+",
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
