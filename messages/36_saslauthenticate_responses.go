package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init36SaslAuthenticateResponse() []schema.Schema {

	return []schema.Schema{

		// Message: SaslAuthenticateResponse, API Key: 36, Version: 0
		schema.NewSchema("SaslAuthenticateResponsev0",
			&schema.Mfield{Name: FieldSaslAuthenticateResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldSaslAuthenticateResponseErrorMessage, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldSaslAuthenticateResponseAuthBytes, Ty: schema.TypeBytes},
		),

		// Message: SaslAuthenticateResponse, API Key: 36, Version: 1
		schema.NewSchema("SaslAuthenticateResponsev1",
			&schema.Mfield{Name: FieldSaslAuthenticateResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldSaslAuthenticateResponseErrorMessage, Ty: schema.TypeStrNullable},
			&schema.Mfield{Name: FieldSaslAuthenticateResponseAuthBytes, Ty: schema.TypeBytes},
			&schema.Mfield{Name: FieldSaslAuthenticateResponseSessionLifetimeMs, Ty: schema.TypeInt64},
		),

		// Message: SaslAuthenticateResponse, API Key: 36, Version: 2
		schema.NewSchema("SaslAuthenticateResponsev2",
			&schema.Mfield{Name: FieldSaslAuthenticateResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldSaslAuthenticateResponseErrorMessage, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldSaslAuthenticateResponseAuthBytes, Ty: schema.TypeBytesCompact},
			&schema.Mfield{Name: FieldSaslAuthenticateResponseSessionLifetimeMs, Ty: schema.TypeInt64},
			&schema.SchemaTaggedFields{Name: FieldSaslAuthenticateResponseTags},
		),
	}
}

const (
	// FieldSaslAuthenticateResponseErrorCode is: The error code, or 0 if there was no error.
	FieldSaslAuthenticateResponseErrorCode = "ErrorCode"
	// FieldSaslAuthenticateResponseErrorMessage is: The error message, or null if there was no error.
	FieldSaslAuthenticateResponseErrorMessage = "ErrorMessage"
	// FieldSaslAuthenticateResponseAuthBytes is: The SASL authentication bytes from the server, as defined by the SASL mechanism.
	FieldSaslAuthenticateResponseAuthBytes = "AuthBytes"
	// FieldSaslAuthenticateResponseSessionLifetimeMs is: The SASL authentication bytes from the server, as defined by the SASL mechanism.
	FieldSaslAuthenticateResponseSessionLifetimeMs = "SessionLifetimeMs"
	// FieldSaslAuthenticateResponseTags is: The tagged fields.
	FieldSaslAuthenticateResponseTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/SaslAuthenticateResponse.json
const originalSaslAuthenticateResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "response",
  "name": "SaslAuthenticateResponse",
  // Version 1 adds the session lifetime.
  // Version 2 adds flexible version support
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The error code, or 0 if there was no error." },
    { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
      "about": "The error message, or null if there was no error." },
    { "name": "AuthBytes", "type": "bytes", "versions": "0+",
      "about": "The SASL authentication bytes from the server, as defined by the SASL mechanism." },
    { "name": "SessionLifetimeMs", "type": "int64", "versions": "1+", "default": "0", "ignorable": true,
      "about": "The SASL authentication bytes from the server, as defined by the SASL mechanism." }
  ]
}
`
