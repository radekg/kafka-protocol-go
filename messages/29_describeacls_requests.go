package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init29DescribeAclsRequest() []schema.Schema {

	return []schema.Schema{
		// Message: DescribeAclsRequest, API Key: 29, Version: 0
		schema.NewSchema("DescribeAclsRequest:v0",
			&schema.Mfield{Name: FieldDescribeAclsRequestResourceTypeFilter, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldDescribeAclsRequestResourceNameFilter, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldDescribeAclsRequestPrincipalFilter, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldDescribeAclsRequestHostFilter, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldDescribeAclsRequestOperation, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldDescribeAclsRequestPermissionType, Ty: schema.TypeInt8},
		),

		// Message: DescribeAclsRequest, API Key: 29, Version: 1
		schema.NewSchema("DescribeAclsRequest:v1",
			&schema.Mfield{Name: FieldDescribeAclsRequestResourceTypeFilter, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldDescribeAclsRequestResourceNameFilter, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldDescribeAclsRequestPatternTypeFilter, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldDescribeAclsRequestPrincipalFilter, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldDescribeAclsRequestHostFilter, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldDescribeAclsRequestOperation, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldDescribeAclsRequestPermissionType, Ty: schema.TypeInt8},
		),

		// Message: DescribeAclsRequest, API Key: 29, Version: 2
		schema.NewSchema("DescribeAclsRequest:v2",
			&schema.Mfield{Name: FieldDescribeAclsRequestResourceTypeFilter, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldDescribeAclsRequestResourceNameFilter, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldDescribeAclsRequestPatternTypeFilter, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldDescribeAclsRequestPrincipalFilter, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldDescribeAclsRequestHostFilter, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldDescribeAclsRequestOperation, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldDescribeAclsRequestPermissionType, Ty: schema.TypeInt8},
			&schema.SchemaTaggedFields{Name: FieldDescribeAclsRequestTags},
		),
	}

}

const (

	// FieldDescribeAclsRequestHostFilter is: The host to match, or null to match any host.
	FieldDescribeAclsRequestHostFilter = "HostFilter"

	// FieldDescribeAclsRequestOperation is: The operation to match.
	FieldDescribeAclsRequestOperation = "Operation"

	// FieldDescribeAclsRequestPatternTypeFilter is: The resource pattern to match.
	FieldDescribeAclsRequestPatternTypeFilter = "PatternTypeFilter"

	// FieldDescribeAclsRequestPermissionType is: The permission type to match.
	FieldDescribeAclsRequestPermissionType = "PermissionType"

	// FieldDescribeAclsRequestPrincipalFilter is: The principal to match, or null to match any principal.
	FieldDescribeAclsRequestPrincipalFilter = "PrincipalFilter"

	// FieldDescribeAclsRequestResourceNameFilter is: The resource name, or null to match any resource name.
	FieldDescribeAclsRequestResourceNameFilter = "ResourceNameFilter"

	// FieldDescribeAclsRequestResourceTypeFilter is: The resource type.
	FieldDescribeAclsRequestResourceTypeFilter = "ResourceTypeFilter"

	// FieldDescribeAclsRequestTags is: The tagged fields.
	FieldDescribeAclsRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeAclsRequest.json
const originalDescribeAclsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 29,
  "type": "request",
  "listeners": ["zkBroker", "broker", "controller"],
  "name": "DescribeAclsRequest",
  // Version 1 adds resource pattern type.
  // Version 2 enables flexible versions.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ResourceTypeFilter", "type": "int8", "versions": "0+",
      "about": "The resource type." },
    { "name": "ResourceNameFilter", "type": "string", "versions": "0+", "nullableVersions": "0+",
      "about": "The resource name, or null to match any resource name." },
    { "name": "PatternTypeFilter", "type": "int8", "versions": "1+", "default": "3", "ignorable": false,
      "about": "The resource pattern to match." },
    { "name": "PrincipalFilter", "type": "string", "versions": "0+", "nullableVersions": "0+",
      "about": "The principal to match, or null to match any principal." },
    { "name": "HostFilter", "type": "string", "versions": "0+", "nullableVersions": "0+",
      "about": "The host to match, or null to match any host." },
    { "name": "Operation", "type": "int8", "versions": "0+",
      "about": "The operation to match." },
    { "name": "PermissionType", "type": "int8", "versions": "0+",
      "about": "The permission type to match." }
  ]
}
`
