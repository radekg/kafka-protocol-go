package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init41DescribeDelegationTokenRequest() []schema.Schema {

	return []schema.Schema{

		// Message: DescribeDelegationTokenRequest, API Key: 41, Version: 0
		schema.NewSchema("DescribeDelegationTokenRequestv0",
			&schema.Array{Name: FieldDescribeDelegationTokenRequestOwners, Ty: schema.NewSchema("OwnersV0",
				&schema.Mfield{Name: FieldDescribeDelegationTokenRequestOwnersPrincipalType, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeDelegationTokenRequestOwnersPrincipalName, Ty: schema.TypeStr},
			)},
		),

		// Message: DescribeDelegationTokenRequest, API Key: 41, Version: 1
		schema.NewSchema("DescribeDelegationTokenRequestv1",
			&schema.Array{Name: FieldDescribeDelegationTokenRequestOwners, Ty: schema.NewSchema("OwnersV1",
				&schema.Mfield{Name: FieldDescribeDelegationTokenRequestOwnersPrincipalType, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeDelegationTokenRequestOwnersPrincipalName, Ty: schema.TypeStr},
			)},
		),

		// Message: DescribeDelegationTokenRequest, API Key: 41, Version: 2
		schema.NewSchema("DescribeDelegationTokenRequestv2",
			&schema.ArrayCompact{Name: FieldDescribeDelegationTokenRequestOwners, Ty: schema.NewSchema("OwnersV2",
				&schema.Mfield{Name: FieldDescribeDelegationTokenRequestOwnersPrincipalType, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeDelegationTokenRequestOwnersPrincipalName, Ty: schema.TypeStrCompact},
				&schema.SchemaTaggedFields{Name: FieldDescribeDelegationTokenRequestOwnersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDescribeDelegationTokenRequestTags},
		),
	}
}

const (
	// FieldDescribeDelegationTokenRequestOwners is: Each owner that we want to describe delegation tokens for, or null to describe all tokens.
	FieldDescribeDelegationTokenRequestOwners = "Owners"
	// FieldDescribeDelegationTokenRequestOwnersPrincipalType is: The owner principal type.
	FieldDescribeDelegationTokenRequestOwnersPrincipalType = "PrincipalType"
	// FieldDescribeDelegationTokenRequestOwnersPrincipalName is: The owner principal name.
	FieldDescribeDelegationTokenRequestOwnersPrincipalName = "PrincipalName"
	// FieldDescribeDelegationTokenRequestOwnersTags is: The tagged fields.
	FieldDescribeDelegationTokenRequestOwnersTags = "Tags"
	// FieldDescribeDelegationTokenRequestTags is: The tagged fields.
	FieldDescribeDelegationTokenRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeDelegationTokenRequest.json
const originalDescribeDelegationTokenRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 41,
  "type": "request",
  "listeners": ["zkBroker"],
  "name": "DescribeDelegationTokenRequest",
  // Version 1 is the same as version 0.
  // Version 2 adds flexible version support
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "Owners", "type": "[]DescribeDelegationTokenOwner", "versions": "0+", "nullableVersions": "0+",
      "about": "Each owner that we want to describe delegation tokens for, or null to describe all tokens.", "fields": [
      { "name": "PrincipalType", "type": "string", "versions": "0+",
        "about": "The owner principal type." },
      { "name": "PrincipalName", "type": "string", "versions": "0+",
        "about": "The owner principal name." }
    ]}
  ]
}
`
