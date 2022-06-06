package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init57UpdateFeaturesResponse() []schema.Schema {

	return []schema.Schema{
		// Message: UpdateFeaturesResponse, API Key: 57, Version: 0
		schema.NewSchema("UpdateFeaturesResponse:v0",
			&schema.Mfield{Name: FieldUpdateFeaturesResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldUpdateFeaturesResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldUpdateFeaturesResponseErrorMessage, Ty: schema.TypeStrCompactNullable},
			&schema.ArrayCompact{Name: FieldUpdateFeaturesResponseResults, Ty: schema.NewSchema("Results:v0",
				&schema.Mfield{Name: FieldUpdateFeaturesResponseResultsFeature, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldUpdateFeaturesResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldUpdateFeaturesResponseResultsErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldUpdateFeaturesResponseResultsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldUpdateFeaturesResponseTags},
		),
	}

}

const (

	// FieldUpdateFeaturesResponseErrorCode is: The top-level error code, or `0` if there was no top-level error.
	FieldUpdateFeaturesResponseErrorCode = "ErrorCode"

	// FieldUpdateFeaturesResponseErrorMessage is: The top-level error message, or `null` if there was no top-level error.
	FieldUpdateFeaturesResponseErrorMessage = "ErrorMessage"

	// FieldUpdateFeaturesResponseResults is: Results for each feature update.
	FieldUpdateFeaturesResponseResults = "Results"

	// FieldUpdateFeaturesResponseResultsErrorCode is: The feature update error code or `0` if the feature update succeeded.
	FieldUpdateFeaturesResponseResultsErrorCode = "ErrorCode"

	// FieldUpdateFeaturesResponseResultsErrorMessage is: The feature update error, or `null` if the feature update succeeded.
	FieldUpdateFeaturesResponseResultsErrorMessage = "ErrorMessage"

	// FieldUpdateFeaturesResponseResultsFeature is: The name of the finalized feature.
	FieldUpdateFeaturesResponseResultsFeature = "Feature"

	// FieldUpdateFeaturesResponseResultsTags is: The tagged fields.
	FieldUpdateFeaturesResponseResultsTags = "Tags"

	// FieldUpdateFeaturesResponseTags is: The tagged fields.
	FieldUpdateFeaturesResponseTags = "Tags"

	// FieldUpdateFeaturesResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldUpdateFeaturesResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/UpdateFeaturesResponse.json
const originalUpdateFeaturesResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 57,
  "type": "response",
  "name": "UpdateFeaturesResponse",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The top-level error code, or '0' if there was no top-level error." },
    { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
      "about": "The top-level error message, or 'null' if there was no top-level error." },
    { "name": "Results", "type": "[]UpdatableFeatureResult", "versions": "0+",
      "about": "Results for each feature update.", "fields": [
      { "name": "Feature", "type": "string", "versions": "0+", "mapKey": true,
        "about": "The name of the finalized feature."},
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The feature update error code or '0' if the feature update succeeded." },
      { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
        "about": "The feature update error, or 'null' if the feature update succeeded." }
    ]}
  ]
}
`
