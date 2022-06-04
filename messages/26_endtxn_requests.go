package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init26EndTxnRequest() []schema.Schema {

	return []schema.Schema{

		// Message: EndTxnRequest, API Key: 26, Version: 0
		schema.NewSchema("EndTxnRequestv0", 
			&schema.Mfield{Name: FieldEndTxnRequestTransactionalId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldEndTxnRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldEndTxnRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldEndTxnRequestCommitted, Ty: schema.TypeBool},
		),

		// Message: EndTxnRequest, API Key: 26, Version: 1
		schema.NewSchema("EndTxnRequestv1", 
			&schema.Mfield{Name: FieldEndTxnRequestTransactionalId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldEndTxnRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldEndTxnRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldEndTxnRequestCommitted, Ty: schema.TypeBool},
		),

		// Message: EndTxnRequest, API Key: 26, Version: 2
		schema.NewSchema("EndTxnRequestv2", 
			&schema.Mfield{Name: FieldEndTxnRequestTransactionalId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldEndTxnRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldEndTxnRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldEndTxnRequestCommitted, Ty: schema.TypeBool},
		),

		// Message: EndTxnRequest, API Key: 26, Version: 3
		schema.NewSchema("EndTxnRequestv3", 
			&schema.Mfield{Name: FieldEndTxnRequestTransactionalId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldEndTxnRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldEndTxnRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldEndTxnRequestCommitted, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldEndTxnRequestTags},
		),

	}
}

const (
	// FieldEndTxnRequestProducerEpoch is a field name that can be used to resolve the correct struct field.
	FieldEndTxnRequestProducerEpoch = "ProducerEpoch"
	// FieldEndTxnRequestCommitted is a field name that can be used to resolve the correct struct field.
	FieldEndTxnRequestCommitted = "Committed"
	// FieldEndTxnRequestTags is a field name that can be used to resolve the correct struct field.
	FieldEndTxnRequestTags = "Tags"
	// FieldEndTxnRequestTransactionalId is a field name that can be used to resolve the correct struct field.
	FieldEndTxnRequestTransactionalId = "TransactionalId"
	// FieldEndTxnRequestProducerId is a field name that can be used to resolve the correct struct field.
	FieldEndTxnRequestProducerId = "ProducerId"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/EndTxnRequest.json
const originalEndTxnRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 26,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "EndTxnRequest",
  // Version 1 is the same as version 0.
  //
  // Version 2 adds the support for new error code PRODUCER_FENCED.
  //
  // Version 3 enables flexible versions.
  "validVersions": "0-3",
  "flexibleVersions": "3+",
  "fields": [
    { "name": "TransactionalId", "type": "string", "versions": "0+", "entityType": "transactionalId",
      "about": "The ID of the transaction to end." },
    { "name": "ProducerId", "type": "int64", "versions": "0+", "entityType": "producerId",
      "about": "The producer ID." },
    { "name": "ProducerEpoch", "type": "int16", "versions": "0+",
      "about": "The current epoch associated with the producer." },
    { "name": "Committed", "type": "bool", "versions": "0+",
      "about": "True if the transaction was committed, false if it was aborted." }
  ]
}
`

