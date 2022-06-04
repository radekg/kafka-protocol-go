package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init25AddOffsetsToTxnRequest() []schema.Schema {

	return []schema.Schema{

		// Message: AddOffsetsToTxnRequest, API Key: 25, Version: 0
		schema.NewSchema("AddOffsetsToTxnRequestv0", 
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestTransactionalId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestGroupId, Ty: schema.TypeStr},
		),

		// Message: AddOffsetsToTxnRequest, API Key: 25, Version: 1
		schema.NewSchema("AddOffsetsToTxnRequestv1", 
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestTransactionalId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestGroupId, Ty: schema.TypeStr},
		),

		// Message: AddOffsetsToTxnRequest, API Key: 25, Version: 2
		schema.NewSchema("AddOffsetsToTxnRequestv2", 
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestTransactionalId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestGroupId, Ty: schema.TypeStr},
		),

		// Message: AddOffsetsToTxnRequest, API Key: 25, Version: 3
		schema.NewSchema("AddOffsetsToTxnRequestv3", 
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestTransactionalId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldAddOffsetsToTxnRequestGroupId, Ty: schema.TypeStrCompact},
			&schema.SchemaTaggedFields{Name: FieldAddOffsetsToTxnRequestTags},
		),

	}
}

const (
	// FieldAddOffsetsToTxnRequestTags is a field name that can be used to resolve the correct struct field.
	FieldAddOffsetsToTxnRequestTags = "Tags"
	// FieldAddOffsetsToTxnRequestTransactionalId is a field name that can be used to resolve the correct struct field.
	FieldAddOffsetsToTxnRequestTransactionalId = "TransactionalId"
	// FieldAddOffsetsToTxnRequestProducerId is a field name that can be used to resolve the correct struct field.
	FieldAddOffsetsToTxnRequestProducerId = "ProducerId"
	// FieldAddOffsetsToTxnRequestProducerEpoch is a field name that can be used to resolve the correct struct field.
	FieldAddOffsetsToTxnRequestProducerEpoch = "ProducerEpoch"
	// FieldAddOffsetsToTxnRequestGroupId is a field name that can be used to resolve the correct struct field.
	FieldAddOffsetsToTxnRequestGroupId = "GroupId"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AddOffsetsToTxnRequest.json
const originalAddOffsetsToTxnRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 25,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "AddOffsetsToTxnRequest",
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
    { "name": "GroupId", "type": "string", "versions": "0+", "entityType": "groupId",
      "about": "The unique group identifier." }
  ]
}
`

