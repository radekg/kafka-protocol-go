package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init66ListTransactionsResponse() []schema.Schema {

	return []schema.Schema{
		// Message: ListTransactionsResponse, API Key: 66, Version: 0
		schema.NewSchema("ListTransactionsResponse:v0",
			&schema.Mfield{Name: FieldListTransactionsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldListTransactionsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldListTransactionsResponseUnknownStateFilters, Ty: schema.TypeStrCompactArray},
			&schema.ArrayCompact{Name: FieldListTransactionsResponseTransactionStates, Ty: schema.NewSchema("TransactionStates:v0",
				&schema.Mfield{Name: FieldListTransactionsResponseTransactionStatesTransactionalId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldListTransactionsResponseTransactionStatesProducerId, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldListTransactionsResponseTransactionStatesTransactionState, Ty: schema.TypeStrCompact},
				&schema.SchemaTaggedFields{Name: FieldListTransactionsResponseTransactionStatesTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldListTransactionsResponseTags},
		),
	}

}

const (

	// FieldListTransactionsResponseErrorCode is:
	FieldListTransactionsResponseErrorCode = "ErrorCode"

	// FieldListTransactionsResponseTags is: The tagged fields.
	FieldListTransactionsResponseTags = "Tags"

	// FieldListTransactionsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldListTransactionsResponseThrottleTimeMs = "ThrottleTimeMs"

	// FieldListTransactionsResponseTransactionStates is:
	FieldListTransactionsResponseTransactionStates = "TransactionStates"

	// FieldListTransactionsResponseTransactionStatesProducerId is:
	FieldListTransactionsResponseTransactionStatesProducerId = "ProducerId"

	// FieldListTransactionsResponseTransactionStatesTags is: The tagged fields.
	FieldListTransactionsResponseTransactionStatesTags = "Tags"

	// FieldListTransactionsResponseTransactionStatesTransactionState is: The current transaction state of the producer
	FieldListTransactionsResponseTransactionStatesTransactionState = "TransactionState"

	// FieldListTransactionsResponseTransactionStatesTransactionalId is:
	FieldListTransactionsResponseTransactionStatesTransactionalId = "TransactionalId"

	// FieldListTransactionsResponseUnknownStateFilters is: Set of state filters provided in the request which were unknown to the transaction coordinator
	FieldListTransactionsResponseUnknownStateFilters = "UnknownStateFilters"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ListTransactionsResponse.json
const originalListTransactionsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "response",
  "name": "ListTransactionsResponse",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
      { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
        "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
      { "name": "ErrorCode", "type": "int16", "versions": "0+" },
      { "name": "UnknownStateFilters", "type": "[]string", "versions": "0+",
        "about": "Set of state filters provided in the request which were unknown to the transaction coordinator" },
      { "name": "TransactionStates", "type": "[]TransactionState", "versions": "0+", "fields": [
        { "name": "TransactionalId", "type": "string", "versions": "0+", "entityType": "transactionalId" },
        { "name": "ProducerId", "type": "int64", "versions": "0+", "entityType": "producerId" },
        { "name": "TransactionState", "type": "string", "versions": "0+",
          "about": "The current transaction state of the producer" }
    ]}
  ]
}
`
