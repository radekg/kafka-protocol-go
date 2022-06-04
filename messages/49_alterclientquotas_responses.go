package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init49AlterClientQuotasResponse() []schema.Schema {

	return []schema.Schema{

		// Message: AlterClientQuotasResponse, API Key: 49, Version: 0
		schema.NewSchema("AlterClientQuotasResponsev0",
			&schema.Mfield{Name: FieldAlterClientQuotasResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldAlterClientQuotasResponseEntries, Ty: schema.NewSchema("EntriesV0",
				&schema.Mfield{Name: FieldAlterClientQuotasResponseEntriesErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldAlterClientQuotasResponseEntriesErrorMessage, Ty: schema.TypeStrNullable},
				&schema.Array{Name: FieldAlterClientQuotasResponseEntriesEntity, Ty: schema.NewSchema("EntityV0",
					&schema.Mfield{Name: FieldAlterClientQuotasResponseEntriesEntityEntityType, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldAlterClientQuotasResponseEntriesEntityEntityName, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: AlterClientQuotasResponse, API Key: 49, Version: 1
		schema.NewSchema("AlterClientQuotasResponsev1",
			&schema.Mfield{Name: FieldAlterClientQuotasResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldAlterClientQuotasResponseEntries, Ty: schema.NewSchema("EntriesV1",
				&schema.Mfield{Name: FieldAlterClientQuotasResponseEntriesErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldAlterClientQuotasResponseEntriesErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.ArrayCompact{Name: FieldAlterClientQuotasResponseEntriesEntity, Ty: schema.NewSchema("EntityV1",
					&schema.Mfield{Name: FieldAlterClientQuotasResponseEntriesEntityEntityType, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldAlterClientQuotasResponseEntriesEntityEntityName, Ty: schema.TypeStrCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldAlterClientQuotasResponseEntriesEntityTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldAlterClientQuotasResponseEntriesTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldAlterClientQuotasResponseTags},
		),
	}
}

const (
	// FieldAlterClientQuotasResponseEntriesEntityEntityType is: The entity type.
	FieldAlterClientQuotasResponseEntriesEntityEntityType = "EntityType"
	// FieldAlterClientQuotasResponseEntriesEntityTags is: The tagged fields.
	FieldAlterClientQuotasResponseEntriesEntityTags = "Tags"
	// FieldAlterClientQuotasResponseEntriesErrorCode is: The error code, or `0` if the quota alteration succeeded.
	FieldAlterClientQuotasResponseEntriesErrorCode = "ErrorCode"
	// FieldAlterClientQuotasResponseEntries is: The quota configuration entries to alter.
	FieldAlterClientQuotasResponseEntries = "Entries"
	// FieldAlterClientQuotasResponseEntriesErrorMessage is: The error message, or `null` if the quota alteration succeeded.
	FieldAlterClientQuotasResponseEntriesErrorMessage = "ErrorMessage"
	// FieldAlterClientQuotasResponseEntriesEntity is: The quota entity to alter.
	FieldAlterClientQuotasResponseEntriesEntity = "Entity"
	// FieldAlterClientQuotasResponseEntriesEntityEntityName is: The name of the entity, or null if the default.
	FieldAlterClientQuotasResponseEntriesEntityEntityName = "EntityName"
	// FieldAlterClientQuotasResponseEntriesTags is: The tagged fields.
	FieldAlterClientQuotasResponseEntriesTags = "Tags"
	// FieldAlterClientQuotasResponseTags is: The tagged fields.
	FieldAlterClientQuotasResponseTags = "Tags"
	// FieldAlterClientQuotasResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldAlterClientQuotasResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AlterClientQuotasResponse.json
const originalAlterClientQuotasResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 49,
  "type": "response",
  "name": "AlterClientQuotasResponse",
  // Version 1 enables flexible versions.
  "validVersions": "0-1",
  "flexibleVersions": "1+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Entries", "type": "[]EntryData", "versions": "0+",
      "about": "The quota configuration entries to alter.", "fields": [
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The error code, or '0' if the quota alteration succeeded." },
      { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
        "about": "The error message, or 'null' if the quota alteration succeeded." },
      { "name": "Entity", "type": "[]EntityData", "versions": "0+",
        "about": "The quota entity to alter.", "fields": [
        { "name": "EntityType", "type": "string", "versions": "0+",
          "about": "The entity type." },
        { "name": "EntityName", "type": "string", "versions": "0+", "nullableVersions": "0+",
          "about": "The name of the entity, or null if the default." }
      ]}
    ]}
  ]
}
`
