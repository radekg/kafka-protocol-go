package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init36SaslAuthenticateRequest() []schema.Schema {

	return []schema.Schema{
		// Message: SaslAuthenticateRequest, API Key: 36, Version: 0
		schema.NewSchema("SaslAuthenticateRequest:v0",
			&schema.Mfield{Name: FieldSaslAuthenticateRequestAuthBytes, Ty: schema.TypeBytes},
		),

		// Message: SaslAuthenticateRequest, API Key: 36, Version: 1
		schema.NewSchema("SaslAuthenticateRequest:v1",
			&schema.Mfield{Name: FieldSaslAuthenticateRequestAuthBytes, Ty: schema.TypeBytes},
		),

		// Message: SaslAuthenticateRequest, API Key: 36, Version: 2
		schema.NewSchema("SaslAuthenticateRequest:v2",
			&schema.Mfield{Name: FieldSaslAuthenticateRequestAuthBytes, Ty: schema.TypeBytesCompact},
			&schema.SchemaTaggedFields{Name: FieldSaslAuthenticateRequestTags},
		),
	}

}

const (

	// FieldSaslAuthenticateRequestAuthBytes is: The SASL authentication bytes from the client, as defined by the SASL mechanism.
	FieldSaslAuthenticateRequestAuthBytes = "AuthBytes"

	// FieldSaslAuthenticateRequestTags is: The tagged fields.
	FieldSaslAuthenticateRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/SaslAuthenticateRequest.json
const originalSaslAuthenticateRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 36,
  "type": "request",
  "listeners": ["zkBroker", "broker", "controller"],
  "name": "SaslAuthenticateRequest",
  // Version 1 is the same as version 0.
  // Version 2 adds flexible version support
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "AuthBytes", "type": "bytes", "versions": "0+",
      "about": "The SASL authentication bytes from the client, as defined by the SASL mechanism." }
  ]
}
`
