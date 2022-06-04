package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init17SaslHandshakeResponse() []schema.Schema {

	return []schema.Schema{

		// Message: SaslHandshakeResponse, API Key: 17, Version: 0
		schema.NewSchema("SaslHandshakeResponsev0",
			&schema.Mfield{Name: FieldSaslHandshakeResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldSaslHandshakeResponseMechanisms, Ty: schema.TypeStrArray},
		),

		// Message: SaslHandshakeResponse, API Key: 17, Version: 1
		schema.NewSchema("SaslHandshakeResponsev1",
			&schema.Mfield{Name: FieldSaslHandshakeResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldSaslHandshakeResponseMechanisms, Ty: schema.TypeStrArray},
		),
	}
}

const (
	// FieldSaslHandshakeResponseMechanisms is: The mechanisms enabled in the server.
	FieldSaslHandshakeResponseMechanisms = "Mechanisms"
	// FieldSaslHandshakeResponseErrorCode is: The error code, or 0 if there was no error.
	FieldSaslHandshakeResponseErrorCode = "ErrorCode"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/SaslHandshakeResponse.json
const originalSaslHandshakeResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 17,
  "type": "response",
  "name": "SaslHandshakeResponse",
  // Version 1 is the same as version 0.
  // NOTE: Version cannot be easily bumped due to incorrect
  // client negotiation for clients <= 2.4.
  // See https://issues.apache.org/jira/browse/KAFKA-9577
  "validVersions": "0-1",
  "flexibleVersions": "none",
  "fields": [
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The error code, or 0 if there was no error." },
    { "name": "Mechanisms", "type": "[]string", "versions": "0+",
      "about": "The mechanisms enabled in the server." }
  ]
}
`
