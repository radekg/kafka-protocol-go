package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init39RenewDelegationTokenRequest() []schema.Schema {

	return []schema.Schema{

		// Message: RenewDelegationTokenRequest, API Key: 39, Version: 0
		schema.NewSchema("RenewDelegationTokenRequestv0",
			&schema.Mfield{Name: FieldRenewDelegationTokenRequestHmac, Ty: schema.TypeBytes},
			&schema.Mfield{Name: FieldRenewDelegationTokenRequestRenewPeriodMs, Ty: schema.TypeInt64},
		),

		// Message: RenewDelegationTokenRequest, API Key: 39, Version: 1
		schema.NewSchema("RenewDelegationTokenRequestv1",
			&schema.Mfield{Name: FieldRenewDelegationTokenRequestHmac, Ty: schema.TypeBytes},
			&schema.Mfield{Name: FieldRenewDelegationTokenRequestRenewPeriodMs, Ty: schema.TypeInt64},
		),

		// Message: RenewDelegationTokenRequest, API Key: 39, Version: 2
		schema.NewSchema("RenewDelegationTokenRequestv2",
			&schema.Mfield{Name: FieldRenewDelegationTokenRequestHmac, Ty: schema.TypeBytesCompact},
			&schema.Mfield{Name: FieldRenewDelegationTokenRequestRenewPeriodMs, Ty: schema.TypeInt64},
			&schema.SchemaTaggedFields{Name: FieldRenewDelegationTokenRequestTags},
		),
	}
}

const (
	// FieldRenewDelegationTokenRequestHmac is: The HMAC of the delegation token to be renewed.
	FieldRenewDelegationTokenRequestHmac = "Hmac"
	// FieldRenewDelegationTokenRequestRenewPeriodMs is: The renewal time period in milliseconds.
	FieldRenewDelegationTokenRequestRenewPeriodMs = "RenewPeriodMs"
	// FieldRenewDelegationTokenRequestTags is: The tagged fields.
	FieldRenewDelegationTokenRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/RenewDelegationTokenRequest.json
const originalRenewDelegationTokenRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 39,
  "type": "request",
  "listeners": ["zkBroker"],
  "name": "RenewDelegationTokenRequest",
  // Version 1 is the same as version 0.
  // Version 2 adds flexible version support
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "Hmac", "type": "bytes", "versions": "0+",
      "about": "The HMAC of the delegation token to be renewed." },
    { "name": "RenewPeriodMs", "type": "int64", "versions": "0+",
      "about": "The renewal time period in milliseconds." }
  ]
}
`
