package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init56AlterPartitionResponse() []schema.Schema {

	return []schema.Schema{
		// Message: AlterPartitionResponse, API Key: 56, Version: 0
		schema.NewSchema("AlterPartitionResponse:v0",
			&schema.Mfield{Name: FieldAlterPartitionResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldAlterPartitionResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldAlterPartitionResponseTopics, Ty: schema.NewSchema("[]TopicData:v0",
				&schema.Mfield{Name: FieldAlterPartitionResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldAlterPartitionResponseTopicsPartitions, Ty: schema.NewSchema("[]PartitionData:v0",
					&schema.Mfield{Name: FieldAlterPartitionResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAlterPartitionResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldAlterPartitionResponseTopicsPartitionsLeaderId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAlterPartitionResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAlterPartitionResponseTopicsPartitionsIsr, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldAlterPartitionResponseTopicsPartitionsPartitionEpoch, Ty: schema.TypeInt32},
					&schema.SchemaTaggedFields{Name: FieldAlterPartitionResponseTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldAlterPartitionResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldAlterPartitionResponseTags},
		),

		// Message: AlterPartitionResponse, API Key: 56, Version: 1
		schema.NewSchema("AlterPartitionResponse:v1",
			&schema.Mfield{Name: FieldAlterPartitionResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldAlterPartitionResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldAlterPartitionResponseTopics, Ty: schema.NewSchema("[]TopicData:v1",
				&schema.Mfield{Name: FieldAlterPartitionResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldAlterPartitionResponseTopicsPartitions, Ty: schema.NewSchema("[]PartitionData:v1",
					&schema.Mfield{Name: FieldAlterPartitionResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAlterPartitionResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldAlterPartitionResponseTopicsPartitionsLeaderId, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAlterPartitionResponseTopicsPartitionsLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAlterPartitionResponseTopicsPartitionsIsr, Ty: schema.TypeInt32CompactArray},
					&schema.Mfield{Name: FieldAlterPartitionResponseTopicsPartitionsLeaderRecoveryState, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldAlterPartitionResponseTopicsPartitionsPartitionEpoch, Ty: schema.TypeInt32},
					&schema.SchemaTaggedFields{Name: FieldAlterPartitionResponseTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldAlterPartitionResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldAlterPartitionResponseTags},
		),
	}

}

const (

	// FieldAlterPartitionResponseErrorCode is: The top level response error code
	FieldAlterPartitionResponseErrorCode = "ErrorCode"

	// FieldAlterPartitionResponseTags is: The tagged fields.
	FieldAlterPartitionResponseTags = "Tags"

	// FieldAlterPartitionResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldAlterPartitionResponseThrottleTimeMs = "ThrottleTimeMs"

	// FieldAlterPartitionResponseTopics is:
	FieldAlterPartitionResponseTopics = "Topics"

	// FieldAlterPartitionResponseTopicsName is: The name of the topic
	FieldAlterPartitionResponseTopicsName = "Name"

	// FieldAlterPartitionResponseTopicsPartitions is:
	FieldAlterPartitionResponseTopicsPartitions = "Partitions"

	// FieldAlterPartitionResponseTopicsPartitionsErrorCode is: The partition level error code
	FieldAlterPartitionResponseTopicsPartitionsErrorCode = "ErrorCode"

	// FieldAlterPartitionResponseTopicsPartitionsIsr is: The in-sync replica IDs.
	FieldAlterPartitionResponseTopicsPartitionsIsr = "Isr"

	// FieldAlterPartitionResponseTopicsPartitionsLeaderEpoch is: The leader epoch.
	FieldAlterPartitionResponseTopicsPartitionsLeaderEpoch = "LeaderEpoch"

	// FieldAlterPartitionResponseTopicsPartitionsLeaderId is: The broker ID of the leader.
	FieldAlterPartitionResponseTopicsPartitionsLeaderId = "LeaderId"

	// FieldAlterPartitionResponseTopicsPartitionsLeaderRecoveryState is: 1 if the partition is recovering from an unclean leader election; 0 otherwise.
	FieldAlterPartitionResponseTopicsPartitionsLeaderRecoveryState = "LeaderRecoveryState"

	// FieldAlterPartitionResponseTopicsPartitionsPartitionEpoch is: The current epoch for the partition for KRaft controllers. The current ZK version for the legacy controllers.
	FieldAlterPartitionResponseTopicsPartitionsPartitionEpoch = "PartitionEpoch"

	// FieldAlterPartitionResponseTopicsPartitionsPartitionIndex is: The partition index
	FieldAlterPartitionResponseTopicsPartitionsPartitionIndex = "PartitionIndex"

	// FieldAlterPartitionResponseTopicsPartitionsTags is: The tagged fields.
	FieldAlterPartitionResponseTopicsPartitionsTags = "Tags"

	// FieldAlterPartitionResponseTopicsTags is: The tagged fields.
	FieldAlterPartitionResponseTopicsTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AlterPartitionResponse.json
const originalAlterPartitionResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 56,
  "type": "response",
  "name": "AlterPartitionResponse",
  "validVersions": "0-1",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The top level response error code" },
    { "name": "Topics", "type": "[]TopicData", "versions": "0+", "fields": [
      { "name":  "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The name of the topic" },
      { "name": "Partitions", "type": "[]PartitionData", "versions": "0+", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The partition index" },
        { "name": "ErrorCode", "type": "int16", "versions": "0+",
          "about": "The partition level error code" },
        { "name": "LeaderId", "type": "int32", "versions": "0+", "entityType": "brokerId",
          "about": "The broker ID of the leader." },
        { "name": "LeaderEpoch", "type": "int32", "versions": "0+",
          "about": "The leader epoch." },
        { "name": "Isr", "type": "[]int32", "versions": "0+", "entityType": "brokerId",
          "about": "The in-sync replica IDs." },
        { "name": "LeaderRecoveryState", "type": "int8", "versions": "1+", "default": "0", "ignorable": true,
          "about": "1 if the partition is recovering from an unclean leader election; 0 otherwise." },
        { "name": "PartitionEpoch", "type": "int32", "versions": "0+",
          "about": "The current epoch for the partition for KRaft controllers. The current ZK version for the legacy controllers." }
      ]}
    ]}
  ]
}
`
