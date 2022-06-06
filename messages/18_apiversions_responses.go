package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init18ApiVersionsResponse() []schema.Schema {

	return []schema.Schema{
		// Message: ApiVersionsResponse, API Key: 18, Version: 0
		schema.NewSchema("ApiVersionsResponse:v0",
			&schema.Mfield{Name: FieldApiVersionsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldApiVersionsResponseApiKeys, Ty: schema.NewSchema("[]ApiVersion:v0",
				&schema.Mfield{Name: FieldApiVersionsResponseApiKeysApiKey, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldApiVersionsResponseApiKeysMinVersion, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldApiVersionsResponseApiKeysMaxVersion, Ty: schema.TypeInt16},
			)},
		),

		// Message: ApiVersionsResponse, API Key: 18, Version: 1
		schema.NewSchema("ApiVersionsResponse:v1",
			&schema.Mfield{Name: FieldApiVersionsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldApiVersionsResponseApiKeys, Ty: schema.NewSchema("[]ApiVersion:v1",
				&schema.Mfield{Name: FieldApiVersionsResponseApiKeysApiKey, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldApiVersionsResponseApiKeysMinVersion, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldApiVersionsResponseApiKeysMaxVersion, Ty: schema.TypeInt16},
			)},
			&schema.Mfield{Name: FieldApiVersionsResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: ApiVersionsResponse, API Key: 18, Version: 2
		schema.NewSchema("ApiVersionsResponse:v2",
			&schema.Mfield{Name: FieldApiVersionsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldApiVersionsResponseApiKeys, Ty: schema.NewSchema("[]ApiVersion:v2",
				&schema.Mfield{Name: FieldApiVersionsResponseApiKeysApiKey, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldApiVersionsResponseApiKeysMinVersion, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldApiVersionsResponseApiKeysMaxVersion, Ty: schema.TypeInt16},
			)},
			&schema.Mfield{Name: FieldApiVersionsResponseThrottleTimeMs, Ty: schema.TypeInt32},
		),

		// Message: ApiVersionsResponse, API Key: 18, Version: 3
		schema.NewSchema("ApiVersionsResponse:v3",
			&schema.Mfield{Name: FieldApiVersionsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldApiVersionsResponseApiKeys, Ty: schema.NewSchema("[]ApiVersion:v3",
				&schema.Mfield{Name: FieldApiVersionsResponseApiKeysApiKey, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldApiVersionsResponseApiKeysMinVersion, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldApiVersionsResponseApiKeysMaxVersion, Ty: schema.TypeInt16},
				&schema.SchemaTaggedFields{Name: FieldApiVersionsResponseApiKeysTags},
			)},
			&schema.Mfield{Name: FieldApiVersionsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldApiVersionsResponseTags},
			/** Applicable tags:

				0: SupportedFeatures (type: []SupportedFeatureKey) =
				&schema.ArrayCompact{Name: FieldApiVersionsResponseSupportedFeatures, Ty: schema.NewSchema("[]SupportedFeatureKey:v3",
					&schema.Mfield{Name: FieldApiVersionsResponseSupportedFeaturesName, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldApiVersionsResponseSupportedFeaturesMinVersion, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldApiVersionsResponseSupportedFeaturesMaxVersion, Ty: schema.TypeInt16},
					&schema.SchemaTaggedFields{Name: FieldApiVersionsResponseSupportedFeaturesTags},

				)},

				1: FinalizedFeaturesEpoch (type: int64) =
				&schema.Mfield{Name: FieldApiVersionsResponseFinalizedFeaturesEpoch, Ty: schema.TypeInt64},

				2: FinalizedFeatures (type: []FinalizedFeatureKey) =
				&schema.ArrayCompact{Name: FieldApiVersionsResponseFinalizedFeatures, Ty: schema.NewSchema("[]FinalizedFeatureKey:v3",
					&schema.Mfield{Name: FieldApiVersionsResponseFinalizedFeaturesName, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldApiVersionsResponseFinalizedFeaturesMaxVersionLevel, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldApiVersionsResponseFinalizedFeaturesMinVersionLevel, Ty: schema.TypeInt16},
					&schema.SchemaTaggedFields{Name: FieldApiVersionsResponseFinalizedFeaturesTags},

				)},

			**/

		),
	}

}

const (

	// FieldApiVersionsResponseApiKeys is: The APIs supported by the broker.
	FieldApiVersionsResponseApiKeys = "ApiKeys"

	// FieldApiVersionsResponseApiKeysApiKey is: The API index.
	FieldApiVersionsResponseApiKeysApiKey = "ApiKey"

	// FieldApiVersionsResponseApiKeysMaxVersion is: The maximum supported version, inclusive.
	FieldApiVersionsResponseApiKeysMaxVersion = "MaxVersion"

	// FieldApiVersionsResponseApiKeysMinVersion is: The minimum supported version, inclusive.
	FieldApiVersionsResponseApiKeysMinVersion = "MinVersion"

	// FieldApiVersionsResponseApiKeysTags is: The tagged fields.
	FieldApiVersionsResponseApiKeysTags = "Tags"

	// FieldApiVersionsResponseErrorCode is: The top-level error code.
	FieldApiVersionsResponseErrorCode = "ErrorCode"

	// FieldApiVersionsResponseFinalizedFeatures is: List of cluster-wide finalized features. The information is valid only if FinalizedFeaturesEpoch >= 0.
	FieldApiVersionsResponseFinalizedFeatures = "FinalizedFeatures"

	// FieldApiVersionsResponseFinalizedFeaturesEpoch is: The monotonically increasing epoch for the finalized features information. Valid values are >= 0. A value of -1 is special and represents unknown epoch.
	FieldApiVersionsResponseFinalizedFeaturesEpoch = "FinalizedFeaturesEpoch"

	// FieldApiVersionsResponseFinalizedFeaturesMaxVersionLevel is: The cluster-wide finalized max version level for the feature.
	FieldApiVersionsResponseFinalizedFeaturesMaxVersionLevel = "MaxVersionLevel"

	// FieldApiVersionsResponseFinalizedFeaturesMinVersionLevel is: The cluster-wide finalized min version level for the feature.
	FieldApiVersionsResponseFinalizedFeaturesMinVersionLevel = "MinVersionLevel"

	// FieldApiVersionsResponseFinalizedFeaturesName is: The name of the feature.
	FieldApiVersionsResponseFinalizedFeaturesName = "Name"

	// FieldApiVersionsResponseFinalizedFeaturesTags is: The tagged fields.
	FieldApiVersionsResponseFinalizedFeaturesTags = "Tags"

	// FieldApiVersionsResponseSupportedFeatures is: Features supported by the broker.
	FieldApiVersionsResponseSupportedFeatures = "SupportedFeatures"

	// FieldApiVersionsResponseSupportedFeaturesMaxVersion is: The maximum supported version for the feature.
	FieldApiVersionsResponseSupportedFeaturesMaxVersion = "MaxVersion"

	// FieldApiVersionsResponseSupportedFeaturesMinVersion is: The minimum supported version for the feature.
	FieldApiVersionsResponseSupportedFeaturesMinVersion = "MinVersion"

	// FieldApiVersionsResponseSupportedFeaturesName is: The name of the feature.
	FieldApiVersionsResponseSupportedFeaturesName = "Name"

	// FieldApiVersionsResponseSupportedFeaturesTags is: The tagged fields.
	FieldApiVersionsResponseSupportedFeaturesTags = "Tags"

	// FieldApiVersionsResponseTags is: The tagged fields.
	FieldApiVersionsResponseTags = "Tags"

	// FieldApiVersionsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldApiVersionsResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ApiVersionsResponse.json
const originalApiVersionsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 18,
  "type": "response",
  "name": "ApiVersionsResponse",
  // Version 1 adds throttle time to the response.
  //
  // Starting in version 2, on quota violation, brokers send out responses before throttling.
  //
  // Version 3 is the first flexible version. Tagged fields are only supported in the body but
  // not in the header. The length of the header must not change in order to guarantee the
  // backward compatibility.
  //
  // Starting from Apache Kafka 2.4 (KIP-511), ApiKeys field is populated with the supported
  // versions of the ApiVersionsRequest when an UNSUPPORTED_VERSION error is returned.
  "validVersions": "0-3",
  "flexibleVersions": "3+",
  "fields": [
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The top-level error code." },
    { "name": "ApiKeys", "type": "[]ApiVersion", "versions": "0+",
      "about": "The APIs supported by the broker.", "fields": [
      { "name": "ApiKey", "type": "int16", "versions": "0+", "mapKey": true,
        "about": "The API index." },
      { "name": "MinVersion", "type": "int16", "versions": "0+",
        "about": "The minimum supported version, inclusive." },
      { "name": "MaxVersion", "type": "int16", "versions": "0+",
        "about": "The maximum supported version, inclusive." }
    ]},
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "1+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name":  "SupportedFeatures", "type": "[]SupportedFeatureKey", "ignorable": true,
      "versions":  "3+", "tag": 0, "taggedVersions": "3+",
      "about": "Features supported by the broker.",
      "fields":  [
        { "name": "Name", "type": "string", "versions": "3+", "mapKey": true,
          "about": "The name of the feature." },
        { "name": "MinVersion", "type": "int16", "versions": "3+",
          "about": "The minimum supported version for the feature." },
        { "name": "MaxVersion", "type": "int16", "versions": "3+",
          "about": "The maximum supported version for the feature." }
      ]
    },
    { "name": "FinalizedFeaturesEpoch", "type": "int64", "versions": "3+",
      "tag": 1, "taggedVersions": "3+", "default": "-1", "ignorable": true,
      "about": "The monotonically increasing epoch for the finalized features information. Valid values are >= 0. A value of -1 is special and represents unknown epoch."},
    { "name":  "FinalizedFeatures", "type": "[]FinalizedFeatureKey", "ignorable": true,
      "versions":  "3+", "tag": 2, "taggedVersions": "3+",
      "about": "List of cluster-wide finalized features. The information is valid only if FinalizedFeaturesEpoch >= 0.",
      "fields":  [
        {"name": "Name", "type": "string", "versions":  "3+", "mapKey": true,
          "about": "The name of the feature."},
        {"name":  "MaxVersionLevel", "type": "int16", "versions":  "3+",
          "about": "The cluster-wide finalized max version level for the feature."},
        {"name":  "MinVersionLevel", "type": "int16", "versions":  "3+",
          "about": "The cluster-wide finalized min version level for the feature."}
      ]
    }
  ]
}
`
