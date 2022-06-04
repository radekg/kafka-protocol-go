package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init51AlterUserScramCredentialsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: AlterUserScramCredentialsRequest, API Key: 51, Version: 0
		schema.NewSchema("AlterUserScramCredentialsRequestv0", 
			&schema.ArrayCompact{Name: FieldAlterUserScramCredentialsRequestDeletions, Ty: schema.NewSchema("DeletionsV0",
				&schema.Mfield{Name: FieldAlterUserScramCredentialsRequestDeletionsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldAlterUserScramCredentialsRequestDeletionsMechanism, Ty: schema.TypeInt8},
				&schema.SchemaTaggedFields{Name: FieldAlterUserScramCredentialsRequestDeletionsTags},
			)},
			&schema.ArrayCompact{Name: FieldAlterUserScramCredentialsRequestUpsertions, Ty: schema.NewSchema("UpsertionsV0",
				&schema.Mfield{Name: FieldAlterUserScramCredentialsRequestUpsertionsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldAlterUserScramCredentialsRequestUpsertionsMechanism, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldAlterUserScramCredentialsRequestUpsertionsIterations, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldAlterUserScramCredentialsRequestUpsertionsSalt, Ty: schema.TypeBytesCompact},
				&schema.Mfield{Name: FieldAlterUserScramCredentialsRequestUpsertionsSaltedPassword, Ty: schema.TypeBytesCompact},
				&schema.SchemaTaggedFields{Name: FieldAlterUserScramCredentialsRequestUpsertionsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldAlterUserScramCredentialsRequestTags},
		),

	}
}

const (
	// FieldAlterUserScramCredentialsRequestUpsertionsName is a field name that can be used to resolve the correct struct field.
	FieldAlterUserScramCredentialsRequestUpsertionsName = "Name"
	// FieldAlterUserScramCredentialsRequestUpsertionsIterations is a field name that can be used to resolve the correct struct field.
	FieldAlterUserScramCredentialsRequestUpsertionsIterations = "Iterations"
	// FieldAlterUserScramCredentialsRequestUpsertionsSalt is a field name that can be used to resolve the correct struct field.
	FieldAlterUserScramCredentialsRequestUpsertionsSalt = "Salt"
	// FieldAlterUserScramCredentialsRequestUpsertionsSaltedPassword is a field name that can be used to resolve the correct struct field.
	FieldAlterUserScramCredentialsRequestUpsertionsSaltedPassword = "SaltedPassword"
	// FieldAlterUserScramCredentialsRequestUpsertionsTags is a field name that can be used to resolve the correct struct field.
	FieldAlterUserScramCredentialsRequestUpsertionsTags = "Tags"
	// FieldAlterUserScramCredentialsRequestTags is a field name that can be used to resolve the correct struct field.
	FieldAlterUserScramCredentialsRequestTags = "Tags"
	// FieldAlterUserScramCredentialsRequestDeletionsName is a field name that can be used to resolve the correct struct field.
	FieldAlterUserScramCredentialsRequestDeletionsName = "Name"
	// FieldAlterUserScramCredentialsRequestDeletionsTags is a field name that can be used to resolve the correct struct field.
	FieldAlterUserScramCredentialsRequestDeletionsTags = "Tags"
	// FieldAlterUserScramCredentialsRequestUpsertions is a field name that can be used to resolve the correct struct field.
	FieldAlterUserScramCredentialsRequestUpsertions = "Upsertions"
	// FieldAlterUserScramCredentialsRequestUpsertionsMechanism is a field name that can be used to resolve the correct struct field.
	FieldAlterUserScramCredentialsRequestUpsertionsMechanism = "Mechanism"
	// FieldAlterUserScramCredentialsRequestDeletions is a field name that can be used to resolve the correct struct field.
	FieldAlterUserScramCredentialsRequestDeletions = "Deletions"
	// FieldAlterUserScramCredentialsRequestDeletionsMechanism is a field name that can be used to resolve the correct struct field.
	FieldAlterUserScramCredentialsRequestDeletionsMechanism = "Mechanism"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AlterUserScramCredentialsRequest.json
const originalAlterUserScramCredentialsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 51,
  "type": "request",
  "listeners": ["zkBroker"],
  "name": "AlterUserScramCredentialsRequest",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "Deletions", "type": "[]ScramCredentialDeletion", "versions": "0+",
      "about": "The SCRAM credentials to remove.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+",
        "about": "The user name." },
      { "name": "Mechanism", "type": "int8", "versions": "0+",
        "about": "The SCRAM mechanism." }
    ]},
    { "name": "Upsertions", "type": "[]ScramCredentialUpsertion", "versions": "0+",
      "about": "The SCRAM credentials to update/insert.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+",
        "about": "The user name." },
      { "name": "Mechanism", "type": "int8", "versions": "0+",
        "about": "The SCRAM mechanism." },
      { "name": "Iterations", "type": "int32", "versions": "0+",
        "about": "The number of iterations." },
      { "name": "Salt", "type": "bytes", "versions": "0+",
        "about": "A random salt generated by the client." },
      { "name": "SaltedPassword", "type": "bytes", "versions": "0+",
        "about": "The salted password." }
    ]}
  ]
}
`

