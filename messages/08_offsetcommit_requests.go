package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init8OffsetCommitRequest() []schema.Schema {

	return []schema.Schema{

		// Message: OffsetCommitRequest, API Key: 8, Version: 0
		schema.NewSchema("OffsetCommitRequestv0", 
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV0",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 1
		schema.NewSchema("OffsetCommitRequestv1", 
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV1",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 2
		schema.NewSchema("OffsetCommitRequestv2", 
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestRetentionTimeMs, Ty: schema.TypeInt64},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV2",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 3
		schema.NewSchema("OffsetCommitRequestv3", 
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestRetentionTimeMs, Ty: schema.TypeInt64},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("TopicsV3",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV3",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 4
		schema.NewSchema("OffsetCommitRequestv4", 
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestRetentionTimeMs, Ty: schema.TypeInt64},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("TopicsV4",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV4",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 5
		schema.NewSchema("OffsetCommitRequestv5", 
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("TopicsV5",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV5",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 6
		schema.NewSchema("OffsetCommitRequestv6", 
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("TopicsV6",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV6",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 7
		schema.NewSchema("OffsetCommitRequestv7", 
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupInstanceId, Ty: schema.TypeStrNullable},
			&schema.Array{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("TopicsV7",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV7",
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: OffsetCommitRequest, API Key: 8, Version: 8
		schema.NewSchema("OffsetCommitRequestv8", 
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGenerationId, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldOffsetCommitRequestMemberId, Ty: schema.TypeStr},
			&schema.Mfield{Name: FieldOffsetCommitRequestGroupInstanceId, Ty: schema.TypeStrCompactNullable},
			&schema.ArrayCompact{Name: FieldOffsetCommitRequestTopics, Ty: schema.NewSchema("TopicsV8",
				&schema.Mfield{Name: FieldOffsetCommitRequestTopicsName, Ty: schema.TypeStr},
				&schema.ArrayCompact{Name: FieldOffsetCommitRequestTopicsPartitions, Ty: schema.NewSchema("PartitionsV8",
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
	// FieldOffsetCommitRequestGroupId is a field name that can be used to resolve the correct struct field.
	FieldOffsetCommitRequestGroupId = "GroupId"
	// FieldOffsetCommitRequestTopics is a field name that can be used to resolve the correct struct field.
	FieldOffsetCommitRequestTopics = "Topics"
	// FieldOffsetCommitRequestTopicsName is a field name that can be used to resolve the correct struct field.
	FieldOffsetCommitRequestTopicsName = "Name"
	// FieldOffsetCommitRequestTopicsPartitions is a field name that can be used to resolve the correct struct field.
	FieldOffsetCommitRequestTopicsPartitions = "Partitions"
	// FieldOffsetCommitRequestMemberId is a field name that can be used to resolve the correct struct field.
	FieldOffsetCommitRequestMemberId = "MemberId"
	// FieldOffsetCommitRequestRetentionTimeMs is a field name that can be used to resolve the correct struct field.
	FieldOffsetCommitRequestRetentionTimeMs = "RetentionTimeMs"
	// FieldOffsetCommitRequestTopicsPartitionsTags is a field name that can be used to resolve the correct struct field.
	FieldOffsetCommitRequestTopicsPartitionsTags = "Tags"
	// FieldOffsetCommitRequestTopicsTags is a field name that can be used to resolve the correct struct field.
	FieldOffsetCommitRequestTopicsTags = "Tags"
	// FieldOffsetCommitRequestTags is a field name that can be used to resolve the correct struct field.
	FieldOffsetCommitRequestTags = "Tags"
	// FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata is a field name that can be used to resolve the correct struct field.
	FieldOffsetCommitRequestTopicsPartitionsCommittedMetadata = "CommittedMetadata"
	// FieldOffsetCommitRequestTopicsPartitionsCommittedLeaderEpoch is a field name that can be used to resolve the correct struct field.
	FieldOffsetCommitRequestTopicsPartitionsCommittedLeaderEpoch = "CommittedLeaderEpoch"
	// FieldOffsetCommitRequestGroupInstanceId is a field name that can be used to resolve the correct struct field.
	FieldOffsetCommitRequestGroupInstanceId = "GroupInstanceId"
	// FieldOffsetCommitRequestTopicsPartitionsPartitionIndex is a field name that can be used to resolve the correct struct field.
	FieldOffsetCommitRequestTopicsPartitionsPartitionIndex = "PartitionIndex"
	// FieldOffsetCommitRequestGenerationId is a field name that can be used to resolve the correct struct field.
	FieldOffsetCommitRequestGenerationId = "GenerationId"
	// FieldOffsetCommitRequestTopicsPartitionsCommittedOffset is a field name that can be used to resolve the correct struct field.
	FieldOffsetCommitRequestTopicsPartitionsCommittedOffset = "CommittedOffset"
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

