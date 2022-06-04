package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init24AddPartitionsToTxnResponse() []schema.Schema {

	return []schema.Schema{

		// Message: AddPartitionsToTxnResponse, API Key: 24, Version: 0
		schema.NewSchema("AddPartitionsToTxnResponsev0",
			&schema.Mfield{Name: FieldAddPartitionsToTxnResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldAddPartitionsToTxnResponseResults, Ty: schema.NewSchema("ResultsV0",
				&schema.Mfield{Name: FieldAddPartitionsToTxnResponseResultsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldAddPartitionsToTxnResponseResultsResults, Ty: schema.NewSchema("ResultsV0",
					&schema.Mfield{Name: FieldAddPartitionsToTxnResponseResultsResultsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAddPartitionsToTxnResponseResultsResultsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: AddPartitionsToTxnResponse, API Key: 24, Version: 1
		schema.NewSchema("AddPartitionsToTxnResponsev1",
			&schema.Mfield{Name: FieldAddPartitionsToTxnResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldAddPartitionsToTxnResponseResults, Ty: schema.NewSchema("ResultsV1",
				&schema.Mfield{Name: FieldAddPartitionsToTxnResponseResultsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldAddPartitionsToTxnResponseResultsResults, Ty: schema.NewSchema("ResultsV1",
					&schema.Mfield{Name: FieldAddPartitionsToTxnResponseResultsResultsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAddPartitionsToTxnResponseResultsResultsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: AddPartitionsToTxnResponse, API Key: 24, Version: 2
		schema.NewSchema("AddPartitionsToTxnResponsev2",
			&schema.Mfield{Name: FieldAddPartitionsToTxnResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldAddPartitionsToTxnResponseResults, Ty: schema.NewSchema("ResultsV2",
				&schema.Mfield{Name: FieldAddPartitionsToTxnResponseResultsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldAddPartitionsToTxnResponseResultsResults, Ty: schema.NewSchema("ResultsV2",
					&schema.Mfield{Name: FieldAddPartitionsToTxnResponseResultsResultsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAddPartitionsToTxnResponseResultsResultsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: AddPartitionsToTxnResponse, API Key: 24, Version: 3
		schema.NewSchema("AddPartitionsToTxnResponsev3",
			&schema.Mfield{Name: FieldAddPartitionsToTxnResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldAddPartitionsToTxnResponseResults, Ty: schema.NewSchema("ResultsV3",
				&schema.Mfield{Name: FieldAddPartitionsToTxnResponseResultsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldAddPartitionsToTxnResponseResultsResults, Ty: schema.NewSchema("ResultsV3",
					&schema.Mfield{Name: FieldAddPartitionsToTxnResponseResultsResultsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAddPartitionsToTxnResponseResultsResultsErrorCode, Ty: schema.TypeInt16},
					&schema.SchemaTaggedFields{Name: FieldAddPartitionsToTxnResponseResultsResultsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldAddPartitionsToTxnResponseResultsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldAddPartitionsToTxnResponseTags},
		),
	}
}

const (
	// FieldAddPartitionsToTxnResponseResults is: The results for each topic.
	FieldAddPartitionsToTxnResponseResults = "Results"
	// FieldAddPartitionsToTxnResponseResultsName is: The topic name.
	FieldAddPartitionsToTxnResponseResultsName = "Name"
	// FieldAddPartitionsToTxnResponseResultsResults is: The results for each partition
	FieldAddPartitionsToTxnResponseResultsResults = "Results"
	// FieldAddPartitionsToTxnResponseResultsResultsErrorCode is: The response error code.
	FieldAddPartitionsToTxnResponseResultsResultsErrorCode = "ErrorCode"
	// FieldAddPartitionsToTxnResponseResultsResultsPartitionIndex is: The partition indexes.
	FieldAddPartitionsToTxnResponseResultsResultsPartitionIndex = "PartitionIndex"
	// FieldAddPartitionsToTxnResponseResultsResultsTags is: The tagged fields.
	FieldAddPartitionsToTxnResponseResultsResultsTags = "Tags"
	// FieldAddPartitionsToTxnResponseResultsTags is: The tagged fields.
	FieldAddPartitionsToTxnResponseResultsTags = "Tags"
	// FieldAddPartitionsToTxnResponseTags is: The tagged fields.
	FieldAddPartitionsToTxnResponseTags = "Tags"
	// FieldAddPartitionsToTxnResponseThrottleTimeMs is: Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldAddPartitionsToTxnResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AddPartitionsToTxnResponse.json
const originalAddPartitionsToTxnResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 24,
  "type": "response",
  "name": "AddPartitionsToTxnResponse",
  // Starting in version 1, on quota violation brokers send out responses before throttling.
  //
  // Version 2 adds the support for new error code PRODUCER_FENCED.
  //
  // Version 3 enables flexible versions.
  "validVersions": "0-3",
  "flexibleVersions": "3+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Results", "type": "[]AddPartitionsToTxnTopicResult", "versions": "0+",
      "about": "The results for each topic.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "mapKey": true, "entityType": "topicName",
        "about": "The topic name." },
      { "name": "Results", "type": "[]AddPartitionsToTxnPartitionResult", "versions": "0+", 
        "about": "The results for each partition", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+", "mapKey": true,
          "about": "The partition indexes." },
        { "name": "ErrorCode", "type": "int16", "versions": "0+",
          "about": "The response error code."}
      ]}
    ]}
  ]
}
`
