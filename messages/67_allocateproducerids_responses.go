package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init67AllocateProducerIdsResponse() []schema.Schema {

	return []schema.Schema{
		// Message: AllocateProducerIdsResponse, API Key: 67, Version: 0
		schema.NewSchema("AllocateProducerIdsResponse:v0",
			&schema.Mfield{Name: FieldAllocateProducerIdsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldAllocateProducerIdsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldAllocateProducerIdsResponseProducerIdStart, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldAllocateProducerIdsResponseProducerIdLen, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldAllocateProducerIdsResponseTags},
		),
	}

}

const (

	// FieldAllocateProducerIdsResponseErrorCode is: The top level response error code
	FieldAllocateProducerIdsResponseErrorCode = "ErrorCode"

	// FieldAllocateProducerIdsResponseProducerIdLen is: The number of producer IDs in this range
	FieldAllocateProducerIdsResponseProducerIdLen = "ProducerIdLen"

	// FieldAllocateProducerIdsResponseProducerIdStart is: The first producer ID in this range, inclusive
	FieldAllocateProducerIdsResponseProducerIdStart = "ProducerIdStart"

	// FieldAllocateProducerIdsResponseTags is: The tagged fields.
	FieldAllocateProducerIdsResponseTags = "Tags"

	// FieldAllocateProducerIdsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldAllocateProducerIdsResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AllocateProducerIdsResponse.json
const originalAllocateProducerIdsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 67,
  "type": "response",
  "name": "AllocateProducerIdsResponse",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The top level response error code" },
    { "name": "ProducerIdStart", "type": "int64", "versions": "0+", "entityType": "producerId",
      "about": "The first producer ID in this range, inclusive"},
    { "name": "ProducerIdLen", "type": "int32", "versions": "0+",
      "about": "The number of producer IDs in this range"}
  ]
}
`
