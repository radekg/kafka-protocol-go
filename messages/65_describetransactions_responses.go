package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init65DescribeTransactionsResponse() []schema.Schema {

	return []schema.Schema{

		// Message: DescribeTransactionsResponse, API Key: 65, Version: 0
		schema.NewSchema("DescribeTransactionsResponsev0",
			&schema.Mfield{Name: FieldDescribeTransactionsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldDescribeTransactionsResponseTransactionStates, Ty: schema.NewSchema("TransactionStatesV0",
				&schema.Mfield{Name: FieldDescribeTransactionsResponseTransactionStatesErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeTransactionsResponseTransactionStatesTransactionalId, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeTransactionsResponseTransactionStatesTransactionState, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeTransactionsResponseTransactionStatesTransactionTimeoutMs, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldDescribeTransactionsResponseTransactionStatesTransactionStartTimeMs, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldDescribeTransactionsResponseTransactionStatesProducerId, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldDescribeTransactionsResponseTransactionStatesProducerEpoch, Ty: schema.TypeInt16},
				&schema.ArrayCompact{Name: FieldDescribeTransactionsResponseTransactionStatesTopics, Ty: schema.NewSchema("TopicsV0",
					&schema.Mfield{Name: FieldDescribeTransactionsResponseTransactionStatesTopicsTopic, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldDescribeTransactionsResponseTransactionStatesTopicsPartitions, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldDescribeTransactionsResponseTransactionStatesTopicsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldDescribeTransactionsResponseTransactionStatesTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDescribeTransactionsResponseTags},
		),
	}
}

const (
	// FieldDescribeTransactionsResponseTags is: The tagged fields.
	FieldDescribeTransactionsResponseTags = "Tags"
	// FieldDescribeTransactionsResponseTransactionStatesErrorCode is:
	FieldDescribeTransactionsResponseTransactionStatesErrorCode = "ErrorCode"
	// FieldDescribeTransactionsResponseTransactionStatesTransactionStartTimeMs is:
	FieldDescribeTransactionsResponseTransactionStatesTransactionStartTimeMs = "TransactionStartTimeMs"
	// FieldDescribeTransactionsResponseTransactionStatesProducerId is:
	FieldDescribeTransactionsResponseTransactionStatesProducerId = "ProducerId"
	// FieldDescribeTransactionsResponseTransactionStatesTopics is: The set of partitions included in the current transaction (if active). When a transaction is preparing to commit or abort, this will include only partitions which do not have markers.
	FieldDescribeTransactionsResponseTransactionStatesTopics = "Topics"
	// FieldDescribeTransactionsResponseTransactionStatesTopicsTags is: The tagged fields.
	FieldDescribeTransactionsResponseTransactionStatesTopicsTags = "Tags"
	// FieldDescribeTransactionsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldDescribeTransactionsResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldDescribeTransactionsResponseTransactionStates is:
	FieldDescribeTransactionsResponseTransactionStates = "TransactionStates"
	// FieldDescribeTransactionsResponseTransactionStatesTransactionalId is:
	FieldDescribeTransactionsResponseTransactionStatesTransactionalId = "TransactionalId"
	// FieldDescribeTransactionsResponseTransactionStatesProducerEpoch is:
	FieldDescribeTransactionsResponseTransactionStatesProducerEpoch = "ProducerEpoch"
	// FieldDescribeTransactionsResponseTransactionStatesTopicsTopic is:
	FieldDescribeTransactionsResponseTransactionStatesTopicsTopic = "Topic"
	// FieldDescribeTransactionsResponseTransactionStatesTags is: The tagged fields.
	FieldDescribeTransactionsResponseTransactionStatesTags = "Tags"
	// FieldDescribeTransactionsResponseTransactionStatesTransactionState is:
	FieldDescribeTransactionsResponseTransactionStatesTransactionState = "TransactionState"
	// FieldDescribeTransactionsResponseTransactionStatesTransactionTimeoutMs is:
	FieldDescribeTransactionsResponseTransactionStatesTransactionTimeoutMs = "TransactionTimeoutMs"
	// FieldDescribeTransactionsResponseTransactionStatesTopicsPartitions is:
	FieldDescribeTransactionsResponseTransactionStatesTopicsPartitions = "Partitions"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeTransactionsResponse.json
const originalDescribeTransactionsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "response",
  "name": "DescribeTransactionsResponse",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
      { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
        "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
      { "name": "TransactionStates", "type": "[]TransactionState", "versions": "0+", "fields": [
        { "name": "ErrorCode", "type": "int16", "versions": "0+" },
        { "name": "TransactionalId", "type": "string", "versions": "0+", "entityType": "transactionalId" },
        { "name": "TransactionState", "type": "string", "versions": "0+" },
        { "name": "TransactionTimeoutMs", "type": "int32", "versions": "0+" },
        { "name": "TransactionStartTimeMs", "type": "int64", "versions": "0+" },
        { "name": "ProducerId", "type": "int64", "versions": "0+", "entityType": "producerId" },
        { "name": "ProducerEpoch", "type": "int16", "versions": "0+" },
        { "name": "Topics", "type": "[]TopicData", "versions": "0+",
          "about": "The set of partitions included in the current transaction (if active). When a transaction is preparing to commit or abort, this will include only partitions which do not have markers.",
          "fields": [
            { "name": "Topic", "type": "string", "versions": "0+", "entityType": "topicName", "mapKey": true },
            { "name": "Partitions", "type": "[]int32", "versions": "0+" }
          ]
        }
      ]}
  ]
}
`
