package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init17SaslHandshakeRequest() []schema.Schema {

	return []schema.Schema{

		// Message: SaslHandshakeRequest, API Key: 17, Version: 0
		schema.NewSchema("SaslHandshakeRequestv0",
			&schema.Mfield{Name: FieldSaslHandshakeRequestMechanism, Ty: schema.TypeStr},
		),

		// Message: SaslHandshakeRequest, API Key: 17, Version: 1
		schema.NewSchema("SaslHandshakeRequestv1",
			&schema.Mfield{Name: FieldSaslHandshakeRequestMechanism, Ty: schema.TypeStr},
		),
	}
}

const (
	// FieldSaslHandshakeRequestMechanism is a field name that can be used to resolve the correct struct field.
	FieldSaslHandshakeRequestMechanism = "Mechanism"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/SaslHandshakeRequest.json
const originalSaslHandshakeRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "request",
  "listeners": ["zkBroker", "broker", "controller"],
  "name": "SaslHandshakeRequest",
  // Version 1 supports SASL_AUTHENTICATE.
  // NOTE: Version cannot be easily bumped due to incorrect
  // client negotiation for clients <= 2.4.
  // See https://issues.apache.org/jira/browse/KAFKA-9577
  "validVersions": "0-1",
  "flexibleVersions": "none",
  "fields": [
    { "name": "Mechanism", "type": "string", "versions": "0+",
      "about": "The SASL mechanism chosen by the client." }
  ]
}
`
