package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init37CreatePartitionsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: CreatePartitionsRequest, API Key: 37, Version: 0
		schema.NewSchema("CreatePartitionsRequestv0",
			&schema.Array{Name: FieldCreatePartitionsRequestTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldCreatePartitionsRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreatePartitionsRequestTopicsCount, Ty: schema.TypeInt32},
				&schema.Array{Name: FieldCreatePartitionsRequestTopicsAssignments, Ty: schema.NewSchema("AssignmentsV0",
					&schema.Mfield{Name: FieldCreatePartitionsRequestTopicsAssignmentsBrokerIds, Ty: schema.TypeInt32Array},
				)},
			)},
			&schema.Mfield{Name: FieldCreatePartitionsRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldCreatePartitionsRequestValidateOnly, Ty: schema.TypeBool},
		),

		// Message: CreatePartitionsRequest, API Key: 37, Version: 1
		schema.NewSchema("CreatePartitionsRequestv1",
			&schema.Array{Name: FieldCreatePartitionsRequestTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldCreatePartitionsRequestTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreatePartitionsRequestTopicsCount, Ty: schema.TypeInt32},
				&schema.Array{Name: FieldCreatePartitionsRequestTopicsAssignments, Ty: schema.NewSchema("AssignmentsV1",
					&schema.Mfield{Name: FieldCreatePartitionsRequestTopicsAssignmentsBrokerIds, Ty: schema.TypeInt32Array},
				)},
			)},
			&schema.Mfield{Name: FieldCreatePartitionsRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldCreatePartitionsRequestValidateOnly, Ty: schema.TypeBool},
		),

		// Message: CreatePartitionsRequest, API Key: 37, Version: 2
		schema.NewSchema("CreatePartitionsRequestv2",
			&schema.ArrayCompact{Name: FieldCreatePartitionsRequestTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldCreatePartitionsRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldCreatePartitionsRequestTopicsCount, Ty: schema.TypeInt32},
				&schema.ArrayCompact{Name: FieldCreatePartitionsRequestTopicsAssignments, Ty: schema.NewSchema("AssignmentsV2",
					&schema.Mfield{Name: FieldCreatePartitionsRequestTopicsAssignmentsBrokerIds, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldCreatePartitionsRequestTopicsAssignmentsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldCreatePartitionsRequestTopicsTags},
			)},
			&schema.Mfield{Name: FieldCreatePartitionsRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldCreatePartitionsRequestValidateOnly, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldCreatePartitionsRequestTags},
		),

		// Message: CreatePartitionsRequest, API Key: 37, Version: 3
		schema.NewSchema("CreatePartitionsRequestv3",
			&schema.ArrayCompact{Name: FieldCreatePartitionsRequestTopics, Ty: schema.NewSchema("TopicsV3",
				&schema.Mfield{Name: FieldCreatePartitionsRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldCreatePartitionsRequestTopicsCount, Ty: schema.TypeInt32},
				&schema.ArrayCompact{Name: FieldCreatePartitionsRequestTopicsAssignments, Ty: schema.NewSchema("AssignmentsV3",
					&schema.Mfield{Name: FieldCreatePartitionsRequestTopicsAssignmentsBrokerIds, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldCreatePartitionsRequestTopicsAssignmentsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldCreatePartitionsRequestTopicsTags},
			)},
			&schema.Mfield{Name: FieldCreatePartitionsRequestTimeoutMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldCreatePartitionsRequestValidateOnly, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldCreatePartitionsRequestTags},
		),
	}
}

const (
	// FieldCreatePartitionsRequestTags is a field name that can be used to resolve the correct struct field.
	FieldCreatePartitionsRequestTags = "Tags"
	// FieldCreatePartitionsRequestTopics is a field name that can be used to resolve the correct struct field.
	FieldCreatePartitionsRequestTopics = "Topics"
	// FieldCreatePartitionsRequestTopicsName is a field name that can be used to resolve the correct struct field.
	FieldCreatePartitionsRequestTopicsName = "Name"
	// FieldCreatePartitionsRequestTimeoutMs is a field name that can be used to resolve the correct struct field.
	FieldCreatePartitionsRequestTimeoutMs = "TimeoutMs"
	// FieldCreatePartitionsRequestValidateOnly is a field name that can be used to resolve the correct struct field.
	FieldCreatePartitionsRequestValidateOnly = "ValidateOnly"
	// FieldCreatePartitionsRequestTopicsTags is a field name that can be used to resolve the correct struct field.
	FieldCreatePartitionsRequestTopicsTags = "Tags"
	// FieldCreatePartitionsRequestTopicsCount is a field name that can be used to resolve the correct struct field.
	FieldCreatePartitionsRequestTopicsCount = "Count"
	// FieldCreatePartitionsRequestTopicsAssignments is a field name that can be used to resolve the correct struct field.
	FieldCreatePartitionsRequestTopicsAssignments = "Assignments"
	// FieldCreatePartitionsRequestTopicsAssignmentsBrokerIds is a field name that can be used to resolve the correct struct field.
	FieldCreatePartitionsRequestTopicsAssignmentsBrokerIds = "BrokerIds"
	// FieldCreatePartitionsRequestTopicsAssignmentsTags is a field name that can be used to resolve the correct struct field.
	FieldCreatePartitionsRequestTopicsAssignmentsTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/CreatePartitionsRequest.json
const originalCreatePartitionsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 37,
  "type": "request",
  "listeners": ["zkBroker", "broker", "controller"],
  "name": "CreatePartitionsRequest",
  // Version 1 is the same as version 0.
  //
  // Version 2 adds flexible version support
  //
  // Version 3 is identical to version 2 but may return a THROTTLING_QUOTA_EXCEEDED error
  // in the response if the partitions creation is throttled (KIP-599).
  "validVersions": "0-3",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "Topics", "type": "[]CreatePartitionsTopic", "versions": "0+",
      "about": "Each topic that we want to create new partitions inside.",  "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "mapKey": true, "entityType": "topicName",
        "about": "The topic name." },
      { "name": "Count", "type": "int32", "versions": "0+",
        "about": "The new partition count." },
      { "name": "Assignments", "type": "[]CreatePartitionsAssignment", "versions": "0+", "nullableVersions": "0+", 
        "about": "The new partition assignments.", "fields": [
        { "name": "BrokerIds", "type": "[]int32", "versions": "0+", "entityType": "brokerId",
          "about": "The assigned broker IDs." }
      ]}
    ]},
    { "name": "TimeoutMs", "type": "int32", "versions": "0+",
      "about": "The time in ms to wait for the partitions to be created." },
    { "name": "ValidateOnly", "type": "bool", "versions": "0+",
      "about": "If true, then validate the request, but don't actually increase the number of partitions." }
  ]
}
`
