package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init37CreatePartitionsResponse() []schema.Schema {

	return []schema.Schema{
		// Message: CreatePartitionsResponse, API Key: 37, Version: 0
		schema.NewSchema("CreatePartitionsResponse:v0",
			&schema.Mfield{Name: FieldCreatePartitionsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldCreatePartitionsResponseResults, Ty: schema.NewSchema("[]CreatePartitionsTopicResult:v0",
				&schema.Mfield{Name: FieldCreatePartitionsResponseResultsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreatePartitionsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreatePartitionsResponseResultsErrorMessage, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: CreatePartitionsResponse, API Key: 37, Version: 1
		schema.NewSchema("CreatePartitionsResponse:v1",
			&schema.Mfield{Name: FieldCreatePartitionsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldCreatePartitionsResponseResults, Ty: schema.NewSchema("[]CreatePartitionsTopicResult:v1",
				&schema.Mfield{Name: FieldCreatePartitionsResponseResultsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreatePartitionsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreatePartitionsResponseResultsErrorMessage, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: CreatePartitionsResponse, API Key: 37, Version: 2
		schema.NewSchema("CreatePartitionsResponse:v2",
			&schema.Mfield{Name: FieldCreatePartitionsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldCreatePartitionsResponseResults, Ty: schema.NewSchema("[]CreatePartitionsTopicResult:v2",
				&schema.Mfield{Name: FieldCreatePartitionsResponseResultsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldCreatePartitionsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreatePartitionsResponseResultsErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldCreatePartitionsResponseResultsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldCreatePartitionsResponseTags},
		),

		// Message: CreatePartitionsResponse, API Key: 37, Version: 3
		schema.NewSchema("CreatePartitionsResponse:v3",
			&schema.Mfield{Name: FieldCreatePartitionsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldCreatePartitionsResponseResults, Ty: schema.NewSchema("[]CreatePartitionsTopicResult:v3",
				&schema.Mfield{Name: FieldCreatePartitionsResponseResultsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldCreatePartitionsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreatePartitionsResponseResultsErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldCreatePartitionsResponseResultsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldCreatePartitionsResponseTags},
		),
	}

}

const (

	// FieldCreatePartitionsResponseResults is: The partition creation results for each topic.
	FieldCreatePartitionsResponseResults = "Results"

	// FieldCreatePartitionsResponseResultsErrorCode is: The result error, or zero if there was no error.
	FieldCreatePartitionsResponseResultsErrorCode = "ErrorCode"

	// FieldCreatePartitionsResponseResultsErrorMessage is: The result message, or null if there was no error.
	FieldCreatePartitionsResponseResultsErrorMessage = "ErrorMessage"

	// FieldCreatePartitionsResponseResultsName is: The topic name.
	FieldCreatePartitionsResponseResultsName = "Name"

	// FieldCreatePartitionsResponseResultsTags is: The tagged fields.
	FieldCreatePartitionsResponseResultsTags = "Tags"

	// FieldCreatePartitionsResponseTags is: The tagged fields.
	FieldCreatePartitionsResponseTags = "Tags"

	// FieldCreatePartitionsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldCreatePartitionsResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/CreatePartitionsResponse.json
const originalCreatePartitionsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 37,
  "type": "response",
  "name": "CreatePartitionsResponse",
  // Starting in version 1, on quota violation, brokers send out responses before throttling.
  //
  // Version 2 adds flexible version support
  //
  // Version 3 is identical to version 2 but may return a THROTTLING_QUOTA_EXCEEDED error
  // in the response if the partitions creation is throttled (KIP-599).
  "validVersions": "0-3",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Results", "type": "[]CreatePartitionsTopicResult", "versions": "0+",
      "about": "The partition creation results for each topic.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The result error, or zero if there was no error."},
      { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
        "default": "null", "about": "The result message, or null if there was no error."}
    ]}
  ]
}
`
