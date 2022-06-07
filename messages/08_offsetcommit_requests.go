package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init8OffsetCommitRequest() []schema.Schema {

	return []schema.Schema{
		// Message: OffsetCommitRequest, API Key: 8, Version: 0
		schema.NewSchema("OffsetCommitRequest:v0",
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("[]OffsetCommitRequestTopic:v0",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitRequestPartition:v0",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 1
		schema.NewSchema("OffsetCommitRequest:v1",
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("[]OffsetCommitRequestTopic:v1",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitRequestPartition:v1",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 2
		schema.NewSchema("OffsetCommitRequest:v2",
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestRetentionTimeMs, Ty: schema.TypeInt64},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("[]OffsetCommitRequestTopic:v2",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitRequestPartition:v2",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 3
		schema.NewSchema("OffsetCommitRequest:v3",
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestRetentionTimeMs, Ty: schema.TypeInt64},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("[]OffsetCommitRequestTopic:v3",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitRequestPartition:v3",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 4
		schema.NewSchema("OffsetCommitRequest:v4",
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestRetentionTimeMs, Ty: schema.TypeInt64},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("[]OffsetCommitRequestTopic:v4",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitRequestPartition:v4",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 5
		schema.NewSchema("OffsetCommitRequest:v5",
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("[]OffsetCommitRequestTopic:v5",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitRequestPartition:v5",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 6
		schema.NewSchema("OffsetCommitRequest:v6",
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("[]OffsetCommitRequestTopic:v6",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitRequestPartition:v6",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 7
		schema.NewSchema("OffsetCommitRequest:v7",
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupInstanceId, Ty: schema.TypeStrNullable},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("[]OffsetCommitRequestTopic:v7",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitRequestPartition:v7",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 8
		schema.NewSchema("OffsetCommitRequest:v8",
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupInstanceId, Ty: schema.TypeStrCompactNullable},
			&schema.ArrayCompact{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("[]OffsetCommitRequestTopic:v8",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("[]OffsetCommitRequestPartition:v8",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldOffsetCommitRequestTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldOffsetCommitRequestTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldOffsetCommitRequestTags},
		),
	}

}

const (

	// FieldOffsetCommitRequestGenerationId is: The generation of the group.
	FieldOffsetCommitRequestGenerationId = "GenerationId"

	// FieldOffsetCommitRequestGroupId is: The unique group identifier.
	FieldOffsetCommitRequestGroupId = "GroupId"

	// FieldOffsetCommitRequestGroupInstanceId is: The unique identifier of the consumer instance provided by end user.
	FieldOffsetCommitRequestGroupInstanceId = "GroupInstanceId"

	// FieldOffsetCommitRequestMemberId is: The member ID assigned by the group coordinator.
	FieldOffsetCommitRequestMemberId = "MemberId"

	// FieldOffsetCommitRequestRetentionTimeMs is: The time period in ms to retain the offset.
	FieldOffsetCommitRequestRetentionTimeMs = "RetentionTimeMs"

	// FieldOffsetCommitRequestTags is: The tagged fields.
	FieldOffsetCommitRequestTags = "Tags"

	// FieldOffsetCommitRequestTopics is: The topics to commit offsets for.
	FieldOffsetCommitRequestTopics = "Topics"

	// FieldOffsetCommitRequestTopicsName is: The topic name.
	FieldOffsetCommitRequestTopicsName = "Name"

	// FieldOffsetCommitRequestTopicsPartitions is: Each partition to commit offsets for.
	FieldOffsetCommitRequestTopicsPartitions = "Partitions"

	// FieldOffsetCommitRequestTopicsPartitionsCommittedLeaderEpoch is: The leader epoch of this partition.
	FieldOffsetCommitRequestTopicsPartitionsCommittedLeaderEpoch = "CommittedLeaderEpoch"

	// FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata is: Any associated metadata the client wants to keep.
	FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata = "CommittedMetadata"

	// FieldOffsetCommitRequestTopicsPartitionsCommittedOffset is: The message offset to be committed.
	FieldOffsetCommitRequestTopicsPartitionsCommittedOffset = "CommittedOffset"

	// FieldOffsetCommitRequestTopicsPartitionsPartitionIndex is: The partition index.
	FieldOffsetCommitRequestTopicsPartitionsPartitionIndex = "PartitionIndex"

	// FieldOffsetCommitRequestTopicsPartitionsTags is: The tagged fields.
	FieldOffsetCommitRequestTopicsPartitionsTags = "Tags"

	// FieldOffsetCommitRequestTopicsTags is: The tagged fields.
	FieldOffsetCommitRequestTopicsTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/OffsetCommitRequest.json
const originalOffsetCommitRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 8,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "OffsetCommitRequest",
  // Version 1 adds timestamp and group membership information, as well as the commit timestamp.
  //
  // Version 2 adds retention time.  It removes the commit timestamp added in version 1.
  //
  // Version 3 and 4 are the same as version 2. 
  //
  // Version 5 removes the retention time, which is now controlled only by a broker configuration.
  //
  // Version 6 adds the leader epoch for fencing.
  //
  // version 7 adds a new field called groupInstanceId to indicate member identity across restarts.
  //
  // Version 8 is the first flexible version.
  "validVersions": "0-8",
  "flexibleVersions": "8+",
  "fields": [
    { "name": "GroupId", "type": "string", "versions": "0+", "entityType": "groupId",
      "about": "The unique group identifier." },
    { "name": "GenerationId", "type": "int32", "versions": "1+", "default": "-1", "ignorable": true,
      "about": "The generation of the group." },
    { "name": "MemberId", "type": "string", "versions": "1+", "ignorable": true,
      "about": "The member ID assigned by the group coordinator." },
    { "name": "GroupInstanceId", "type": "string", "versions": "7+",
      "nullableVersions": "7+", "default": "null",
      "about": "The unique identifier of the consumer instance provided by end user." },
    { "name": "RetentionTimeMs", "type": "int64", "versions": "2-4", "default": "-1", "ignorable": true,
      "about": "The time period in ms to retain the offset." },
    { "name": "Topics", "type": "[]OffsetCommitRequestTopic", "versions": "0+",
      "about": "The topics to commit offsets for.",  "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "Partitions", "type": "[]OffsetCommitRequestPartition", "versions": "0+",
        "about": "Each partition to commit offsets for.", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "CommittedOffset", "type": "int64", "versions": "0+",
          "about": "The message offset to be committed." },
        { "name": "CommittedLeaderEpoch", "type": "int32", "versions": "6+", "default": "-1", "ignorable": true,
          "about": "The leader epoch of this partition." },
        // CommitTimestamp has been removed from v2 and later.
        { "name": "CommitTimestamp", "type": "int64", "versions": "1", "default": "-1",
          "about": "The timestamp of the commit." },
        { "name": "CommittedMetadata", "type": "string", "versions": "0+", "nullableVersions": "0+",
          "about": "Any associated metadata the client wants to keep." }
      ]}
    ]}
  ]
}
`
