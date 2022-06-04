package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init61DescribeProducersRequest() []schema.Schema {

	return []schema.Schema{

		// Message: DescribeProducersRequest, API Key: 61, Version: 0
		schema.NewSchema("DescribeProducersRequestv0",
			&schema.ArrayCompact{Name: FieldDescribeProducersRequestTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldDescribeProducersRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeProducersRequestTopicsPartitionIndexes, Ty: schema.TypeInt32CompactArray},
				&schema.SchemaTaggedFields{Name: FieldDescribeProducersRequestTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDescribeProducersRequestTags},
		),
	}
}

const (
	// FieldDescribeProducersRequestTags is: The tagged fields.
	FieldDescribeProducersRequestTags = "Tags"
	// FieldDescribeProducersRequestTopics is:
	FieldDescribeProducersRequestTopics = "Topics"
	// FieldDescribeProducersRequestTopicsName is: The topic name.
	FieldDescribeProducersRequestTopicsName = "Name"
	// FieldDescribeProducersRequestTopicsPartitionIndexes is: The indexes of the partitions to list producers for.
	FieldDescribeProducersRequestTopicsPartitionIndexes = "PartitionIndexes"
	// FieldDescribeProducersRequestTopicsTags is: The tagged fields.
	FieldDescribeProducersRequestTopicsTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeProducersRequest.json
const originalDescribeProducersRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 61,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "DescribeProducersRequest",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "Topics", "type": "[]TopicRequest", "versions": "0+", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The topic name." },
      { "name": "PartitionIndexes", "type": "[]int32", "versions": "0+",
        "about": "The indexes of the partitions to list producers for." }
     ]}
  ]
}
`
