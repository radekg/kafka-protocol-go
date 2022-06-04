package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init22InitProducerIdRequest() []schema.Schema {

	return []schema.Schema{

		// Message: InitProducerIdRequest, API Key: 22, Version: 0
		schema.NewSchema("InitProducerIdRequestv0", 
			&schema.Mfield{Name: FieldInitProducerIdRequestTransactionalId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldInitProducerIdRequestTransactionTimeoutMs, Ty: schema.TypeInt32},
		),

		// Message: InitProducerIdRequest, API Key: 22, Version: 1
		schema.NewSchema("InitProducerIdRequestv1", 
			&schema.Mfield{Name: FieldInitProducerIdRequestTransactionalId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldInitProducerIdRequestTransactionTimeoutMs, Ty: schema.TypeInt32},
		),

		// Message: InitProducerIdRequest, API Key: 22, Version: 2
		schema.NewSchema("InitProducerIdRequestv2", 
			&schema.Mfield{Name: FieldInitProducerIdRequestTransactionalId, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldInitProducerIdRequestTransactionTimeoutMs, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldInitProducerIdRequestTags},
		),

		// Message: InitProducerIdRequest, API Key: 22, Version: 3
		schema.NewSchema("InitProducerIdRequestv3", 
			&schema.Mfield{Name: FieldInitProducerIdRequestTransactionalId, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldInitProducerIdRequestTransactionTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldInitProducerIdRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldInitProducerIdRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.SchemaTaggedFields{Name: FieldInitProducerIdRequestTags},
		),

		// Message: InitProducerIdRequest, API Key: 22, Version: 4
		schema.NewSchema("InitProducerIdRequestv4", 
			&schema.Mfield{Name: FieldInitProducerIdRequestTransactionalId, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldInitProducerIdRequestTransactionTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldInitProducerIdRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldInitProducerIdRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.SchemaTaggedFields{Name: FieldInitProducerIdRequestTags},
		),

	}
}

const (
	// FieldInitProducerIdRequestTransactionTimeoutMs is a field name that can be used to resolve the correct struct field.
	FieldInitProducerIdRequestTransactionTimeoutMs = "TransactionTimeoutMs"
	// FieldInitProducerIdRequestTags is a field name that can be used to resolve the correct struct field.
	FieldInitProducerIdRequestTags = "Tags"
	// FieldInitProducerIdRequestProducerId is a field name that can be used to resolve the correct struct field.
	FieldInitProducerIdRequestProducerId = "ProducerId"
	// FieldInitProducerIdRequestProducerEpoch is a field name that can be used to resolve the correct struct field.
	FieldInitProducerIdRequestProducerEpoch = "ProducerEpoch"
	// FieldInitProducerIdRequestTransactionalId is a field name that can be used to resolve the correct struct field.
	FieldInitProducerIdRequestTransactionalId = "TransactionalId"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/InitProducerIdRequest.json
const originalInitProducerIdRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 22,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "InitProducerIdRequest",
  // Version 1 is the same as version 0.
  //
  // Version 2 is the first flexible version.
  //
  // Version 3 adds ProducerId and ProducerEpoch, allowing producers to try to resume after an INVALID_PRODUCER_EPOCH error
  //
  // Version 4 adds the support for new error code PRODUCER_FENCED.
  "validVersions": "0-4",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "TransactionalId", "type": "string", "versions": "0+", "nullableVersions": "0+", "entityType": "transactionalId",
      "about": "The transactional id, or null if the producer is not transactional." },
    { "name": "TransactionTimeoutMs", "type": "int32", "versions": "0+",
      "about": "The time in ms to wait before aborting idle transactions sent by this producer. This is only relevant if a TransactionalId has been defined." },
    { "name": "ProducerId", "type": "int64", "versions": "3+", "default": "-1", "entityType": "producerId",
      "about": "The producer id. This is used to disambiguate requests if a transactional id is reused following its expiration." },
    { "name": "ProducerEpoch", "type": "int16", "versions": "3+", "default": "-1",
      "about": "The producer's current epoch. This will be checked against the producer epoch on the broker, and the request will return an error if they do not match." }
  ]
}
`

