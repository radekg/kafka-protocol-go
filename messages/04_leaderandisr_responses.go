package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init4LeaderAndIsrResponse() []schema.Schema {

	return []schema.Schema{

		// Message: LeaderAndIsrResponse, API Key: 4, Version: 0
		schema.NewSchema("LeaderAndIsrResponsev0",
			&schema.Mfield{Name: FieldLeaderAndIsrResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldLeaderAndIsrResponsePartitionErrors, Ty: schema.NewSchema("PartitionErrorsV0")},
		),

		// Message: LeaderAndIsrResponse, API Key: 4, Version: 1
		schema.NewSchema("LeaderAndIsrResponsev1",
			&schema.Mfield{Name: FieldLeaderAndIsrResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldLeaderAndIsrResponsePartitionErrors, Ty: schema.NewSchema("PartitionErrorsV1")},
		),

		// Message: LeaderAndIsrResponse, API Key: 4, Version: 2
		schema.NewSchema("LeaderAndIsrResponsev2",
			&schema.Mfield{Name: FieldLeaderAndIsrResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldLeaderAndIsrResponsePartitionErrors, Ty: schema.NewSchema("PartitionErrorsV2")},
		),

		// Message: LeaderAndIsrResponse, API Key: 4, Version: 3
		schema.NewSchema("LeaderAndIsrResponsev3",
			&schema.Mfield{Name: FieldLeaderAndIsrResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldLeaderAndIsrResponsePartitionErrors, Ty: schema.NewSchema("PartitionErrorsV3")},
		),

		// Message: LeaderAndIsrResponse, API Key: 4, Version: 4
		schema.NewSchema("LeaderAndIsrResponsev4",
			&schema.Mfield{Name: FieldLeaderAndIsrResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldLeaderAndIsrResponsePartitionErrors, Ty: schema.NewSchema("PartitionErrorsV4")},
			&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrResponseTags},
		),

		// Message: LeaderAndIsrResponse, API Key: 4, Version: 5
		schema.NewSchema("LeaderAndIsrResponsev5",
			&schema.Mfield{Name: FieldLeaderAndIsrResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldLeaderAndIsrResponseTopics, Ty: schema.NewSchema("TopicsV5",
				&schema.Mfield{Name: FieldLeaderAndIsrResponseTopicsTopicId, Ty: schema.TypeUuid},
				&schema.ArrayCompact{Name: FieldLeaderAndIsrResponseTopicsPartitionErrors, Ty: schema.NewSchema("PartitionErrorsV5")},
				&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrResponseTags},
		),

		// Message: LeaderAndIsrResponse, API Key: 4, Version: 6
		schema.NewSchema("LeaderAndIsrResponsev6",
			&schema.Mfield{Name: FieldLeaderAndIsrResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldLeaderAndIsrResponseTopics, Ty: schema.NewSchema("TopicsV6",
				&schema.Mfield{Name: FieldLeaderAndIsrResponseTopicsTopicId, Ty: schema.TypeUuid},
				&schema.ArrayCompact{Name: FieldLeaderAndIsrResponseTopicsPartitionErrors, Ty: schema.NewSchema("PartitionErrorsV6")},
				&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldLeaderAndIsrResponseTags},
		),
	}
}

const (
	// FieldLeaderAndIsrResponseTopicsTopicId is: The unique topic ID
	FieldLeaderAndIsrResponseTopicsTopicId = "TopicId"
	// FieldLeaderAndIsrResponseTopicsPartitionErrors is: Each partition.
	FieldLeaderAndIsrResponseTopicsPartitionErrors = "PartitionErrors"
	// FieldLeaderAndIsrResponseTopicsTags is: The tagged fields.
	FieldLeaderAndIsrResponseTopicsTags = "Tags"
	// FieldLeaderAndIsrResponseErrorCode is: The error code, or 0 if there was no error.
	FieldLeaderAndIsrResponseErrorCode = "ErrorCode"
	// FieldLeaderAndIsrResponsePartitionErrors is: Each partition in v0 to v4 message.
	FieldLeaderAndIsrResponsePartitionErrors = "PartitionErrors"
	// FieldLeaderAndIsrResponseTags is: The tagged fields.
	FieldLeaderAndIsrResponseTags = "Tags"
	// FieldLeaderAndIsrResponseTopics is: Each topic
	FieldLeaderAndIsrResponseTopics = "Topics"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/LeaderAndIsrResponse.json
const originalLeaderAndIsrResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 4,
  "type": "response",
  "name": "LeaderAndIsrResponse",
  // Version 1 adds KAFKA_STORAGE_ERROR as a valid error code.
  //
  // Version 2 is the same as version 1.
  //
  // Version 3 is the same as version 2.
  //
  // Version 4 is the first flexible version.
  //
  // Version 5 removes TopicName and replaces it with TopicId and reorganizes
  // the partitions by topic, as described by KIP-516.
  "validVersions": "0-6",
  "flexibleVersions": "4+",
  "fields": [
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The error code, or 0 if there was no error." },
    { "name": "PartitionErrors", "type": "[]LeaderAndIsrPartitionError", "versions": "0-4",
      "about": "Each partition in v0 to v4 message."},
    { "name":  "Topics", "type": "[]LeaderAndIsrTopicError", "versions": "5+",
      "about": "Each topic", "fields": [
      { "name": "TopicId", "type": "uuid", "versions": "5+", "mapKey": true,
        "about": "The unique topic ID" },
      { "name": "PartitionErrors", "type": "[]LeaderAndIsrPartitionError", "versions": "5+",
        "about": "Each partition."}
      ]}
    ],
    "commonStructs": [
    { "name": "LeaderAndIsrPartitionError", "versions": "0+", "fields": [
      { "name": "TopicName", "type": "string", "versions": "0-4", "entityType": "topicName", "ignorable": true,
        "about": "The topic name."},
      { "name": "PartitionIndex", "type": "int32", "versions": "0+",
        "about": "The partition index." },
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The partition error code, or 0 if there was no error." }
    ]}
  ]
}
`
