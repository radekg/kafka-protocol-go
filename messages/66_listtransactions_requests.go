package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init66ListTransactionsRequest() []schema.Schema {

	return []schema.Schema{
		// Message: ListTransactionsRequest, API Key: 66, Version: 0
		schema.NewSchema("ListTransactionsRequest:v0",
			&schema.Mfield{Name: FieldListTransactionsRequestStateFilters, Ty: schema.TypeStrCompactArray},
			&schema.Mfield{Name: FieldListTransactionsRequestProducerIdFilters, Ty: schema.TypeInt64CompactArray},
			&schema.SchemaTaggedFields{Name: FieldListTransactionsRequestTags},
		),
	}

}

const (

	// FieldListTransactionsRequestProducerIdFilters is: The producerIds to filter by: if empty, all transactions will be returned; if non-empty, only transactions which match one of the filtered producerIds will be returned
	FieldListTransactionsRequestProducerIdFilters = "ProducerIdFilters"

	// FieldListTransactionsRequestStateFilters is: The transaction states to filter by: if empty, all transactions are returned; if non-empty, then only transactions matching one of the filtered states will be returned
	FieldListTransactionsRequestStateFilters = "StateFilters"

	// FieldListTransactionsRequestTags is: The tagged fields.
	FieldListTransactionsRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ListTransactionsRequest.json
const originalListTransactionsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 66,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "ListTransactionsRequest",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "StateFilters", "type": "[]string", "versions": "0+",
      "about": "The transaction states to filter by: if empty, all transactions are returned; if non-empty, then only transactions matching one of the filtered states will be returned"
    },
    { "name": "ProducerIdFilters", "type": "[]int64", "versions": "0+", "entityType": "producerId",
      "about": "The producerIds to filter by: if empty, all transactions will be returned; if non-empty, only transactions which match one of the filtered producerIds will be returned"
    }
  ]
}
`
