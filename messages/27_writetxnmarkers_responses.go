package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init27WriteTxnMarkersResponse() []schema.Schema {

	return []schema.Schema{
		// Message: WriteTxnMarkersResponse, API Key: 27, Version: 0
		schema.NewSchema("WriteTxnMarkersResponse:v0",
			&schema.Array{Name: FieldWriteTxnMarkersResponseMarkers, Ty: schema.NewSchema("Markers:v0",
				&schema.Mfield{Name: FieldWriteTxnMarkersResponseMarkersProducerId, Ty: schema.TypeInt64},
				&schema.Array{Name: FieldWriteTxnMarkersResponseMarkersTopics, Ty: schema.NewSchema("Topics:v0",
					&schema.Mfield{Name: FieldWriteTxnMarkersResponseMarkersTopicsName, Ty: schema.TypeStr},
					&schema.Array{Name: FieldWriteTxnMarkersResponseMarkersTopicsPartitions, Ty: schema.NewSchema("Partitions:v0",
						&schema.Mfield{Name: FieldWriteTxnMarkersResponseMarkersTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
						&schema.Mfield{Name: FieldWriteTxnMarkersResponseMarkersTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
					)},
				)},
			)},
		),

		// Message: WriteTxnMarkersResponse, API Key: 27, Version: 1
		schema.NewSchema("WriteTxnMarkersResponse:v1",
			&schema.ArrayCompact{Name: FieldWriteTxnMarkersResponseMarkers, Ty: schema.NewSchema("Markers:v1",
				&schema.Mfield{Name: FieldWriteTxnMarkersResponseMarkersProducerId, Ty: schema.TypeInt64},
				&schema.ArrayCompact{Name: FieldWriteTxnMarkersResponseMarkersTopics, Ty: schema.NewSchema("Topics:v1",
					&schema.Mfield{Name: FieldWriteTxnMarkersResponseMarkersTopicsName, Ty: schema.TypeStrCompact},
					&schema.ArrayCompact{Name: FieldWriteTxnMarkersResponseMarkersTopicsPartitions, Ty: schema.NewSchema("Partitions:v1",
						&schema.Mfield{Name: FieldWriteTxnMarkersResponseMarkersTopicsPartitionsPartitionIndex, Ty: schema.TypeInt32},
						&schema.Mfield{Name: FieldWriteTxnMarkersResponseMarkersTopicsPartitionsErrorCode, Ty: schema.TypeInt16},
						&schema.SchemaTaggedFields{Name: FieldWriteTxnMarkersResponseMarkersTopicsPartitionsTags},
					)},
					&schema.SchemaTaggedFields{Name: FieldWriteTxnMarkersResponseMarkersTopicsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldWriteTxnMarkersResponseMarkersTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldWriteTxnMarkersResponseTags},
		),
	}

}

const (

	// FieldWriteTxnMarkersResponseMarkers is: The results for writing makers.
	FieldWriteTxnMarkersResponseMarkers = "Markers"

	// FieldWriteTxnMarkersResponseMarkersProducerId is: The current producer ID in use by the transactional ID.
	FieldWriteTxnMarkersResponseMarkersProducerId = "ProducerId"

	// FieldWriteTxnMarkersResponseMarkersTags is: The tagged fields.
	FieldWriteTxnMarkersResponseMarkersTags = "Tags"

	// FieldWriteTxnMarkersResponseMarkersTopics is: The results by topic.
	FieldWriteTxnMarkersResponseMarkersTopics = "Topics"

	// FieldWriteTxnMarkersResponseMarkersTopicsName is: The topic name.
	FieldWriteTxnMarkersResponseMarkersTopicsName = "Name"

	// FieldWriteTxnMarkersResponseMarkersTopicsPartitions is: The results by partition.
	FieldWriteTxnMarkersResponseMarkersTopicsPartitions = "Partitions"

	// FieldWriteTxnMarkersResponseMarkersTopicsPartitionsErrorCode is: The error code, or 0 if there was no error.
	FieldWriteTxnMarkersResponseMarkersTopicsPartitionsErrorCode = "ErrorCode"

	// FieldWriteTxnMarkersResponseMarkersTopicsPartitionsPartitionIndex is: The partition index.
	FieldWriteTxnMarkersResponseMarkersTopicsPartitionsPartitionIndex = "PartitionIndex"

	// FieldWriteTxnMarkersResponseMarkersTopicsPartitionsTags is: The tagged fields.
	FieldWriteTxnMarkersResponseMarkersTopicsPartitionsTags = "Tags"

	// FieldWriteTxnMarkersResponseMarkersTopicsTags is: The tagged fields.
	FieldWriteTxnMarkersResponseMarkersTopicsTags = "Tags"

	// FieldWriteTxnMarkersResponseTags is: The tagged fields.
	FieldWriteTxnMarkersResponseTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/WriteTxnMarkersResponse.json
const originalWriteTxnMarkersResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 27,
  "type": "response",
  "name": "WriteTxnMarkersResponse",
  "validVersions": "0-1",
  // Version 1 enables flexible versions.
  "flexibleVersions": "1+",
  "fields": [
    { "name": "Markers", "type": "[]WritableTxnMarkerResult", "versions": "0+",
      "about": "The results for writing makers.", "fields": [
      { "name": "ProducerId", "type": "int64", "versions": "0+", "entityType": "producerId",
        "about": "The current producer ID in use by the transactional ID." },
      { "name": "Topics", "type": "[]WritableTxnMarkerTopicResult", "versions": "0+",
        "about": "The results by topic.", "fields": [
        { "name": "Name", "type": "string", "versions": "0+", "entityType": "topicName",
          "about": "The topic name." },
        { "name": "Partitions", "type": "[]WritableTxnMarkerPartitionResult", "versions": "0+",
          "about": "The results by partition.", "fields": [
          { "name": "PartitionIndex", "type": "int32", "versions": "0+",
            "about": "The partition index." },
          { "name": "ErrorCode", "type": "int16", "versions": "0+",
            "about": "The error code, or 0 if there was no error." }
        ]}
      ]}
    ]}
  ]
}
`
