package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init26EndTxnResponse() []schema.Schema {

	return []schema.Schema{

		// Message: EndTxnResponse, API Key: 26, Version: 0
		schema.NewSchema("EndTxnResponsev0",
			&schema.Mfield{Name: FieldEndTxnResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldEndTxnResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: EndTxnResponse, API Key: 26, Version: 1
		schema.NewSchema("EndTxnResponsev1",
			&schema.Mfield{Name: FieldEndTxnResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldEndTxnResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: EndTxnResponse, API Key: 26, Version: 2
		schema.NewSchema("EndTxnResponsev2",
			&schema.Mfield{Name: FieldEndTxnResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldEndTxnResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: EndTxnResponse, API Key: 26, Version: 3
		schema.NewSchema("EndTxnResponsev3",
			&schema.Mfield{Name: FieldEndTxnResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldEndTxnResponseErrorCode, Ty: schema.TypeInt16},
			&schema.SchemaTaggedFields{Name: FieldEndTxnResponseTags},
		),
	}
}

const (
	// FieldEndTxnResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldEndTxnResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldEndTxnResponseErrorCode is: The error code, or 0 if there was no error.
	FieldEndTxnResponseErrorCode = "ErrorCode"
	// FieldEndTxnResponseTags is: The tagged fields.
	FieldEndTxnResponseTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/EndTxnResponse.json
const originalEndTxnResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 26,
  "type": "response",
  "name": "EndTxnResponse",
  // Starting in version 1, on quota violation, brokers send out responses before throttling.
  //
  // Version 2 adds the support for new error code PRODUCER_FENCED.
  //
  // Version 3 enables flexible versions.
  "validVersions": "0-3",
  "flexibleVersions": "3+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The error code, or 0 if there was no error." }
  ]
}
`
