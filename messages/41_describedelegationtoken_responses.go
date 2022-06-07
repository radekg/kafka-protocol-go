package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init41DescribeDelegationTokenResponse() []schema.Schema {

	return []schema.Schema{
		// Message: DescribeDelegationTokenResponse, API Key: 41, Version: 0
		schema.NewSchema("DescribeDelegationTokenResponse:v0",
			&schema.Mfield{Name: FieldDescribeDelegationTokenResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldDescribeDelegationTokenResponseTokens, Ty: schema.NewSchema("[]DescribedDelegationToken:v0",
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensPrincipalType, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensPrincipalName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensIssueTimestamp, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensExpiryTimestamp, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensMaxTimestamp, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensTokenId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensHmac, Ty: schema.TypeBytes},
				&schema.Array{Name: FieldDescribeDelegationTokenResponseTokensRenewers, Ty: schema.NewSchema("[]DescribedDelegationTokenRenewer:v0",
					&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensRenewersPrincipalType, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensRenewersPrincipalName, Ty: schema.TypeStr},
				)},
			)},
			&schema.Mfield{Name: FieldDescribeDelegationTokenResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: DescribeDelegationTokenResponse, API Key: 41, Version: 1
		schema.NewSchema("DescribeDelegationTokenResponse:v1",
			&schema.Mfield{Name: FieldDescribeDelegationTokenResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldDescribeDelegationTokenResponseTokens, Ty: schema.NewSchema("[]DescribedDelegationToken:v1",
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensPrincipalType, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensPrincipalName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensIssueTimestamp, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensExpiryTimestamp, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensMaxTimestamp, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensTokenId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensHmac, Ty: schema.TypeBytes},
				&schema.Array{Name: FieldDescribeDelegationTokenResponseTokensRenewers, Ty: schema.NewSchema("[]DescribedDelegationTokenRenewer:v1",
					&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensRenewersPrincipalType, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensRenewersPrincipalName, Ty: schema.TypeStr},
				)},
			)},
			&schema.Mfield{Name: FieldDescribeDelegationTokenResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: DescribeDelegationTokenResponse, API Key: 41, Version: 2
		schema.NewSchema("DescribeDelegationTokenResponse:v2",
			&schema.Mfield{Name: FieldDescribeDelegationTokenResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldDescribeDelegationTokenResponseTokens, Ty: schema.NewSchema("[]DescribedDelegationToken:v2",
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensPrincipalType, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensPrincipalName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensIssueTimestamp, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensExpiryTimestamp, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensMaxTimestamp, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensTokenId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensHmac, Ty: schema.TypeBytesCompact},
				&schema.ArrayCompact{Name: FieldDescribeDelegationTokenResponseTokensRenewers, Ty: schema.NewSchema("[]DescribedDelegationTokenRenewer:v2",
					&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensRenewersPrincipalType, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldDescribeDelegationTokenResponseTokensRenewersPrincipalName, Ty: schema.TypeStrCompact},
					&schema.SchemaTaggedFields{Name: FieldDescribeDelegationTokenResponseTokensRenewersTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldDescribeDelegationTokenResponseTokensTags},
			)},
			&schema.Mfield{Name: FieldDescribeDelegationTokenResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldDescribeDelegationTokenResponseTags},
		),
	}

}

const (

	// FieldDescribeDelegationTokenResponseErrorCode is: The error code, or 0 if there was no error.
	FieldDescribeDelegationTokenResponseErrorCode = "ErrorCode"

	// FieldDescribeDelegationTokenResponseTags is: The tagged fields.
	FieldDescribeDelegationTokenResponseTags = "Tags"

	// FieldDescribeDelegationTokenResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldDescribeDelegationTokenResponseThrottleTimeMs = "ThrottleTimeMs"

	// FieldDescribeDelegationTokenResponseTokens is: The tokens.
	FieldDescribeDelegationTokenResponseTokens = "Tokens"

	// FieldDescribeDelegationTokenResponseTokensExpiryTimestamp is: The token expiry timestamp in milliseconds.
	FieldDescribeDelegationTokenResponseTokensExpiryTimestamp = "ExpiryTimestamp"

	// FieldDescribeDelegationTokenResponseTokensHmac is: The token HMAC.
	FieldDescribeDelegationTokenResponseTokensHmac = "Hmac"

	// FieldDescribeDelegationTokenResponseTokensIssueTimestamp is: The token issue timestamp in milliseconds.
	FieldDescribeDelegationTokenResponseTokensIssueTimestamp = "IssueTimestamp"

	// FieldDescribeDelegationTokenResponseTokensMaxTimestamp is: The token maximum timestamp length in milliseconds.
	FieldDescribeDelegationTokenResponseTokensMaxTimestamp = "MaxTimestamp"

	// FieldDescribeDelegationTokenResponseTokensPrincipalName is: The token principal name.
	FieldDescribeDelegationTokenResponseTokensPrincipalName = "PrincipalName"

	// FieldDescribeDelegationTokenResponseTokensPrincipalType is: The token principal type.
	FieldDescribeDelegationTokenResponseTokensPrincipalType = "PrincipalType"

	// FieldDescribeDelegationTokenResponseTokensRenewers is: Those who are able to renew this token before it expires.
	FieldDescribeDelegationTokenResponseTokensRenewers = "Renewers"

	// FieldDescribeDelegationTokenResponseTokensRenewersPrincipalName is: The renewer principal name
	FieldDescribeDelegationTokenResponseTokensRenewersPrincipalName = "PrincipalName"

	// FieldDescribeDelegationTokenResponseTokensRenewersPrincipalType is: The renewer principal type
	FieldDescribeDelegationTokenResponseTokensRenewersPrincipalType = "PrincipalType"

	// FieldDescribeDelegationTokenResponseTokensRenewersTags is: The tagged fields.
	FieldDescribeDelegationTokenResponseTokensRenewersTags = "Tags"

	// FieldDescribeDelegationTokenResponseTokensTags is: The tagged fields.
	FieldDescribeDelegationTokenResponseTokensTags = "Tags"

	// FieldDescribeDelegationTokenResponseTokensTokenId is: The token ID.
	FieldDescribeDelegationTokenResponseTokensTokenId = "TokenId"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeDelegationTokenResponse.json
const originalDescribeDelegationTokenResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 41,
  "type": "response",
  "name": "DescribeDelegationTokenResponse",
  // Starting in version 1, on quota violation, brokers send out responses before throttling.
  // Version 2 adds flexible version support
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The error code, or 0 if there was no error." },
    { "name": "Tokens", "type": "[]DescribedDelegationToken", "versions": "0+",
      "about": "The tokens.", "fields": [
      { "name": "PrincipalType", "type": "string", "versions": "0+",
        "about": "The token principal type." },
      { "name": "PrincipalName", "type": "string", "versions": "0+",
        "about": "The token principal name." },
      { "name": "IssueTimestamp", "type": "int64", "versions": "0+",
        "about": "The token issue timestamp in milliseconds." },
      { "name": "ExpiryTimestamp", "type": "int64", "versions": "0+",
        "about": "The token expiry timestamp in milliseconds." },
      { "name": "MaxTimestamp", "type": "int64", "versions": "0+",
        "about": "The token maximum timestamp length in milliseconds." },
      { "name": "TokenId", "type": "string", "versions": "0+",
        "about": "The token ID." },
      { "name": "Hmac", "type": "bytes", "versions": "0+",
        "about": "The token HMAC." },
      { "name": "Renewers", "type": "[]DescribedDelegationTokenRenewer", "versions": "0+",
        "about": "Those who are able to renew this token before it expires.", "fields": [
        { "name": "PrincipalType", "type": "string", "versions": "0+",
          "about": "The renewer principal type" },
        { "name": "PrincipalName", "type": "string", "versions": "0+",
          "about": "The renewer principal name" }
      ]}
    ]},
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." }
  ]
}
`
