package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init60DescribeClusterRequest() []schema.Schema {

	return []schema.Schema{

		// Message: DescribeClusterRequest, API Key: 60, Version: 0
		schema.NewSchema("DescribeClusterRequestv0",
			&schema.Mfield{Name: FieldDescribeClusterRequestIncludeClusterAuthorizedOperations, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldDescribeClusterRequestTags},
		),
	}
}

const (
	// FieldDescribeClusterRequestIncludeClusterAuthorizedOperations is a field name that can be used to resolve the correct struct field.
	FieldDescribeClusterRequestIncludeClusterAuthorizedOperations = "IncludeClusterAuthorizedOperations"
	// FieldDescribeClusterRequestTags is a field name that can be used to resolve the correct struct field.
	FieldDescribeClusterRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeClusterRequest.json
const originalDescribeClusterRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 60,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "DescribeClusterRequest",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "IncludeClusterAuthorizedOperations", "type": "bool", "versions": "0+",
      "about": "Whether to include cluster authorized operations." }
  ]
}
`
