package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init6UpdateMetadataResponse() []schema.Schema {

	return []schema.Schema{

		// Message: UpdateMetadataResponse, API Key: 6, Version: 0
		schema.NewSchema("UpdateMetadataResponsev0",
			&schema.Mfield{Name: FieldUpdateMetadataResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: UpdateMetadataResponse, API Key: 6, Version: 1
		schema.NewSchema("UpdateMetadataResponsev1",
			&schema.Mfield{Name: FieldUpdateMetadataResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: UpdateMetadataResponse, API Key: 6, Version: 2
		schema.NewSchema("UpdateMetadataResponsev2",
			&schema.Mfield{Name: FieldUpdateMetadataResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: UpdateMetadataResponse, API Key: 6, Version: 3
		schema.NewSchema("UpdateMetadataResponsev3",
			&schema.Mfield{Name: FieldUpdateMetadataResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: UpdateMetadataResponse, API Key: 6, Version: 4
		schema.NewSchema("UpdateMetadataResponsev4",
			&schema.Mfield{Name: FieldUpdateMetadataResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: UpdateMetadataResponse, API Key: 6, Version: 5
		schema.NewSchema("UpdateMetadataResponsev5",
			&schema.Mfield{Name: FieldUpdateMetadataResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: UpdateMetadataResponse, API Key: 6, Version: 6
		schema.NewSchema("UpdateMetadataResponsev6",
			&schema.Mfield{Name: FieldUpdateMetadataResponseErrorCode, Ty: schema.TypeInt16},
			&schema.SchemaTaggedFields{Name: FieldUpdateMetadataResponseTags},
		),

		// Message: UpdateMetadataResponse, API Key: 6, Version: 7
		schema.NewSchema("UpdateMetadataResponsev7",
			&schema.Mfield{Name: FieldUpdateMetadataResponseErrorCode, Ty: schema.TypeInt16},
			&schema.SchemaTaggedFields{Name: FieldUpdateMetadataResponseTags},
		),
	}
}

const (
	// FieldUpdateMetadataResponseErrorCode is: The error code, or 0 if there was no error.
	FieldUpdateMetadataResponseErrorCode = "ErrorCode"
	// FieldUpdateMetadataResponseTags is: The tagged fields.
	FieldUpdateMetadataResponseTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/UpdateMetadataResponse.json
const originalUpdateMetadataResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 6,
  "type": "response",
  "name": "UpdateMetadataResponse",
  // Versions 1, 2, 3, 4, and 5 are the same as version 0
  "validVersions": "0-7",
  "flexibleVersions": "6+",
  "fields": [
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The error code, or 0 if there was no error." }
  ]
}
`
