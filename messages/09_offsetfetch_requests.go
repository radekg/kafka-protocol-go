package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init9OffsetFetchRequest() []schema.Schema {

	return []schema.Schema{

		// Message: OffsetFetchRequest, API Key: 9, Version: 0
		schema.NewSchema("OffsetFetchRequestv0",
			&schema.Mfield{Name: FieldOffsetFetchRequestGroupId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldOffsetFetchRequestTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsPartitionIndexes, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: OffsetFetchRequest, API Key: 9, Version: 1
		schema.NewSchema("OffsetFetchRequestv1",
			&schema.Mfield{Name: FieldOffsetFetchRequestGroupId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldOffsetFetchRequestTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsPartitionIndexes, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: OffsetFetchRequest, API Key: 9, Version: 2
		schema.NewSchema("OffsetFetchRequestv2",
			&schema.Mfield{Name: FieldOffsetFetchRequestGroupId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldOffsetFetchRequestTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsPartitionIndexes, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: OffsetFetchRequest, API Key: 9, Version: 3
		schema.NewSchema("OffsetFetchRequestv3",
			&schema.Mfield{Name: FieldOffsetFetchRequestGroupId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldOffsetFetchRequestTopics, Ty: schema.NewSchema("TopicsV3",
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsPartitionIndexes, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: OffsetFetchRequest, API Key: 9, Version: 4
		schema.NewSchema("OffsetFetchRequestv4",
			&schema.Mfield{Name: FieldOffsetFetchRequestGroupId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldOffsetFetchRequestTopics, Ty: schema.NewSchema("TopicsV4",
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsPartitionIndexes, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: OffsetFetchRequest, API Key: 9, Version: 5
		schema.NewSchema("OffsetFetchRequestv5",
			&schema.Mfield{Name: FieldOffsetFetchRequestGroupId, Ty: schema.TypeStr},
			&schema.Array{Name: FieldOffsetFetchRequestTopics, Ty: schema.NewSchema("TopicsV5",
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsPartitionIndexes, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: OffsetFetchRequest, API Key: 9, Version: 6
		schema.NewSchema("OffsetFetchRequestv6",
			&schema.Mfield{Name: FieldOffsetFetchRequestGroupId, Ty: schema.TypeStrCompact},
			&schema.ArrayCompact{Name: FieldOffsetFetchRequestTopics, Ty: schema.NewSchema("TopicsV6",
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsPartitionIndexes, Ty: schema.TypeInt32CompactArray},
				&schema.SchemaTaggedFields{Name: FieldOffsetFetchRequestTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldOffsetFetchRequestTags},
		),

		// Message: OffsetFetchRequest, API Key: 9, Version: 7
		schema.NewSchema("OffsetFetchRequestv7",
			&schema.Mfield{Name: FieldOffsetFetchRequestGroupId, Ty: schema.TypeStrCompact},
			&schema.ArrayCompact{Name: FieldOffsetFetchRequestTopics, Ty: schema.NewSchema("TopicsV7",
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldOffsetFetchRequestTopicsPartitionIndexes, Ty: schema.TypeInt32CompactArray},
				&schema.SchemaTaggedFields{Name: FieldOffsetFetchRequestTopicsTags},
			)},
			&schema.Mfield{Name: FieldOffsetFetchRequestRequireStable, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldOffsetFetchRequestTags},
		),

		// Message: OffsetFetchRequest, API Key: 9, Version: 8
		schema.NewSchema("OffsetFetchRequestv8",
			&schema.ArrayCompact{Name: FieldOffsetFetchRequestGroups, Ty: schema.NewSchema("GroupsV8",
				&schema.Mfield{Name: FieldOffsetFetchRequestGroupsgroupId, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldOffsetFetchRequestGroupsTopics, Ty: schema.NewSchema("TopicsV8",
					&schema.Mfield{Name: FieldOffsetFetchRequestGroupsTopicsName, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldOffsetFetchRequestGroupsTopicsPartitionIndexes, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldOffsetFetchRequestGroupsTopicsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldOffsetFetchRequestGroupsTags},
			)},
			&schema.Mfield{Name: FieldOffsetFetchRequestRequireStable, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldOffsetFetchRequestTags},
		),
	}
}

const (
	// FieldOffsetFetchRequestGroupId is a field name that can be used to resolve the correct struct field.
	FieldOffsetFetchRequestGroupId = "GroupId"
	// FieldOffsetFetchRequestTopicsName is a field name that can be used to resolve the correct struct field.
	FieldOffsetFetchRequestTopicsName = "Name"
	// FieldOffsetFetchRequestTopicsPartitionIndexes is a field name that can be used to resolve the correct struct field.
	FieldOffsetFetchRequestTopicsPartitionIndexes = "PartitionIndexes"
	// FieldOffsetFetchRequestRequireStable is a field name that can be used to resolve the correct struct field.
	FieldOffsetFetchRequestRequireStable = "RequireStable"
	// FieldOffsetFetchRequestGroupsTags is a field name that can be used to resolve the correct struct field.
	FieldOffsetFetchRequestGroupsTags = "Tags"
	// FieldOffsetFetchRequestTopicsTags is a field name that can be used to resolve the correct struct field.
	FieldOffsetFetchRequestTopicsTags = "Tags"
	// FieldOffsetFetchRequestTags is a field name that can be used to resolve the correct struct field.
	FieldOffsetFetchRequestTags = "Tags"
	// FieldOffsetFetchRequestGroupsgroupId is a field name that can be used to resolve the correct struct field.
	FieldOffsetFetchRequestGroupsgroupId = "groupId"
	// FieldOffsetFetchRequestGroupsTopicsName is a field name that can be used to resolve the correct struct field.
	FieldOffsetFetchRequestGroupsTopicsName = "Name"
	// FieldOffsetFetchRequestGroupsTopicsPartitionIndexes is a field name that can be used to resolve the correct struct field.
	FieldOffsetFetchRequestGroupsTopicsPartitionIndexes = "PartitionIndexes"
	// FieldOffsetFetchRequestGroups is a field name that can be used to resolve the correct struct field.
	FieldOffsetFetchRequestGroups = "Groups"
	// FieldOffsetFetchRequestTopics is a field name that can be used to resolve the correct struct field.
	FieldOffsetFetchRequestTopics = "Topics"
	// FieldOffsetFetchRequestGroupsTopics is a field name that can be used to resolve the correct struct field.
	FieldOffsetFetchRequestGroupsTopics = "Topics"
	// FieldOffsetFetchRequestGroupsTopicsTags is a field name that can be used to resolve the correct struct field.
	FieldOffsetFetchRequestGroupsTopicsTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/OffsetFetchRequest.json
const originalOffsetFetchRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "OffsetFetchRequest",
  // In version 0, the request read offsets from ZK.
  //
  // Starting in version 1, the broker supports fetching offsets from the internal __consumer_offsets topic.
  //
  // Starting in version 2, the request can contain a null topics array to indicate that offsets
  // for all topics should be fetched. It also returns a top level error code
  // for group or coordinator level errors.
  //
  // Version 3, 4, and 5 are the same as version 2.
  //
  // Version 6 is the first flexible version.
  //
  // Version 7 is adding the require stable flag.
  //
  // Version 8 is adding support for fetching offsets for multiple groups at a time
  "validVersions": "0-8",
  "flexibleVersions": "6+",
  "fields": [
    { "name": "GroupId", "type": "string", "versions": "0-7", "entityType": "groupId",
      "about": "The group to fetch offsets for." },
    { "name": "Topics", "type": "[]OffsetFetchRequestTopic", "versions": "0-7", "nullableVersions": "2-7",
      "about": "Each topic we would like to fetch offsets for, or null to fetch offsets for all topics.", "fields": [
      { "name": "Name", "type": "string", "versions": "0-7", "entityType": "topicName",
        "about": "The topic name."},
      { "name": "PartitionIndexes", "type": "[]int32", "versions": "0-7",
        "about": "The partition indexes we would like to fetch offsets for." }
    ]},
    { "name": "Groups", "type": "[]OffsetFetchRequestGroup", "versions": "8+",
      "about": "Each group we would like to fetch offsets for", "fields": [
      { "name": "groupId", "type": "string", "versions": "8+", "entityType": "groupId",
        "about": "The group ID."},
      { "name": "Topics", "type": "[]OffsetFetchRequestTopics", "versions": "8+", "nullableVersions": "8+",
        "about": "Each topic we would like to fetch offsets for, or null to fetch offsets for all topics.", "fields": [
        { "name": "Name", "type": "string", "versions": "8+", "entityType": "topicName",
          "about": "The topic name."},
        { "name": "PartitionIndexes", "type": "[]int32", "versions": "8+",
          "about": "The partition indexes we would like to fetch offsets for." }
      ]}
    ]},
    {"name": "RequireStable", "type": "bool", "versions": "7+", "default": "false",
      "about": "Whether broker should hold on returning unstable offsets but set a retriable error code for the partitions."}
  ]
}
`
