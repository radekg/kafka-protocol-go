package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init0ProduceRequest() []schema.Schema {

	return []schema.Schema{

		// Message: ProduceRequest, API Key: 0, Version: 0
		schema.NewSchema("ProduceRequestv0", 
			&schema.Mfield{Name: FieldProduceRequestAcks, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldProduceRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldProduceRequestTopicData, Ty: schema.NewSchema("TopicDataV0",
				&schema.Mfield{Name: FieldProduceRequestTopicDataName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceRequestTopicDataPartitionData, Ty: schema.NewSchema("PartitionDataV0",
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: ProduceRequest, API Key: 0, Version: 1
		schema.NewSchema("ProduceRequestv1", 
			&schema.Mfield{Name: FieldProduceRequestAcks, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldProduceRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldProduceRequestTopicData, Ty: schema.NewSchema("TopicDataV1",
				&schema.Mfield{Name: FieldProduceRequestTopicDataName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceRequestTopicDataPartitionData, Ty: schema.NewSchema("PartitionDataV1",
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: ProduceRequest, API Key: 0, Version: 2
		schema.NewSchema("ProduceRequestv2", 
			&schema.Mfield{Name: FieldProduceRequestAcks, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldProduceRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldProduceRequestTopicData, Ty: schema.NewSchema("TopicDataV2",
				&schema.Mfield{Name: FieldProduceRequestTopicDataName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceRequestTopicDataPartitionData, Ty: schema.NewSchema("PartitionDataV2",
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: ProduceRequest, API Key: 0, Version: 3
		schema.NewSchema("ProduceRequestv3", 
			&schema.Mfield{Name: FieldProduceRequestTransactionalId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldProduceRequestAcks, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldProduceRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldProduceRequestTopicData, Ty: schema.NewSchema("TopicDataV3",
				&schema.Mfield{Name: FieldProduceRequestTopicDataName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceRequestTopicDataPartitionData, Ty: schema.NewSchema("PartitionDataV3",
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: ProduceRequest, API Key: 0, Version: 4
		schema.NewSchema("ProduceRequestv4", 
			&schema.Mfield{Name: FieldProduceRequestTransactionalId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldProduceRequestAcks, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldProduceRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldProduceRequestTopicData, Ty: schema.NewSchema("TopicDataV4",
				&schema.Mfield{Name: FieldProduceRequestTopicDataName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceRequestTopicDataPartitionData, Ty: schema.NewSchema("PartitionDataV4",
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: ProduceRequest, API Key: 0, Version: 5
		schema.NewSchema("ProduceRequestv5", 
			&schema.Mfield{Name: FieldProduceRequestTransactionalId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldProduceRequestAcks, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldProduceRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldProduceRequestTopicData, Ty: schema.NewSchema("TopicDataV5",
				&schema.Mfield{Name: FieldProduceRequestTopicDataName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceRequestTopicDataPartitionData, Ty: schema.NewSchema("PartitionDataV5",
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: ProduceRequest, API Key: 0, Version: 6
		schema.NewSchema("ProduceRequestv6", 
			&schema.Mfield{Name: FieldProduceRequestTransactionalId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldProduceRequestAcks, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldProduceRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldProduceRequestTopicData, Ty: schema.NewSchema("TopicDataV6",
				&schema.Mfield{Name: FieldProduceRequestTopicDataName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceRequestTopicDataPartitionData, Ty: schema.NewSchema("PartitionDataV6",
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: ProduceRequest, API Key: 0, Version: 7
		schema.NewSchema("ProduceRequestv7", 
			&schema.Mfield{Name: FieldProduceRequestTransactionalId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldProduceRequestAcks, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldProduceRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldProduceRequestTopicData, Ty: schema.NewSchema("TopicDataV7",
				&schema.Mfield{Name: FieldProduceRequestTopicDataName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceRequestTopicDataPartitionData, Ty: schema.NewSchema("PartitionDataV7",
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: ProduceRequest, API Key: 0, Version: 8
		schema.NewSchema("ProduceRequestv8", 
			&schema.Mfield{Name: FieldProduceRequestTransactionalId, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldProduceRequestAcks, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldProduceRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldProduceRequestTopicData, Ty: schema.NewSchema("TopicDataV8",
				&schema.Mfield{Name: FieldProduceRequestTopicDataName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldProduceRequestTopicDataPartitionData, Ty: schema.NewSchema("PartitionDataV8",
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: ProduceRequest, API Key: 0, Version: 9
		schema.NewSchema("ProduceRequestv9", 
			&schema.Mfield{Name: FieldProduceRequestTransactionalId, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldProduceRequestAcks, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldProduceRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldProduceRequestTopicData, Ty: schema.NewSchema("TopicDataV9",
				&schema.Mfield{Name: FieldProduceRequestTopicDataName, Ty: schema.TypeStr},
				&schema.ArrayCompact{Name: FieldProduceRequestTopicDataPartitionData, Ty: schema.NewSchema("PartitionDataV9",
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldProduceRequestTopicDataPartitionDataRecords, Ty: schema.TypeBytesCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldProduceRequestTopicDataPartitionDataTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldProduceRequestTopicDataTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldProduceRequestTags},
		),

	}
}

const (
	// FieldProduceRequestTopicDataName is a field name that can be used to resolve the correct struct field.
	FieldProduceRequestTopicDataName = "Name"
	// FieldProduceRequestTopicDataPartitionData is a field name that can be used to resolve the correct struct field.
	FieldProduceRequestTopicDataPartitionData = "PartitionData"
	// FieldProduceRequestTopicDataPartitionDataRecords is a field name that can be used to resolve the correct struct field.
	FieldProduceRequestTopicDataPartitionDataRecords = "Records"
	// FieldProduceRequestTopicDataPartitionDataTags is a field name that can be used to resolve the correct struct field.
	FieldProduceRequestTopicDataPartitionDataTags = "Tags"
	// FieldProduceRequestTags is a field name that can be used to resolve the correct struct field.
	FieldProduceRequestTags = "Tags"
	// FieldProduceRequestAcks is a field name that can be used to resolve the correct struct field.
	FieldProduceRequestAcks = "Acks"
	// FieldProduceRequestTimeoutMs is a field name that can be used to resolve the correct struct field.
	FieldProduceRequestTimeoutMs = "TimeoutMs"
	// FieldProduceRequestTopicData is a field name that can be used to resolve the correct struct field.
	FieldProduceRequestTopicData = "TopicData"
	// FieldProduceRequestTopicDataPartitionDataIndex is a field name that can be used to resolve the correct struct field.
	FieldProduceRequestTopicDataPartitionDataIndex = "Index"
	// FieldProduceRequestTransactionalId is a field name that can be used to resolve the correct struct field.
	FieldProduceRequestTransactionalId = "TransactionalId"
	// FieldProduceRequestTopicDataTags is a field name that can be used to resolve the correct struct field.
	FieldProduceRequestTopicDataTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ProduceRequest.json
const originalProduceRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 0,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "ProduceRequest",
  // Version 1 and 2 are the same as version 0.
  //
  // Version 3 adds the transactional ID, which is used for authorization when attempting to write
  // transactional data.  Version 3 also adds support for Kafka Message Format v2.
  //
  // Version 4 is the same as version 3, but the requestor must be prepared to handle a
  // KAFKA_STORAGE_ERROR. 
  //
  // Version 5 and 6 are the same as version 3.
  //
  // Starting in version 7, records can be produced using ZStandard compression.  See KIP-110.
  //
  // Starting in Version 8, response has RecordErrors and ErrorMEssage. See KIP-467.
  //
  // Version 9 enables flexible versions.
  "validVersions": "0-9",
  "flexibleVersions": "9+",
  "fields": [
    { "name": "TransactionalId", "type": "string", "versions": "3+", "nullableVersions": "3+", "default": "null", "entityType": "transactionalId",
      "about": "The transactional ID, or null if the producer is not transactional." },
    { "name": "Acks", "type": "int16", "versions": "0+",
      "about": "The number of acknowledgments the producer requires the leader to have received before considering a request complete. Allowed values: 0 for no acknowledgments, 1 for only the leader and -1 for the full ISR." },
    { "name": "TimeoutMs", "type": "int32", "versions": "0+",
      "about": "The timeout to await a response in milliseconds." },
    { "name": "TopicData", "type": "[]TopicProduceData", "versions": "0+",
      "about": "Each topic to produce to.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName", "mapKey": true,
        "about": "The topic name." },
      { "name": "PartitionData", "type": "[]PartitionProduceData", "versions": "0+",
        "about": "Each partition to produce to.", "fields": [
        { "name": "Index", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "Records", "type": "records", "versions": "0+", "nullableVersions": "0+",
          "about": "The record data to be produced." }
      ]}
    ]}
  ]
}
`

