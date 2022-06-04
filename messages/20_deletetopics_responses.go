package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init20DeleteTopicsResponse() []schema.Schema {

	return []schema.Schema{

		// Message: DeleteTopicsResponse, API Key: 20, Version: 0
		schema.NewSchema("DeleteTopicsResponsev0",
			&schema.Array{Name: FieldDeleteTopicsResponseResponses, Ty: schema.NewSchema("ResponsesV0",
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesErrorCode, Ty: schema.TypeInt16},
			)},
		),

		// Message: DeleteTopicsResponse, API Key: 20, Version: 1
		schema.NewSchema("DeleteTopicsResponsev1",
			&schema.Mfield{Name: FieldDeleteTopicsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDeleteTopicsResponseResponses, Ty: schema.NewSchema("ResponsesV1",
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesErrorCode, Ty: schema.TypeInt16},
			)},
		),

		// Message: DeleteTopicsResponse, API Key: 20, Version: 2
		schema.NewSchema("DeleteTopicsResponsev2",
			&schema.Mfield{Name: FieldDeleteTopicsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDeleteTopicsResponseResponses, Ty: schema.NewSchema("ResponsesV2",
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesErrorCode, Ty: schema.TypeInt16},
			)},
		),

		// Message: DeleteTopicsResponse, API Key: 20, Version: 3
		schema.NewSchema("DeleteTopicsResponsev3",
			&schema.Mfield{Name: FieldDeleteTopicsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDeleteTopicsResponseResponses, Ty: schema.NewSchema("ResponsesV3",
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesErrorCode, Ty: schema.TypeInt16},
			)},
		),

		// Message: DeleteTopicsResponse, API Key: 20, Version: 4
		schema.NewSchema("DeleteTopicsResponsev4",
			&schema.Mfield{Name: FieldDeleteTopicsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldDeleteTopicsResponseResponses, Ty: schema.NewSchema("ResponsesV4",
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesErrorCode, Ty: schema.TypeInt16},
				&schema.SchemaTaggedFields{Name: FieldDeleteTopicsResponseResponsesTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDeleteTopicsResponseTags},
		),

		// Message: DeleteTopicsResponse, API Key: 20, Version: 5
		schema.NewSchema("DeleteTopicsResponsev5",
			&schema.Mfield{Name: FieldDeleteTopicsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldDeleteTopicsResponseResponses, Ty: schema.NewSchema("ResponsesV5",
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldDeleteTopicsResponseResponsesTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDeleteTopicsResponseTags},
		),

		// Message: DeleteTopicsResponse, API Key: 20, Version: 6
		schema.NewSchema("DeleteTopicsResponsev6",
			&schema.Mfield{Name: FieldDeleteTopicsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldDeleteTopicsResponseResponses, Ty: schema.NewSchema("ResponsesV6",
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesName, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesTopicId, Ty: schema.TypeUuid},
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDeleteTopicsResponseResponsesErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldDeleteTopicsResponseResponsesTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDeleteTopicsResponseTags},
		),
	}
}

const (
	// FieldDeleteTopicsResponseResponsesName is: The topic name
	FieldDeleteTopicsResponseResponsesName = "Name"
	// FieldDeleteTopicsResponseResponsesErrorCode is: The deletion error, or 0 if the deletion succeeded.
	FieldDeleteTopicsResponseResponsesErrorCode = "ErrorCode"
	// FieldDeleteTopicsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldDeleteTopicsResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldDeleteTopicsResponseResponsesTags is: The tagged fields.
	FieldDeleteTopicsResponseResponsesTags = "Tags"
	// FieldDeleteTopicsResponseTags is: The tagged fields.
	FieldDeleteTopicsResponseTags = "Tags"
	// FieldDeleteTopicsResponseResponsesErrorMessage is: The error message, or null if there was no error.
	FieldDeleteTopicsResponseResponsesErrorMessage = "ErrorMessage"
	// FieldDeleteTopicsResponseResponsesTopicId is: the unique topic ID
	FieldDeleteTopicsResponseResponsesTopicId = "TopicId"
	// FieldDeleteTopicsResponseResponses is: The results for each topic we tried to delete.
	FieldDeleteTopicsResponseResponses = "Responses"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DeleteTopicsResponse.json
const originalDeleteTopicsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 20,
  "type": "response",
  "name": "DeleteTopicsResponse",
  // Version 1 adds the throttle time.
  //
  // Starting in version 2, on quota violation, brokers send out responses before throttling.
  //
  // Starting in version 3, a TOPIC_DELETION_DISABLED error code may be returned.
  //
  // Version 4 is the first flexible version.
  //
  // Version 5 adds ErrorMessage in the response and may return a THROTTLING_QUOTA_EXCEEDED error
  // in the response if the topics deletion is throttled (KIP-599).
  //
  // Version 6 adds topic ID to responses. An UNSUPPORTED_VERSION error code will be returned when attempting to
  // delete using topic IDs when IBP < 2.8. UNKNOWN_TOPIC_ID error code will be returned when IBP is at least 2.8, but
  // the topic ID was not found.
  "validVersions": "0-6",
  "flexibleVersions": "4+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "1+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Responses", "type": "[]DeletableTopicResult", "versions": "0+",
      "about": "The results for each topic we tried to delete.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "nullableVersions": "6+", "mapKey": true, "entityType": "topicName",
        "about": "The topic name" },
      {"name": "TopicId", "type": "uuid", "versions": "6+", "ignorable": true, "about": "the unique topic ID"},
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The deletion error, or 0 if the deletion succeeded." },
      { "name": "ErrorMessage", "type": "string", "versions": "5+", "nullableVersions": "5+", "ignorable": true, "default": "null",
        "about": "The error message, or null if there was no error." }
    ]}
  ]
}
`
