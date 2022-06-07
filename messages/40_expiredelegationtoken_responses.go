package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init40ExpireDelegationTokenResponse() []schema.Schema {

	return []schema.Schema{
		// Message: ExpireDelegationTokenResponse, API Key: 40, Version: 0
		schema.NewSchema("ExpireDelegationTokenResponse:v0",
			&schema.Mfield{Name: FieldExpireDelegationTokenResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldExpireDelegationTokenResponseExpiryTimestampMs, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldExpireDelegationTokenResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: ExpireDelegationTokenResponse, API Key: 40, Version: 1
		schema.NewSchema("ExpireDelegationTokenResponse:v1",
			&schema.Mfield{Name: FieldExpireDelegationTokenResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldExpireDelegationTokenResponseExpiryTimestampMs, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldExpireDelegationTokenResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: ExpireDelegationTokenResponse, API Key: 40, Version: 2
		schema.NewSchema("ExpireDelegationTokenResponse:v2",
			&schema.Mfield{Name: FieldExpireDelegationTokenResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldExpireDelegationTokenResponseExpiryTimestampMs, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldExpireDelegationTokenResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldExpireDelegationTokenResponseTags},
		),
	}

}

const (

	// FieldExpireDelegationTokenResponseErrorCode is: The error code, or 0 if there was no error.
	FieldExpireDelegationTokenResponseErrorCode = "ErrorCode"

	// FieldExpireDelegationTokenResponseExpiryTimestampMs is: The timestamp in milliseconds at which this token expires.
	FieldExpireDelegationTokenResponseExpiryTimestampMs = "ExpiryTimestampMs"

	// FieldExpireDelegationTokenResponseTags is: The tagged fields.
	FieldExpireDelegationTokenResponseTags = "Tags"

	// FieldExpireDelegationTokenResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldExpireDelegationTokenResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ExpireDelegationTokenResponse.json
const originalExpireDelegationTokenResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 40,
  "type": "response",
  "name": "ExpireDelegationTokenResponse",
  // Starting in version 1, on quota violation, brokers send out responses before throttling.
  // Version 2 adds flexible version support
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The error code, or 0 if there was no error." },
    { "name": "ExpiryTimestampMs", "type": "int64", "versions": "0+",
      "about": "The timestamp in milliseconds at which this token expires." },
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." }
  ]
}
`
