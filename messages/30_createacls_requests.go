package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init30CreateAclsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: CreateAclsRequest, API Key: 30, Version: 0
		schema.NewSchema("CreateAclsRequestv0",
			&schema.Array{Name: FieldCreateAclsRequestCreations, Ty: schema.NewSchema("CreationsV0",
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsResourceName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsPrincipal, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsHost, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsOperation, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsPermissionType, Ty: schema.TypeInt8},
			)},
		),

		// Message: CreateAclsRequest, API Key: 30, Version: 1
		schema.NewSchema("CreateAclsRequestv1",
			&schema.Array{Name: FieldCreateAclsRequestCreations, Ty: schema.NewSchema("CreationsV1",
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsResourceName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsResourcePatternType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsPrincipal, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsHost, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsOperation, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsPermissionType, Ty: schema.TypeInt8},
			)},
		),

		// Message: CreateAclsRequest, API Key: 30, Version: 2
		schema.NewSchema("CreateAclsRequestv2",
			&schema.ArrayCompact{Name: FieldCreateAclsRequestCreations, Ty: schema.NewSchema("CreationsV2",
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsResourceName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsResourcePatternType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsPrincipal, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsHost, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsOperation, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldCreateAclsRequestCreationsPermissionType, Ty: schema.TypeInt8},
				&schema.SchemaTaggedFields{Name: FieldCreateAclsRequestCreationsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldCreateAclsRequestTags},
		),
	}
}

const (
	// FieldCreateAclsRequestCreationsOperation is a field name that can be used to resolve the correct struct field.
	FieldCreateAclsRequestCreationsOperation = "Operation"
	// FieldCreateAclsRequestCreationsResourcePatternType is a field name that can be used to resolve the correct struct field.
	FieldCreateAclsRequestCreationsResourcePatternType = "ResourcePatternType"
	// FieldCreateAclsRequestTags is a field name that can be used to resolve the correct struct field.
	FieldCreateAclsRequestTags = "Tags"
	// FieldCreateAclsRequestCreations is a field name that can be used to resolve the correct struct field.
	FieldCreateAclsRequestCreations = "Creations"
	// FieldCreateAclsRequestCreationsResourceType is a field name that can be used to resolve the correct struct field.
	FieldCreateAclsRequestCreationsResourceType = "ResourceType"
	// FieldCreateAclsRequestCreationsResourceName is a field name that can be used to resolve the correct struct field.
	FieldCreateAclsRequestCreationsResourceName = "ResourceName"
	// FieldCreateAclsRequestCreationsPrincipal is a field name that can be used to resolve the correct struct field.
	FieldCreateAclsRequestCreationsPrincipal = "Principal"
	// FieldCreateAclsRequestCreationsHost is a field name that can be used to resolve the correct struct field.
	FieldCreateAclsRequestCreationsHost = "Host"
	// FieldCreateAclsRequestCreationsPermissionType is a field name that can be used to resolve the correct struct field.
	FieldCreateAclsRequestCreationsPermissionType = "PermissionType"
	// FieldCreateAclsRequestCreationsTags is a field name that can be used to resolve the correct struct field.
	FieldCreateAclsRequestCreationsTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/CreateAclsRequest.json
const originalCreateAclsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 30,
  "type": "request",
  "listeners": ["zkBroker", "broker", "controller"],
  "name": "CreateAclsRequest",
  // Version 1 adds resource pattern type.
  // Version 2 enables flexible versions.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "Creations", "type": "[]AclCreation", "versions": "0+",
      "about": "The ACLs that we want to create.", "fields": [
      { "name": "ResourceType", "type": "int8", "versions": "0+",
        "about": "The type of the resource." },
      { "name": "ResourceName", "type": "string", "versions": "0+",
        "about": "The resource name for the ACL." },
      { "name": "ResourcePatternType", "type": "int8", "versions": "1+", "default": "3",
        "about": "The pattern type for the ACL." },
      { "name": "Principal", "type": "string", "versions": "0+",
        "about": "The principal for the ACL." },
      { "name": "Host", "type": "string", "versions": "0+",
        "about": "The host for the ACL." },
      { "name": "Operation", "type": "int8", "versions": "0+",
        "about": "The operation type for the ACL (read, write, etc.)." },
      { "name": "PermissionType", "type": "int8", "versions": "0+",
        "about": "The permission type for the ACL (allow, deny, etc.)." }
    ]}
  ]
}
`
