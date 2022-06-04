package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init42DeleteGroupsResponse() []schema.Schema {

	return []schema.Schema{

		// Message: DeleteGroupsResponse, API Key: 42, Version: 0
		schema.NewSchema("DeleteGroupsResponsev0",
			&schema.Mfield{Name: FieldDeleteGroupsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDeleteGroupsResponseResults, Ty: schema.NewSchema("ResultsV0",
				&schema.Mfield{Name: FieldDeleteGroupsResponseResultsGroupId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDeleteGroupsResponseResultsErrorCode, Ty: schema.TypeInt16},
			)},
		),

		// Message: DeleteGroupsResponse, API Key: 42, Version: 1
		schema.NewSchema("DeleteGroupsResponsev1",
			&schema.Mfield{Name: FieldDeleteGroupsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDeleteGroupsResponseResults, Ty: schema.NewSchema("ResultsV1",
				&schema.Mfield{Name: FieldDeleteGroupsResponseResultsGroupId, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDeleteGroupsResponseResultsErrorCode, Ty: schema.TypeInt16},
			)},
		),

		// Message: DeleteGroupsResponse, API Key: 42, Version: 2
		schema.NewSchema("DeleteGroupsResponsev2",
			&schema.Mfield{Name: FieldDeleteGroupsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldDeleteGroupsResponseResults, Ty: schema.NewSchema("ResultsV2",
				&schema.Mfield{Name: FieldDeleteGroupsResponseResultsGroupId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDeleteGroupsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.SchemaTaggedFields{Name: FieldDeleteGroupsResponseResultsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDeleteGroupsResponseTags},
		),
	}
}

const (
	// FieldDeleteGroupsResponseResults is: The deletion results
	FieldDeleteGroupsResponseResults = "Results"
	// FieldDeleteGroupsResponseResultsGroupId is: The group id
	FieldDeleteGroupsResponseResultsGroupId = "GroupId"
	// FieldDeleteGroupsResponseResultsErrorCode is: The deletion error, or 0 if the deletion succeeded.
	FieldDeleteGroupsResponseResultsErrorCode = "ErrorCode"
	// FieldDeleteGroupsResponseResultsTags is: The tagged fields.
	FieldDeleteGroupsResponseResultsTags = "Tags"
	// FieldDeleteGroupsResponseTags is: The tagged fields.
	FieldDeleteGroupsResponseTags = "Tags"
	// FieldDeleteGroupsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldDeleteGroupsResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DeleteGroupsResponse.json
const originalDeleteGroupsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 42,
  "type": "response",
  "name": "DeleteGroupsResponse",
  // Starting in version 1, on quota violation, brokers send out responses before throttling.
  //
  // Version 2 is the first flexible version.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Results", "type": "[]DeletableGroupResult", "versions": "0+",
      "about": "The deletion results", "fields": [
      { "name": "GroupId", "type": "string", "versions": "0+", "mapKey": true, "entityType": "groupId",
        "about": "The group id" },
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The deletion error, or 0 if the deletion succeeded." }
    ]}
  ]
}
`
