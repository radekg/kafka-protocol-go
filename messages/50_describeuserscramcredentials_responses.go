package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init50DescribeUserScramCredentialsResponse() []schema.Schema {

	return []schema.Schema{
		// Message: DescribeUserScramCredentialsResponse, API Key: 50, Version: 0
		schema.NewSchema("DescribeUserScramCredentialsResponse:v0",
			&schema.Mfield{Name: FieldDescribeUserScramCredentialsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldDescribeUserScramCredentialsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldDescribeUserScramCredentialsResponseErrorMessage, Ty: schema.TypeStrCompactNullable},
			&schema.ArrayCompact{Name: FieldDescribeUserScramCredentialsResponseResults, Ty: schema.NewSchema("[]DescribeUserScramCredentialsResult:v0",
				&schema.Mfield{Name: FieldDescribeUserScramCredentialsResponseResultsUser, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeUserScramCredentialsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeUserScramCredentialsResponseResultsErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.ArrayCompact{Name: FieldDescribeUserScramCredentialsResponseResultsCredentialInfos, Ty: schema.NewSchema("[]CredentialInfo:v0",
					&schema.Mfield{Name: FieldDescribeUserScramCredentialsResponseResultsCredentialInfosMechanism, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDescribeUserScramCredentialsResponseResultsCredentialInfosIterations, Ty: schema.TypeInt32},
					&schema.SchemaTaggedFields{Name: FieldDescribeUserScramCredentialsResponseResultsCredentialInfosTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldDescribeUserScramCredentialsResponseResultsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDescribeUserScramCredentialsResponseTags},
		),
	}

}

const (

	// FieldDescribeUserScramCredentialsResponseErrorCode is: The message-level error code, 0 except for user authorization or infrastructure issues.
	FieldDescribeUserScramCredentialsResponseErrorCode = "ErrorCode"

	// FieldDescribeUserScramCredentialsResponseErrorMessage is: The message-level error message, if any.
	FieldDescribeUserScramCredentialsResponseErrorMessage = "ErrorMessage"

	// FieldDescribeUserScramCredentialsResponseResults is: The results for descriptions, one per user.
	FieldDescribeUserScramCredentialsResponseResults = "Results"

	// FieldDescribeUserScramCredentialsResponseResultsCredentialInfos is: The mechanism and related information associated with the user's SCRAM credentials.
	FieldDescribeUserScramCredentialsResponseResultsCredentialInfos = "CredentialInfos"

	// FieldDescribeUserScramCredentialsResponseResultsCredentialInfosIterations is: The number of iterations used in the SCRAM credential.
	FieldDescribeUserScramCredentialsResponseResultsCredentialInfosIterations = "Iterations"

	// FieldDescribeUserScramCredentialsResponseResultsCredentialInfosMechanism is: The SCRAM mechanism.
	FieldDescribeUserScramCredentialsResponseResultsCredentialInfosMechanism = "Mechanism"

	// FieldDescribeUserScramCredentialsResponseResultsCredentialInfosTags is: The tagged fields.
	FieldDescribeUserScramCredentialsResponseResultsCredentialInfosTags = "Tags"

	// FieldDescribeUserScramCredentialsResponseResultsErrorCode is: The user-level error code.
	FieldDescribeUserScramCredentialsResponseResultsErrorCode = "ErrorCode"

	// FieldDescribeUserScramCredentialsResponseResultsErrorMessage is: The user-level error message, if any.
	FieldDescribeUserScramCredentialsResponseResultsErrorMessage = "ErrorMessage"

	// FieldDescribeUserScramCredentialsResponseResultsTags is: The tagged fields.
	FieldDescribeUserScramCredentialsResponseResultsTags = "Tags"

	// FieldDescribeUserScramCredentialsResponseResultsUser is: The user name.
	FieldDescribeUserScramCredentialsResponseResultsUser = "User"

	// FieldDescribeUserScramCredentialsResponseTags is: The tagged fields.
	FieldDescribeUserScramCredentialsResponseTags = "Tags"

	// FieldDescribeUserScramCredentialsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldDescribeUserScramCredentialsResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeUserScramCredentialsResponse.json
const originalDescribeUserScramCredentialsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 50,
  "type": "response",
  "name": "DescribeUserScramCredentialsResponse",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The message-level error code, 0 except for user authorization or infrastructure issues." },
    { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
      "about": "The message-level error message, if any." },
    { "name": "Results", "type": "[]DescribeUserScramCredentialsResult", "versions": "0+",
      "about": "The results for descriptions, one per user.", "fields": [
      { "name": "User", "type": "string", "versions": "0+",
        "about": "The user name." },
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The user-level error code." },
      { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
        "about": "The user-level error message, if any." },
      { "name": "CredentialInfos", "type": "[]CredentialInfo", "versions": "0+",
        "about": "The mechanism and related information associated with the user's SCRAM credentials.", "fields": [
        { "name": "Mechanism", "type": "int8", "versions": "0+",
          "about": "The SCRAM mechanism." },
        { "name": "Iterations", "type": "int32", "versions": "0+",
          "about": "The number of iterations used in the SCRAM credential." }]}
    ]}
  ]
}
`
