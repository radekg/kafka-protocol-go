package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init49AlterClientQuotasRequest() []schema.Schema {

	return []schema.Schema{
		// Message: AlterClientQuotasRequest, API Key: 49, Version: 0
		schema.NewSchema("AlterClientQuotasRequest:v0",
			&schema.Array{Name: FieldAlterClientQuotasRequestEntries, Ty: schema.NewSchema("Entries:v0",
				&schema.Array{Name: FieldAlterClientQuotasRequestEntriesEntity, Ty: schema.NewSchema("Entity:v0",
					&schema.Mfield{Name: FieldAlterClientQuotasRequestEntriesEntityEntityType, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldAlterClientQuotasRequestEntriesEntityEntityName, Ty: schema.TypeStrNullable},
				)},
				&schema.Array{Name: FieldAlterClientQuotasRequestEntriesOps, Ty: schema.NewSchema("Ops:v0",
					&schema.Mfield{Name: FieldAlterClientQuotasRequestEntriesOpsKey, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldAlterClientQuotasRequestEntriesOpsValue, Ty: schema.TypeFloat64},
					&schema.Mfield{Name: FieldAlterClientQuotasRequestEntriesOpsRemove, Ty: schema.TypeBool},
				)},
			)},
			&schema.Mfield{Name: FieldAlterClientQuotasRequestValidateOnly, Ty: schema.TypeBool},
		),

		// Message: AlterClientQuotasRequest, API Key: 49, Version: 1
		schema.NewSchema("AlterClientQuotasRequest:v1",
			&schema.ArrayCompact{Name: FieldAlterClientQuotasRequestEntries, Ty: schema.NewSchema("Entries:v1",
				&schema.ArrayCompact{Name: FieldAlterClientQuotasRequestEntriesEntity, Ty: schema.NewSchema("Entity:v1",
					&schema.Mfield{Name: FieldAlterClientQuotasRequestEntriesEntityEntityType, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldAlterClientQuotasRequestEntriesEntityEntityName, Ty: schema.TypeStrCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldAlterClientQuotasRequestEntriesEntityTags},
				)},
				&schema.ArrayCompact{Name: FieldAlterClientQuotasRequestEntriesOps, Ty: schema.NewSchema("Ops:v1",
					&schema.Mfield{Name: FieldAlterClientQuotasRequestEntriesOpsKey, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldAlterClientQuotasRequestEntriesOpsValue, Ty: schema.TypeFloat64},
					&schema.Mfield{Name: FieldAlterClientQuotasRequestEntriesOpsRemove, Ty: schema.TypeBool},
					&schema.SchemaTaggedFields{Name: FieldAlterClientQuotasRequestEntriesOpsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldAlterClientQuotasRequestEntriesTags},
			)},
			&schema.Mfield{Name: FieldAlterClientQuotasRequestValidateOnly, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldAlterClientQuotasRequestTags},
		),
	}

}

const (

	// FieldAlterClientQuotasRequestEntries is: The quota configuration entries to alter.
	FieldAlterClientQuotasRequestEntries = "Entries"

	// FieldAlterClientQuotasRequestEntriesEntity is: The quota entity to alter.
	FieldAlterClientQuotasRequestEntriesEntity = "Entity"

	// FieldAlterClientQuotasRequestEntriesEntityEntityName is: The name of the entity, or null if the default.
	FieldAlterClientQuotasRequestEntriesEntityEntityName = "EntityName"

	// FieldAlterClientQuotasRequestEntriesEntityEntityType is: The entity type.
	FieldAlterClientQuotasRequestEntriesEntityEntityType = "EntityType"

	// FieldAlterClientQuotasRequestEntriesEntityTags is: The tagged fields.
	FieldAlterClientQuotasRequestEntriesEntityTags = "Tags"

	// FieldAlterClientQuotasRequestEntriesOps is: An individual quota configuration entry to alter.
	FieldAlterClientQuotasRequestEntriesOps = "Ops"

	// FieldAlterClientQuotasRequestEntriesOpsKey is: The quota configuration key.
	FieldAlterClientQuotasRequestEntriesOpsKey = "Key"

	// FieldAlterClientQuotasRequestEntriesOpsRemove is: Whether the quota configuration value should be removed, otherwise set.
	FieldAlterClientQuotasRequestEntriesOpsRemove = "Remove"

	// FieldAlterClientQuotasRequestEntriesOpsTags is: The tagged fields.
	FieldAlterClientQuotasRequestEntriesOpsTags = "Tags"

	// FieldAlterClientQuotasRequestEntriesOpsValue is: The value to set, otherwise ignored if the value is to be removed.
	FieldAlterClientQuotasRequestEntriesOpsValue = "Value"

	// FieldAlterClientQuotasRequestEntriesTags is: The tagged fields.
	FieldAlterClientQuotasRequestEntriesTags = "Tags"

	// FieldAlterClientQuotasRequestTags is: The tagged fields.
	FieldAlterClientQuotasRequestTags = "Tags"

	// FieldAlterClientQuotasRequestValidateOnly is: Whether the alteration should be validated, but not performed.
	FieldAlterClientQuotasRequestValidateOnly = "ValidateOnly"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AlterClientQuotasRequest.json
const originalAlterClientQuotasRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "request",
  "listeners": ["zkBroker", "broker", "controller"],
  "name": "AlterClientQuotasRequest",
  "validVersions": "0-1",
  // Version 1 enables flexible versions.
  "flexibleVersions": "1+",
  "fields": [
    { "name": "Entries", "type": "[]EntryData", "versions": "0+",
      "about": "The quota configuration entries to alter.", "fields": [
      { "name": "Entity", "type": "[]EntityData", "versions": "0+",
        "about": "The quota entity to alter.", "fields": [
        { "name": "EntityType", "type": "string", "versions": "0+",
          "about": "The entity type." },
        { "name": "EntityName", "type": "string", "versions": "0+", "nullableVersions": "0+",
          "about": "The name of the entity, or null if the default." }
      ]},
      { "name": "Ops", "type": "[]OpData", "versions": "0+",
        "about": "An individual quota configuration entry to alter.", "fields": [
        { "name": "Key", "type": "string", "versions": "0+",
          "about": "The quota configuration key." },
        { "name": "Value", "type": "float64", "versions": "0+",
          "about": "The value to set, otherwise ignored if the value is to be removed." },
        { "name": "Remove", "type": "bool", "versions": "0+",
          "about": "Whether the quota configuration value should be removed, otherwise set." }
      ]}
    ]},
    { "name": "ValidateOnly", "type": "bool", "versions": "0+",
      "about": "Whether the alteration should be validated, but not performed." }
  ]
}
`
