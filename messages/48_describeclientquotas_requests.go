package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init48DescribeClientQuotasRequest() []schema.Schema {

	return []schema.Schema{
		// Message: DescribeClientQuotasRequest, API Key: 48, Version: 0
		schema.NewSchema("DescribeClientQuotasRequest:v0",
			&schema.Array{Name: FieldDescribeClientQuotasRequestComponents, Ty: schema.NewSchema("Components:v0",
				&schema.Mfield{Name: FieldDescribeClientQuotasRequestComponentsEntityType, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeClientQuotasRequestComponentsMatchType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDescribeClientQuotasRequestComponentsMatch, Ty: schema.TypeStrNullable},
			)},
			&schema.Mfield{Name: FieldDescribeClientQuotasRequestStrict, Ty: schema.TypeBool},
		),

		// Message: DescribeClientQuotasRequest, API Key: 48, Version: 1
		schema.NewSchema("DescribeClientQuotasRequest:v1",
			&schema.ArrayCompact{Name: FieldDescribeClientQuotasRequestComponents, Ty: schema.NewSchema("Components:v1",
				&schema.Mfield{Name: FieldDescribeClientQuotasRequestComponentsEntityType, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeClientQuotasRequestComponentsMatchType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDescribeClientQuotasRequestComponentsMatch, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldDescribeClientQuotasRequestComponentsTags},
			)},
			&schema.Mfield{Name: FieldDescribeClientQuotasRequestStrict, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldDescribeClientQuotasRequestTags},
		),
	}

}

const (

	// FieldDescribeClientQuotasRequestComponents is: Filter components to apply to quota entities.
	FieldDescribeClientQuotasRequestComponents = "Components"

	// FieldDescribeClientQuotasRequestComponentsEntityType is: The entity type that the filter component applies to.
	FieldDescribeClientQuotasRequestComponentsEntityType = "EntityType"

	// FieldDescribeClientQuotasRequestComponentsMatch is: The string to match against, or null if unused for the match type.
	FieldDescribeClientQuotasRequestComponentsMatch = "Match"

	// FieldDescribeClientQuotasRequestComponentsMatchType is: How to match the entity {0 = exact name, 1 = default name, 2 = any specified name}.
	FieldDescribeClientQuotasRequestComponentsMatchType = "MatchType"

	// FieldDescribeClientQuotasRequestComponentsTags is: The tagged fields.
	FieldDescribeClientQuotasRequestComponentsTags = "Tags"

	// FieldDescribeClientQuotasRequestStrict is: Whether the match is strict, i.e. should exclude entities with unspecified entity types.
	FieldDescribeClientQuotasRequestStrict = "Strict"

	// FieldDescribeClientQuotasRequestTags is: The tagged fields.
	FieldDescribeClientQuotasRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeClientQuotasRequest.json
const originalDescribeClientQuotasRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "DescribeClientQuotasRequest",
  // Version 1 enables flexible versions.
  "validVersions": "0-1",
  "flexibleVersions": "1+",
  "fields": [
    { "name": "Components", "type": "[]ComponentData", "versions": "0+",
      "about": "Filter components to apply to quota entities.", "fields": [
      { "name": "EntityType", "type": "string", "versions": "0+",
        "about": "The entity type that the filter component applies to." },
      { "name": "MatchType", "type": "int8", "versions": "0+",
        "about": "How to match the entity {0 = exact name, 1 = default name, 2 = any specified name}." },
      { "name": "Match", "type": "string", "versions": "0+", "nullableVersions": "0+",
        "about": "The string to match against, or null if unused for the match type." }
    ]},
    { "name": "Strict", "type": "bool", "versions": "0+",
      "about": "Whether the match is strict, i.e. should exclude entities with unspecified entity types." }
  ]
}
`
