package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init42DeleteGroupsRequest() []schema.Schema {

	return []schema.Schema{
		// Message: DeleteGroupsRequest, API Key: 42, Version: 0
		schema.NewSchema("DeleteGroupsRequest:v0",
			&schema.Mfield{Name: FieldDeleteGroupsRequestGroupsNames, Ty: schema.TypeStrArray},
		),

		// Message: DeleteGroupsRequest, API Key: 42, Version: 1
		schema.NewSchema("DeleteGroupsRequest:v1",
			&schema.Mfield{Name: FieldDeleteGroupsRequestGroupsNames, Ty: schema.TypeStrArray},
		),

		// Message: DeleteGroupsRequest, API Key: 42, Version: 2
		schema.NewSchema("DeleteGroupsRequest:v2",
			&schema.Mfield{Name: FieldDeleteGroupsRequestGroupsNames, Ty: schema.TypeStrCompactArray},
			&schema.SchemaTaggedFields{Name: FieldDeleteGroupsRequestTags},
		),
	}

}

const (

	// FieldDeleteGroupsRequestGroupsNames is: The group names to delete.
	FieldDeleteGroupsRequestGroupsNames = "GroupsNames"

	// FieldDeleteGroupsRequestTags is: The tagged fields.
	FieldDeleteGroupsRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DeleteGroupsRequest.json
const originalDeleteGroupsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 42,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "DeleteGroupsRequest",
  // Version 1 is the same as version 0.
  //
  // Version 2 is the first flexible version.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "GroupsNames", "type": "[]string", "versions": "0+", "entityType": "groupId",
      "about": "The group names to delete." }
  ]
}
`
