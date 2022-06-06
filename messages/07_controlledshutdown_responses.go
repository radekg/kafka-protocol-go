package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init7ControlledShutdownResponse() []schema.Schema {

	return []schema.Schema{
		// Message: ControlledShutdownResponse, API Key: 7, Version: 0
		schema.NewSchema("ControlledShutdownResponse:v0",
			&schema.Mfield{Name: FieldControlledShutdownResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldControlledShutdownResponseRemainingPartitions, Ty: schema.NewSchema("RemainingPartitions:v0",
				&schema.Mfield{Name: FieldControlledShutdownResponseRemainingPartitionsTopicName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldControlledShutdownResponseRemainingPartitionsPartitionIndex, Ty: schema.TypeInt32},
			)},
		),

		// Message: ControlledShutdownResponse, API Key: 7, Version: 1
		schema.NewSchema("ControlledShutdownResponse:v1",
			&schema.Mfield{Name: FieldControlledShutdownResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldControlledShutdownResponseRemainingPartitions, Ty: schema.NewSchema("RemainingPartitions:v1",
				&schema.Mfield{Name: FieldControlledShutdownResponseRemainingPartitionsTopicName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldControlledShutdownResponseRemainingPartitionsPartitionIndex, Ty: schema.TypeInt32},
			)},
		),

		// Message: ControlledShutdownResponse, API Key: 7, Version: 2
		schema.NewSchema("ControlledShutdownResponse:v2",
			&schema.Mfield{Name: FieldControlledShutdownResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldControlledShutdownResponseRemainingPartitions, Ty: schema.NewSchema("RemainingPartitions:v2",
				&schema.Mfield{Name: FieldControlledShutdownResponseRemainingPartitionsTopicName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldControlledShutdownResponseRemainingPartitionsPartitionIndex, Ty: schema.TypeInt32},
			)},
		),

		// Message: ControlledShutdownResponse, API Key: 7, Version: 3
		schema.NewSchema("ControlledShutdownResponse:v3",
			&schema.Mfield{Name: FieldControlledShutdownResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldControlledShutdownResponseRemainingPartitions, Ty: schema.NewSchema("RemainingPartitions:v3",
				&schema.Mfield{Name: FieldControlledShutdownResponseRemainingPartitionsTopicName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldControlledShutdownResponseRemainingPartitionsPartitionIndex, Ty: schema.TypeInt32},
				&schema.SchemaTaggedFields{Name: FieldControlledShutdownResponseRemainingPartitionsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldControlledShutdownResponseTags},
		),
	}

}

const (

	// FieldControlledShutdownResponseErrorCode is: The top-level error code.
	FieldControlledShutdownResponseErrorCode = "ErrorCode"

	// FieldControlledShutdownResponseRemainingPartitions is: The partitions that the broker still leads.
	FieldControlledShutdownResponseRemainingPartitions = "RemainingPartitions"

	// FieldControlledShutdownResponseRemainingPartitionsPartitionIndex is: The index of the partition.
	FieldControlledShutdownResponseRemainingPartitionsPartitionIndex = "PartitionIndex"

	// FieldControlledShutdownResponseRemainingPartitionsTags is: The tagged fields.
	FieldControlledShutdownResponseRemainingPartitionsTags = "Tags"

	// FieldControlledShutdownResponseRemainingPartitionsTopicName is: The name of the topic.
	FieldControlledShutdownResponseRemainingPartitionsTopicName = "TopicName"

	// FieldControlledShutdownResponseTags is: The tagged fields.
	FieldControlledShutdownResponseTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ControlledShutdownResponse.json
const originalControlledShutdownResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 7,
  "type": "response",
  "name": "ControlledShutdownResponse",
  // Versions 1 and 2 are the same as version 0.
  "validVersions": "0-3",
  "flexibleVersions": "3+",
  "fields": [
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The top-level error code." },
    { "name": "RemainingPartitions", "type": "[]RemainingPartition", "versions": "0+",
      "about": "The partitions that the broker still leads.", "fields": [
      { "name": "TopicName", "type": "string", "versions": "0+", "mapKey": true, "entityType": "topicName",
        "about": "The name of the topic." },
      { "name": "PartitionIndex", "type": "int32", "versions": "0+", "mapKey": true,
        "about": "The index of the partition." }
    ]}
  ]
}
`
