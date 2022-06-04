package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init28TxnOffsetCommitRequest() []schema.Schema {

	return []schema.Schema{

		// Message: TxnOffsetCommitRequest, API Key: 28, Version: 0
		schema.NewSchema("TxnOffsetCommitRequestv0", 
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestTransactionalId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldTxnOffsetCommitRequestTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldTxnOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV0",
					&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: TxnOffsetCommitRequest, API Key: 28, Version: 1
		schema.NewSchema("TxnOffsetCommitRequestv1", 
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestTransactionalId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldTxnOffsetCommitRequestTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldTxnOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV1",
					&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: TxnOffsetCommitRequest, API Key: 28, Version: 2
		schema.NewSchema("TxnOffsetCommitRequestv2", 
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestTransactionalId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.Array{Name: FieldTxnOffsetCommitRequestTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldTxnOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV2",
					&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsPartitionsCommittedLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: TxnOffsetCommitRequest, API Key: 28, Version: 3
		schema.NewSchema("TxnOffsetCommitRequestv3", 
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestTransactionalId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestGroupId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestProducerId, Ty: schema.TypeInt64},
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestProducerEpoch, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestMemberId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldTxnOffsetCommitRequestGroupInstanceId, Ty: schema.TypeStrCompactNullable},
			&schema.ArrayCompact{Name: FieldTxnOffsetCommitRequestTopics, Ty: schema.NewSchema("TopicsV3",
				&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldTxnOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV3",
					&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsPartitionsCommittedLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldTxnOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldTxnOffsetCommitRequestTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldTxnOffsetCommitRequestTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldTxnOffsetCommitRequestTags},
		),

	}
}

const (
	// FieldTxnOffsetCommitRequestTopics is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestTopics = "Topics"
	// FieldTxnOffsetCommitRequestTopicsPartitions is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestTopicsPartitions = "Partitions"
	// FieldTxnOffsetCommitRequestGenerationId is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestGenerationId = "GenerationId"
	// FieldTxnOffsetCommitRequestGroupId is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestGroupId = "GroupId"
	// FieldTxnOffsetCommitRequestProducerId is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestProducerId = "ProducerId"
	// FieldTxnOffsetCommitRequestProducerEpoch is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestProducerEpoch = "ProducerEpoch"
	// FieldTxnOffsetCommitRequestTopicsPartitionsCommittedMetadata is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestTopicsPartitionsCommittedMetadata = "CommittedMetadata"
	// FieldTxnOffsetCommitRequestGroupInstanceId is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestGroupInstanceId = "GroupInstanceId"
	// FieldTxnOffsetCommitRequestTopicsTags is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestTopicsTags = "Tags"
	// FieldTxnOffsetCommitRequestTransactionalId is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestTransactionalId = "TransactionalId"
	// FieldTxnOffsetCommitRequestTopicsPartitionsPartitionIndex is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestTopicsPartitionsPartitionIndex = "PartitionIndex"
	// FieldTxnOffsetCommitRequestTopicsPartitionsCommittedOffset is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestTopicsPartitionsCommittedOffset = "CommittedOffset"
	// FieldTxnOffsetCommitRequestTopicsName is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestTopicsName = "Name"
	// FieldTxnOffsetCommitRequestMemberId is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestMemberId = "MemberId"
	// FieldTxnOffsetCommitRequestTopicsPartitionsTags is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestTopicsPartitionsTags = "Tags"
	// FieldTxnOffsetCommitRequestTags is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestTags = "Tags"
	// FieldTxnOffsetCommitRequestTopicsPartitionsCommittedLeaderEpoch is a field name that can be used to resolve the correct struct field.
	FieldTxnOffsetCommitRequestTopicsPartitionsCommittedLeaderEpoch = "CommittedLeaderEpoch"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/TxnOffsetCommitRequest.json
const originalTxnOffsetCommitRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 28,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "TxnOffsetCommitRequest",
  // Version 1 is the same as version 0.
  //
  // Version 2 adds the committed leader epoch.
  //
  // Version 3 adds the member.id, group.instance.id and generation.id.
  "validVersions": "0-3",
  "flexibleVersions": "3+",
  "fields": [
    { "name": "TransactionalId", "type": "string", "versions": "0+", "entityType": "transactionalId",
      "about": "The ID of the transaction." },
    { "name": "GroupId", "type": "string", "versions": "0+", "entityType": "groupId",
      "about": "The ID of the group." },
    { "name": "ProducerId", "type": "int64", "versions": "0+", "entityType": "producerId",
      "about": "The current producer ID in use by the transactional ID." },
    { "name": "ProducerEpoch", "type": "int16", "versions": "0+",
      "about": "The current epoch associated with the producer ID." },
    { "name": "GenerationId", "type": "int32", "versions": "3+", "default": "-1",
      "about": "The generation of the consumer." },
    { "name": "MemberId", "type": "string", "versions": "3+", "default": "",
      "about": "The member ID assigned by the group coordinator." },
    { "name": "GroupInstanceId", "type": "string", "versions": "3+",
      "nullableVersions": "3+", "default": "null",
      "about": "The unique identifier of the consumer instance provided by end user." },
    { "name": "Topics", "type" : "[]TxnOffsetCommitRequestTopic", "versions": "0+",
      "about": "Each topic that we want to commit offsets for.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "Partitions", "type": "[]TxnOffsetCommitRequestPartition", "versions": "0+",
        "about": "The partitions inside the topic that we want to committ offsets for.", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The index of the partition within the topic." },
        { "name": "CommittedOffset", "type": "int64", "versions": "0+",
          "about": "The message offset to be committed." },
        { "name": "CommittedLeaderEpoch", "type": "int32", "versions": "2+", "default": "-1", "ignorable": true,
          "about": "The leader epoch of the last consumed record." },
        { "name": "CommittedMetadata", "type": "string", "versions": "0+", "nullableVersions": "0+",
          "about": "Any associated metadata the client wants to keep." }
      ]}
    ]}
  ]
}
`

