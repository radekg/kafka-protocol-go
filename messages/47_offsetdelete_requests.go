package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init47OffsetDeleteRequest() []schema.Schema {

	return []schema.Schema{

		// Message: OffsetDeleteRequest, API Key: 47, Version: 0
		schema.NewSchema("OffsetDeleteRequestv0",
			&schema.Mfield{Name: FieldOffsetDeleteRequestGroupId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldOffsetDeleteRequestTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldOffsetDeleteRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetDeleteRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV0",
					&schema.Mfield{Name: FieldOffsetDeleteRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
				)},
			)},
		),
	}
}

const (
	// FieldOffsetDeleteRequestGroupId is a field name that can be used to resolve the correct struct field.
	FieldOffsetDeleteRequestGroupId = "GroupId"
	// FieldOffsetDeleteRequestTopics is a field name that can be used to resolve the correct struct field.
	FieldOffsetDeleteRequestTopics = "Topics"
	// FieldOffsetDeleteRequestTopicsName is a field name that can be used to resolve the correct struct field.
	FieldOffsetDeleteRequestTopicsName = "Name"
	// FieldOffsetDeleteRequestTopicsPartitions is a field name that can be used to resolve the correct struct field.
	FieldOffsetDeleteRequestTopicsPartitions = "Partitions"
	// FieldOffsetDeleteRequestTopicsPartitionsPartitionIndex is a field name that can be used to resolve the correct struct field.
	FieldOffsetDeleteRequestTopicsPartitionsPartitionIndex = "PartitionIndex"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/OffsetDeleteRequest.json
const originalOffsetDeleteRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "OffsetDeleteRequest",
  "validVersions": "0",
  "flexibleVersions": "none",
  "fields": [
    { "name": "GroupId", "type": "string", "versions": "0+", "entityType": "groupId",
      "about": "The unique group identifier." },
    { "name": "Topics", "type": "[]OffsetDeleteRequestTopic", "versions": "0+",
      "about": "The topics to delete offsets for", "fields": [
        { "name": "Name",  "type": "string",  "versions": "0+", "mapKey": true, "entityType": "topicName",
          "about": "The topic name." },
        { "name": "Partitions", "type": "[]OffsetDeleteRequestPartition", "versions": "0+",
          "about": "Each partition to delete offsets for.", "fields": [
            { "name": "PartitionIndex", "type": "int32", "versions": "0+",
              "about": "The partition index." }
          ]
        }
      ]
    }
  ]
}
`
