package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init29DescribeAclsResponse() []schema.Schema {

	return []schema.Schema{

		// Message: DescribeAclsResponse, API Key: 29, Version: 0
		schema.NewSchema("DescribeAclsResponsev0",
			&schema.Mfield{Name: FieldDescribeAclsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldDescribeAclsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldDescribeAclsResponseErrorMessage, Ty: schema.TypeStrNullable},
			&schema.Array{Name: FieldDescribeAclsResponseResources, Ty: schema.NewSchema("ResourcesV0",
				&schema.Mfield{Name: FieldDescribeAclsResponseResourcesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDescribeAclsResponseResourcesResourceName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDescribeAclsResponseResourcesAcls, Ty: schema.NewSchema("AclsV0",
					&schema.Mfield{Name: FieldDescribeAclsResponseResourcesAclsPrincipal, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeAclsResponseResourcesAclsHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeAclsResponseResourcesAclsOperation, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDescribeAclsResponseResourcesAclsPermissionType, Ty: schema.TypeInt8},
				)},
			)},
		),

		// Message: DescribeAclsResponse, API Key: 29, Version: 1
		schema.NewSchema("DescribeAclsResponsev1",
			&schema.Mfield{Name: FieldDescribeAclsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldDescribeAclsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldDescribeAclsResponseErrorMessage, Ty: schema.TypeStrNullable},
			&schema.Array{Name: FieldDescribeAclsResponseResources, Ty: schema.NewSchema("ResourcesV1",
				&schema.Mfield{Name: FieldDescribeAclsResponseResourcesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDescribeAclsResponseResourcesResourceName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeAclsResponseResourcesPatternType, Ty: schema.TypeInt8},
				&schema.Array{Name: FieldDescribeAclsResponseResourcesAcls, Ty: schema.NewSchema("AclsV1",
					&schema.Mfield{Name: FieldDescribeAclsResponseResourcesAclsPrincipal, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeAclsResponseResourcesAclsHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeAclsResponseResourcesAclsOperation, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDescribeAclsResponseResourcesAclsPermissionType, Ty: schema.TypeInt8},
				)},
			)},
		),

		// Message: DescribeAclsResponse, API Key: 29, Version: 2
		schema.NewSchema("DescribeAclsResponsev2",
			&schema.Mfield{Name: FieldDescribeAclsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldDescribeAclsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldDescribeAclsResponseErrorMessage, Ty: schema.TypeStrCompactNullable},
			&schema.ArrayCompact{Name: FieldDescribeAclsResponseResources, Ty: schema.NewSchema("ResourcesV2",
				&schema.Mfield{Name: FieldDescribeAclsResponseResourcesResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDescribeAclsResponseResourcesResourceName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeAclsResponseResourcesPatternType, Ty: schema.TypeInt8},
				&schema.ArrayCompact{Name: FieldDescribeAclsResponseResourcesAcls, Ty: schema.NewSchema("AclsV2",
					&schema.Mfield{Name: FieldDescribeAclsResponseResourcesAclsPrincipal, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldDescribeAclsResponseResourcesAclsHost, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldDescribeAclsResponseResourcesAclsOperation, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDescribeAclsResponseResourcesAclsPermissionType, Ty: schema.TypeInt8},
					&schema.SchemaTaggedFields{Name: FieldDescribeAclsResponseResourcesAclsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldDescribeAclsResponseResourcesTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDescribeAclsResponseTags},
		),
	}
}

const (
	// FieldDescribeAclsResponseErrorCode is: The error code, or 0 if there was no error.
	FieldDescribeAclsResponseErrorCode = "ErrorCode"
	// FieldDescribeAclsResponseErrorMessage is: The error message, or null if there was no error.
	FieldDescribeAclsResponseErrorMessage = "ErrorMessage"
	// FieldDescribeAclsResponseResources is: Each Resource that is referenced in an ACL.
	FieldDescribeAclsResponseResources = "Resources"
	// FieldDescribeAclsResponseResourcesAcls is: The ACLs.
	FieldDescribeAclsResponseResourcesAcls = "Acls"
	// FieldDescribeAclsResponseResourcesAclsHost is: The ACL host.
	FieldDescribeAclsResponseResourcesAclsHost = "Host"
	// FieldDescribeAclsResponseResourcesAclsOperation is: The ACL operation.
	FieldDescribeAclsResponseResourcesAclsOperation = "Operation"
	// FieldDescribeAclsResponseResourcesAclsPermissionType is: The ACL permission type.
	FieldDescribeAclsResponseResourcesAclsPermissionType = "PermissionType"
	// FieldDescribeAclsResponseResourcesAclsPrincipal is: The ACL principal.
	FieldDescribeAclsResponseResourcesAclsPrincipal = "Principal"
	// FieldDescribeAclsResponseResourcesAclsTags is: The tagged fields.
	FieldDescribeAclsResponseResourcesAclsTags = "Tags"
	// FieldDescribeAclsResponseResourcesPatternType is: The resource pattern type.
	FieldDescribeAclsResponseResourcesPatternType = "PatternType"
	// FieldDescribeAclsResponseResourcesResourceName is: The resource name.
	FieldDescribeAclsResponseResourcesResourceName = "ResourceName"
	// FieldDescribeAclsResponseResourcesResourceType is: The resource type.
	FieldDescribeAclsResponseResourcesResourceType = "ResourceType"
	// FieldDescribeAclsResponseResourcesTags is: The tagged fields.
	FieldDescribeAclsResponseResourcesTags = "Tags"
	// FieldDescribeAclsResponseTags is: The tagged fields.
	FieldDescribeAclsResponseTags = "Tags"
	// FieldDescribeAclsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldDescribeAclsResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeAclsResponse.json
const originalDescribeAclsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "response",
  "name": "DescribeAclsResponse",
  // Version 1 adds PatternType.
  // Starting in version 1, on quota violation, brokers send out responses before throttling.
  // Version 2 enables flexible versions.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The error code, or 0 if there was no error." },
    { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
      "about": "The error message, or null if there was no error." },
    { "name": "Resources", "type": "[]DescribeAclsResource", "versions": "0+",
      "about": "Each Resource that is referenced in an ACL.", "fields": [
      { "name": "ResourceType", "type": "int8", "versions": "0+",
        "about": "The resource type." },
      { "name": "ResourceName", "type": "string", "versions": "0+",
        "about": "The resource name." },
      { "name": "PatternType", "type": "int8", "versions": "1+", "default": "3", "ignorable": false,
        "about": "The resource pattern type." },
      { "name": "Acls", "type": "[]AclDescription", "versions": "0+",
        "about": "The ACLs.", "fields": [
        { "name": "Principal", "type": "string", "versions": "0+",
          "about": "The ACL principal." },
        { "name": "Host", "type": "string", "versions": "0+",
          "about": "The ACL host." },
        { "name": "Operation", "type": "int8", "versions": "0+",
          "about": "The ACL operation." },
        { "name": "PermissionType", "type": "int8", "versions": "0+",
          "about": "The ACL permission type." }
      ]}
    ]}
  ]
}
`
