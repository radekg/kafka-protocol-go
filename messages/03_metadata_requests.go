package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init3MetadataRequest() []schema.Schema {

	return []schema.Schema{

		// Message: MetadataRequest, API Key: 3, Version: 0
		schema.NewSchema("MetadataRequestv0", 
			&schema.Array{Name: FieldMetadataRequestTopics, Ty: schema.NewSchema("TopicsV0",
				&schema.Mfield{Name: FieldMetadataRequestTopicsName, Ty: schema.TypeStr},
			)},
		),

		// Message: MetadataRequest, API Key: 3, Version: 1
		schema.NewSchema("MetadataRequestv1", 
			&schema.Array{Name: FieldMetadataRequestTopics, Ty: schema.NewSchema("TopicsV1",
				&schema.Mfield{Name: FieldMetadataRequestTopicsName, Ty: schema.TypeStr},
			)},
		),

		// Message: MetadataRequest, API Key: 3, Version: 2
		schema.NewSchema("MetadataRequestv2", 
			&schema.Array{Name: FieldMetadataRequestTopics, Ty: schema.NewSchema("TopicsV2",
				&schema.Mfield{Name: FieldMetadataRequestTopicsName, Ty: schema.TypeStr},
			)},
		),

		// Message: MetadataRequest, API Key: 3, Version: 3
		schema.NewSchema("MetadataRequestv3", 
			&schema.Array{Name: FieldMetadataRequestTopics, Ty: schema.NewSchema("TopicsV3",
				&schema.Mfield{Name: FieldMetadataRequestTopicsName, Ty: schema.TypeStr},
			)},
		),

		// Message: MetadataRequest, API Key: 3, Version: 4
		schema.NewSchema("MetadataRequestv4", 
			&schema.Array{Name: FieldMetadataRequestTopics, Ty: schema.NewSchema("TopicsV4",
				&schema.Mfield{Name: FieldMetadataRequestTopicsName, Ty: schema.TypeStr},
			)},
			&schema.Mfield{Name: FieldMetadataRequestAllowAutoTopicCreation, Ty: schema.TypeBool},
		),

		// Message: MetadataRequest, API Key: 3, Version: 5
		schema.NewSchema("MetadataRequestv5", 
			&schema.Array{Name: FieldMetadataRequestTopics, Ty: schema.NewSchema("TopicsV5",
				&schema.Mfield{Name: FieldMetadataRequestTopicsName, Ty: schema.TypeStr},
			)},
			&schema.Mfield{Name: FieldMetadataRequestAllowAutoTopicCreation, Ty: schema.TypeBool},
		),

		// Message: MetadataRequest, API Key: 3, Version: 6
		schema.NewSchema("MetadataRequestv6", 
			&schema.Array{Name: FieldMetadataRequestTopics, Ty: schema.NewSchema("TopicsV6",
				&schema.Mfield{Name: FieldMetadataRequestTopicsName, Ty: schema.TypeStr},
			)},
			&schema.Mfield{Name: FieldMetadataRequestAllowAutoTopicCreation, Ty: schema.TypeBool},
		),

		// Message: MetadataRequest, API Key: 3, Version: 7
		schema.NewSchema("MetadataRequestv7", 
			&schema.Array{Name: FieldMetadataRequestTopics, Ty: schema.NewSchema("TopicsV7",
				&schema.Mfield{Name: FieldMetadataRequestTopicsName, Ty: schema.TypeStr},
			)},
			&schema.Mfield{Name: FieldMetadataRequestAllowAutoTopicCreation, Ty: schema.TypeBool},
		),

		// Message: MetadataRequest, API Key: 3, Version: 8
		schema.NewSchema("MetadataRequestv8", 
			&schema.Array{Name: FieldMetadataRequestTopics, Ty: schema.NewSchema("TopicsV8",
				&schema.Mfield{Name: FieldMetadataRequestTopicsName, Ty: schema.TypeStr},
			)},
			&schema.Mfield{Name: FieldMetadataRequestAllowAutoTopicCreation, Ty: schema.TypeBool},
			&schema.Mfield{Name: FieldMetadataRequestIncludeClusterAuthorizedOperations, Ty: schema.TypeBool},
			&schema.Mfield{Name: FieldMetadataRequestIncludeTopicAuthorizedOperations, Ty: schema.TypeBool},
		),

		// Message: MetadataRequest, API Key: 3, Version: 9
		schema.NewSchema("MetadataRequestv9", 
			&schema.ArrayCompact{Name: FieldMetadataRequestTopics, Ty: schema.NewSchema("TopicsV9",
				&schema.Mfield{Name: FieldMetadataRequestTopicsName, Ty: schema.TypeStrCompact},
				&schema.SchemaTaggedFields{Name: FieldMetadataRequestTopicsTags},
			)},
			&schema.Mfield{Name: FieldMetadataRequestAllowAutoTopicCreation, Ty: schema.TypeBool},
			&schema.Mfield{Name: FieldMetadataRequestIncludeClusterAuthorizedOperations, Ty: schema.TypeBool},
			&schema.Mfield{Name: FieldMetadataRequestIncludeTopicAuthorizedOperations, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldMetadataRequestTags},
		),

		// Message: MetadataRequest, API Key: 3, Version: 10
		schema.NewSchema("MetadataRequestv10", 
			&schema.ArrayCompact{Name: FieldMetadataRequestTopics, Ty: schema.NewSchema("TopicsV10",
				&schema.Mfield{Name: FieldMetadataRequestTopicsTopicId, Ty: schema.TypeUuid},
				&schema.Mfield{Name: FieldMetadataRequestTopicsName, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldMetadataRequestTopicsTags},
			)},
			&schema.Mfield{Name: FieldMetadataRequestAllowAutoTopicCreation, Ty: schema.TypeBool},
			&schema.Mfield{Name: FieldMetadataRequestIncludeClusterAuthorizedOperations, Ty: schema.TypeBool},
			&schema.Mfield{Name: FieldMetadataRequestIncludeTopicAuthorizedOperations, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldMetadataRequestTags},
		),

		// Message: MetadataRequest, API Key: 3, Version: 11
		schema.NewSchema("MetadataRequestv11", 
			&schema.ArrayCompact{Name: FieldMetadataRequestTopics, Ty: schema.NewSchema("TopicsV11",
				&schema.Mfield{Name: FieldMetadataRequestTopicsTopicId, Ty: schema.TypeUuid},
				&schema.Mfield{Name: FieldMetadataRequestTopicsName, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldMetadataRequestTopicsTags},
			)},
			&schema.Mfield{Name: FieldMetadataRequestAllowAutoTopicCreation, Ty: schema.TypeBool},
			&schema.Mfield{Name: FieldMetadataRequestIncludeTopicAuthorizedOperations, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldMetadataRequestTags},
		),

		// Message: MetadataRequest, API Key: 3, Version: 12
		schema.NewSchema("MetadataRequestv12", 
			&schema.ArrayCompact{Name: FieldMetadataRequestTopics, Ty: schema.NewSchema("TopicsV12",
				&schema.Mfield{Name: FieldMetadataRequestTopicsTopicId, Ty: schema.TypeUuid},
				&schema.Mfield{Name: FieldMetadataRequestTopicsName, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldMetadataRequestTopicsTags},
			)},
			&schema.Mfield{Name: FieldMetadataRequestAllowAutoTopicCreation, Ty: schema.TypeBool},
			&schema.Mfield{Name: FieldMetadataRequestIncludeTopicAuthorizedOperations, Ty: schema.TypeBool},
			&schema.SchemaTaggedFields{Name: FieldMetadataRequestTags},
		),

	}
}

const (
	// FieldMetadataRequestIncludeTopicAuthorizedOperations is a field name that can be used to resolve the correct struct field.
	FieldMetadataRequestIncludeTopicAuthorizedOperations = "IncludeTopicAuthorizedOperations"
	// FieldMetadataRequestTopicsTags is a field name that can be used to resolve the correct struct field.
	FieldMetadataRequestTopicsTags = "Tags"
	// FieldMetadataRequestTags is a field name that can be used to resolve the correct struct field.
	FieldMetadataRequestTags = "Tags"
	// FieldMetadataRequestTopicsTopicId is a field name that can be used to resolve the correct struct field.
	FieldMetadataRequestTopicsTopicId = "TopicId"
	// FieldMetadataRequestTopics is a field name that can be used to resolve the correct struct field.
	FieldMetadataRequestTopics = "Topics"
	// FieldMetadataRequestTopicsName is a field name that can be used to resolve the correct struct field.
	FieldMetadataRequestTopicsName = "Name"
	// FieldMetadataRequestAllowAutoTopicCreation is a field name that can be used to resolve the correct struct field.
	FieldMetadataRequestAllowAutoTopicCreation = "AllowAutoTopicCreation"
	// FieldMetadataRequestIncludeClusterAuthorizedOperations is a field name that can be used to resolve the correct struct field.
	FieldMetadataRequestIncludeClusterAuthorizedOperations = "IncludeClusterAuthorizedOperations"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/MetadataRequest.json
const originalMetadataRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 3,
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "MetadataRequest",
  "validVersions": "0-12",
  "flexibleVersions": "9+",
  "fields": [
    // In version 0, an empty array indicates "request metadata for all topics."  In version 1 and
    // higher, an empty array indicates "request metadata for no topics," and a null array is used to
    // indiate "request metadata for all topics."
    //
    // Version 2 and 3 are the same as version 1.
    //
    // Version 4 adds AllowAutoTopicCreation.
    //
    // Starting in version 8, authorized operations can be requested for cluster and topic resource.
    //
    // Version 9 is the first flexible version.
    //
    // Version 10 adds topicId and allows name field to be null. However, this functionality was not implemented on the server.
    // Versions 10 and 11 should not use the topicId field or set topic name to null.
    //
    // Version 11 deprecates IncludeClusterAuthorizedOperations field. This is now exposed
    // by the DescribeCluster API (KIP-700).
    // Version 12 supports topic Id.
    { "name": "Topics", "type": "[]MetadataRequestTopic", "versions": "0+", "nullableVersions": "1+",
      "about": "The topics to fetch metadata for.", "fields": [
      { "name": "TopicId", "type": "uuid", "versions": "10+", "ignorable": true, "about": "The topic id." },
      { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName", "nullableVersions": "10+",
        "about": "The topic name." }
    ]},
    { "name": "AllowAutoTopicCreation", "type": "bool", "versions": "4+", "default": "true", "ignorable": false,
      "about": "If this is true, the broker may auto-create topics that we requested which do not already exist, if it is configured to do so." },
    { "name": "IncludeClusterAuthorizedOperations", "type": "bool", "versions": "8-10",
      "about": "Whether to include cluster authorized operations." },
    { "name": "IncludeTopicAuthorizedOperations", "type": "bool", "versions": "8+",
      "about": "Whether to include topic authorized operations." }
  ]
}
`

