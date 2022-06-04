package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init24AddPartitionsToTxnRequest() []schema.Schema {

	return []schema.Schema{

		// Message: AddPartitionsToTxnRequest, API Key: 24, Version: 0
		schema.NewSchema("AddPartitionsToTxnRequestv0", 
			&schema.Mfield{Name: FieldAddPartitionsToTxnRequestTransactionalId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldAddPartitionsToTxnRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldAddPartitionsToTxnRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldAddPartitionsToTxnRequestTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldAddPartitionsToTxnRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldAddPartitionsToTxnRequestTopicsPartitions, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: AddPartitionsToTxnRequest, API Key: 24, Version: 1
		schema.NewSchema("AddPartitionsToTxnRequestv1", 
			&schema.Mfield{Name: FieldAddPartitionsToTxnRequestTransactionalId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldAddPartitionsToTxnRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldAddPartitionsToTxnRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldAddPartitionsToTxnRequestTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldAddPartitionsToTxnRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldAddPartitionsToTxnRequestTopicsPartitions, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: AddPartitionsToTxnRequest, API Key: 24, Version: 2
		schema.NewSchema("AddPartitionsToTxnRequestv2", 
			&schema.Mfield{Name: FieldAddPartitionsToTxnRequestTransactionalId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldAddPartitionsToTxnRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldAddPartitionsToTxnRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldAddPartitionsToTxnRequestTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldAddPartitionsToTxnRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldAddPartitionsToTxnRequestTopicsPartitions, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: AddPartitionsToTxnRequest, API Key: 24, Version: 3
		schema.NewSchema("AddPartitionsToTxnRequestv3", 
			&schema.Mfield{Name: FieldAddPartitionsToTxnRequestTransactionalId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldAddPartitionsToTxnRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldAddPartitionsToTxnRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldAddPartitionsToTxnRequestTopics, Ty: schema.NewSchema("TopicsV3",
				&schema.Mfield{Name: FieldAddPartitionsToTxnRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldAddPartitionsToTxnRequestTopicsPartitions, Ty: schema.TypeInt32CompactArray},
				&schema.SchemaTaggedFields{Name: FieldAddPartitionsToTxnRequestTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldAddPartitionsToTxnRequestTags},
		),

	}
}

const (
	// FieldAddPartitionsToTxnRequestTopics is a field name that can be used to resolve the correct struct field.
	FieldAddPartitionsToTxnRequestTopics = "Topics"
	// FieldAddPartitionsToTxnRequestTopicsName is a field name that can be used to resolve the correct struct field.
	FieldAddPartitionsToTxnRequestTopicsName = "Name"
	// FieldAddPartitionsToTxnRequestTopicsPartitions is a field name that can be used to resolve the correct struct field.
	FieldAddPartitionsToTxnRequestTopicsPartitions = "Partitions"
	// FieldAddPartitionsToTxnRequestTopicsTags is a field name that can be used to resolve the correct struct field.
	FieldAddPartitionsToTxnRequestTopicsTags = "Tags"
	// FieldAddPartitionsToTxnRequestTags is a field name that can be used to resolve the correct struct field.
	FieldAddPartitionsToTxnRequestTags = "Tags"
	// FieldAddPartitionsToTxnRequestTransactionalId is a field name that can be used to resolve the correct struct field.
	FieldAddPartitionsToTxnRequestTransactionalId = "TransactionalId"
	// FieldAddPartitionsToTxnRequestProducerId is a field name that can be used to resolve the correct struct field.
	FieldAddPartitionsToTxnRequestProducerId = "ProducerId"
	// FieldAddPartitionsToTxnRequestProducerEpoch is a field name that can be used to resolve the correct struct field.
	FieldAddPartitionsToTxnRequestProducerEpoch = "ProducerEpoch"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AddPartitionsToTxnRequest.json
const originalAddPartitionsToTxnRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 24,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "AddPartitionsToTxnRequest",
  // Version 1 is the same as version 0.
  //
  // Version 2 adds the support for new error code PRODUCER_FENCED.
  //
  // Version 3 enables flexible versions.
  "validVersions": "0-3",
  "flexibleVersions": "3+",
  "fields": [
    { "name": "TransactionalId", "type": "string", "versions": "0+", "entityType": "transactionalId",
      "about": "The transactional id corresponding to the transaction."},
    { "name": "ProducerId", "type": "int64", "versions": "0+", "entityType": "producerId",
      "about": "Current producer id in use by the transactional id." },
    { "name": "ProducerEpoch", "type": "int16", "versions": "0+",
      "about": "Current epoch associated with the producer id." },
    { "name": "Topics", "type": "[]AddPartitionsToTxnTopic", "versions": "0+",
      "about": "The partitions to add to the transaction.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "mapKey": true, "entityType": "topicName",
        "about": "The name of the topic." },
      { "name": "Partitions", "type": "[]int32", "versions": "0+",
        "about": "The partition indexes to add to the transaction" }
    ]}
  ]
}
`

