package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init33AlterConfigsRequest() []schema.Schema {

	return []schema.Schema{
		// Message: AlterConfigsRequest, API Key: 33, Version: 0
		schema.NewSchema("AlterConfigsRequest:v0",
			&schema.Array{Name: FieldAlterConfigsRequestResources, Ty: schema.NewSchema("Resources:v0",
				&schema.Mfield{Name: FieldAlterConfigsRequestResourcesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldAlterConfigsRequestResourcesResourceName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldAlterConfigsRequestResourcesConfigs, Ty: schema.NewSchema("Configs:v0",
					&schema.Mfield{Name: FieldAlterConfigsRequestResourcesConfigsName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldAlterConfigsRequestResourcesConfigsValue, Ty: schema.TypeStrNullable},
				)},
			)},
			&schema.Mfield{Name: FieldAlterConfigsRequestValidateOnly, Ty: schema.TypeBool},
		),

		// Message: AlterConfigsRequest, API Key: 33, Version: 1
		schema.NewSchema("AlterConfigsRequest:v1",
			&schema.Array{Name: FieldAlterConfigsRequestResources, Ty: schema.NewSchema("Resources:v1",
				&schema.Mfield{Name: FieldAlterConfigsRequestResourcesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldAlterConfigsRequestResourcesResourceName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldAlterConfigsRequestResourcesConfigs, Ty: schema.NewSchema("Configs:v1",
					&schema.Mfield{Name: FieldAlterConfigsRequestResourcesConfigsName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldAlterConfigsRequestResourcesConfigsValue, Ty: schema.TypeStrNullable},
				)},
			)},
			&schema.Mfield{Name: FieldAlterConfigsRequestValidateOnly, Ty: schema.TypeBool},
		),

		// Message: AlterConfigsRequest, API Key: 33, Version: 2
		schema.NewSchema("AlterConfigsRequest:v2",
			&schema.ArrayCompact{Name: FieldAlterConfigsRequestResources, Ty: schema.NewSchema("Resources:v2",
				&schema.Mfield{Name: FieldAlterConfigsRequestResourcesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldAlterConfigsRequestResourcesResourceName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldAlterConfigsRequestResourcesConfigs, Ty: schema.NewSchema("Configs:v2",
					&schema.Mfield{Name: FieldAlterConfigsRequestResourcesConfigsName, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldAlterConfigsRequestResourcesConfigsValue, Ty: schema.TypeStrCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldAlterConfigsRequestResourcesConfigsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldAlterConfigsRequestResourcesTags},
			)},
			&schema.Mfield{Name: FieldAlterConfigsRequestValidateOnly, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldAlterConfigsRequestTags},
		),
	}

}

const (

	// FieldAlterConfigsRequestResources is: The updates for each resource.
	FieldAlterConfigsRequestResources = "Resources"

	// FieldAlterConfigsRequestResourcesConfigs is: The configurations.
	FieldAlterConfigsRequestResourcesConfigs = "Configs"

	// FieldAlterConfigsRequestResourcesConfigsName is: The configuration key name.
	FieldAlterConfigsRequestResourcesConfigsName = "Name"

	// FieldAlterConfigsRequestResourcesConfigsTags is: The tagged fields.
	FieldAlterConfigsRequestResourcesConfigsTags = "Tags"

	// FieldAlterConfigsRequestResourcesConfigsValue is: The value to set for the configuration key.
	FieldAlterConfigsRequestResourcesConfigsValue = "Value"

	// FieldAlterConfigsRequestResourcesResourceName is: The resource name.
	FieldAlterConfigsRequestResourcesResourceName = "ResourceName"

	// FieldAlterConfigsRequestResourcesResourceType is: The resource type.
	FieldAlterConfigsRequestResourcesResourceType = "ResourceType"

	// FieldAlterConfigsRequestResourcesTags is: The tagged fields.
	FieldAlterConfigsRequestResourcesTags = "Tags"

	// FieldAlterConfigsRequestTags is: The tagged fields.
	FieldAlterConfigsRequestTags = "Tags"

	// FieldAlterConfigsRequestValidateOnly is: True if we should validate the request, but not change the configurations.
	FieldAlterConfigsRequestValidateOnly = "ValidateOnly"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AlterConfigsRequest.json
const originalAlterConfigsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 33,
  "type": "request",
  "listeners": ["zkBroker", "broker", "controller"],
  "name": "AlterConfigsRequest",
  // Version 1 is the same as version 0.
  // Version 2 enables flexible versions.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "Resources", "type": "[]AlterConfigsResource", "versions": "0+", 
      "about": "The updates for each resource.", "fields": [
      { "name": "ResourceType", "type": "int8", "versions": "0+", "mapKey": true,
        "about": "The resource type." },
      { "name": "ResourceName", "type": "string", "versions": "0+", "mapKey": true,
        "about": "The resource name." },
      { "name": "Configs", "type": "[]AlterableConfig", "versions": "0+",
        "about": "The configurations.",  "fields": [
        { "name": "Name", "type": "string", "versions": "0+", "mapKey": true,
          "about": "The configuration key name." },
        { "name": "Value", "type": "string", "versions": "0+", "nullableVersions": "0+",
          "about": "The value to set for the configuration key."}
      ]}
    ]},
    { "name": "ValidateOnly", "type": "bool", "versions": "0+",
      "about": "True if we should validate the request, but not change the configurations."}
  ]
}
`
