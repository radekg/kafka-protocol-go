package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init34AlterReplicaLogDirsResponse() []schema.Schema {

	return []schema.Schema{

		// Message: AlterReplicaLogDirsResponse, API Key: 34, Version: 0
		schema.NewSchema("AlterReplicaLogDirsResponsev0",
			&schema.Mfield{Name: FieldAlterReplicaLogDirsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldAlterReplicaLogDirsResponseResults, Ty: schema.NewSchema("ResultsV0",
				&schema.Mfield{Name: FieldAlterReplicaLogDirsResponseResultsTopicName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldAlterReplicaLogDirsResponseResultsPartitions, Ty: schema.NewSchema("PartitionsV0",
					&schema.Mfield{Name: FieldAlterReplicaLogDirsResponseResultsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAlterReplicaLogDirsResponseResultsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: AlterReplicaLogDirsResponse, API Key: 34, Version: 1
		schema.NewSchema("AlterReplicaLogDirsResponsev1",
			&schema.Mfield{Name: FieldAlterReplicaLogDirsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldAlterReplicaLogDirsResponseResults, Ty: schema.NewSchema("ResultsV1",
				&schema.Mfield{Name: FieldAlterReplicaLogDirsResponseResultsTopicName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldAlterReplicaLogDirsResponseResultsPartitions, Ty: schema.NewSchema("PartitionsV1",
					&schema.Mfield{Name: FieldAlterReplicaLogDirsResponseResultsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAlterReplicaLogDirsResponseResultsPartitionsErrorCode, Ty: schema.TypeInt16},
				)},
			)},
		),

		// Message: AlterReplicaLogDirsResponse, API Key: 34, Version: 2
		schema.NewSchema("AlterReplicaLogDirsResponsev2",
			&schema.Mfield{Name: FieldAlterReplicaLogDirsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldAlterReplicaLogDirsResponseResults, Ty: schema.NewSchema("ResultsV2",
				&schema.Mfield{Name: FieldAlterReplicaLogDirsResponseResultsTopicName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldAlterReplicaLogDirsResponseResultsPartitions, Ty: schema.NewSchema("PartitionsV2",
					&schema.Mfield{Name: FieldAlterReplicaLogDirsResponseResultsPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldAlterReplicaLogDirsResponseResultsPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.SchemaTaggedFields{Name: FieldAlterReplicaLogDirsResponseResultsPartitionsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldAlterReplicaLogDirsResponseResultsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldAlterReplicaLogDirsResponseTags},
		),
	}
}

const (
	// FieldAlterReplicaLogDirsResponseResults is: The results for each topic.
	FieldAlterReplicaLogDirsResponseResults = "Results"
	// FieldAlterReplicaLogDirsResponseResultsPartitions is: The results for each partition.
	FieldAlterReplicaLogDirsResponseResultsPartitions = "Partitions"
	// FieldAlterReplicaLogDirsResponseResultsPartitionsErrorCode is: The error code, or 0 if there was no error.
	FieldAlterReplicaLogDirsResponseResultsPartitionsErrorCode = "ErrorCode"
	// FieldAlterReplicaLogDirsResponseResultsPartitionsPartitionIndex is: The partition index.
	FieldAlterReplicaLogDirsResponseResultsPartitionsPartitionIndex = "PartitionIndex"
	// FieldAlterReplicaLogDirsResponseResultsPartitionsTags is: The tagged fields.
	FieldAlterReplicaLogDirsResponseResultsPartitionsTags = "Tags"
	// FieldAlterReplicaLogDirsResponseResultsTags is: The tagged fields.
	FieldAlterReplicaLogDirsResponseResultsTags = "Tags"
	// FieldAlterReplicaLogDirsResponseResultsTopicName is: The name of the topic.
	FieldAlterReplicaLogDirsResponseResultsTopicName = "TopicName"
	// FieldAlterReplicaLogDirsResponseTags is: The tagged fields.
	FieldAlterReplicaLogDirsResponseTags = "Tags"
	// FieldAlterReplicaLogDirsResponseThrottleTimeMs is: Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldAlterReplicaLogDirsResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AlterReplicaLogDirsResponse.json
const originalAlterReplicaLogDirsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 34,
  "type": "response",
  "name": "AlterReplicaLogDirsResponse",
  // Starting in version 1, on quota violation brokers send out responses before throttling.
  // Version 2 enables flexible versions.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Results", "type": "[]AlterReplicaLogDirTopicResult", "versions": "0+",
      "about": "The results for each topic.", "fields": [
      { "name": "TopicName", "type": "string", "versions": "0+", "entityType": "topicName",
        "about": "The name of the topic." },
      { "name": "Partitions", "type": "[]AlterReplicaLogDirPartitionResult", "versions": "0+",
        "about": "The results for each partition.", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The partition index."},
        { "name": "ErrorCode", "type": "int16", "versions": "0+",
          "about": "The error code, or 0 if there was no error." }
      ]}
    ]}
  ]
}
`
