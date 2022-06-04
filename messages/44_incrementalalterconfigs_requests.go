package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init44IncrementalAlterConfigsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: IncrementalAlterConfigsRequest, API Key: 44, Version: 0
		schema.NewSchema("IncrementalAlterConfigsRequestv0",
			&schema.Array{Name: FieldIncrementalAlterConfigsRequestResources, Ty: schema.NewSchema("ResourcesV0",
				&schema.Mfield{Name: FieldIncrementalAlterConfigsRequestResourcesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldIncrementalAlterConfigsRequestResourcesResourceName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldIncrementalAlterConfigsRequestResourcesConfigs, Ty: schema.NewSchema("ConfigsV0",
					&schema.Mfield{Name: FieldIncrementalAlterConfigsRequestResourcesConfigsName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldIncrementalAlterConfigsRequestResourcesConfigsConfigOperation, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldIncrementalAlterConfigsRequestResourcesConfigsValue, Ty: schema.TypeStrNullable},
				)},
			)},
			&schema.Mfield{Name: FieldIncrementalAlterConfigsRequestValidateOnly, Ty: schema.TypeBool},
		),

		// Message: IncrementalAlterConfigsRequest, API Key: 44, Version: 1
		schema.NewSchema("IncrementalAlterConfigsRequestv1",
			&schema.ArrayCompact{Name: FieldIncrementalAlterConfigsRequestResources, Ty: schema.NewSchema("ResourcesV1",
				&schema.Mfield{Name: FieldIncrementalAlterConfigsRequestResourcesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldIncrementalAlterConfigsRequestResourcesResourceName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldIncrementalAlterConfigsRequestResourcesConfigs, Ty: schema.NewSchema("ConfigsV1",
					&schema.Mfield{Name: FieldIncrementalAlterConfigsRequestResourcesConfigsName, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldIncrementalAlterConfigsRequestResourcesConfigsConfigOperation, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldIncrementalAlterConfigsRequestResourcesConfigsValue, Ty: schema.TypeStrCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldIncrementalAlterConfigsRequestResourcesConfigsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldIncrementalAlterConfigsRequestResourcesTags},
			)},
			&schema.Mfield{Name: FieldIncrementalAlterConfigsRequestValidateOnly, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldIncrementalAlterConfigsRequestTags},
		),
	}
}

const (
	// FieldIncrementalAlterConfigsRequestResourcesResourceType is a field name that can be used to resolve the correct struct field.
	FieldIncrementalAlterConfigsRequestResourcesResourceType = "ResourceType"
	// FieldIncrementalAlterConfigsRequestResourcesResourceName is a field name that can be used to resolve the correct struct field.
	FieldIncrementalAlterConfigsRequestResourcesResourceName = "ResourceName"
	// FieldIncrementalAlterConfigsRequestResourcesConfigsName is a field name that can be used to resolve the correct struct field.
	FieldIncrementalAlterConfigsRequestResourcesConfigsName = "Name"
	// FieldIncrementalAlterConfigsRequestResourcesConfigsTags is a field name that can be used to resolve the correct struct field.
	FieldIncrementalAlterConfigsRequestResourcesConfigsTags = "Tags"
	// FieldIncrementalAlterConfigsRequestResources is a field name that can be used to resolve the correct struct field.
	FieldIncrementalAlterConfigsRequestResources = "Resources"
	// FieldIncrementalAlterConfigsRequestResourcesConfigs is a field name that can be used to resolve the correct struct field.
	FieldIncrementalAlterConfigsRequestResourcesConfigs = "Configs"
	// FieldIncrementalAlterConfigsRequestResourcesConfigsConfigOperation is a field name that can be used to resolve the correct struct field.
	FieldIncrementalAlterConfigsRequestResourcesConfigsConfigOperation = "ConfigOperation"
	// FieldIncrementalAlterConfigsRequestResourcesConfigsValue is a field name that can be used to resolve the correct struct field.
	FieldIncrementalAlterConfigsRequestResourcesConfigsValue = "Value"
	// FieldIncrementalAlterConfigsRequestValidateOnly is a field name that can be used to resolve the correct struct field.
	FieldIncrementalAlterConfigsRequestValidateOnly = "ValidateOnly"
	// FieldIncrementalAlterConfigsRequestResourcesTags is a field name that can be used to resolve the correct struct field.
	FieldIncrementalAlterConfigsRequestResourcesTags = "Tags"
	// FieldIncrementalAlterConfigsRequestTags is a field name that can be used to resolve the correct struct field.
	FieldIncrementalAlterConfigsRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/IncrementalAlterConfigsRequest.json
const originalIncrementalAlterConfigsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 44,
  "type": "request",
  "listeners": ["zkBroker", "broker", "controller"],
  "name": "IncrementalAlterConfigsRequest",
  // Version 1 is the first flexible version.
  "validVersions": "0-1",
  "flexibleVersions": "1+",
  "fields": [
    { "name": "Resources", "type": "[]AlterConfigsResource", "versions": "0+",
      "about": "The incremental updates for each resource.", "fields": [
      { "name": "ResourceType", "type": "int8", "versions": "0+", "mapKey": true,
        "about": "The resource type." },
      { "name": "ResourceName", "type": "string", "versions": "0+", "mapKey": true,
        "about": "The resource name." },
      { "name": "Configs", "type": "[]AlterableConfig", "versions": "0+",
        "about": "The configurations.",  "fields": [
        { "name": "Name", "type": "string", "versions": "0+", "mapKey": true,
          "about": "The configuration key name." },
        { "name": "ConfigOperation", "type": "int8", "versions": "0+", "mapKey": true,
          "about": "The type (Set, Delete, Append, Subtract) of operation." },
        { "name": "Value", "type": "string", "versions": "0+", "nullableVersions": "0+",
          "about": "The value to set for the configuration key."}
      ]}
    ]},
    { "name": "ValidateOnly", "type": "bool", "versions": "0+",
      "about": "True if we should validate the request, but not change the configurations."}
  ]
}
`
