package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init57UpdateFeaturesRequest() []schema.Schema {

	return []schema.Schema{

		// Message: UpdateFeaturesRequest, API Key: 57, Version: 0
		schema.NewSchema("UpdateFeaturesRequestv0",
			&schema.Mfield{Name: FieldUpdateFeaturesRequesttimeoutMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldUpdateFeaturesRequestFeatureUpdates, Ty: schema.NewSchema("FeatureUpdatesV0",
				&schema.Mfield{Name: FieldUpdateFeaturesRequestFeatureUpdatesFeature, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldUpdateFeaturesRequestFeatureUpdatesMaxVersionLevel, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldUpdateFeaturesRequestFeatureUpdatesAllowDowngrade, Ty: schema.TypeBool},
				&schema.SchemaTaggedFields{Name: FieldUpdateFeaturesRequestFeatureUpdatesTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldUpdateFeaturesRequestTags},
		),
	}
}

const (
	// FieldUpdateFeaturesRequestFeatureUpdates is: The list of updates to finalized features.
	FieldUpdateFeaturesRequestFeatureUpdates = "FeatureUpdates"
	// FieldUpdateFeaturesRequestFeatureUpdatesAllowDowngrade is: When set to true, the finalized feature version level is allowed to be downgraded/deleted. The downgrade request will fail if the new maximum version level is a value that's not lower than the existing maximum finalized version level.
	FieldUpdateFeaturesRequestFeatureUpdatesAllowDowngrade = "AllowDowngrade"
	// FieldUpdateFeaturesRequestFeatureUpdatesFeature is: The name of the finalized feature to be updated.
	FieldUpdateFeaturesRequestFeatureUpdatesFeature = "Feature"
	// FieldUpdateFeaturesRequestFeatureUpdatesMaxVersionLevel is: The new maximum version level for the finalized feature. A value >= 1 is valid. A value < 1, is special, and can be used to request the deletion of the finalized feature.
	FieldUpdateFeaturesRequestFeatureUpdatesMaxVersionLevel = "MaxVersionLevel"
	// FieldUpdateFeaturesRequestFeatureUpdatesTags is: The tagged fields.
	FieldUpdateFeaturesRequestFeatureUpdatesTags = "Tags"
	// FieldUpdateFeaturesRequestTags is: The tagged fields.
	FieldUpdateFeaturesRequestTags = "Tags"
	// FieldUpdateFeaturesRequesttimeoutMs is: How long to wait in milliseconds before timing out the request.
	FieldUpdateFeaturesRequesttimeoutMs = "timeoutMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/UpdateFeaturesRequest.json
const originalUpdateFeaturesRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "UpdateFeaturesRequest",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "timeoutMs", "type": "int32", "versions": "0+", "default": "60000",
      "about": "How long to wait in milliseconds before timing out the request." },
    { "name": "FeatureUpdates", "type": "[]FeatureUpdateKey", "versions": "0+",
      "about": "The list of updates to finalized features.", "fields": [
      {"name": "Feature", "type": "string", "versions": "0+", "mapKey": true,
        "about": "The name of the finalized feature to be updated."},
      {"name": "MaxVersionLevel", "type": "int16", "versions": "0+",
        "about": "The new maximum version level for the finalized feature. A value >= 1 is valid. A value < 1, is special, and can be used to request the deletion of the finalized feature."},
      {"name": "AllowDowngrade", "type": "bool", "versions": "0+",
        "about": "When set to true, the finalized feature version level is allowed to be downgraded/deleted. The downgrade request will fail if the new maximum version level is a value that's not lower than the existing maximum finalized version level."}
    ]}
  ]
}
`
