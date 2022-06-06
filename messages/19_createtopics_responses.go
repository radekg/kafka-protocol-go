package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init19CreateTopicsResponse() []schema.Schema {

	return []schema.Schema{
		// Message: CreateTopicsResponse, API Key: 19, Version: 0
		schema.NewSchema("CreateTopicsResponse:v0",
			&schema.Array{Name: FieldCreateTopicsResponseTopics, Ty: schema.NewSchema("Topics:v0",
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsErrorCode, Ty: schema.TypeInt16},
			)},
		),

		// Message: CreateTopicsResponse, API Key: 19, Version: 1
		schema.NewSchema("CreateTopicsResponse:v1",
			&schema.Array{Name: FieldCreateTopicsResponseTopics, Ty: schema.NewSchema("Topics:v1",
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsErrorMessage, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: CreateTopicsResponse, API Key: 19, Version: 2
		schema.NewSchema("CreateTopicsResponse:v2",
			&schema.Mfield{Name: FieldCreateTopicsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldCreateTopicsResponseTopics, Ty: schema.NewSchema("Topics:v2",
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsErrorMessage, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: CreateTopicsResponse, API Key: 19, Version: 3
		schema.NewSchema("CreateTopicsResponse:v3",
			&schema.Mfield{Name: FieldCreateTopicsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldCreateTopicsResponseTopics, Ty: schema.NewSchema("Topics:v3",
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsErrorMessage, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: CreateTopicsResponse, API Key: 19, Version: 4
		schema.NewSchema("CreateTopicsResponse:v4",
			&schema.Mfield{Name: FieldCreateTopicsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldCreateTopicsResponseTopics, Ty: schema.NewSchema("Topics:v4",
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsName, Ty: schema.TypeStr},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsErrorMessage, Ty: schema.TypeStrNullable},
			)},
		),

		// Message: CreateTopicsResponse, API Key: 19, Version: 5
		schema.NewSchema("CreateTopicsResponse:v5",
			&schema.Mfield{Name: FieldCreateTopicsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldCreateTopicsResponseTopics, Ty: schema.NewSchema("Topics:v5",
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsTopicConfigErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsNumPartitions, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsReplicationFactor, Ty: schema.TypeInt16},
				&schema.ArrayCompact{Name: FieldCreateTopicsResponseTopicsConfigs, Ty: schema.NewSchema("Configs:v5",
					&schema.Mfield{Name: FieldCreateTopicsResponseTopicsConfigsName, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldCreateTopicsResponseTopicsConfigsValue, Ty: schema.TypeStrCompactNullable},
					&schema.Mfield{Name: FieldCreateTopicsResponseTopicsConfigsReadOnly, Ty: schema.TypeBool},
					&schema.Mfield{Name: FieldCreateTopicsResponseTopicsConfigsConfigSource, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldCreateTopicsResponseTopicsConfigsIsSensitive, Ty: schema.TypeBool},
					&schema.SchemaTaggedFields{Name: FieldCreateTopicsResponseTopicsConfigsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldCreateTopicsResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldCreateTopicsResponseTags},
		),

		// Message: CreateTopicsResponse, API Key: 19, Version: 6
		schema.NewSchema("CreateTopicsResponse:v6",
			&schema.Mfield{Name: FieldCreateTopicsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldCreateTopicsResponseTopics, Ty: schema.NewSchema("Topics:v6",
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsTopicConfigErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsNumPartitions, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsReplicationFactor, Ty: schema.TypeInt16},
				&schema.ArrayCompact{Name: FieldCreateTopicsResponseTopicsConfigs, Ty: schema.NewSchema("Configs:v6",
					&schema.Mfield{Name: FieldCreateTopicsResponseTopicsConfigsName, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldCreateTopicsResponseTopicsConfigsValue, Ty: schema.TypeStrCompactNullable},
					&schema.Mfield{Name: FieldCreateTopicsResponseTopicsConfigsReadOnly, Ty: schema.TypeBool},
					&schema.Mfield{Name: FieldCreateTopicsResponseTopicsConfigsConfigSource, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldCreateTopicsResponseTopicsConfigsIsSensitive, Ty: schema.TypeBool},
					&schema.SchemaTaggedFields{Name: FieldCreateTopicsResponseTopicsConfigsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldCreateTopicsResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldCreateTopicsResponseTags},
		),

		// Message: CreateTopicsResponse, API Key: 19, Version: 7
		schema.NewSchema("CreateTopicsResponse:v7",
			&schema.Mfield{Name: FieldCreateTopicsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldCreateTopicsResponseTopics, Ty: schema.NewSchema("Topics:v7",
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsName, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsTopicId, Ty: schema.TypeUuid},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsTopicConfigErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsNumPartitions, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldCreateTopicsResponseTopicsReplicationFactor, Ty: schema.TypeInt16},
				&schema.ArrayCompact{Name: FieldCreateTopicsResponseTopicsConfigs, Ty: schema.NewSchema("Configs:v7",
					&schema.Mfield{Name: FieldCreateTopicsResponseTopicsConfigsName, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldCreateTopicsResponseTopicsConfigsValue, Ty: schema.TypeStrCompactNullable},
					&schema.Mfield{Name: FieldCreateTopicsResponseTopicsConfigsReadOnly, Ty: schema.TypeBool},
					&schema.Mfield{Name: FieldCreateTopicsResponseTopicsConfigsConfigSource, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldCreateTopicsResponseTopicsConfigsIsSensitive, Ty: schema.TypeBool},
					&schema.SchemaTaggedFields{Name: FieldCreateTopicsResponseTopicsConfigsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldCreateTopicsResponseTopicsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldCreateTopicsResponseTags},
		),
	}

}

const (

	// FieldCreateTopicsResponseTags is: The tagged fields.
	FieldCreateTopicsResponseTags = "Tags"

	// FieldCreateTopicsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldCreateTopicsResponseThrottleTimeMs = "ThrottleTimeMs"

	// FieldCreateTopicsResponseTopics is: Results for each topic we tried to create.
	FieldCreateTopicsResponseTopics = "Topics"

	// FieldCreateTopicsResponseTopicsConfigs is: Configuration of the topic.
	FieldCreateTopicsResponseTopicsConfigs = "Configs"

	// FieldCreateTopicsResponseTopicsConfigsConfigSource is: The configuration source.
	FieldCreateTopicsResponseTopicsConfigsConfigSource = "ConfigSource"

	// FieldCreateTopicsResponseTopicsConfigsIsSensitive is: True if this configuration is sensitive.
	FieldCreateTopicsResponseTopicsConfigsIsSensitive = "IsSensitive"

	// FieldCreateTopicsResponseTopicsConfigsName is: The configuration name.
	FieldCreateTopicsResponseTopicsConfigsName = "Name"

	// FieldCreateTopicsResponseTopicsConfigsReadOnly is: True if the configuration is read-only.
	FieldCreateTopicsResponseTopicsConfigsReadOnly = "ReadOnly"

	// FieldCreateTopicsResponseTopicsConfigsTags is: The tagged fields.
	FieldCreateTopicsResponseTopicsConfigsTags = "Tags"

	// FieldCreateTopicsResponseTopicsConfigsValue is: The configuration value.
	FieldCreateTopicsResponseTopicsConfigsValue = "Value"

	// FieldCreateTopicsResponseTopicsErrorCode is: The error code, or 0 if there was no error.
	FieldCreateTopicsResponseTopicsErrorCode = "ErrorCode"

	// FieldCreateTopicsResponseTopicsErrorMessage is: The error message, or null if there was no error.
	FieldCreateTopicsResponseTopicsErrorMessage = "ErrorMessage"

	// FieldCreateTopicsResponseTopicsName is: The topic name.
	FieldCreateTopicsResponseTopicsName = "Name"

	// FieldCreateTopicsResponseTopicsNumPartitions is: Number of partitions of the topic.
	FieldCreateTopicsResponseTopicsNumPartitions = "NumPartitions"

	// FieldCreateTopicsResponseTopicsReplicationFactor is: Replication factor of the topic.
	FieldCreateTopicsResponseTopicsReplicationFactor = "ReplicationFactor"

	// FieldCreateTopicsResponseTopicsTags is: The tagged fields.
	FieldCreateTopicsResponseTopicsTags = "Tags"

	// FieldCreateTopicsResponseTopicsTopicConfigErrorCode is: Optional topic config error returned if configs are not returned in the response.
	FieldCreateTopicsResponseTopicsTopicConfigErrorCode = "TopicConfigErrorCode"

	// FieldCreateTopicsResponseTopicsTopicId is: The unique topic ID
	FieldCreateTopicsResponseTopicsTopicId = "TopicId"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/CreateTopicsResponse.json
const originalCreateTopicsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "response",
  "name": "CreateTopicsResponse",
  // Version 1 adds a per-topic error message string.
  //
  // Version 2 adds the throttle time.
  //
  // Starting in version 3, on quota violation, brokers send out responses before throttling.
  //
  // Version 4 makes partitions/replicationFactor optional even when assignments are not present (KIP-464).
  //
  // Version 5 is the first flexible version.
  // Version 5 also returns topic configs in the response (KIP-525).
  //
  // Version 6 is identical to version 5 but may return a THROTTLING_QUOTA_EXCEEDED error
  // in the response if the topics creation is throttled (KIP-599).
  //
  // Version 7 returns the topic ID of the newly created topic if creation is sucessful.
  "validVersions": "0-7",
  "flexibleVersions": "5+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "2+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Topics", "type": "[]CreatableTopicResult", "versions": "0+",
      "about": "Results for each topic we tried to create.", "fields": [
      { "name": "Name", "type": "string", "versions": "0+", "mapKey": true, "entityType": "topicName",
        "about": "The topic name." },
      { "name": "TopicId", "type": "uuid", "versions": "7+", "ignorable": true, "about": "The unique topic ID"},
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The error code, or 0 if there was no error." },
      { "name": "ErrorMessage", "type": "string", "versions": "1+", "nullableVersions": "0+", "ignorable": true,
        "about": "The error message, or null if there was no error." },
      { "name": "TopicConfigErrorCode", "type": "int16", "versions": "5+", "tag": 0, "taggedVersions": "5+", "ignorable": true,
        "about": "Optional topic config error returned if configs are not returned in the response." },
      { "name": "NumPartitions", "type": "int32", "versions": "5+", "default": "-1", "ignorable": true,
        "about": "Number of partitions of the topic." },
      { "name": "ReplicationFactor", "type": "int16", "versions": "5+", "default": "-1", "ignorable": true,
        "about": "Replication factor of the topic." },
      { "name": "Configs", "type": "[]CreatableTopicConfigs", "versions": "5+", "nullableVersions": "5+", "ignorable": true,
        "about": "Configuration of the topic.", "fields": [
        { "name": "Name", "type": "string", "versions": "5+",
          "about": "The configuration name." },
        { "name": "Value", "type": "string", "versions": "5+", "nullableVersions": "5+",
          "about": "The configuration value." },
        { "name": "ReadOnly", "type": "bool", "versions": "5+",
          "about": "True if the configuration is read-only." },
        { "name": "ConfigSource", "type": "int8", "versions": "5+", "default": "-1", "ignorable": true,
          "about": "The configuration source." },
        { "name": "IsSensitive", "type": "bool", "versions": "5+",
          "about": "True if this configuration is sensitive." }
      ]}
    ]}
  ]
}
`
