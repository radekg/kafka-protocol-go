package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init18ApiVersionsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: ApiVersionsRequest, API Key: 18, Version: 0
		schema.NewSchema("ApiVersionsRequestv0"),
		// Message: ApiVersionsRequest, API Key: 18, Version: 1
		schema.NewSchema("ApiVersionsRequestv1"),
		// Message: ApiVersionsRequest, API Key: 18, Version: 2
		schema.NewSchema("ApiVersionsRequestv2"),
		// Message: ApiVersionsRequest, API Key: 18, Version: 3
		schema.NewSchema("ApiVersionsRequestv3",
			&schema.Mfield{Name: FieldApiVersionsRequestClientSoftwareName, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldApiVersionsRequestClientSoftwareVersion, Ty: schema.TypeStrCompact},
			&schema.SchemaTaggedFields{Name: FieldApiVersionsRequestTags},
		),
	}
}

const (
	// FieldApiVersionsRequestClientSoftwareName is: The name of the client.
	FieldApiVersionsRequestClientSoftwareName = "ClientSoftwareName"
	// FieldApiVersionsRequestClientSoftwareVersion is: The version of the client.
	FieldApiVersionsRequestClientSoftwareVersion = "ClientSoftwareVersion"
	// FieldApiVersionsRequestTags is: The tagged fields.
	FieldApiVersionsRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ApiVersionsRequest.json
const originalApiVersionsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 18,
  "type": "request",
  "listeners": ["zkBroker", "broker", "controller"],
  "name": "ApiVersionsRequest",
  // Versions 0 through 2 of ApiVersionsRequest are the same.
  //
  // Version 3 is the first flexible version and adds ClientSoftwareName and ClientSoftwareVersion.
  "validVersions": "0-3",
  "flexibleVersions": "3+",
  "fields": [
    { "name": "ClientSoftwareName", "type": "string", "versions": "3+",
      "ignorable": true, "about": "The name of the client." },
    { "name": "ClientSoftwareVersion", "type": "string", "versions": "3+",
      "ignorable": true, "about": "The version of the client." }
  ]
}
`
