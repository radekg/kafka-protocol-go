package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init15DescribeGroupsRequest() []schema.Schema {

	return []schema.Schema{
		// Message: DescribeGroupsRequest, API Key: 15, Version: 0
		schema.NewSchema("DescribeGroupsRequest:v0",
			&schema.Mfield{Name: FieldDescribeGroupsRequestGroups, Ty: schema.TypeStrArray},
		),

		// Message: DescribeGroupsRequest, API Key: 15, Version: 1
		schema.NewSchema("DescribeGroupsRequest:v1",
			&schema.Mfield{Name: FieldDescribeGroupsRequestGroups, Ty: schema.TypeStrArray},
		),

		// Message: DescribeGroupsRequest, API Key: 15, Version: 2
		schema.NewSchema("DescribeGroupsRequest:v2",
			&schema.Mfield{Name: FieldDescribeGroupsRequestGroups, Ty: schema.TypeStrArray},
		),

		// Message: DescribeGroupsRequest, API Key: 15, Version: 3
		schema.NewSchema("DescribeGroupsRequest:v3",
			&schema.Mfield{Name: FieldDescribeGroupsRequestGroups, Ty: schema.TypeStrArray},
			&schema.Mfield{Name: FieldDescribeGroupsRequestIncludeAuthorizedOperations, Ty: schema.TypeBool},
		),

		// Message: DescribeGroupsRequest, API Key: 15, Version: 4
		schema.NewSchema("DescribeGroupsRequest:v4",
			&schema.Mfield{Name: FieldDescribeGroupsRequestGroups, Ty: schema.TypeStrArray},
			&schema.Mfield{Name: FieldDescribeGroupsRequestIncludeAuthorizedOperations, Ty: schema.TypeBool},
		),

		// Message: DescribeGroupsRequest, API Key: 15, Version: 5
		schema.NewSchema("DescribeGroupsRequest:v5",
			&schema.Mfield{Name: FieldDescribeGroupsRequestGroups, Ty: schema.TypeStrCompactArray},
			&schema.Mfield{Name: FieldDescribeGroupsRequestIncludeAuthorizedOperations, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldDescribeGroupsRequestTags},
		),
	}

}

const (

	// FieldDescribeGroupsRequestGroups is: The names of the groups to describe
	FieldDescribeGroupsRequestGroups = "Groups"

	// FieldDescribeGroupsRequestIncludeAuthorizedOperations is: Whether to include authorized operations.
	FieldDescribeGroupsRequestIncludeAuthorizedOperations = "IncludeAuthorizedOperations"

	// FieldDescribeGroupsRequestTags is: The tagged fields.
	FieldDescribeGroupsRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeGroupsRequest.json
const originalDescribeGroupsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 15,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "DescribeGroupsRequest",
  // Versions 1 and 2 are the same as version 0.
  //
  // Starting in version 3, authorized operations can be requested.
  //
  // Starting in version 4, the response will include group.instance.id info for members.
  //
  // Version 5 is the first flexible version.
  "validVersions": "0-5",
  "flexibleVersions": "5+",
  "fields": [
    { "name": "Groups", "type": "[]string", "versions": "0+", "entityType": "groupId",
      "about": "The names of the groups to describe" },
    { "name": "IncludeAuthorizedOperations", "type": "bool", "versions": "3+",
      "about": "Whether to include authorized operations." }
  ]
}
`
