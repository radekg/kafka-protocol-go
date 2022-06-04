package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init38CreateDelegationTokenResponse() []schema.Schema {

	return []schema.Schema{

		// Message: CreateDelegationTokenResponse, API Key: 38, Version: 0
		schema.NewSchema("CreateDelegationTokenResponsev0",
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponsePrincipalType, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponsePrincipalName, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseIssueTimestampMs, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseExpiryTimestampMs, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseMaxTimestampMs, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseTokenId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseHmac, Ty: schema.TypeBytes},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: CreateDelegationTokenResponse, API Key: 38, Version: 1
		schema.NewSchema("CreateDelegationTokenResponsev1",
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponsePrincipalType, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponsePrincipalName, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseIssueTimestampMs, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseExpiryTimestampMs, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseMaxTimestampMs, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseTokenId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseHmac, Ty: schema.TypeBytes},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: CreateDelegationTokenResponse, API Key: 38, Version: 2
		schema.NewSchema("CreateDelegationTokenResponsev2",
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponsePrincipalType, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponsePrincipalName, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseIssueTimestampMs, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseExpiryTimestampMs, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseMaxTimestampMs, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseTokenId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseHmac, Ty: schema.TypeBytesCompact},
			&schema.Mfield{Name: FieldCreateDelegationTokenResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldCreateDelegationTokenResponseTags},
		),
	}
}

const (
	// FieldCreateDelegationTokenResponseErrorCode is: The top-level error, or zero if there was no error.
	FieldCreateDelegationTokenResponseErrorCode = "ErrorCode"
	// FieldCreateDelegationTokenResponseExpiryTimestampMs is: When this token expires.
	FieldCreateDelegationTokenResponseExpiryTimestampMs = "ExpiryTimestampMs"
	// FieldCreateDelegationTokenResponseHmac is: HMAC of the delegation token.
	FieldCreateDelegationTokenResponseHmac = "Hmac"
	// FieldCreateDelegationTokenResponseIssueTimestampMs is: When this token was generated.
	FieldCreateDelegationTokenResponseIssueTimestampMs = "IssueTimestampMs"
	// FieldCreateDelegationTokenResponseMaxTimestampMs is: The maximum lifetime of this token.
	FieldCreateDelegationTokenResponseMaxTimestampMs = "MaxTimestampMs"
	// FieldCreateDelegationTokenResponsePrincipalName is: The name of the token owner.
	FieldCreateDelegationTokenResponsePrincipalName = "PrincipalName"
	// FieldCreateDelegationTokenResponsePrincipalType is: The principal type of the token owner.
	FieldCreateDelegationTokenResponsePrincipalType = "PrincipalType"
	// FieldCreateDelegationTokenResponseTags is: The tagged fields.
	FieldCreateDelegationTokenResponseTags = "Tags"
	// FieldCreateDelegationTokenResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldCreateDelegationTokenResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldCreateDelegationTokenResponseTokenId is: The token UUID.
	FieldCreateDelegationTokenResponseTokenId = "TokenId"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/CreateDelegationTokenResponse.json
const originalCreateDelegationTokenResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 38,
  "type": "response",
  "name": "CreateDelegationTokenResponse",
  // Starting in version 1, on quota violation, brokers send out responses before throttling.
  //
  // Version 2 is the first flexible version.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The top-level error, or zero if there was no error."},
    { "name": "PrincipalType", "type": "string", "versions": "0+",
      "about": "The principal type of the token owner." },
    { "name": "PrincipalName", "type": "string", "versions": "0+",
      "about": "The name of the token owner." },
    { "name": "IssueTimestampMs", "type": "int64", "versions": "0+",
      "about": "When this token was generated." },
    { "name": "ExpiryTimestampMs", "type": "int64", "versions": "0+",
      "about": "When this token expires." },
    { "name": "MaxTimestampMs", "type": "int64", "versions": "0+",
      "about": "The maximum lifetime of this token." },
    { "name": "TokenId", "type": "string", "versions": "0+",
      "about": "The token UUID." },
    { "name": "Hmac", "type": "bytes", "versions": "0+",
      "about": "HMAC of the delegation token." },
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." }
  ]
}
`
