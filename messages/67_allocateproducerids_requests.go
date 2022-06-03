package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init67AllocateProducerIdsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: AllocateProducerIdsRequest, API Key: 67, Version: 0
		schema.NewSchema("AllocateProducerIdsRequestv0", 
			&schema.Mfield{Name: FieldAllocateProducerIdsRequestBrokerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldAllocateProducerIdsRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.SchemaTaggedFields{Name: FieldAllocateProducerIdsRequestTags},
		),

	}
}

const (
	// FieldAllocateProducerIdsRequestBrokerEpoch is a field name that can be used to resolve the correct struct field.
	FieldAllocateProducerIdsRequestBrokerEpoch = "BrokerEpoch"
	// FieldAllocateProducerIdsRequestTags is a field name that can be used to resolve the correct struct field.
	FieldAllocateProducerIdsRequestTags = "Tags"
	// FieldAllocateProducerIdsRequestBrokerId is a field name that can be used to resolve the correct struct field.
	FieldAllocateProducerIdsRequestBrokerId = "BrokerId"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AllocateProducerIdsRequest.json
const originalAllocateProducerIdsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implie
// See the License for the specific language governing permissions and
// limitations under the License.

{
  "apiKey": 67,
  "type": "request",
  "listeners": ["zkBroker", "controller"],
  "name": "AllocateProducerIdsRequest",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "BrokerId", "type": "int32", "versions": "0+", "entityType": "brokerId",
      "about": "The ID of the requesting broker" },
    { "name": "BrokerEpoch", "type": "int64", "versions": "0+", "default": "-1",
      "about": "The epoch of the requesting broker" }
  ]
}
`

