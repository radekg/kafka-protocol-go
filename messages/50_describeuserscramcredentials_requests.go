package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init50DescribeUserScramCredentialsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: DescribeUserScramCredentialsRequest, API Key: 50, Version: 0
		schema.NewSchema("DescribeUserScramCredentialsRequestv0", 
			&schema.ArrayCompact{Name: FieldDescribeUserScramCredentialsRequestUsers, Ty: schema.NewSchema("UsersV0",
				&schema.Mfield{Name: FieldDescribeUserScramCredentialsRequestUsersName, Ty: schema.TypeStrCompact},
				&schema.SchemaTaggedFields{Name: FieldDescribeUserScramCredentialsRequestUsersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDescribeUserScramCredentialsRequestTags},
		),

	}
}

const (
	// FieldDescribeUserScramCredentialsRequestUsers is a field name that can be used to resolve the correct struct field.
	FieldDescribeUserScramCredentialsRequestUsers = "Users"
	// FieldDescribeUserScramCredentialsRequestUsersName is a field name that can be used to resolve the correct struct field.
	FieldDescribeUserScramCredentialsRequestUsersName = "Name"
	// FieldDescribeUserScramCredentialsRequestUsersTags is a field name that can be used to resolve the correct struct field.
	FieldDescribeUserScramCredentialsRequestUsersTags = "Tags"
	// FieldDescribeUserScramCredentialsRequestTags is a field name that can be used to resolve the correct struct field.
	FieldDescribeUserScramCredentialsRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeUserScramCredentialsRequest.json
const originalDescribeUserScramCredentialsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 50,
  "type": "request",
  "listeners": ["zkBroker"],
  "name": "DescribeUserScramCredentialsRequest",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "Users", "type": "[]UserName", "versions": "0+", "nullableVersions": "0+",
      "about": "The users to describe, or null/empty to describe all users.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+",
        "about": "The user name." }
    ]}
  ]
}
`

