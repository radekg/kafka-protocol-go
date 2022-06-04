package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init43ElectLeadersRequest() []schema.Schema {

	return []schema.Schema{

		// Message: ElectLeadersRequest, API Key: 43, Version: 0
		schema.NewSchema("ElectLeadersRequestv0",
			&schema.Array{Name: FieldElectLeadersRequestTopicPartitions, Ty: schema.NewSchema("TopicPartitionsV0",
				&schema.Mfield{Name: FieldElectLeadersRequestTopicPartitionsTopic, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldElectLeadersRequestTopicPartitionsPartitions, Ty: schema.TypeInt32Array},
			)},
			&schema.Mfield{Name: FieldElectLeadersRequestTimeoutMs, Ty: schema.TypeInt32},
		),

		// Message: ElectLeadersRequest, API Key: 43, Version: 1
		schema.NewSchema("ElectLeadersRequestv1",
			&schema.Mfield{Name: FieldElectLeadersRequestElectionType, Ty: schema.TypeInt8},
			&schema.Array{Name: FieldElectLeadersRequestTopicPartitions, Ty: schema.NewSchema("TopicPartitionsV1",
				&schema.Mfield{Name: FieldElectLeadersRequestTopicPartitionsTopic, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldElectLeadersRequestTopicPartitionsPartitions, Ty: schema.TypeInt32Array},
			)},
			&schema.Mfield{Name: FieldElectLeadersRequestTimeoutMs, Ty: schema.TypeInt32},
		),

		// Message: ElectLeadersRequest, API Key: 43, Version: 2
		schema.NewSchema("ElectLeadersRequestv2",
			&schema.Mfield{Name: FieldElectLeadersRequestElectionType, Ty: schema.TypeInt8},
			&schema.ArrayCompact{Name: FieldElectLeadersRequestTopicPartitions, Ty: schema.NewSchema("TopicPartitionsV2",
				&schema.Mfield{Name: FieldElectLeadersRequestTopicPartitionsTopic, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldElectLeadersRequestTopicPartitionsPartitions, Ty: schema.TypeInt32CompactArray},
				&schema.SchemaTaggedFields{Name: FieldElectLeadersRequestTopicPartitionsTags},
			)},
			&schema.Mfield{Name: FieldElectLeadersRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldElectLeadersRequestTags},
		),
	}
}

const (
	// FieldElectLeadersRequestTimeoutMs is: The time in ms to wait for the election to complete.
	FieldElectLeadersRequestTimeoutMs = "TimeoutMs"
	// FieldElectLeadersRequestElectionType is: Type of elections to conduct for the partition. A value of '0' elects the preferred replica. A value of '1' elects the first live replica if there are no in-sync replica.
	FieldElectLeadersRequestElectionType = "ElectionType"
	// FieldElectLeadersRequestTopicPartitionsTags is: The tagged fields.
	FieldElectLeadersRequestTopicPartitionsTags = "Tags"
	// FieldElectLeadersRequestTags is: The tagged fields.
	FieldElectLeadersRequestTags = "Tags"
	// FieldElectLeadersRequestTopicPartitions is: The topic partitions to elect leaders.
	FieldElectLeadersRequestTopicPartitions = "TopicPartitions"
	// FieldElectLeadersRequestTopicPartitionsTopic is: The name of a topic.
	FieldElectLeadersRequestTopicPartitionsTopic = "Topic"
	// FieldElectLeadersRequestTopicPartitionsPartitions is: The partitions of this topic whose leader should be elected.
	FieldElectLeadersRequestTopicPartitionsPartitions = "Partitions"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/ElectLeadersRequest.json
const originalElectLeadersRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 43,
  "type": "request",
  "listeners": ["zkBroker", "broker", "controller"],
  "name": "ElectLeadersRequest",
  // Version 1 implements multiple leader election types, as described by KIP-460.
  //
  // Version 2 is the first flexible version.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ElectionType", "type": "int8", "versions": "1+",
      "about": "Type of elections to conduct for the partition. A value of '0' elects the preferred replica. A value of '1' elects the first live replica if there are no in-sync replica." },
    { "name": "TopicPartitions", "type": "[]TopicPartitions", "versions": "0+", "nullableVersions": "0+",
      "about": "The topic partitions to elect leaders.",
      "fields": [
        { "name": "Topic", "type": "string", "versions": "0+", "entityType": "topicName", "mapKey": true,
          "about": "The name of a topic." },
        { "name": "Partitions", "type": "[]int32", "versions": "0+",
          "about": "The partitions of this topic whose leader should be elected." }
      ]
    },
    { "name": "TimeoutMs", "type": "int32", "versions": "0+", "default": "60000",
      "about": "The time in ms to wait for the election to complete." }
  ]
}
`
