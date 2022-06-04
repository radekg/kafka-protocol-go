package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init16ListGroupsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: ListGroupsRequest, API Key: 16, Version: 0
		schema.NewSchema("ListGroupsRequestv0"),
		// Message: ListGroupsRequest, API Key: 16, Version: 1
		schema.NewSchema("ListGroupsRequestv1"),
		// Message: ListGroupsRequest, API Key: 16, Version: 2
		schema.NewSchema("ListGroupsRequestv2"),
		// Message: ListGroupsRequest, API Key: 16, Version: 3
		schema.NewSchema("ListGroupsRequestv3"),
		// Message: ListGroupsRequest, API Key: 16, Version: 4
		schema.NewSchema("ListGroupsRequestv4",
			&schema.Mfield{Name: FieldListGroupsRequestStatesFilter, Ty: schema.TypeStrCompactArray},
			&schema.SchemaTaggedFields{Name: FieldListGroupsRequestTags},
		),
	}
}

const (
	// FieldListGroupsRequestStatesFilter is a field name that can be used to resolve the correct struct field.
	FieldListGroupsRequestStatesFilter = "StatesFilter"
	// FieldListGroupsRequestTags is a field name that can be used to resolve the correct struct field.
	FieldListGroupsRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ListGroupsRequest.json
const originalListGroupsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 16,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "ListGroupsRequest",
  // Version 1 and 2 are the same as version 0.
  //
  // Version 3 is the first flexible version.
  //
  // Version 4 adds the StatesFilter field (KIP-518).
  "validVersions": "0-4",
  "flexibleVersions": "3+",
  "fields": [
    { "name": "StatesFilter", "type": "[]string", "versions": "4+",
      "about": "The states of the groups we want to list. If empty all groups are returned with their state."
    }
  ]
}
`
