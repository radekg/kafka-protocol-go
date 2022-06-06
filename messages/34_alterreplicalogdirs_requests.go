package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init34AlterReplicaLogDirsRequest() []schema.Schema {

	return []schema.Schema{
		// Message: AlterReplicaLogDirsRequest, API Key: 34, Version: 0
		schema.NewSchema("AlterReplicaLogDirsRequest:v0",
			&schema.Array{Name: FieldAlterReplicaLogDirsRequestDirs, Ty: schema.NewSchema("[]AlterReplicaLogDir:v0",
				&schema.Mfield{Name: FieldAlterReplicaLogDirsRequestDirsPath, Ty: schema.TypeStr},
				&schema.Array{Name: FieldAlterReplicaLogDirsRequestDirsTopics, Ty: schema.NewSchema("[]AlterReplicaLogDirTopic:v0",
					&schema.Mfield{Name: FieldAlterReplicaLogDirsRequestDirsTopicsName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldAlterReplicaLogDirsRequestDirsTopicsPartitions, Ty: schema.TypeInt32Array},
				)},
			)},
		),

		// Message: AlterReplicaLogDirsRequest, API Key: 34, Version: 1
		schema.NewSchema("AlterReplicaLogDirsRequest:v1",
			&schema.Array{Name: FieldAlterReplicaLogDirsRequestDirs, Ty: schema.NewSchema("[]AlterReplicaLogDir:v1",
				&schema.Mfield{Name: FieldAlterReplicaLogDirsRequestDirsPath, Ty: schema.TypeStr},
				&schema.Array{Name: FieldAlterReplicaLogDirsRequestDirsTopics, Ty: schema.NewSchema("[]AlterReplicaLogDirTopic:v1",
					&schema.Mfield{Name: FieldAlterReplicaLogDirsRequestDirsTopicsName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldAlterReplicaLogDirsRequestDirsTopicsPartitions, Ty: schema.TypeInt32Array},
				)},
			)},
		),

		// Message: AlterReplicaLogDirsRequest, API Key: 34, Version: 2
		schema.NewSchema("AlterReplicaLogDirsRequest:v2",
			&schema.ArrayCompact{Name: FieldAlterReplicaLogDirsRequestDirs, Ty: schema.NewSchema("[]AlterReplicaLogDir:v2",
				&schema.Mfield{Name: FieldAlterReplicaLogDirsRequestDirsPath, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldAlterReplicaLogDirsRequestDirsTopics, Ty: schema.NewSchema("[]AlterReplicaLogDirTopic:v2",
					&schema.Mfield{Name: FieldAlterReplicaLogDirsRequestDirsTopicsName, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldAlterReplicaLogDirsRequestDirsTopicsPartitions, Ty: schema.TypeInt32CompactArray},
					&schema.SchemaTaggedFields{Name: FieldAlterReplicaLogDirsRequestDirsTopicsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldAlterReplicaLogDirsRequestDirsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldAlterReplicaLogDirsRequestTags},
		),
	}

}

const (

	// FieldAlterReplicaLogDirsRequestDirs is: The alterations to make for each directory.
	FieldAlterReplicaLogDirsRequestDirs = "Dirs"

	// FieldAlterReplicaLogDirsRequestDirsPath is: The absolute directory path.
	FieldAlterReplicaLogDirsRequestDirsPath = "Path"

	// FieldAlterReplicaLogDirsRequestDirsTags is: The tagged fields.
	FieldAlterReplicaLogDirsRequestDirsTags = "Tags"

	// FieldAlterReplicaLogDirsRequestDirsTopics is: The topics to add to the directory.
	FieldAlterReplicaLogDirsRequestDirsTopics = "Topics"

	// FieldAlterReplicaLogDirsRequestDirsTopicsName is: The topic name.
	FieldAlterReplicaLogDirsRequestDirsTopicsName = "Name"

	// FieldAlterReplicaLogDirsRequestDirsTopicsPartitions is: The partition indexes.
	FieldAlterReplicaLogDirsRequestDirsTopicsPartitions = "Partitions"

	// FieldAlterReplicaLogDirsRequestDirsTopicsTags is: The tagged fields.
	FieldAlterReplicaLogDirsRequestDirsTopicsTags = "Tags"

	// FieldAlterReplicaLogDirsRequestTags is: The tagged fields.
	FieldAlterReplicaLogDirsRequestTags = "Tags"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/AlterReplicaLogDirsRequest.json
const originalAlterReplicaLogDirsRequestInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "type": "request",
  "listeners": ["zkBroker", "broker"],
  "name": "AlterReplicaLogDirsRequest",
  // Version 1 is the same as version 0.
  // Version 2 enables flexible versions.
  "validVersions": "0-2",
  "flexibleVersions": "2+",
  "fields": [
    { "name": "Dirs", "type": "[]AlterReplicaLogDir", "versions": "0+", 
      "about": "The alterations to make for each directory.", "fields": [
      { "name": "Path", "type": "string", "versions": "0+", "mapKey": true,
        "about": "The absolute directory path." },
      { "name": "Topics", "type": "[]AlterReplicaLogDirTopic", "versions": "0+",
        "about": "The topics to add to the directory.",  "fields": [
        { "name": "Name", "type": "string", "versions": "0+", "mapKey": true, "entityType": "topicName",
          "about": "The topic name." },
        { "name": "Partitions", "type": "[]int32", "versions": "0+",
          "about": "The partition indexes." }
      ]}
    ]}
  ]
}
`
