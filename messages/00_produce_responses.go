package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init0ProduceResponse() []schema.Schema {

	return []schema.Schema{
		// Message: ProduceResponse, API Key: 0, Version: 0
		schema.NewSchema("ProduceResponse:v0",
			&schema.Array{Name: FieldProduceResponseResponses, Ty: schema.NewSchema("Responses:v0",
				&schema.Mfield{Name: FieldProduceResponseResponsesName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceResponseResponsesPartitionResponses, Ty: schema.NewSchema("PartitionResponses:v0",
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesBaseOffset, Ty: schema.TypeInt64},
				)},
			)},
		),

		// Message: ProduceResponse, API Key: 0, Version: 1
		schema.NewSchema("ProduceResponse:v1",
			&schema.Array{Name: FieldProduceResponseResponses, Ty: schema.NewSchema("Responses:v1",
				&schema.Mfield{Name: FieldProduceResponseResponsesName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceResponseResponsesPartitionResponses, Ty: schema.NewSchema("PartitionResponses:v1",
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesBaseOffset, Ty: schema.TypeInt64},
				)},
			)},
			&schema.Mfield{Name: FieldProduceResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: ProduceResponse, API Key: 0, Version: 2
		schema.NewSchema("ProduceResponse:v2",
			&schema.Array{Name: FieldProduceResponseResponses, Ty: schema.NewSchema("Responses:v2",
				&schema.Mfield{Name: FieldProduceResponseResponsesName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceResponseResponsesPartitionResponses, Ty: schema.NewSchema("PartitionResponses:v2",
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesBaseOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesLogAppendTimeMs, Ty: schema.TypeInt64},
				)},
			)},
			&schema.Mfield{Name: FieldProduceResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: ProduceResponse, API Key: 0, Version: 3
		schema.NewSchema("ProduceResponse:v3",
			&schema.Array{Name: FieldProduceResponseResponses, Ty: schema.NewSchema("Responses:v3",
				&schema.Mfield{Name: FieldProduceResponseResponsesName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceResponseResponsesPartitionResponses, Ty: schema.NewSchema("PartitionResponses:v3",
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesBaseOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesLogAppendTimeMs, Ty: schema.TypeInt64},
				)},
			)},
			&schema.Mfield{Name: FieldProduceResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: ProduceResponse, API Key: 0, Version: 4
		schema.NewSchema("ProduceResponse:v4",
			&schema.Array{Name: FieldProduceResponseResponses, Ty: schema.NewSchema("Responses:v4",
				&schema.Mfield{Name: FieldProduceResponseResponsesName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceResponseResponsesPartitionResponses, Ty: schema.NewSchema("PartitionResponses:v4",
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesBaseOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesLogAppendTimeMs, Ty: schema.TypeInt64},
				)},
			)},
			&schema.Mfield{Name: FieldProduceResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: ProduceResponse, API Key: 0, Version: 5
		schema.NewSchema("ProduceResponse:v5",
			&schema.Array{Name: FieldProduceResponseResponses, Ty: schema.NewSchema("Responses:v5",
				&schema.Mfield{Name: FieldProduceResponseResponsesName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceResponseResponsesPartitionResponses, Ty: schema.NewSchema("PartitionResponses:v5",
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesBaseOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesLogAppendTimeMs, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesLogStartOffset, Ty: schema.TypeInt64},
				)},
			)},
			&schema.Mfield{Name: FieldProduceResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: ProduceResponse, API Key: 0, Version: 6
		schema.NewSchema("ProduceResponse:v6",
			&schema.Array{Name: FieldProduceResponseResponses, Ty: schema.NewSchema("Responses:v6",
				&schema.Mfield{Name: FieldProduceResponseResponsesName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceResponseResponsesPartitionResponses, Ty: schema.NewSchema("PartitionResponses:v6",
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesBaseOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesLogAppendTimeMs, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesLogStartOffset, Ty: schema.TypeInt64},
				)},
			)},
			&schema.Mfield{Name: FieldProduceResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: ProduceResponse, API Key: 0, Version: 7
		schema.NewSchema("ProduceResponse:v7",
			&schema.Array{Name: FieldProduceResponseResponses, Ty: schema.NewSchema("Responses:v7",
				&schema.Mfield{Name: FieldProduceResponseResponsesName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceResponseResponsesPartitionResponses, Ty: schema.NewSchema("PartitionResponses:v7",
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesBaseOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesLogAppendTimeMs, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesLogStartOffset, Ty: schema.TypeInt64},
				)},
			)},
			&schema.Mfield{Name: FieldProduceResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: ProduceResponse, API Key: 0, Version: 8
		schema.NewSchema("ProduceResponse:v8",
			&schema.Array{Name: FieldProduceResponseResponses, Ty: schema.NewSchema("Responses:v8",
				&schema.Mfield{Name: FieldProduceResponseResponsesName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceResponseResponsesPartitionResponses, Ty: schema.NewSchema("PartitionResponses:v8",
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesBaseOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesLogAppendTimeMs, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesLogStartOffset, Ty: schema.TypeInt64},
					&schema.Array{Name: FieldProduceResponseResponsesPartitionResponsesRecordErrors, Ty: schema.NewSchema("RecordErrors:v8",
						&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesRecordErrorsBatchIndex, Ty: schema.TypeInt32},
						&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesRecordErrorsBatchIndexErrorMessage, Ty: schema.TypeStrNullable},
					)},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesErrorMessage, Ty: schema.TypeStrNullable},
				)},
			)},
			&schema.Mfield{Name: FieldProduceResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: ProduceResponse, API Key: 0, Version: 9
		schema.NewSchema("ProduceResponse:v9",
			&schema.ArrayCompact{Name: FieldProduceResponseResponses, Ty: schema.NewSchema("Responses:v9",
				&schema.Mfield{Name: FieldProduceResponseResponsesName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldProduceResponseResponsesPartitionResponses, Ty: schema.NewSchema("PartitionResponses:v9",
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesBaseOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesLogAppendTimeMs, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesLogStartOffset, Ty: schema.TypeInt64},
					&schema.ArrayCompact{Name: FieldProduceResponseResponsesPartitionResponsesRecordErrors, Ty: schema.NewSchema("RecordErrors:v9",
						&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesRecordErrorsBatchIndex, Ty: schema.TypeInt32},
						&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesRecordErrorsBatchIndexErrorMessage, Ty: schema.TypeStrCompactNullable},
						&schema.SchemaTaggedFields{Name: FieldProduceResponseResponsesPartitionResponsesRecordErrorsTags},
					)},
					&schema.Mfield{Name: FieldProduceResponseResponsesPartitionResponsesErrorMessage, Ty: schema.TypeStrCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldProduceResponseResponsesPartitionResponsesTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldProduceResponseResponsesTags},
			)},
			&schema.Mfield{Name: FieldProduceResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldProduceResponseTags},
		),
	}

}

const (

	// FieldProduceResponseResponses is: Each produce response
	FieldProduceResponseResponses = "Responses"

	// FieldProduceResponseResponsesName is: The topic name
	FieldProduceResponseResponsesName = "Name"

	// FieldProduceResponseResponsesPartitionResponses is: Each partition that we produced to within the topic.
	FieldProduceResponseResponsesPartitionResponses = "PartitionResponses"

	// FieldProduceResponseResponsesPartitionResponsesBaseOffset is: The base offset.
	FieldProduceResponseResponsesPartitionResponsesBaseOffset = "BaseOffset"

	// FieldProduceResponseResponsesPartitionResponsesErrorCode is: The error code, or 0 if there was no error.
	FieldProduceResponseResponsesPartitionResponsesErrorCode = "ErrorCode"

	// FieldProduceResponseResponsesPartitionResponsesErrorMessage is: The global error message summarizing the common root cause of the records that caused the batch to be dropped
	FieldProduceResponseResponsesPartitionResponsesErrorMessage = "ErrorMessage"

	// FieldProduceResponseResponsesPartitionResponsesIndex is: The partition index.
	FieldProduceResponseResponsesPartitionResponsesIndex = "Index"

	// FieldProduceResponseResponsesPartitionResponsesLogAppendTimeMs is: The timestamp returned by broker after appending the messages. If CreateTime is used for the topic, the timestamp will be -1.  If LogAppendTime is used for the topic, the timestamp will be the broker local time when the messages are appended.
	FieldProduceResponseResponsesPartitionResponsesLogAppendTimeMs = "LogAppendTimeMs"

	// FieldProduceResponseResponsesPartitionResponsesLogStartOffset is: The log start offset.
	FieldProduceResponseResponsesPartitionResponsesLogStartOffset = "LogStartOffset"

	// FieldProduceResponseResponsesPartitionResponsesRecordErrors is: The batch indices of records that caused the batch to be dropped
	FieldProduceResponseResponsesPartitionResponsesRecordErrors = "RecordErrors"

	// FieldProduceResponseResponsesPartitionResponsesRecordErrorsBatchIndex is: The batch index of the record that cause the batch to be dropped
	FieldProduceResponseResponsesPartitionResponsesRecordErrorsBatchIndex = "BatchIndex"

	// FieldProduceResponseResponsesPartitionResponsesRecordErrorsBatchIndexErrorMessage is: The error message of the record that caused the batch to be dropped
	FieldProduceResponseResponsesPartitionResponsesRecordErrorsBatchIndexErrorMessage = "BatchIndexErrorMessage"

	// FieldProduceResponseResponsesPartitionResponsesRecordErrorsTags is: The tagged fields.
	FieldProduceResponseResponsesPartitionResponsesRecordErrorsTags = "Tags"

	// FieldProduceResponseResponsesPartitionResponsesTags is: The tagged fields.
	FieldProduceResponseResponsesPartitionResponsesTags = "Tags"

	// FieldProduceResponseResponsesTags is: The tagged fields.
	FieldProduceResponseResponsesTags = "Tags"

	// FieldProduceResponseTags is: The tagged fields.
	FieldProduceResponseTags = "Tags"

	// FieldProduceResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldProduceResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ProduceResponse.json
const originalProduceResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 0,
  "type": "response",
  "name": "ProduceResponse",
  // Version 1 added the throttle time.
  //
  // Version 2 added the log append time.
  //
  // Version 3 is the same as version 2.
  //
  // Version 4 added KAFKA_STORAGE_ERROR as a possible error code.
  //
  // Version 5 added LogStartOffset to filter out spurious
  // OutOfOrderSequenceExceptions on the client.
  //
  // Version 8 added RecordErrors and ErrorMessage to include information about
  // records that cause the whole batch to be dropped.  See KIP-467 for details.
  //
  // Version 9 enables flexible versions.
  "validVersions": "0-9",
  "flexibleVersions": "9+",
  "fields": [
    { "name": "Responses", "type": "[]TopicProduceResponse", "versions": "0+",
      "about": "Each produce response", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName", "mapKey": true,
        "about": "The topic name" },
      { "name": "PartitionResponses", "type": "[]PartitionProduceResponse", "versions": "0+",
        "about": "Each partition that we produced to within the topic.", "fields": [
        { "name": "Index", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "ErrorCode", "type": "int16", "versions": "0+",
          "about": "The error code, or 0 if there was no error." },
        { "name": "BaseOffset", "type": "int64", "versions": "0+",
          "about": "The base offset." },
        { "name": "LogAppendTimeMs", "type": "int64", "versions": "2+", "default": "-1", "ignorable": true,
          "about": "The timestamp returned by broker after appending the messages. If CreateTime is used for the topic, the timestamp will be -1.  If LogAppendTime is used for the topic, the timestamp will be the broker local time when the messages are appended." },
        { "name": "LogStartOffset", "type": "int64", "versions": "5+", "default": "-1", "ignorable": true,
          "about": "The log start offset." },
        { "name": "RecordErrors", "type": "[]BatchIndexAndErrorMessage", "versions": "8+", "ignorable": true,
          "about": "The batch indices of records that caused the batch to be dropped", "fields": [
          { "name": "BatchIndex", "type": "int32", "versions":  "8+",
            "about": "The batch index of the record that cause the batch to be dropped" },
          { "name": "BatchIndexErrorMessage", "type": "string", "default": "null", "versions": "8+", "nullableVersions": "8+",
            "about": "The error message of the record that caused the batch to be dropped"}
        ]},
        { "name":  "ErrorMessage", "type": "string", "default": "null", "versions": "8+", "nullableVersions": "8+", "ignorable":  true,
          "about":  "The global error message summarizing the common root cause of the records that caused the batch to be dropped"}
      ]}
    ]},
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "1+", "ignorable": true, "default": "0",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." }
  ]
}
`
