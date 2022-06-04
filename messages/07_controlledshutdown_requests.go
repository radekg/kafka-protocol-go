package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init7ControlledShutdownRequest() []schema.Schema {

	return []schema.Schema{

		// Message: ControlledShutdownRequest, API Key: 7, Version: 0
		schema.NewSchema("ControlledShutdownRequestv0",
			&schema.Mfield{Name: FieldControlledShutdownRequestBrokerId, Ty: schema.TypeInt32},
		),

		// Message: ControlledShutdownRequest, API Key: 7, Version: 1
		schema.NewSchema("ControlledShutdownRequestv1",
			&schema.Mfield{Name: FieldControlledShutdownRequestBrokerId, Ty: schema.TypeInt32},
		),

		// Message: ControlledShutdownRequest, API Key: 7, Version: 2
		schema.NewSchema("ControlledShutdownRequestv2",
			&schema.Mfield{Name: FieldControlledShutdownRequestBrokerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldControlledShutdownRequestBrokerEpoch, Ty: schema.TypeInt64},
		),

		// Message: ControlledShutdownRequest, API Key: 7, Version: 3
		schema.NewSchema("ControlledShutdownRequestv3",
			&schema.Mfield{Name: FieldControlledShutdownRequestBrokerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldControlledShutdownRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.SchemaTaggedFields{Name: FieldControlledShutdownRequestTags},
		),
	}
}

const (
	// FieldControlledShutdownRequestBrokerEpoch is: The broker epoch.
	FieldControlledShutdownRequestBrokerEpoch = "BrokerEpoch"
	// FieldControlledShutdownRequestBrokerId is: The id of the broker for which controlled shutdown has been requested.
	FieldControlledShutdownRequestBrokerId = "BrokerId"
	// FieldControlledShutdownRequestTags is: The tagged fields.
	FieldControlledShutdownRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ControlledShutdownRequest.json
const originalControlledShutdownRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 7,
  "type": "request",
  "listeners": ["zkBroker", "controller"],
  "name": "ControlledShutdownRequest",
  // Version 0 of ControlledShutdownRequest has a non-standard request header
  // which does not include clientId.  Version 1 and later use the standard
  // request header.
  //
  // Version 1 is the same as version 0.
  //
  // Version 2 adds BrokerEpoch.
  "validVersions": "0-3",
  "flexibleVersions": "3+",
  "fields": [
    { "name": "BrokerId", "type": "int32", "versions": "0+", "entityType": "brokerId",
      "about": "The id of the broker for which controlled shutdown has been requested." },
    { "name": "BrokerEpoch", "type": "int64", "versions": "2+", "default": "-1", "ignorable": true,
      "about": "The broker epoch." }
  ]
}
`
