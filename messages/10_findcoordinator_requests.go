package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init10FindCoordinatorRequest() []schema.Schema {

	return []schema.Schema{

		// Message: FindCoordinatorRequest, API Key: 10, Version: 0
		schema.NewSchema("FindCoordinatorRequestv0", 
			&schema.Mfield{Name: FieldFindCoordinatorRequestKey, Ty: schema.TypeStr},
		),

		// Message: FindCoordinatorRequest, API Key: 10, Version: 1
		schema.NewSchema("FindCoordinatorRequestv1", 
			&schema.Mfield{Name: FieldFindCoordinatorRequestKey, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldFindCoordinatorRequestKeyType, Ty: schema.TypeInt8},
		),

		// Message: FindCoordinatorRequest, API Key: 10, Version: 2
		schema.NewSchema("FindCoordinatorRequestv2", 
			&schema.Mfield{Name: FieldFindCoordinatorRequestKey, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldFindCoordinatorRequestKeyType, Ty: schema.TypeInt8},
		),

		// Message: FindCoordinatorRequest, API Key: 10, Version: 3
		schema.NewSchema("FindCoordinatorRequestv3", 
			&schema.Mfield{Name: FieldFindCoordinatorRequestKey, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldFindCoordinatorRequestKeyType, Ty: schema.TypeInt8},
			&schema.SchemaTaggedFields{Name: FieldFindCoordinatorRequestTags},
		),

		// Message: FindCoordinatorRequest, API Key: 10, Version: 4
		schema.NewSchema("FindCoordinatorRequestv4", 
			&schema.Mfield{Name: FieldFindCoordinatorRequestKeyType, Ty: schema.TypeInt8},
			&schema.Mfield{Name: FieldFindCoordinatorRequestCoordinatorKeys, Ty: schema.TypeStrCompactArray},
			&schema.SchemaTaggedFields{Name: FieldFindCoordinatorRequestTags},
		),

	}
}

const (
	// FieldFindCoordinatorRequestKey is a field name that can be used to resolve the correct struct field.
	FieldFindCoordinatorRequestKey = "Key"
	// FieldFindCoordinatorRequestKeyType is a field name that can be used to resolve the correct struct field.
	FieldFindCoordinatorRequestKeyType = "KeyType"
	// FieldFindCoordinatorRequestTags is a field name that can be used to resolve the correct struct field.
	FieldFindCoordinatorRequestTags = "Tags"
	// FieldFindCoordinatorRequestCoordinatorKeys is a field name that can be used to resolve the correct struct field.
	FieldFindCoordinatorRequestCoordinatorKeys = "CoordinatorKeys"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/FindCoordinatorRequest.json
const originalFindCoordinatorRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 10,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "FindCoordinatorRequest",
  // Version 1 adds KeyType.
  //
  // Version 2 is the same as version 1.
  //
  // Version 3 is the first flexible version.
  //
  // Version 4 adds support for batching via CoordinatorKeys (KIP-699)
  "validVersions": "0-4",
  "flexibleVersions": "3+",
  "fields": [
    { "name": "Key", "type": "string", "versions": "0-3",
      "about": "The coordinator key." },
    { "name": "KeyType", "type": "int8", "versions": "1+", "default": "0", "ignorable": false,
      "about": "The coordinator key type. (Group, transaction, etc.)" },
    { "name": "CoordinatorKeys", "type": "[]string", "versions": "4+",
      "about": "The coordinator keys." }
  ]
}
`

