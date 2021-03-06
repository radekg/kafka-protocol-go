package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init30CreateAclsResponse() []schema.Schema {

	return []schema.Schema{

		// Message: CreateAclsResponse, API Key: 30, Version: 0
		schema.NewSchema("CreateAclsResponsev0",
			&schema.Mfield{Name: FieldCreateAclsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldCreateAclsResponseResults, Ty: schema.NewSchema("ResultsV0",
				&schema.Mfield{Name: FieldCreateAclsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreateAclsResponseResultsErrorMessage, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: CreateAclsResponse, API Key: 30, Version: 1
		schema.NewSchema("CreateAclsResponsev1",
			&schema.Mfield{Name: FieldCreateAclsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldCreateAclsResponseResults, Ty: schema.NewSchema("ResultsV1",
				&schema.Mfield{Name: FieldCreateAclsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreateAclsResponseResultsErrorMessage, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: CreateAclsResponse, API Key: 30, Version: 2
		schema.NewSchema("CreateAclsResponsev2",
			&schema.Mfield{Name: FieldCreateAclsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldCreateAclsResponseResults, Ty: schema.NewSchema("ResultsV2",
				&schema.Mfield{Name: FieldCreateAclsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreateAclsResponseResultsErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldCreateAclsResponseResultsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldCreateAclsResponseTags},
		),
	}
}

const (
	// FieldCreateAclsResponseResults is: The results for each ACL creation.
	FieldCreateAclsResponseResults = "Results"
	// FieldCreateAclsResponseResultsErrorCode is: The result error, or zero if there was no error.
	FieldCreateAclsResponseResultsErrorCode = "ErrorCode"
	// FieldCreateAclsResponseResultsErrorMessage is: The result message, or null if there was no error.
	FieldCreateAclsResponseResultsErrorMessage = "ErrorMessage"
	// FieldCreateAclsResponseResultsTags is: The tagged fields.
	FieldCreateAclsResponseResultsTags = "Tags"
	// FieldCreateAclsResponseTags is: The tagged fields.
	FieldCreateAclsResponseTags = "Tags"
	// FieldCreateAclsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldCreateAclsResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/CreateAclsResponse.json
const originalCreateAclsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 30,
  "type": "response",
  "name": "CreateAclsResponse",
  // Starting in version 1, on quota violation, brokers send out responses before throttling.
  // Version 2 enables flexible versions.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Results", "type": "[]AclCreationResult", "versions": "0+",
      "about": "The results for each ACL creation.", "fields": [
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The result error, or zero if there was no error." },
      { "name": "ErrorMessage", "type": "string", "nullableVersions": "0+", "versions": "0+",
        "about": "The result message, or null if there was no error." }
    ]}
  ]
}
`
