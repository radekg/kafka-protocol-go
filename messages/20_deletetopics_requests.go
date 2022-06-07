package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init20DeleteTopicsRequest() []schema.Schema {

	return []schema.Schema{
		// Message: DeleteTopicsRequest, API Key: 20, Version: 0
		schema.NewSchema("DeleteTopicsRequest:v0",
			&schema.Mfield{Name: FieldDeleteTopicsRequestTopicNames, Ty: schema.TypeStrArray},
			&schema.Mfield{Name: FieldDeleteTopicsRequestTimeoutMs, Ty: schema.TypeInt32},
		),

		// Message: DeleteTopicsRequest, API Key: 20, Version: 1
		schema.NewSchema("DeleteTopicsRequest:v1",
			&schema.Mfield{Name: FieldDeleteTopicsRequestTopicNames, Ty: schema.TypeStrArray},
			&schema.Mfield{Name: FieldDeleteTopicsRequestTimeoutMs, Ty: schema.TypeInt32},
		),

		// Message: DeleteTopicsRequest, API Key: 20, Version: 2
		schema.NewSchema("DeleteTopicsRequest:v2",
			&schema.Mfield{Name: FieldDeleteTopicsRequestTopicNames, Ty: schema.TypeStrArray},
			&schema.Mfield{Name: FieldDeleteTopicsRequestTimeoutMs, Ty: schema.TypeInt32},
		),

		// Message: DeleteTopicsRequest, API Key: 20, Version: 3
		schema.NewSchema("DeleteTopicsRequest:v3",
			&schema.Mfield{Name: FieldDeleteTopicsRequestTopicNames, Ty: schema.TypeStrArray},
			&schema.Mfield{Name: FieldDeleteTopicsRequestTimeoutMs, Ty: schema.TypeInt32},
		),

		// Message: DeleteTopicsRequest, API Key: 20, Version: 4
		schema.NewSchema("DeleteTopicsRequest:v4",
			&schema.Mfield{Name: FieldDeleteTopicsRequestTopicNames, Ty: schema.TypeStrCompactArray},
			&schema.Mfield{Name: FieldDeleteTopicsRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldDeleteTopicsRequestTags},
		),

		// Message: DeleteTopicsRequest, API Key: 20, Version: 5
		schema.NewSchema("DeleteTopicsRequest:v5",
			&schema.Mfield{Name: FieldDeleteTopicsRequestTopicNames, Ty: schema.TypeStrCompactArray},
			&schema.Mfield{Name: FieldDeleteTopicsRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldDeleteTopicsRequestTags},
		),

		// Message: DeleteTopicsRequest, API Key: 20, Version: 6
		schema.NewSchema("DeleteTopicsRequest:v6",
			&schema.ArrayCompact{Name: FieldDeleteTopicsRequestTopics, Ty: schema.NewSchema("[]DeleteTopicState:v6",
				&schema.Mfield{Name: FieldDeleteTopicsRequestTopicsName, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldDeleteTopicsRequestTopicsTopicId, Ty: schema.TypeUuid},
				&schema.SchemaTaggedFields{Name: FieldDeleteTopicsRequestTopicsTags},
			)},
			&schema.Mfield{Name: FieldDeleteTopicsRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldDeleteTopicsRequestTags},
		),
	}

}

const (

	// FieldDeleteTopicsRequestTags is: The tagged fields.
	FieldDeleteTopicsRequestTags = "Tags"

	// FieldDeleteTopicsRequestTimeoutMs is: The length of time in milliseconds to wait for the deletions to complete.
	FieldDeleteTopicsRequestTimeoutMs = "TimeoutMs"

	// FieldDeleteTopicsRequestTopicNames is: The names of the topics to delete
	FieldDeleteTopicsRequestTopicNames = "TopicNames"

	// FieldDeleteTopicsRequestTopics is: The name or topic ID of the topic
	FieldDeleteTopicsRequestTopics = "Topics"

	// FieldDeleteTopicsRequestTopicsName is: The topic name
	FieldDeleteTopicsRequestTopicsName = "Name"

	// FieldDeleteTopicsRequestTopicsTags is: The tagged fields.
	FieldDeleteTopicsRequestTopicsTags = "Tags"

	// FieldDeleteTopicsRequestTopicsTopicId is: The unique topic ID
	FieldDeleteTopicsRequestTopicsTopicId = "TopicId"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DeleteTopicsRequest.json
const originalDeleteTopicsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 20,
  "type": "request",
  "listeners": ["zkBroker", "broker", "controller"],
  "name": "DeleteTopicsRequest",
  // Versions 0, 1, 2, and 3 are the same.
  //
  // Version 4 is the first flexible version.
  //
  // Version 5 adds ErrorMessage in the response and may return a THROTTLING_QUOTA_EXCEEDED error
  // in the response if the topics deletion is throttled (KIP-599).
  //
  // Version 6 reorganizes topics, adds topic IDs and allows topic names to be null.
  "validVersions": "0-6",
  "flexibleVersions": "4+",
  "fields": [
    { "name": "Topics", "type": "[]DeleteTopicState", "versions": "6+", "about": "The name or topic ID of the topic",
      "fields": [
        {"name": "Name", "type": "string", "versions": "6+", "nullableVersions": "6+", "default": "null", "entityType": "topicName", "about": "The topic name"},
        {"name": "TopicId", "type": "uuid", "versions": "6+", "about": "The unique topic ID"}
    ]},
    { "name": "TopicNames", "type": "[]string", "versions": "0-5", "entityType": "topicName", "ignorable": true,
      "about": "The names of the topics to delete" },
    { "name": "TimeoutMs", "type": "int32", "versions": "0+",
      "about": "The length of time in milliseconds to wait for the deletions to complete." }
  ]
}
`
