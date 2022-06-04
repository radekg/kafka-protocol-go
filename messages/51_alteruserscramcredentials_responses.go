package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init51AlterUserScramCredentialsResponse() []schema.Schema {

	return []schema.Schema{

		// Message: AlterUserScramCredentialsResponse, API Key: 51, Version: 0
		schema.NewSchema("AlterUserScramCredentialsResponsev0",
			&schema.Mfield{Name: FieldAlterUserScramCredentialsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldAlterUserScramCredentialsResponseResults, Ty: schema.NewSchema("ResultsV0",
				&schema.Mfield{Name: FieldAlterUserScramCredentialsResponseResultsUser, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldAlterUserScramCredentialsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldAlterUserScramCredentialsResponseResultsErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldAlterUserScramCredentialsResponseResultsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldAlterUserScramCredentialsResponseTags},
		),
	}
}

const (
	// FieldAlterUserScramCredentialsResponseTags is: The tagged fields.
	FieldAlterUserScramCredentialsResponseTags = "Tags"
	// FieldAlterUserScramCredentialsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldAlterUserScramCredentialsResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldAlterUserScramCredentialsResponseResults is: The results for deletions and alterations, one per affected user.
	FieldAlterUserScramCredentialsResponseResults = "Results"
	// FieldAlterUserScramCredentialsResponseResultsUser is: The user name.
	FieldAlterUserScramCredentialsResponseResultsUser = "User"
	// FieldAlterUserScramCredentialsResponseResultsErrorCode is: The error code.
	FieldAlterUserScramCredentialsResponseResultsErrorCode = "ErrorCode"
	// FieldAlterUserScramCredentialsResponseResultsErrorMessage is: The error message, if any.
	FieldAlterUserScramCredentialsResponseResultsErrorMessage = "ErrorMessage"
	// FieldAlterUserScramCredentialsResponseResultsTags is: The tagged fields.
	FieldAlterUserScramCredentialsResponseResultsTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AlterUserScramCredentialsResponse.json
const originalAlterUserScramCredentialsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 51,
  "type": "response",
  "name": "AlterUserScramCredentialsResponse",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Results", "type": "[]AlterUserScramCredentialsResult", "versions": "0+",
      "about": "The results for deletions and alterations, one per affected user.", "fields": [
      { "name": "User", "type": "string", "versions": "0+",
        "about": "The user name." },
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The error code." },
      { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
        "about": "The error message, if any." }
    ]}
  ]
}
`
