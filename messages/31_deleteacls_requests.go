package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init31DeleteAclsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: DeleteAclsRequest, API Key: 31, Version: 0
		schema.NewSchema("DeleteAclsRequestv0",
			&schema.Array{Name: FieldDeleteAclsRequestFilters, Ty: schema.NewSchema("FiltersV0",
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersResourceTypeFilter, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersResourceNameFilter, Ty: schema.TypeStrNullable},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersPrincipalFilter, Ty: schema.TypeStrNullable},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersHostFilter, Ty: schema.TypeStrNullable},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersOperation, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersPermissionType, Ty: schema.TypeInt8},
			)},
		),

		// Message: DeleteAclsRequest, API Key: 31, Version: 1
		schema.NewSchema("DeleteAclsRequestv1",
			&schema.Array{Name: FieldDeleteAclsRequestFilters, Ty: schema.NewSchema("FiltersV1",
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersResourceTypeFilter, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersResourceNameFilter, Ty: schema.TypeStrNullable},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersPatternTypeFilter, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersPrincipalFilter, Ty: schema.TypeStrNullable},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersHostFilter, Ty: schema.TypeStrNullable},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersOperation, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersPermissionType, Ty: schema.TypeInt8},
			)},
		),

		// Message: DeleteAclsRequest, API Key: 31, Version: 2
		schema.NewSchema("DeleteAclsRequestv2",
			&schema.ArrayCompact{Name: FieldDeleteAclsRequestFilters, Ty: schema.NewSchema("FiltersV2",
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersResourceTypeFilter, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersResourceNameFilter, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersPatternTypeFilter, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersPrincipalFilter, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersHostFilter, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersOperation, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDeleteAclsRequestFiltersPermissionType, Ty: schema.TypeInt8},
				&schema.SchemaTaggedFields{Name: FieldDeleteAclsRequestFiltersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDeleteAclsRequestTags},
		),
	}
}

const (
	// FieldDeleteAclsRequestTags is a field name that can be used to resolve the correct struct field.
	FieldDeleteAclsRequestTags = "Tags"
	// FieldDeleteAclsRequestFiltersResourceTypeFilter is a field name that can be used to resolve the correct struct field.
	FieldDeleteAclsRequestFiltersResourceTypeFilter = "ResourceTypeFilter"
	// FieldDeleteAclsRequestFiltersResourceNameFilter is a field name that can be used to resolve the correct struct field.
	FieldDeleteAclsRequestFiltersResourceNameFilter = "ResourceNameFilter"
	// FieldDeleteAclsRequestFiltersPrincipalFilter is a field name that can be used to resolve the correct struct field.
	FieldDeleteAclsRequestFiltersPrincipalFilter = "PrincipalFilter"
	// FieldDeleteAclsRequestFiltersHostFilter is a field name that can be used to resolve the correct struct field.
	FieldDeleteAclsRequestFiltersHostFilter = "HostFilter"
	// FieldDeleteAclsRequestFiltersOperation is a field name that can be used to resolve the correct struct field.
	FieldDeleteAclsRequestFiltersOperation = "Operation"
	// FieldDeleteAclsRequestFiltersPermissionType is a field name that can be used to resolve the correct struct field.
	FieldDeleteAclsRequestFiltersPermissionType = "PermissionType"
	// FieldDeleteAclsRequestFilters is a field name that can be used to resolve the correct struct field.
	FieldDeleteAclsRequestFilters = "Filters"
	// FieldDeleteAclsRequestFiltersPatternTypeFilter is a field name that can be used to resolve the correct struct field.
	FieldDeleteAclsRequestFiltersPatternTypeFilter = "PatternTypeFilter"
	// FieldDeleteAclsRequestFiltersTags is a field name that can be used to resolve the correct struct field.
	FieldDeleteAclsRequestFiltersTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DeleteAclsRequest.json
const originalDeleteAclsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "request",
  "listeners": ["zkBroker", "broker", "controller"],
  "name": "DeleteAclsRequest",
  // Version 1 adds the pattern type.
  // Version 2 enables flexible versions.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "Filters", "type": "[]DeleteAclsFilter", "versions": "0+",
      "about": "The filters to use when deleting ACLs.", "fields": [
      { "name": "ResourceTypeFilter", "type": "int8", "versions": "0+",
        "about": "The resource type." },
      { "name": "ResourceNameFilter", "type": "string", "versions": "0+", "nullableVersions": "0+",
        "about": "The resource name." },
      { "name": "PatternTypeFilter", "type": "int8", "versions": "1+", "default": "3", "ignorable": false,
        "about": "The pattern type." },
      { "name": "PrincipalFilter", "type": "string", "versions": "0+", "nullableVersions": "0+",
        "about": "The principal filter, or null to accept all principals." },
      { "name": "HostFilter", "type": "string", "versions": "0+", "nullableVersions": "0+",
        "about": "The host filter, or null to accept all hosts." },
      { "name": "Operation", "type": "int8", "versions": "0+",
        "about": "The ACL operation." },
      { "name": "PermissionType", "type": "int8", "versions": "0+",
        "about": "The permission type." }
    ]}
  ]
}
`
