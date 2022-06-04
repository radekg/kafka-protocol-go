package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init35DescribeLogDirsRequest() []schema.Schema {

	return []schema.Schema{

		// Message: DescribeLogDirsRequest, API Key: 35, Version: 0
		schema.NewSchema("DescribeLogDirsRequestv0",
			&schema.Array{Name: FieldDescribeLogDirsRequestTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldDescribeLogDirsRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeLogDirsRequestTopicsPartitions, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: DescribeLogDirsRequest, API Key: 35, Version: 1
		schema.NewSchema("DescribeLogDirsRequestv1",
			&schema.Array{Name: FieldDescribeLogDirsRequestTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldDescribeLogDirsRequestTopicsTopic, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldDescribeLogDirsRequestTopicsPartitions, Ty: schema.TypeInt32Array},
			)},
		),

		// Message: DescribeLogDirsRequest, API Key: 35, Version: 2
		schema.NewSchema("DescribeLogDirsRequestv2",
			&schema.ArrayCompact{Name: FieldDescribeLogDirsRequestTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldDescribeLogDirsRequestTopicsTopic, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeLogDirsRequestTopicsPartitions, Ty: schema.TypeInt32CompactArray},
				&schema.SchemaTaggedFields{Name: FieldDescribeLogDirsRequestTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDescribeLogDirsRequestTags},
		),

		// Message: DescribeLogDirsRequest, API Key: 35, Version: 3
		schema.NewSchema("DescribeLogDirsRequestv3",
			&schema.ArrayCompact{Name: FieldDescribeLogDirsRequestTopics, Ty: schema.NewSchema("TopicsV3",
				&schema.Mfield{Name: FieldDescribeLogDirsRequestTopicsTopic, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeLogDirsRequestTopicsPartitions, Ty: schema.TypeInt32CompactArray},
				&schema.SchemaTaggedFields{Name: FieldDescribeLogDirsRequestTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDescribeLogDirsRequestTags},
		),
	}
}

const (
	// FieldDescribeLogDirsRequestTopics is: Each topic that we want to describe log directories for, or null for all topics.
	FieldDescribeLogDirsRequestTopics = "Topics"
	// FieldDescribeLogDirsRequestTopicsTopic is: The topic name
	FieldDescribeLogDirsRequestTopicsTopic = "Topic"
	// FieldDescribeLogDirsRequestTopicsPartitions is: The partition indexes.
	FieldDescribeLogDirsRequestTopicsPartitions = "Partitions"
	// FieldDescribeLogDirsRequestTopicsTags is: The tagged fields.
	FieldDescribeLogDirsRequestTopicsTags = "Tags"
	// FieldDescribeLogDirsRequestTags is: The tagged fields.
	FieldDescribeLogDirsRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeLogDirsRequest.json
const originalDescribeLogDirsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 35,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "DescribeLogDirsRequest",
  // Version 1 is the same as version 0.
  "validVersions": "0-3",
  // Version 2 is the first flexible version.
  // Version 3 is the same as version 2 (new field in response).
  "flexibleVersions": "2+",
  "fields": [
    { "name": "Topics", "type": "[]DescribableLogDirTopic", "versions": "0+", "nullableVersions": "0+",
      "about": "Each topic that we want to describe log directories for, or null for all topics.", "fields": [
      { "name": "Topic", "type": "string", "versions": "0+", "entityType": "topicName", "mapKey": true,
        "about": "The topic name" },
      { "name": "Partitions", "type": "[]int32", "versions": "0+",
        "about": "The partition indexes." }
    ]}
  ]
}
`
