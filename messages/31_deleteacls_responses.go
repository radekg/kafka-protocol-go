package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init31DeleteAclsResponse() []schema.Schema {

	return []schema.Schema{
		// Message: DeleteAclsResponse, API Key: 31, Version: 0
		schema.NewSchema("DeleteAclsResponse:v0",
			&schema.Mfield{Name: FieldDeleteAclsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDeleteAclsResponseFilterResults, Ty: schema.NewSchema("[]DeleteAclsFilterResult:v0",
				&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsErrorMessage, Ty: schema.TypeStrNullable},
				&schema.Array{Name: FieldDeleteAclsResponseFilterResultsMatchingAcls, Ty: schema.NewSchema("[]DeleteAclsMatchingAcl:v0",
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsErrorMessage, Ty: schema.TypeStrNullable},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsResourceType, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsResourceName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsPrincipal, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsOperation, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsPermissionType, Ty: schema.TypeInt8},
				)},
			)},
		),

		// Message: DeleteAclsResponse, API Key: 31, Version: 1
		schema.NewSchema("DeleteAclsResponse:v1",
			&schema.Mfield{Name: FieldDeleteAclsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDeleteAclsResponseFilterResults, Ty: schema.NewSchema("[]DeleteAclsFilterResult:v1",
				&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsErrorMessage, Ty: schema.TypeStrNullable},
				&schema.Array{Name: FieldDeleteAclsResponseFilterResultsMatchingAcls, Ty: schema.NewSchema("[]DeleteAclsMatchingAcl:v1",
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsErrorMessage, Ty: schema.TypeStrNullable},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsResourceType, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsResourceName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsPatternType, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsPrincipal, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsHost, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsOperation, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsPermissionType, Ty: schema.TypeInt8},
				)},
			)},
		),

		// Message: DeleteAclsResponse, API Key: 31, Version: 2
		schema.NewSchema("DeleteAclsResponse:v2",
			&schema.Mfield{Name: FieldDeleteAclsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldDeleteAclsResponseFilterResults, Ty: schema.NewSchema("[]DeleteAclsFilterResult:v2",
				&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.ArrayCompact{Name: FieldDeleteAclsResponseFilterResultsMatchingAcls, Ty: schema.NewSchema("[]DeleteAclsMatchingAcl:v2",
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsErrorMessage, Ty: schema.TypeStrCompactNullable},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsResourceType, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsResourceName, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsPatternType, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsPrincipal, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsHost, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsOperation, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsPermissionType, Ty: schema.TypeInt8},
					&schema.SchemaTaggedFields{Name: FieldDeleteAclsResponseFilterResultsMatchingAclsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldDeleteAclsResponseFilterResultsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDeleteAclsResponseTags},
		),
	}

}

const (

	// FieldDeleteAclsResponseFilterResults is: The results for each filter.
	FieldDeleteAclsResponseFilterResults = "FilterResults"

	// FieldDeleteAclsResponseFilterResultsErrorCode is: The error code, or 0 if the filter succeeded.
	FieldDeleteAclsResponseFilterResultsErrorCode = "ErrorCode"

	// FieldDeleteAclsResponseFilterResultsErrorMessage is: The error message, or null if the filter succeeded.
	FieldDeleteAclsResponseFilterResultsErrorMessage = "ErrorMessage"

	// FieldDeleteAclsResponseFilterResultsMatchingAcls is: The ACLs which matched this filter.
	FieldDeleteAclsResponseFilterResultsMatchingAcls = "MatchingAcls"

	// FieldDeleteAclsResponseFilterResultsMatchingAclsErrorCode is: The deletion error code, or 0 if the deletion succeeded.
	FieldDeleteAclsResponseFilterResultsMatchingAclsErrorCode = "ErrorCode"

	// FieldDeleteAclsResponseFilterResultsMatchingAclsErrorMessage is: The deletion error message, or null if the deletion succeeded.
	FieldDeleteAclsResponseFilterResultsMatchingAclsErrorMessage = "ErrorMessage"

	// FieldDeleteAclsResponseFilterResultsMatchingAclsHost is: The ACL host.
	FieldDeleteAclsResponseFilterResultsMatchingAclsHost = "Host"

	// FieldDeleteAclsResponseFilterResultsMatchingAclsOperation is: The ACL operation.
	FieldDeleteAclsResponseFilterResultsMatchingAclsOperation = "Operation"

	// FieldDeleteAclsResponseFilterResultsMatchingAclsPatternType is: The ACL resource pattern type.
	FieldDeleteAclsResponseFilterResultsMatchingAclsPatternType = "PatternType"

	// FieldDeleteAclsResponseFilterResultsMatchingAclsPermissionType is: The ACL permission type.
	FieldDeleteAclsResponseFilterResultsMatchingAclsPermissionType = "PermissionType"

	// FieldDeleteAclsResponseFilterResultsMatchingAclsPrincipal is: The ACL principal.
	FieldDeleteAclsResponseFilterResultsMatchingAclsPrincipal = "Principal"

	// FieldDeleteAclsResponseFilterResultsMatchingAclsResourceName is: The ACL resource name.
	FieldDeleteAclsResponseFilterResultsMatchingAclsResourceName = "ResourceName"

	// FieldDeleteAclsResponseFilterResultsMatchingAclsResourceType is: The ACL resource type.
	FieldDeleteAclsResponseFilterResultsMatchingAclsResourceType = "ResourceType"

	// FieldDeleteAclsResponseFilterResultsMatchingAclsTags is: The tagged fields.
	FieldDeleteAclsResponseFilterResultsMatchingAclsTags = "Tags"

	// FieldDeleteAclsResponseFilterResultsTags is: The tagged fields.
	FieldDeleteAclsResponseFilterResultsTags = "Tags"

	// FieldDeleteAclsResponseTags is: The tagged fields.
	FieldDeleteAclsResponseTags = "Tags"

	// FieldDeleteAclsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldDeleteAclsResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DeleteAclsResponse.json
const originalDeleteAclsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 31,
  "type": "response",
  "name": "DeleteAclsResponse",
  // Version 1 adds the resource pattern type.
  // Starting in version 1, on quota violation, brokers send out responses before throttling.
  // Version 2 enables flexible versions.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "FilterResults", "type": "[]DeleteAclsFilterResult", "versions": "0+",
      "about": "The results for each filter.", "fields": [
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The error code, or 0 if the filter succeeded." },
      { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
        "about": "The error message, or null if the filter succeeded." },
      { "name": "MatchingAcls", "type": "[]DeleteAclsMatchingAcl", "versions": "0+",
        "about": "The ACLs which matched this filter.", "fields": [
        { "name": "ErrorCode", "type": "int16", "versions": "0+",
          "about": "The deletion error code, or 0 if the deletion succeeded." },
        { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
          "about": "The deletion error message, or null if the deletion succeeded." },
        { "name": "ResourceType", "type": "int8", "versions": "0+",
          "about": "The ACL resource type." },
        { "name": "ResourceName", "type": "string", "versions": "0+",
          "about": "The ACL resource name." },
        { "name": "PatternType", "type": "int8", "versions": "1+", "default": "3", "ignorable": false,
          "about": "The ACL resource pattern type." },
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
