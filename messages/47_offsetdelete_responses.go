package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init47OffsetDeleteResponse() []schema.Schema {

	return []schema.Schema{

		// Message: OffsetDeleteResponse, API Key: 47, Version: 0
		schema.NewSchema("OffsetDeleteResponsev0",
			&schema.Mfield{Name: FieldOffsetDeleteResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldOffsetDeleteResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldOffsetDeleteResponseTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldOffsetDeleteResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetDeleteResponseTopicsPartitions, Ty: schema.NewSchema("PartitionsV0",
					&schema.Mfield{Name: FieldOffsetDeleteResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetDeleteResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),
	}
}

const (
	// FieldOffsetDeleteResponseErrorCode is: The top-level error code, or 0 if there was no error.
	FieldOffsetDeleteResponseErrorCode = "ErrorCode"
	// FieldOffsetDeleteResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldOffsetDeleteResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldOffsetDeleteResponseTopics is: The responses for each topic.
	FieldOffsetDeleteResponseTopics = "Topics"
	// FieldOffsetDeleteResponseTopicsName is: The topic name.
	FieldOffsetDeleteResponseTopicsName = "Name"
	// FieldOffsetDeleteResponseTopicsPartitions is: The responses for each partition in the topic.
	FieldOffsetDeleteResponseTopicsPartitions = "Partitions"
	// FieldOffsetDeleteResponseTopicsPartitionsErrorCode is: The error code, or 0 if there was no error.
	FieldOffsetDeleteResponseTopicsPartitionsErrorCode = "ErrorCode"
	// FieldOffsetDeleteResponseTopicsPartitionsPartitionIndex is: The partition index.
	FieldOffsetDeleteResponseTopicsPartitionsPartitionIndex = "PartitionIndex"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/OffsetDeleteResponse.json
const originalOffsetDeleteResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 47,
  "type": "response",
  "name": "OffsetDeleteResponse",
  "validVersions": "0",
  "flexibleVersions": "none",
  "fields": [
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The top-level error code, or 0 if there was no error." },
    { "name": "ThrottleTimeMs",  "type": "int32",  "versions": "0+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Topics", "type": "[]OffsetDeleteResponseTopic", "versions": "0+",
      "about": "The responses for each topic.", "fields": [
        { "name": "Name", "type": "string", "versions": "0+", "mapKey": true, "entityType": "topicName",
          "about": "The topic name." },
        { "name": "Partitions", "type": "[]OffsetDeleteResponsePartition", "versions": "0+",
          "about": "The responses for each partition in the topic.", "fields": [
            { "name": "PartitionIndex", "type": "int32", "versions": "0+", "mapKey": true,
              "about": "The partition index." },
            { "name": "ErrorCode", "type": "int16", "versions": "0+",
              "about": "The error code, or 0 if there was no error." }
          ]
        }
      ]
    }
  ]
}
`
