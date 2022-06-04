package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init35DescribeLogDirsResponse() []schema.Schema {

	return []schema.Schema{

		// Message: DescribeLogDirsResponse, API Key: 35, Version: 0
		schema.NewSchema("DescribeLogDirsResponsev0",
			&schema.Mfield{Name: FieldDescribeLogDirsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDescribeLogDirsResponseResults, Ty: schema.NewSchema("ResultsV0",
				&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsLogDir, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDescribeLogDirsResponseResultsTopics, Ty: schema.NewSchema("TopicsV0",
					&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsName, Ty: schema.TypeStr},
					&schema.Array{Name: FieldDescribeLogDirsResponseResultsTopicsPartitions, Ty: schema.NewSchema("PartitionsV0",
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsPartitionSize, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsOffsetLag, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsIsFutureKey, Ty: schema.TypeBool},
					)},
				)},
			)},
		),

		// Message: DescribeLogDirsResponse, API Key: 35, Version: 1
		schema.NewSchema("DescribeLogDirsResponsev1",
			&schema.Mfield{Name: FieldDescribeLogDirsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDescribeLogDirsResponseResults, Ty: schema.NewSchema("ResultsV1",
				&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsLogDir, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDescribeLogDirsResponseResultsTopics, Ty: schema.NewSchema("TopicsV1",
					&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsName, Ty: schema.TypeStr},
					&schema.Array{Name: FieldDescribeLogDirsResponseResultsTopicsPartitions, Ty: schema.NewSchema("PartitionsV1",
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsPartitionSize, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsOffsetLag, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsIsFutureKey, Ty: schema.TypeBool},
					)},
				)},
			)},
		),

		// Message: DescribeLogDirsResponse, API Key: 35, Version: 2
		schema.NewSchema("DescribeLogDirsResponsev2",
			&schema.Mfield{Name: FieldDescribeLogDirsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldDescribeLogDirsResponseResults, Ty: schema.NewSchema("ResultsV2",
				&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsLogDir, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldDescribeLogDirsResponseResultsTopics, Ty: schema.NewSchema("TopicsV2",
					&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsName, Ty: schema.TypeStrCompact},
					&schema.ArrayCompact{Name: FieldDescribeLogDirsResponseResultsTopicsPartitions, Ty: schema.NewSchema("PartitionsV2",
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsPartitionSize, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsOffsetLag, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsIsFutureKey, Ty: schema.TypeBool},
						&schema.SchemaTaggedFields{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsTags},
					)},
					&schema.SchemaTaggedFields{Name: FieldDescribeLogDirsResponseResultsTopicsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldDescribeLogDirsResponseResultsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDescribeLogDirsResponseTags},
		),

		// Message: DescribeLogDirsResponse, API Key: 35, Version: 3
		schema.NewSchema("DescribeLogDirsResponsev3",
			&schema.Mfield{Name: FieldDescribeLogDirsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldDescribeLogDirsResponseErrorCode, Ty: schema.TypeInt16},
			&schema.ArrayCompact{Name: FieldDescribeLogDirsResponseResults, Ty: schema.NewSchema("ResultsV3",
				&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsLogDir, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldDescribeLogDirsResponseResultsTopics, Ty: schema.NewSchema("TopicsV3",
					&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsName, Ty: schema.TypeStrCompact},
					&schema.ArrayCompact{Name: FieldDescribeLogDirsResponseResultsTopicsPartitions, Ty: schema.NewSchema("PartitionsV3",
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsPartitionSize, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsOffsetLag, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsIsFutureKey, Ty: schema.TypeBool},
						&schema.SchemaTaggedFields{Name: FieldDescribeLogDirsResponseResultsTopicsPartitionsTags},
					)},
					&schema.SchemaTaggedFields{Name: FieldDescribeLogDirsResponseResultsTopicsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldDescribeLogDirsResponseResultsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDescribeLogDirsResponseTags},
		),
	}
}

const (
	// FieldDescribeLogDirsResponseErrorCode is: The error code, or 0 if there was no error.
	FieldDescribeLogDirsResponseErrorCode = "ErrorCode"
	// FieldDescribeLogDirsResponseResults is: The log directories.
	FieldDescribeLogDirsResponseResults = "Results"
	// FieldDescribeLogDirsResponseResultsErrorCode is: The error code, or 0 if there was no error.
	FieldDescribeLogDirsResponseResultsErrorCode = "ErrorCode"
	// FieldDescribeLogDirsResponseResultsLogDir is: The absolute log directory path.
	FieldDescribeLogDirsResponseResultsLogDir = "LogDir"
	// FieldDescribeLogDirsResponseResultsTags is: The tagged fields.
	FieldDescribeLogDirsResponseResultsTags = "Tags"
	// FieldDescribeLogDirsResponseResultsTopics is: Each topic.
	FieldDescribeLogDirsResponseResultsTopics = "Topics"
	// FieldDescribeLogDirsResponseResultsTopicsName is: The topic name.
	FieldDescribeLogDirsResponseResultsTopicsName = "Name"
	// FieldDescribeLogDirsResponseResultsTopicsPartitions is:
	FieldDescribeLogDirsResponseResultsTopicsPartitions = "Partitions"
	// FieldDescribeLogDirsResponseResultsTopicsPartitionsIsFutureKey is: True if this log is created by AlterReplicaLogDirsRequest and will replace the current log of the replica in the future.
	FieldDescribeLogDirsResponseResultsTopicsPartitionsIsFutureKey = "IsFutureKey"
	// FieldDescribeLogDirsResponseResultsTopicsPartitionsOffsetLag is: The lag of the log's LEO w.r.t. partition's HW (if it is the current log for the partition) or current replica's LEO (if it is the future log for the partition)
	FieldDescribeLogDirsResponseResultsTopicsPartitionsOffsetLag = "OffsetLag"
	// FieldDescribeLogDirsResponseResultsTopicsPartitionsPartitionIndex is: The partition index.
	FieldDescribeLogDirsResponseResultsTopicsPartitionsPartitionIndex = "PartitionIndex"
	// FieldDescribeLogDirsResponseResultsTopicsPartitionsPartitionSize is: The size of the log segments in this partition in bytes.
	FieldDescribeLogDirsResponseResultsTopicsPartitionsPartitionSize = "PartitionSize"
	// FieldDescribeLogDirsResponseResultsTopicsPartitionsTags is: The tagged fields.
	FieldDescribeLogDirsResponseResultsTopicsPartitionsTags = "Tags"
	// FieldDescribeLogDirsResponseResultsTopicsTags is: The tagged fields.
	FieldDescribeLogDirsResponseResultsTopicsTags = "Tags"
	// FieldDescribeLogDirsResponseTags is: The tagged fields.
	FieldDescribeLogDirsResponseTags = "Tags"
	// FieldDescribeLogDirsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldDescribeLogDirsResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeLogDirsResponse.json
const originalDescribeLogDirsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "response",
  "name": "DescribeLogDirsResponse",
  // Starting in version 1, on quota violation, brokers send out responses before throttling.
  "validVersions": "0-3",
  // Version 2 is the first flexible version.
  // Version 3 adds the top-level ErrorCode field
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "3+",
      "ignorable": true, "about": "The error code, or 0 if there was no error." },
    { "name": "Results", "type": "[]DescribeLogDirsResult", "versions": "0+",
      "about": "The log directories.", "fields": [
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The error code, or 0 if there was no error." },
      { "name": "LogDir", "type": "string", "versions": "0+",
        "about": "The absolute log directory path." },
      { "name": "Topics", "type": "[]DescribeLogDirsTopic", "versions": "0+",
        "about": "Each topic.", "fields": [
        { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
          "about": "The topic name." },
        { "name": "Partitions", "type": "[]DescribeLogDirsPartition", "versions": "0+", "fields": [
          { "name": "PartitionIndex", "type": "int32", "versions": "0+",
            "about": "The partition index." },
          { "name": "PartitionSize", "type": "int64", "versions": "0+",
            "about": "The size of the log segments in this partition in bytes." },
          { "name": "OffsetLag", "type": "int64", "versions": "0+",
            "about": "The lag of the log's LEO w.r.t. partition's HW (if it is the current log for the partition) or current replica's LEO (if it is the future log for the partition)" },
          { "name": "IsFutureKey", "type": "bool", "versions": "0+",
            "about": "True if this log is created by AlterReplicaLogDirsRequest and will replace the current log of the replica in the future." }
        ]}
      ]}
    ]}
  ]
}
`
