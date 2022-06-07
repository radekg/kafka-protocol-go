package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init27WriteTxnMarkersRequest() []schema.Schema {

	return []schema.Schema{
		// Message: WriteTxnMarkersRequest, API Key: 27, Version: 0
		schema.NewSchema("WriteTxnMarkersRequest:v0",
			&schema.Array{Name: FieldWriteTxnMarkersRequestMarkers, Ty: schema.NewSchema("[]WritableTxnMarker:v0",
				&schema.Mfield{Name: FieldWriteTxnMarkersRequestMarkersProducerId, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldWriteTxnMarkersRequestMarkersProducerEpoch, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldWriteTxnMarkersRequestMarkersTransactionResult, Ty: schema.TypeBool},
				&schema.Array{Name: FieldWriteTxnMarkersRequestMarkersTopics, Ty: schema.NewSchema("[]WritableTxnMarkerTopic:v0",
					&schema.Mfield{Name: FieldWriteTxnMarkersRequestMarkersTopicsName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldWriteTxnMarkersRequestMarkersTopicsPartitionIndexes, Ty: schema.TypeInt32Array},
				)},
				&schema.Mfield{Name: FieldWriteTxnMarkersRequestMarkersCoordinatorEpoch, Ty: schema.TypeInt32},
			)},
		),

		// Message: WriteTxnMarkersRequest, API Key: 27, Version: 1
		schema.NewSchema("WriteTxnMarkersRequest:v1",
			&schema.ArrayCompact{Name: FieldWriteTxnMarkersRequestMarkers, Ty: schema.NewSchema("[]WritableTxnMarker:v1",
				&schema.Mfield{Name: FieldWriteTxnMarkersRequestMarkersProducerId, Ty: schema.TypeInt64},
				&schema.Mfield{Name: FieldWriteTxnMarkersRequestMarkersProducerEpoch, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldWriteTxnMarkersRequestMarkersTransactionResult, Ty: schema.TypeBool},
				&schema.ArrayCompact{Name: FieldWriteTxnMarkersRequestMarkersTopics, Ty: schema.NewSchema("[]WritableTxnMarkerTopic:v1",
					&schema.Mfield{Name: FieldWriteTxnMarkersRequestMarkersTopicsName, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldWriteTxnMarkersRequestMarkersTopicsPartitionIndexes, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldWriteTxnMarkersRequestMarkersTopicsTags},
				)},
				&schema.Mfield{Name: FieldWriteTxnMarkersRequestMarkersCoordinatorEpoch, Ty: schema.TypeInt32},
				&schema.SchemaTaggedFields{Name: FieldWriteTxnMarkersRequestMarkersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldWriteTxnMarkersRequestTags},
		),
	}

}

const (

	// FieldWriteTxnMarkersRequestMarkers is: The transaction markers to be written.
	FieldWriteTxnMarkersRequestMarkers = "Markers"

	// FieldWriteTxnMarkersRequestMarkersCoordinatorEpoch is: Epoch associated with the transaction state partition hosted by this transaction coordinator
	FieldWriteTxnMarkersRequestMarkersCoordinatorEpoch = "CoordinatorEpoch"

	// FieldWriteTxnMarkersRequestMarkersProducerEpoch is: The current epoch associated with the producer ID.
	FieldWriteTxnMarkersRequestMarkersProducerEpoch = "ProducerEpoch"

	// FieldWriteTxnMarkersRequestMarkersProducerId is: The current producer ID.
	FieldWriteTxnMarkersRequestMarkersProducerId = "ProducerId"

	// FieldWriteTxnMarkersRequestMarkersTags is: The tagged fields.
	FieldWriteTxnMarkersRequestMarkersTags = "Tags"

	// FieldWriteTxnMarkersRequestMarkersTopics is: Each topic that we want to write transaction marker(s) for.
	FieldWriteTxnMarkersRequestMarkersTopics = "Topics"

	// FieldWriteTxnMarkersRequestMarkersTopicsName is: The topic name.
	FieldWriteTxnMarkersRequestMarkersTopicsName = "Name"

	// FieldWriteTxnMarkersRequestMarkersTopicsPartitionIndexes is: The indexes of the partitions to write transaction markers for.
	FieldWriteTxnMarkersRequestMarkersTopicsPartitionIndexes = "PartitionIndexes"

	// FieldWriteTxnMarkersRequestMarkersTopicsTags is: The tagged fields.
	FieldWriteTxnMarkersRequestMarkersTopicsTags = "Tags"

	// FieldWriteTxnMarkersRequestMarkersTransactionResult is: The result of the transaction to write to the partitions (false = ABORT, true = COMMIT).
	FieldWriteTxnMarkersRequestMarkersTransactionResult = "TransactionResult"

	// FieldWriteTxnMarkersRequestTags is: The tagged fields.
	FieldWriteTxnMarkersRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/WriteTxnMarkersRequest.json
const originalWriteTxnMarkersRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 27,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "WriteTxnMarkersRequest",
  // Version 1 enables flexible versions.
  "validVersions": "0-1",
  "flexibleVersions": "1+",
  "fields": [
    { "name": "Markers", "type": "[]WritableTxnMarker", "versions": "0+",
      "about": "The transaction markers to be written.", "fields": [
      { "name": "ProducerId", "type": "int64", "versions": "0+", "entityType": "producerId",
        "about": "The current producer ID."},
      { "name": "ProducerEpoch", "type": "int16", "versions": "0+",
        "about": "The current epoch associated with the producer ID." },
      { "name": "TransactionResult", "type": "bool", "versions": "0+",
        "about": "The result of the transaction to write to the partitions (false = ABORT, true = COMMIT)." },
      { "name": "Topics", "type": "[]WritableTxnMarkerTopic", "versions": "0+",
        "about": "Each topic that we want to write transaction marker(s) for.", "fields": [
        { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
          "about": "The topic name." },
        { "name": "PartitionIndexes", "type": "[]int32", "versions": "0+",
          "about": "The indexes of the partitions to write transaction markers for." }
      ]},
      { "name": "CoordinatorEpoch", "type": "int32", "versions": "0+",
        "about": "Epoch associated with the transaction state partition hosted by this transaction coordinator" }
    ]}
  ]
}
`
