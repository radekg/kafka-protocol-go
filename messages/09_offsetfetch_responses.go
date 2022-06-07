package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init9OffsetFetchResponse() []schema.Schema {

	return []schema.Schema{
		// Message: OffsetFetchResponse, API Key: 9, Version: 0
		schema.NewSchema("OffsetFetchResponse:v0",
			&schema.Array{Name: FieldOffsetFetchResponseTopics, Ty: schema.NewSchema("[]OffsetFetchResponseTopic:v0",
				&schema.Mfield{Name: FieldOffsetFetchResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetFetchResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetFetchResponsePartition:v0",
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsMetadata, Ty: schema.TypeStrNullable},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: OffsetFetchResponse, API Key: 9, Version: 1
		schema.NewSchema("OffsetFetchResponse:v1",
			&schema.Array{Name: FieldOffsetFetchResponseTopics, Ty: schema.NewSchema("[]OffsetFetchResponseTopic:v1",
				&schema.Mfield{Name: FieldOffsetFetchResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetFetchResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetFetchResponsePartition:v1",
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsMetadata, Ty: schema.TypeStrNullable},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: OffsetFetchResponse, API Key: 9, Version: 2
		schema.NewSchema("OffsetFetchResponse:v2",
			&schema.Array{Name: FieldOffsetFetchResponseTopics, Ty: schema.NewSchema("[]OffsetFetchResponseTopic:v2",
				&schema.Mfield{Name: FieldOffsetFetchResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetFetchResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetFetchResponsePartition:v2",
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsMetadata, Ty: schema.TypeStrNullable},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
			&schema.Mfield{Name: FieldOffsetFetchResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: OffsetFetchResponse, API Key: 9, Version: 3
		schema.NewSchema("OffsetFetchResponse:v3",
			&schema.Mfield{Name: FieldOffsetFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldOffsetFetchResponseTopics, Ty: schema.NewSchema("[]OffsetFetchResponseTopic:v3",
				&schema.Mfield{Name: FieldOffsetFetchResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetFetchResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetFetchResponsePartition:v3",
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsMetadata, Ty: schema.TypeStrNullable},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
			&schema.Mfield{Name: FieldOffsetFetchResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: OffsetFetchResponse, API Key: 9, Version: 4
		schema.NewSchema("OffsetFetchResponse:v4",
			&schema.Mfield{Name: FieldOffsetFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldOffsetFetchResponseTopics, Ty: schema.NewSchema("[]OffsetFetchResponseTopic:v4",
				&schema.Mfield{Name: FieldOffsetFetchResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetFetchResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetFetchResponsePartition:v4",
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsMetadata, Ty: schema.TypeStrNullable},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
			&schema.Mfield{Name: FieldOffsetFetchResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: OffsetFetchResponse, API Key: 9, Version: 5
		schema.NewSchema("OffsetFetchResponse:v5",
			&schema.Mfield{Name: FieldOffsetFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldOffsetFetchResponseTopics, Ty: schema.NewSchema("[]OffsetFetchResponseTopic:v5",
				&schema.Mfield{Name: FieldOffsetFetchResponseTopicsName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldOffsetFetchResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetFetchResponsePartition:v5",
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsCommittedLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsMetadata, Ty: schema.TypeStrNullable},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
			&schema.Mfield{Name: FieldOffsetFetchResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: OffsetFetchResponse, API Key: 9, Version: 6
		schema.NewSchema("OffsetFetchResponse:v6",
			&schema.Mfield{Name: FieldOffsetFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldOffsetFetchResponseTopics, Ty: schema.NewSchema("[]OffsetFetchResponseTopic:v6",
				&schema.Mfield{Name: FieldOffsetFetchResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldOffsetFetchResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetFetchResponsePartition:v6",
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsCommittedLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsMetadata, Ty: schema.TypeStrCompactNullable},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.SchemaTaggedFields{Name: FieldOffsetFetchResponseTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldOffsetFetchResponseTopicsTags},
			)},
			&schema.Mfield{Name: FieldOffsetFetchResponseErrorCode, Ty: schema.TypeInt16},
			&schema.SchemaTaggedFields{Name: FieldOffsetFetchResponseTags},
		),

		// Message: OffsetFetchResponse, API Key: 9, Version: 7
		schema.NewSchema("OffsetFetchResponse:v7",
			&schema.Mfield{Name: FieldOffsetFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldOffsetFetchResponseTopics, Ty: schema.NewSchema("[]OffsetFetchResponseTopic:v7",
				&schema.Mfield{Name: FieldOffsetFetchResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldOffsetFetchResponseTopicsPartitions, Ty: schema.NewSchema("[]OffsetFetchResponsePartition:v7",
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsCommittedLeaderEpoch, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsMetadata, Ty: schema.TypeStrCompactNullable},
					&schema.Mfield{Name: FieldOffsetFetchResponseTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.SchemaTaggedFields{Name: FieldOffsetFetchResponseTopicsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldOffsetFetchResponseTopicsTags},
			)},
			&schema.Mfield{Name: FieldOffsetFetchResponseErrorCode, Ty: schema.TypeInt16},
			&schema.SchemaTaggedFields{Name: FieldOffsetFetchResponseTags},
		),

		// Message: OffsetFetchResponse, API Key: 9, Version: 8
		schema.NewSchema("OffsetFetchResponse:v8",
			&schema.Mfield{Name: FieldOffsetFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldOffsetFetchResponseGroups, Ty: schema.NewSchema("[]OffsetFetchResponseGroup:v8",
				&schema.Mfield{Name: FieldOffsetFetchResponseGroupsgroupId, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldOffsetFetchResponseGroupsTopics, Ty: schema.NewSchema("[]OffsetFetchResponseTopics:v8",
					&schema.Mfield{Name: FieldOffsetFetchResponseGroupsTopicsName, Ty: schema.TypeStrCompact},
					&schema.ArrayCompact{Name: FieldOffsetFetchResponseGroupsTopicsPartitions, Ty: schema.NewSchema("[]OffsetFetchResponsePartitions:v8",
						&schema.Mfield{Name: FieldOffsetFetchResponseGroupsTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
						&schema.Mfield{Name: FieldOffsetFetchResponseGroupsTopicsPartitionsCommittedOffset, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldOffsetFetchResponseGroupsTopicsPartitionsCommittedLeaderEpoch, Ty: schema.TypeInt32},
						&schema.Mfield{Name: FieldOffsetFetchResponseGroupsTopicsPartitionsMetadata, Ty: schema.TypeStrCompactNullable},
						&schema.Mfield{Name: FieldOffsetFetchResponseGroupsTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
						&schema.SchemaTaggedFields{Name: FieldOffsetFetchResponseGroupsTopicsPartitionsTags},
					)},
					&schema.SchemaTaggedFields{Name: FieldOffsetFetchResponseGroupsTopicsTags},
				)},
				&schema.Mfield{Name: FieldOffsetFetchResponseGroupsErrorCode, Ty: schema.TypeInt16},
				&schema.SchemaTaggedFields{Name: FieldOffsetFetchResponseGroupsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldOffsetFetchResponseTags},
		),
	}

}

const (

	// FieldOffsetFetchResponseErrorCode is: The top-level error code, or 0 if there was no error.
	FieldOffsetFetchResponseErrorCode = "ErrorCode"

	// FieldOffsetFetchResponseGroups is: The responses per group id.
	FieldOffsetFetchResponseGroups = "Groups"

	// FieldOffsetFetchResponseGroupsErrorCode is: The group-level error code, or 0 if there was no error.
	FieldOffsetFetchResponseGroupsErrorCode = "ErrorCode"

	// FieldOffsetFetchResponseGroupsTags is: The tagged fields.
	FieldOffsetFetchResponseGroupsTags = "Tags"

	// FieldOffsetFetchResponseGroupsTopics is: The responses per topic.
	FieldOffsetFetchResponseGroupsTopics = "Topics"

	// FieldOffsetFetchResponseGroupsTopicsName is: The topic name.
	FieldOffsetFetchResponseGroupsTopicsName = "Name"

	// FieldOffsetFetchResponseGroupsTopicsPartitions is: The responses per partition
	FieldOffsetFetchResponseGroupsTopicsPartitions = "Partitions"

	// FieldOffsetFetchResponseGroupsTopicsPartitionsCommittedLeaderEpoch is: The leader epoch.
	FieldOffsetFetchResponseGroupsTopicsPartitionsCommittedLeaderEpoch = "CommittedLeaderEpoch"

	// FieldOffsetFetchResponseGroupsTopicsPartitionsCommittedOffset is: The committed message offset.
	FieldOffsetFetchResponseGroupsTopicsPartitionsCommittedOffset = "CommittedOffset"

	// FieldOffsetFetchResponseGroupsTopicsPartitionsErrorCode is: The partition-level error code, or 0 if there was no error.
	FieldOffsetFetchResponseGroupsTopicsPartitionsErrorCode = "ErrorCode"

	// FieldOffsetFetchResponseGroupsTopicsPartitionsMetadata is: The partition metadata.
	FieldOffsetFetchResponseGroupsTopicsPartitionsMetadata = "Metadata"

	// FieldOffsetFetchResponseGroupsTopicsPartitionsPartitionIndex is: The partition index.
	FieldOffsetFetchResponseGroupsTopicsPartitionsPartitionIndex = "PartitionIndex"

	// FieldOffsetFetchResponseGroupsTopicsPartitionsTags is: The tagged fields.
	FieldOffsetFetchResponseGroupsTopicsPartitionsTags = "Tags"

	// FieldOffsetFetchResponseGroupsTopicsTags is: The tagged fields.
	FieldOffsetFetchResponseGroupsTopicsTags = "Tags"

	// FieldOffsetFetchResponseGroupsgroupId is: The group ID.
	FieldOffsetFetchResponseGroupsgroupId = "groupId"

	// FieldOffsetFetchResponseTags is: The tagged fields.
	FieldOffsetFetchResponseTags = "Tags"

	// FieldOffsetFetchResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldOffsetFetchResponseThrottleTimeMs = "ThrottleTimeMs"

	// FieldOffsetFetchResponseTopics is: The responses per topic.
	FieldOffsetFetchResponseTopics = "Topics"

	// FieldOffsetFetchResponseTopicsName is: The topic name.
	FieldOffsetFetchResponseTopicsName = "Name"

	// FieldOffsetFetchResponseTopicsPartitions is: The responses per partition
	FieldOffsetFetchResponseTopicsPartitions = "Partitions"

	// FieldOffsetFetchResponseTopicsPartitionsCommittedLeaderEpoch is: The leader epoch.
	FieldOffsetFetchResponseTopicsPartitionsCommittedLeaderEpoch = "CommittedLeaderEpoch"

	// FieldOffsetFetchResponseTopicsPartitionsCommittedOffset is: The committed message offset.
	FieldOffsetFetchResponseTopicsPartitionsCommittedOffset = "CommittedOffset"

	// FieldOffsetFetchResponseTopicsPartitionsErrorCode is: The error code, or 0 if there was no error.
	FieldOffsetFetchResponseTopicsPartitionsErrorCode = "ErrorCode"

	// FieldOffsetFetchResponseTopicsPartitionsMetadata is: The partition metadata.
	FieldOffsetFetchResponseTopicsPartitionsMetadata = "Metadata"

	// FieldOffsetFetchResponseTopicsPartitionsPartitionIndex is: The partition index.
	FieldOffsetFetchResponseTopicsPartitionsPartitionIndex = "PartitionIndex"

	// FieldOffsetFetchResponseTopicsPartitionsTags is: The tagged fields.
	FieldOffsetFetchResponseTopicsPartitionsTags = "Tags"

	// FieldOffsetFetchResponseTopicsTags is: The tagged fields.
	FieldOffsetFetchResponseTopicsTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/OffsetFetchResponse.json
const originalOffsetFetchResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 9,
  "type": "response",
  "name": "OffsetFetchResponse",
  // Version 1 is the same as version 0.
  //
  // Version 2 adds a top-level error code.
  //
  // Version 3 adds the throttle time.
  //
  // Starting in version 4, on quota violation, brokers send out responses before throttling.
  //
  // Version 5 adds the leader epoch to the committed offset.
  //
  // Version 6 is the first flexible version.
  //
  // Version 7 adds pending offset commit as new error response on partition level.
  //
  // Version 8 is adding support for fetching offsets for multiple groups
  "validVersions": "0-8",
  "flexibleVersions": "6+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "3+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Topics", "type": "[]OffsetFetchResponseTopic", "versions": "0-7",
      "about": "The responses per topic.", "fields": [
      { "name": "Name", "type": "string", "versions": "0-7", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "Partitions", "type": "[]OffsetFetchResponsePartition", "versions": "0-7",
        "about": "The responses per partition", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0-7",
          "about": "The partition index." },
        { "name": "CommittedOffset", "type": "int64", "versions": "0-7",
          "about": "The committed message offset." },
        { "name": "CommittedLeaderEpoch", "type": "int32", "versions": "5-7", "default": "-1",
          "ignorable": true, "about": "The leader epoch." },
        { "name": "Metadata", "type": "string", "versions": "0-7", "nullableVersions": "0-7",
          "about": "The partition metadata." },
        { "name": "ErrorCode", "type": "int16", "versions": "0-7",
          "about": "The error code, or 0 if there was no error." }
      ]}
    ]},
    { "name": "ErrorCode", "type": "int16", "versions": "2-7", "default": "0", "ignorable": true,
      "about": "The top-level error code, or 0 if there was no error." },
    {"name": "Groups", "type": "[]OffsetFetchResponseGroup", "versions": "8+",
      "about": "The responses per group id.", "fields": [
      { "name": "groupId", "type": "string", "versions": "8+", "entityType": "groupId",
        "about": "The group ID." },
      { "name": "Topics", "type": "[]OffsetFetchResponseTopics", "versions": "8+",
        "about": "The responses per topic.", "fields": [
        { "name": "Name", "type": "string", "versions": "8+", "entityType": "topicName",
          "about": "The topic name." },
        { "name": "Partitions", "type": "[]OffsetFetchResponsePartitions", "versions": "8+",
          "about": "The responses per partition", "fields": [
          { "name": "PartitionIndex", "type": "int32", "versions": "8+",
            "about": "The partition index." },
          { "name": "CommittedOffset", "type": "int64", "versions": "8+",
            "about": "The committed message offset." },
          { "name": "CommittedLeaderEpoch", "type": "int32", "versions": "8+", "default": "-1",
            "ignorable": true, "about": "The leader epoch." },
          { "name": "Metadata", "type": "string", "versions": "8+", "nullableVersions": "8+",
            "about": "The partition metadata." },
          { "name": "ErrorCode", "type": "int16", "versions": "8+",
            "about": "The partition-level error code, or 0 if there was no error." }
        ]}
      ]},
      { "name": "ErrorCode", "type": "int16", "versions": "8+", "default": "0",
        "about": "The group-level error code, or 0 if there was no error." }
    ]}
  ]
}
`
