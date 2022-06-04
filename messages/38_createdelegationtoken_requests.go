package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init38CreateDelegationTokenRequest() []schema.Schema {

	return []schema.Schema{

		// Message: CreateDelegationTokenRequest, API Key: 38, Version: 0
		schema.NewSchema("CreateDelegationTokenRequestv0",
			&schema.Array{Name: FieldCreateDelegationTokenRequestRenewers, Ty: schema.NewSchema("RenewersV0",
				&schema.Mfield{Name: FieldCreateDelegationTokenRequestRenewersPrincipalType, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateDelegationTokenRequestRenewersPrincipalName, Ty: schema.TypeStr},
			)},
			&schema.Mfield{Name: FieldCreateDelegationTokenRequestMaxLifetimeMs, Ty: schema.TypeInt64},
		),

		// Message: CreateDelegationTokenRequest, API Key: 38, Version: 1
		schema.NewSchema("CreateDelegationTokenRequestv1",
			&schema.Array{Name: FieldCreateDelegationTokenRequestRenewers, Ty: schema.NewSchema("RenewersV1",
				&schema.Mfield{Name: FieldCreateDelegationTokenRequestRenewersPrincipalType, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateDelegationTokenRequestRenewersPrincipalName, Ty: schema.TypeStr},
			)},
			&schema.Mfield{Name: FieldCreateDelegationTokenRequestMaxLifetimeMs, Ty: schema.TypeInt64},
		),

		// Message: CreateDelegationTokenRequest, API Key: 38, Version: 2
		schema.NewSchema("CreateDelegationTokenRequestv2",
			&schema.ArrayCompact{Name: FieldCreateDelegationTokenRequestRenewers, Ty: schema.NewSchema("RenewersV2",
				&schema.Mfield{Name: FieldCreateDelegationTokenRequestRenewersPrincipalType, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldCreateDelegationTokenRequestRenewersPrincipalName, Ty: schema.TypeStrCompact},
				&schema.SchemaTaggedFields{Name: FieldCreateDelegationTokenRequestRenewersTags},
			)},
			&schema.Mfield{Name: FieldCreateDelegationTokenRequestMaxLifetimeMs, Ty: schema.TypeInt64},
			&schema.SchemaTaggedFields{Name: FieldCreateDelegationTokenRequestTags},
		),
	}
}

const (
	// FieldCreateDelegationTokenRequestTags is a field name that can be used to resolve the correct struct field.
	FieldCreateDelegationTokenRequestTags = "Tags"
	// FieldCreateDelegationTokenRequestRenewers is a field name that can be used to resolve the correct struct field.
	FieldCreateDelegationTokenRequestRenewers = "Renewers"
	// FieldCreateDelegationTokenRequestRenewersPrincipalType is a field name that can be used to resolve the correct struct field.
	FieldCreateDelegationTokenRequestRenewersPrincipalType = "PrincipalType"
	// FieldCreateDelegationTokenRequestRenewersPrincipalName is a field name that can be used to resolve the correct struct field.
	FieldCreateDelegationTokenRequestRenewersPrincipalName = "PrincipalName"
	// FieldCreateDelegationTokenRequestMaxLifetimeMs is a field name that can be used to resolve the correct struct field.
	FieldCreateDelegationTokenRequestMaxLifetimeMs = "MaxLifetimeMs"
	// FieldCreateDelegationTokenRequestRenewersTags is a field name that can be used to resolve the correct struct field.
	FieldCreateDelegationTokenRequestRenewersTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/CreateDelegationTokenRequest.json
const originalCreateDelegationTokenRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 38,
  "type": "request",
  "listeners": ["zkBroker"],
  "name": "CreateDelegationTokenRequest",
  // Version 1 is the same as version 0.
  //
  // Version 2 is the first flexible version.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "Renewers", "type": "[]CreatableRenewers", "versions": "0+",
      "about": "A list of those who are allowed to renew this token before it expires.", "fields": [
      { "name": "PrincipalType", "type": "string", "versions": "0+",
        "about": "The type of the Kafka principal." },
      { "name": "PrincipalName", "type": "string", "versions": "0+",
        "about": "The name of the Kafka principal." }
    ]},
    { "name": "MaxLifetimeMs", "type": "int64", "versions": "0+",
      "about": "The maximum lifetime of the token in milliseconds, or -1 to use the server side default." }
  ]
}
`
