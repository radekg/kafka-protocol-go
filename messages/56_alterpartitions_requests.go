package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init56AlterPartitionRequest() []schema.Schema {

	return []schema.Schema{
		// Message: AlterPartitionRequest, API Key: 56, Version: 0
		schema.NewSchema("AlterPartitionRequest:v0",
			&schema.Mfield{Name: FieldAlterPartitionRequestBrokerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldAlterPartitionRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.ArrayCompact{Name: FieldAlterPartitionRequestTopics, Ty: schema.NewSchema("[]TopicData:v0",
				&schema.Mfield{Name: FieldAlterPartitionRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldAlterPartitionRequestTopicsPartitions, Ty: schema.NewSchema("[]PartitionData:v0",
					&schema.Mfield{Name: FieldAlterPartitionRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAlterPartitionRequestTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAlterPartitionRequestTopicsPartitionsNewIsr, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldAlterPartitionRequestTopicsPartitionsPartitionEpoch, Ty: schema.TypeInt32},
					&schema.SchemaTaggedFields{Name: FieldAlterPartitionRequestTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldAlterPartitionRequestTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldAlterPartitionRequestTags},
		),

		// Message: AlterPartitionRequest, API Key: 56, Version: 1
		schema.NewSchema("AlterPartitionRequest:v1",
			&schema.Mfield{Name: FieldAlterPartitionRequestBrokerId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldAlterPartitionRequestBrokerEpoch, Ty: schema.TypeInt64},
			&schema.ArrayCompact{Name: FieldAlterPartitionRequestTopics, Ty: schema.NewSchema("[]TopicData:v1",
				&schema.Mfield{Name: FieldAlterPartitionRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldAlterPartitionRequestTopicsPartitions, Ty: schema.NewSchema("[]PartitionData:v1",
					&schema.Mfield{Name: FieldAlterPartitionRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAlterPartitionRequestTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAlterPartitionRequestTopicsPartitionsNewIsr, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldAlterPartitionRequestTopicsPartitionsLeaderRecoveryState, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldAlterPartitionRequestTopicsPartitionsPartitionEpoch, Ty: schema.TypeInt32},
					&schema.SchemaTaggedFields{Name: FieldAlterPartitionRequestTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldAlterPartitionRequestTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldAlterPartitionRequestTags},
		),
	}

}

const (

	// FieldAlterPartitionRequestBrokerEpoch is: The epoch of the requesting broker
	FieldAlterPartitionRequestBrokerEpoch = "BrokerEpoch"

	// FieldAlterPartitionRequestBrokerId is: The ID of the requesting broker
	FieldAlterPartitionRequestBrokerId = "BrokerId"

	// FieldAlterPartitionRequestTags is: The tagged fields.
	FieldAlterPartitionRequestTags = "Tags"

	// FieldAlterPartitionRequestTopics is:
	FieldAlterPartitionRequestTopics = "Topics"

	// FieldAlterPartitionRequestTopicsName is: The name of the topic to alter ISRs for
	FieldAlterPartitionRequestTopicsName = "Name"

	// FieldAlterPartitionRequestTopicsPartitions is:
	FieldAlterPartitionRequestTopicsPartitions = "Partitions"

	// FieldAlterPartitionRequestTopicsPartitionsLeaderEpoch is: The leader epoch of this partition
	FieldAlterPartitionRequestTopicsPartitionsLeaderEpoch = "LeaderEpoch"

	// FieldAlterPartitionRequestTopicsPartitionsLeaderRecoveryState is: 1 if the partition is recovering from an unclean leader election; 0 otherwise.
	FieldAlterPartitionRequestTopicsPartitionsLeaderRecoveryState = "LeaderRecoveryState"

	// FieldAlterPartitionRequestTopicsPartitionsNewIsr is: The ISR for this partition
	FieldAlterPartitionRequestTopicsPartitionsNewIsr = "NewIsr"

	// FieldAlterPartitionRequestTopicsPartitionsPartitionEpoch is: The expected epoch of the partition which is being updated. For legacy cluster this is the ZkVersion in the LeaderAndIsr request.
	FieldAlterPartitionRequestTopicsPartitionsPartitionEpoch = "PartitionEpoch"

	// FieldAlterPartitionRequestTopicsPartitionsPartitionIndex is: The partition index
	FieldAlterPartitionRequestTopicsPartitionsPartitionIndex = "PartitionIndex"

	// FieldAlterPartitionRequestTopicsPartitionsTags is: The tagged fields.
	FieldAlterPartitionRequestTopicsPartitionsTags = "Tags"

	// FieldAlterPartitionRequestTopicsTags is: The tagged fields.
	FieldAlterPartitionRequestTopicsTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AlterPartitionRequest.json
const originalAlterPartitionRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implie
// See the License for the specific language governing permissions and
// limitations under the License.

{
  "apiKey": 56,
  "type": "request",
  "listeners": ["zkBroker", "controller"],
  "name": "AlterPartitionRequest",
  "validVersions": "0-1",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "BrokerId", "type": "int32", "versions": "0+", "entityType": "brokerId",
      "about": "The ID of the requesting broker" },
    { "name": "BrokerEpoch", "type": "int64", "versions": "0+", "default": "-1",
      "about": "The epoch of the requesting broker" },
    { "name": "Topics", "type": "[]TopicData", "versions": "0+", "fields": [
      { "name":  "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The name of the topic to alter ISRs for" },
      { "name": "Partitions", "type": "[]PartitionData", "versions": "0+", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The partition index" },
        { "name": "LeaderEpoch", "type": "int32", "versions": "0+",
          "about": "The leader epoch of this partition" },
        { "name": "NewIsr", "type": "[]int32", "versions": "0+", "entityType": "brokerId",
          "about": "The ISR for this partition" },
        { "name": "LeaderRecoveryState", "type": "int8", "versions": "1+", "default": "0",
          "about": "1 if the partition is recovering from an unclean leader election; 0 otherwise." },
        { "name": "PartitionEpoch", "type": "int32", "versions": "0+",
          "about": "The expected epoch of the partition which is being updated. For legacy cluster this is the ZkVersion in the LeaderAndIsr request." }
      ]}
    ]}
  ]
}
`
