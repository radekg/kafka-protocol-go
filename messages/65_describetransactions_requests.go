package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init65DescribeTransactionsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: DescribeTransactionsRequest, API Key: 65, Version: 0
		schema.NewSchema("DescribeTransactionsRequestv0",
			&schema.Mfield{Name: FieldDescribeTransactionsRequestTransactionalIds, Ty: schema.TypeStrCompactArray},
			&schema.SchemaTaggedFields{Name: FieldDescribeTransactionsRequestTags},
		),
	}
}

const (
	// FieldDescribeTransactionsRequestTags is: The tagged fields.
	FieldDescribeTransactionsRequestTags = "Tags"
	// FieldDescribeTransactionsRequestTransactionalIds is: Array of transactionalIds to include in describe results. If empty, then no results will be returned.
	FieldDescribeTransactionsRequestTransactionalIds = "TransactionalIds"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeTransactionsRequest.json
const originalDescribeTransactionsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 65,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "DescribeTransactionsRequest",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
      { "name": "TransactionalIds", "entityType": "transactionalId", "type": "[]string", "versions": "0+",
        "about": "Array of transactionalIds to include in describe results. If empty, then no results will be returned." }
  ]
}
`
