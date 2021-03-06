package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init19CreateTopicsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: CreateTopicsRequest, API Key: 19, Version: 0
		schema.NewSchema("CreateTopicsRequestv0",
			&schema.Array{Name: FieldCreateTopicsRequestTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsNumPartitions, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsReplicationFactor, Ty: schema.TypeInt16},
				&schema.Array{Name: FieldCreateTopicsRequestTopicsAssignments, Ty: schema.NewSchema("AssignmentsV0",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsBrokerIds, Ty: schema.TypeInt32Array},
				)},
				&schema.Array{Name: FieldCreateTopicsRequestTopicsConfigs, Ty: schema.NewSchema("ConfigsV0",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsValue, Ty: schema.TypeStrNullable},
				)},
			)},
			&schema.Mfield{Name: FieldCreateTopicsRequesttimeoutMs, Ty: schema.TypeInt32},
		),

		// Message: CreateTopicsRequest, API Key: 19, Version: 1
		schema.NewSchema("CreateTopicsRequestv1",
			&schema.Array{Name: FieldCreateTopicsRequestTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsNumPartitions, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsReplicationFactor, Ty: schema.TypeInt16},
				&schema.Array{Name: FieldCreateTopicsRequestTopicsAssignments, Ty: schema.NewSchema("AssignmentsV1",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsBrokerIds, Ty: schema.TypeInt32Array},
				)},
				&schema.Array{Name: FieldCreateTopicsRequestTopicsConfigs, Ty: schema.NewSchema("ConfigsV1",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsValue, Ty: schema.TypeStrNullable},
				)},
			)},
			&schema.Mfield{Name: FieldCreateTopicsRequesttimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldCreateTopicsRequestvalidateOnly, Ty: schema.TypeBool},
		),

		// Message: CreateTopicsRequest, API Key: 19, Version: 2
		schema.NewSchema("CreateTopicsRequestv2",
			&schema.Array{Name: FieldCreateTopicsRequestTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsNumPartitions, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsReplicationFactor, Ty: schema.TypeInt16},
				&schema.Array{Name: FieldCreateTopicsRequestTopicsAssignments, Ty: schema.NewSchema("AssignmentsV2",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsBrokerIds, Ty: schema.TypeInt32Array},
				)},
				&schema.Array{Name: FieldCreateTopicsRequestTopicsConfigs, Ty: schema.NewSchema("ConfigsV2",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsValue, Ty: schema.TypeStrNullable},
				)},
			)},
			&schema.Mfield{Name: FieldCreateTopicsRequesttimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldCreateTopicsRequestvalidateOnly, Ty: schema.TypeBool},
		),

		// Message: CreateTopicsRequest, API Key: 19, Version: 3
		schema.NewSchema("CreateTopicsRequestv3",
			&schema.Array{Name: FieldCreateTopicsRequestTopics, Ty: schema.NewSchema("TopicsV3",
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsNumPartitions, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsReplicationFactor, Ty: schema.TypeInt16},
				&schema.Array{Name: FieldCreateTopicsRequestTopicsAssignments, Ty: schema.NewSchema("AssignmentsV3",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsBrokerIds, Ty: schema.TypeInt32Array},
				)},
				&schema.Array{Name: FieldCreateTopicsRequestTopicsConfigs, Ty: schema.NewSchema("ConfigsV3",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsValue, Ty: schema.TypeStrNullable},
				)},
			)},
			&schema.Mfield{Name: FieldCreateTopicsRequesttimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldCreateTopicsRequestvalidateOnly, Ty: schema.TypeBool},
		),

		// Message: CreateTopicsRequest, API Key: 19, Version: 4
		schema.NewSchema("CreateTopicsRequestv4",
			&schema.Array{Name: FieldCreateTopicsRequestTopics, Ty: schema.NewSchema("TopicsV4",
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsNumPartitions, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsReplicationFactor, Ty: schema.TypeInt16},
				&schema.Array{Name: FieldCreateTopicsRequestTopicsAssignments, Ty: schema.NewSchema("AssignmentsV4",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsBrokerIds, Ty: schema.TypeInt32Array},
				)},
				&schema.Array{Name: FieldCreateTopicsRequestTopicsConfigs, Ty: schema.NewSchema("ConfigsV4",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsValue, Ty: schema.TypeStrNullable},
				)},
			)},
			&schema.Mfield{Name: FieldCreateTopicsRequesttimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldCreateTopicsRequestvalidateOnly, Ty: schema.TypeBool},
		),

		// Message: CreateTopicsRequest, API Key: 19, Version: 5
		schema.NewSchema("CreateTopicsRequestv5",
			&schema.ArrayCompact{Name: FieldCreateTopicsRequestTopics, Ty: schema.NewSchema("TopicsV5",
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsNumPartitions, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsReplicationFactor, Ty: schema.TypeInt16},
				&schema.ArrayCompact{Name: FieldCreateTopicsRequestTopicsAssignments, Ty: schema.NewSchema("AssignmentsV5",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsBrokerIds, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldCreateTopicsRequestTopicsAssignmentsTags},
				)},
				&schema.ArrayCompact{Name: FieldCreateTopicsRequestTopicsConfigs, Ty: schema.NewSchema("ConfigsV5",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsName, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsValue, Ty: schema.TypeStrCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldCreateTopicsRequestTopicsConfigsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldCreateTopicsRequestTopicsTags},
			)},
			&schema.Mfield{Name: FieldCreateTopicsRequesttimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldCreateTopicsRequestvalidateOnly, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldCreateTopicsRequestTags},
		),

		// Message: CreateTopicsRequest, API Key: 19, Version: 6
		schema.NewSchema("CreateTopicsRequestv6",
			&schema.ArrayCompact{Name: FieldCreateTopicsRequestTopics, Ty: schema.NewSchema("TopicsV6",
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsNumPartitions, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsReplicationFactor, Ty: schema.TypeInt16},
				&schema.ArrayCompact{Name: FieldCreateTopicsRequestTopicsAssignments, Ty: schema.NewSchema("AssignmentsV6",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsBrokerIds, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldCreateTopicsRequestTopicsAssignmentsTags},
				)},
				&schema.ArrayCompact{Name: FieldCreateTopicsRequestTopicsConfigs, Ty: schema.NewSchema("ConfigsV6",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsName, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsValue, Ty: schema.TypeStrCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldCreateTopicsRequestTopicsConfigsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldCreateTopicsRequestTopicsTags},
			)},
			&schema.Mfield{Name: FieldCreateTopicsRequesttimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldCreateTopicsRequestvalidateOnly, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldCreateTopicsRequestTags},
		),

		// Message: CreateTopicsRequest, API Key: 19, Version: 7
		schema.NewSchema("CreateTopicsRequestv7",
			&schema.ArrayCompact{Name: FieldCreateTopicsRequestTopics, Ty: schema.NewSchema("TopicsV7",
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsNumPartitions, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldCreateTopicsRequestTopicsReplicationFactor, Ty: schema.TypeInt16},
				&schema.ArrayCompact{Name: FieldCreateTopicsRequestTopicsAssignments, Ty: schema.NewSchema("AssignmentsV7",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsAssignmentsBrokerIds, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldCreateTopicsRequestTopicsAssignmentsTags},
				)},
				&schema.ArrayCompact{Name: FieldCreateTopicsRequestTopicsConfigs, Ty: schema.NewSchema("ConfigsV7",
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsName, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldCreateTopicsRequestTopicsConfigsValue, Ty: schema.TypeStrCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldCreateTopicsRequestTopicsConfigsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldCreateTopicsRequestTopicsTags},
			)},
			&schema.Mfield{Name: FieldCreateTopicsRequesttimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldCreateTopicsRequestvalidateOnly, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldCreateTopicsRequestTags},
		),
	}
}

const (
	// FieldCreateTopicsRequestTags is: The tagged fields.
	FieldCreateTopicsRequestTags = "Tags"
	// FieldCreateTopicsRequestTopics is: The topics to create.
	FieldCreateTopicsRequestTopics = "Topics"
	// FieldCreateTopicsRequestTopicsAssignments is: The manual partition assignment, or the empty array if we are using automatic assignment.
	FieldCreateTopicsRequestTopicsAssignments = "Assignments"
	// FieldCreateTopicsRequestTopicsAssignmentsBrokerIds is: The brokers to place the partition on.
	FieldCreateTopicsRequestTopicsAssignmentsBrokerIds = "BrokerIds"
	// FieldCreateTopicsRequestTopicsAssignmentsPartitionIndex is: The partition index.
	FieldCreateTopicsRequestTopicsAssignmentsPartitionIndex = "PartitionIndex"
	// FieldCreateTopicsRequestTopicsAssignmentsTags is: The tagged fields.
	FieldCreateTopicsRequestTopicsAssignmentsTags = "Tags"
	// FieldCreateTopicsRequestTopicsConfigs is: The custom topic configurations to set.
	FieldCreateTopicsRequestTopicsConfigs = "Configs"
	// FieldCreateTopicsRequestTopicsConfigsName is: The configuration name.
	FieldCreateTopicsRequestTopicsConfigsName = "Name"
	// FieldCreateTopicsRequestTopicsConfigsTags is: The tagged fields.
	FieldCreateTopicsRequestTopicsConfigsTags = "Tags"
	// FieldCreateTopicsRequestTopicsConfigsValue is: The configuration value.
	FieldCreateTopicsRequestTopicsConfigsValue = "Value"
	// FieldCreateTopicsRequestTopicsName is: The topic name.
	FieldCreateTopicsRequestTopicsName = "Name"
	// FieldCreateTopicsRequestTopicsNumPartitions is: The number of partitions to create in the topic, or -1 if we are either specifying a manual partition assignment or using the default partitions.
	FieldCreateTopicsRequestTopicsNumPartitions = "NumPartitions"
	// FieldCreateTopicsRequestTopicsReplicationFactor is: The number of replicas to create for each partition in the topic, or -1 if we are either specifying a manual partition assignment or using the default replication factor.
	FieldCreateTopicsRequestTopicsReplicationFactor = "ReplicationFactor"
	// FieldCreateTopicsRequestTopicsTags is: The tagged fields.
	FieldCreateTopicsRequestTopicsTags = "Tags"
	// FieldCreateTopicsRequesttimeoutMs is: How long to wait in milliseconds before timing out the request.
	FieldCreateTopicsRequesttimeoutMs = "timeoutMs"
	// FieldCreateTopicsRequestvalidateOnly is: If true, check that the topics can be created as specified, but don't create anything.
	FieldCreateTopicsRequestvalidateOnly = "validateOnly"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/CreateTopicsRequest.json
const originalCreateTopicsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 19,
  "type": "request",
  "listeners": ["zkBroker", "broker", "controller"],
  "name": "CreateTopicsRequest",
  // Version 1 adds validateOnly.
  //
  // Version 4 makes partitions/replicationFactor optional even when assignments are not present (KIP-464)
  //
  // Version 5 is the first flexible version.
  // Version 5 also returns topic configs in the response (KIP-525).
  //
  // Version 6 is identical to version 5 but may return a THROTTLING_QUOTA_EXCEEDED error
  // in the response if the topics creation is throttled (KIP-599).
  //
  // Version 7 is the same as version 6.
  "validVersions": "0-7",
  "flexibleVersions": "5+",
  "fields": [
    { "name": "Topics", "type": "[]CreatableTopic", "versions": "0+",
      "about": "The topics to create.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "mapKey": true, "entityType": "topicName",
        "about": "The topic name." },
      { "name": "NumPartitions", "type": "int32", "versions": "0+",
        "about": "The number of partitions to create in the topic, or -1 if we are either specifying a manual partition assignment or using the default partitions." },
      { "name": "ReplicationFactor", "type": "int16", "versions": "0+",
        "about": "The number of replicas to create for each partition in the topic, or -1 if we are either specifying a manual partition assignment or using the default replication factor." },
      { "name": "Assignments", "type": "[]CreatableReplicaAssignment", "versions": "0+",
        "about": "The manual partition assignment, or the empty array if we are using automatic assignment.", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+", "mapKey": true,
          "about": "The partition index." },
        { "name": "BrokerIds", "type": "[]int32", "versions": "0+", "entityType": "brokerId",
          "about": "The brokers to place the partition on." }
      ]},
      { "name": "Configs", "type": "[]CreateableTopicConfig", "versions": "0+",
        "about": "The custom topic configurations to set.", "fields": [
        { "name": "Name", "type": "string", "versions": "0+" , "mapKey": true,
          "about": "The configuration name." },
        { "name": "Value", "type": "string", "versions": "0+", "nullableVersions": "0+",
          "about": "The configuration value." }
      ]}
    ]},
    { "name": "timeoutMs", "type": "int32", "versions": "0+", "default": "60000",
      "about": "How long to wait in milliseconds before timing out the request." },
    { "name": "validateOnly", "type": "bool", "versions": "1+", "default": "false", "ignorable": false,
      "about": "If true, check that the topics can be created as specified, but don't create anything." }
  ]
}
`
