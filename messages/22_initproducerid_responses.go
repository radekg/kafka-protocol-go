package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init22InitProducerIdResponse() []schema.Schema {

	return []schema.Schema{
		// Message: InitProducerIdResponse, API Key: 22, Version: 0
		schema.NewSchema("InitProducerIdResponse:v0",
			&schema.Mfield{Name: FieldInitProducerIdResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldInitProducerIdResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldInitProducerIdResponseProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldInitProducerIdResponseProducerEpoch, Ty: schema.TypeInt16},
		),

		// Message: InitProducerIdResponse, API Key: 22, Version: 1
		schema.NewSchema("InitProducerIdResponse:v1",
			&schema.Mfield{Name: FieldInitProducerIdResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldInitProducerIdResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldInitProducerIdResponseProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldInitProducerIdResponseProducerEpoch, Ty: schema.TypeInt16},
		),

		// Message: InitProducerIdResponse, API Key: 22, Version: 2
		schema.NewSchema("InitProducerIdResponse:v2",
			&schema.Mfield{Name: FieldInitProducerIdResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldInitProducerIdResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldInitProducerIdResponseProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldInitProducerIdResponseProducerEpoch, Ty: schema.TypeInt16},
			&schema.SchemaTaggedFields{Name: FieldInitProducerIdResponseTags},
		),

		// Message: InitProducerIdResponse, API Key: 22, Version: 3
		schema.NewSchema("InitProducerIdResponse:v3",
			&schema.Mfield{Name: FieldInitProducerIdResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldInitProducerIdResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldInitProducerIdResponseProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldInitProducerIdResponseProducerEpoch, Ty: schema.TypeInt16},
			&schema.SchemaTaggedFields{Name: FieldInitProducerIdResponseTags},
		),

		// Message: InitProducerIdResponse, API Key: 22, Version: 4
		schema.NewSchema("InitProducerIdResponse:v4",
			&schema.Mfield{Name: FieldInitProducerIdResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldInitProducerIdResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldInitProducerIdResponseProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldInitProducerIdResponseProducerEpoch, Ty: schema.TypeInt16},
			&schema.SchemaTaggedFields{Name: FieldInitProducerIdResponseTags},
		),
	}

}

const (

	// FieldInitProducerIdResponseErrorCode is: The error code, or 0 if there was no error.
	FieldInitProducerIdResponseErrorCode = "ErrorCode"

	// FieldInitProducerIdResponseProducerEpoch is: The current epoch associated with the producer id.
	FieldInitProducerIdResponseProducerEpoch = "ProducerEpoch"

	// FieldInitProducerIdResponseProducerId is: The current producer id.
	FieldInitProducerIdResponseProducerId = "ProducerId"

	// FieldInitProducerIdResponseTags is: The tagged fields.
	FieldInitProducerIdResponseTags = "Tags"

	// FieldInitProducerIdResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldInitProducerIdResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/InitProducerIdResponse.json
const originalInitProducerIdResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 22,
  "type": "response",
  "name": "InitProducerIdResponse",
  // Starting in version 1, on quota violation, brokers send out responses before throttling.
  //
  // Version 2 is the first flexible version.
  //
  // Version 3 is the same as version 2.
  //
  // Version 4 adds the support for new error code PRODUCER_FENCED.
  "validVersions": "0-4",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The error code, or 0 if there was no error." },
    { "name": "ProducerId", "type": "int64", "versions": "0+", "entityType": "producerId",
      "default": -1, "about": "The current producer id." },
    { "name": "ProducerEpoch", "type": "int16", "versions": "0+",
      "about": "The current epoch associated with the producer id." }
  ]
}
`
