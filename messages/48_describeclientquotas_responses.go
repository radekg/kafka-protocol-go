package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init48DescribeClientQuotasResponse() []schema.Schema {

	return []schema.Schema{
		// Message: DescribeClientQuotasResponse, API Key: 48, Version: 0
		schema.NewSchema("DescribeClientQuotasResponse:v0",
			&schema.Mfield{Name: FieldDescribeClientQuotasResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldDescribeClientQuotasResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldDescribeClientQuotasResponseErrorMessage, Ty: schema.TypeStrNullable},
			&schema.Array{Name: FieldDescribeClientQuotasResponseEntries, Ty: schema.NewSchema("[]EntryData:v0",
				&schema.Array{Name: FieldDescribeClientQuotasResponseEntriesEntity, Ty: schema.NewSchema("[]EntityData:v0",
					&schema.Mfield{Name: FieldDescribeClientQuotasResponseEntriesEntityEntityType, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeClientQuotasResponseEntriesEntityEntityName, Ty: schema.TypeStrNullable},
				)},
				&schema.Array{Name: FieldDescribeClientQuotasResponseEntriesValues, Ty: schema.NewSchema("[]ValueData:v0",
					&schema.Mfield{Name: FieldDescribeClientQuotasResponseEntriesValuesKey, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeClientQuotasResponseEntriesValuesValue, Ty: schema.TypeFloat64},
				)},
			)},
		),

		// Message: DescribeClientQuotasResponse, API Key: 48, Version: 1
		schema.NewSchema("DescribeClientQuotasResponse:v1",
			&schema.Mfield{Name: FieldDescribeClientQuotasResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldDescribeClientQuotasResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldDescribeClientQuotasResponseErrorMessage, Ty: schema.TypeStrCompactNullable},
			&schema.ArrayCompact{Name: FieldDescribeClientQuotasResponseEntries, Ty: schema.NewSchema("[]EntryData:v1",
				&schema.ArrayCompact{Name: FieldDescribeClientQuotasResponseEntriesEntity, Ty: schema.NewSchema("[]EntityData:v1",
					&schema.Mfield{Name: FieldDescribeClientQuotasResponseEntriesEntityEntityType, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldDescribeClientQuotasResponseEntriesEntityEntityName, Ty: schema.TypeStrCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldDescribeClientQuotasResponseEntriesEntityTags},
				)},
				&schema.ArrayCompact{Name: FieldDescribeClientQuotasResponseEntriesValues, Ty: schema.NewSchema("[]ValueData:v1",
					&schema.Mfield{Name: FieldDescribeClientQuotasResponseEntriesValuesKey, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldDescribeClientQuotasResponseEntriesValuesValue, Ty: schema.TypeFloat64},
					&schema.SchemaTaggedFields{Name: FieldDescribeClientQuotasResponseEntriesValuesTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldDescribeClientQuotasResponseEntriesTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDescribeClientQuotasResponseTags},
		),
	}

}

const (

	// FieldDescribeClientQuotasResponseEntries is: A result entry.
	FieldDescribeClientQuotasResponseEntries = "Entries"

	// FieldDescribeClientQuotasResponseEntriesEntity is: The quota entity description.
	FieldDescribeClientQuotasResponseEntriesEntity = "Entity"

	// FieldDescribeClientQuotasResponseEntriesEntityEntityName is: The entity name, or null if the default.
	FieldDescribeClientQuotasResponseEntriesEntityEntityName = "EntityName"

	// FieldDescribeClientQuotasResponseEntriesEntityEntityType is: The entity type.
	FieldDescribeClientQuotasResponseEntriesEntityEntityType = "EntityType"

	// FieldDescribeClientQuotasResponseEntriesEntityTags is: The tagged fields.
	FieldDescribeClientQuotasResponseEntriesEntityTags = "Tags"

	// FieldDescribeClientQuotasResponseEntriesTags is: The tagged fields.
	FieldDescribeClientQuotasResponseEntriesTags = "Tags"

	// FieldDescribeClientQuotasResponseEntriesValues is: The quota values for the entity.
	FieldDescribeClientQuotasResponseEntriesValues = "Values"

	// FieldDescribeClientQuotasResponseEntriesValuesKey is: The quota configuration key.
	FieldDescribeClientQuotasResponseEntriesValuesKey = "Key"

	// FieldDescribeClientQuotasResponseEntriesValuesTags is: The tagged fields.
	FieldDescribeClientQuotasResponseEntriesValuesTags = "Tags"

	// FieldDescribeClientQuotasResponseEntriesValuesValue is: The quota configuration value.
	FieldDescribeClientQuotasResponseEntriesValuesValue = "Value"

	// FieldDescribeClientQuotasResponseErrorCode is: The error code, or `0` if the quota description succeeded.
	FieldDescribeClientQuotasResponseErrorCode = "ErrorCode"

	// FieldDescribeClientQuotasResponseErrorMessage is: The error message, or `null` if the quota description succeeded.
	FieldDescribeClientQuotasResponseErrorMessage = "ErrorMessage"

	// FieldDescribeClientQuotasResponseTags is: The tagged fields.
	FieldDescribeClientQuotasResponseTags = "Tags"

	// FieldDescribeClientQuotasResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldDescribeClientQuotasResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeClientQuotasResponse.json
const originalDescribeClientQuotasResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 48,
  "type": "response",
  "name": "DescribeClientQuotasResponse",
  // Version 1 enables flexible versions.
  "validVersions": "0-1",
  "flexibleVersions": "1+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The error code, or '0' if the quota description succeeded." },
    { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
      "about": "The error message, or 'null' if the quota description succeeded." },
    { "name": "Entries", "type": "[]EntryData", "versions": "0+", "nullableVersions": "0+",
      "about": "A result entry.", "fields": [
      { "name": "Entity", "type": "[]EntityData", "versions": "0+",
        "about": "The quota entity description.", "fields": [
        { "name": "EntityType", "type": "string", "versions": "0+",
          "about": "The entity type." },
        { "name": "EntityName", "type": "string", "versions": "0+", "nullableVersions": "0+",
          "about": "The entity name, or null if the default." }
      ]},
      { "name": "Values", "type": "[]ValueData", "versions": "0+",
	"about": "The quota values for the entity.", "fields": [
        { "name": "Key", "type": "string", "versions": "0+",
          "about": "The quota configuration key." },
        { "name": "Value", "type": "float64", "versions": "0+",
          "about": "The quota configuration value." }
      ]}
    ]}
  ]
}
`
