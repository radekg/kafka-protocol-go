package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init32DescribeConfigsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: DescribeConfigsRequest, API Key: 32, Version: 0
		schema.NewSchema("DescribeConfigsRequestv0",
			&schema.Array{Name: FieldDescribeConfigsRequestResources, Ty: schema.NewSchema("ResourcesV0",
				&schema.Mfield{Name: FieldDescribeConfigsRequestResourcesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDescribeConfigsRequestResourcesResourceName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeConfigsRequestResourcesConfigurationKeys, Ty: schema.TypeStrArray},
			)},
		),

		// Message: DescribeConfigsRequest, API Key: 32, Version: 1
		schema.NewSchema("DescribeConfigsRequestv1",
			&schema.Array{Name: FieldDescribeConfigsRequestResources, Ty: schema.NewSchema("ResourcesV1",
				&schema.Mfield{Name: FieldDescribeConfigsRequestResourcesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDescribeConfigsRequestResourcesResourceName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeConfigsRequestResourcesConfigurationKeys, Ty: schema.TypeStrArray},
			)},
			&schema.Mfield{Name: FieldDescribeConfigsRequestIncludeSynonyms, Ty: schema.TypeBool},
		),

		// Message: DescribeConfigsRequest, API Key: 32, Version: 2
		schema.NewSchema("DescribeConfigsRequestv2",
			&schema.Array{Name: FieldDescribeConfigsRequestResources, Ty: schema.NewSchema("ResourcesV2",
				&schema.Mfield{Name: FieldDescribeConfigsRequestResourcesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDescribeConfigsRequestResourcesResourceName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeConfigsRequestResourcesConfigurationKeys, Ty: schema.TypeStrArray},
			)},
			&schema.Mfield{Name: FieldDescribeConfigsRequestIncludeSynonyms, Ty: schema.TypeBool},
		),

		// Message: DescribeConfigsRequest, API Key: 32, Version: 3
		schema.NewSchema("DescribeConfigsRequestv3",
			&schema.Array{Name: FieldDescribeConfigsRequestResources, Ty: schema.NewSchema("ResourcesV3",
				&schema.Mfield{Name: FieldDescribeConfigsRequestResourcesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDescribeConfigsRequestResourcesResourceName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeConfigsRequestResourcesConfigurationKeys, Ty: schema.TypeStrArray},
			)},
			&schema.Mfield{Name: FieldDescribeConfigsRequestIncludeSynonyms, Ty: schema.TypeBool},
			&schema.Mfield{Name: FieldDescribeConfigsRequestIncludeDocumentation, Ty: schema.TypeBool},
		),

		// Message: DescribeConfigsRequest, API Key: 32, Version: 4
		schema.NewSchema("DescribeConfigsRequestv4",
			&schema.ArrayCompact{Name: FieldDescribeConfigsRequestResources, Ty: schema.NewSchema("ResourcesV4",
				&schema.Mfield{Name: FieldDescribeConfigsRequestResourcesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDescribeConfigsRequestResourcesResourceName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeConfigsRequestResourcesConfigurationKeys, Ty: schema.TypeStrCompactArray},
				&schema.SchemaTaggedFields{Name: FieldDescribeConfigsRequestResourcesTags},
			)},
			&schema.Mfield{Name: FieldDescribeConfigsRequestIncludeSynonyms, Ty: schema.TypeBool},
			&schema.Mfield{Name: FieldDescribeConfigsRequestIncludeDocumentation, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldDescribeConfigsRequestTags},
		),
	}
}

const (
	// FieldDescribeConfigsRequestIncludeDocumentation is: True if we should include configuration documentation.
	FieldDescribeConfigsRequestIncludeDocumentation = "IncludeDocumentation"
	// FieldDescribeConfigsRequestIncludeSynonyms is: True if we should include all synonyms.
	FieldDescribeConfigsRequestIncludeSynonyms = "IncludeSynonyms"
	// FieldDescribeConfigsRequestResources is: The resources whose configurations we want to describe.
	FieldDescribeConfigsRequestResources = "Resources"
	// FieldDescribeConfigsRequestResourcesConfigurationKeys is: The configuration keys to list, or null to list all configuration keys.
	FieldDescribeConfigsRequestResourcesConfigurationKeys = "ConfigurationKeys"
	// FieldDescribeConfigsRequestResourcesResourceName is: The resource name.
	FieldDescribeConfigsRequestResourcesResourceName = "ResourceName"
	// FieldDescribeConfigsRequestResourcesResourceType is: The resource type.
	FieldDescribeConfigsRequestResourcesResourceType = "ResourceType"
	// FieldDescribeConfigsRequestResourcesTags is: The tagged fields.
	FieldDescribeConfigsRequestResourcesTags = "Tags"
	// FieldDescribeConfigsRequestTags is: The tagged fields.
	FieldDescribeConfigsRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeConfigsRequest.json
const originalDescribeConfigsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 32,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "DescribeConfigsRequest",
  // Version 1 adds IncludeSynonyms.
  // Version 2 is the same as version 1.
  // Version 4 enables flexible versions.
  "validVersions": "0-4",
  "flexibleVersions": "4+",
  "fields": [
    { "name": "Resources", "type": "[]DescribeConfigsResource", "versions": "0+",
      "about": "The resources whose configurations we want to describe.", "fields": [
      { "name": "ResourceType", "type": "int8", "versions": "0+",
        "about": "The resource type." },
      { "name": "ResourceName", "type": "string", "versions": "0+",
        "about": "The resource name." },
      { "name": "ConfigurationKeys", "type": "[]string", "versions": "0+", "nullableVersions": "0+",
        "about": "The configuration keys to list, or null to list all configuration keys." }
    ]},
    { "name": "IncludeSynonyms", "type": "bool", "versions": "1+", "default": "false", "ignorable": false,
      "about": "True if we should include all synonyms." },
    { "name": "IncludeDocumentation", "type": "bool", "versions": "3+", "default": "false", "ignorable": false,
      "about": "True if we should include configuration documentation." }
  ]
}
`
