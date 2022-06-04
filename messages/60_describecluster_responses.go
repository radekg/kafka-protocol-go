package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init60DescribeClusterResponse() []schema.Schema {

	return []schema.Schema{

		// Message: DescribeClusterResponse, API Key: 60, Version: 0
		schema.NewSchema("DescribeClusterResponsev0",
			&schema.Mfield{Name: FieldDescribeClusterResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldDescribeClusterResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldDescribeClusterResponseErrorMessage, Ty: schema.TypeStrCompactNullable},
			&schema.Mfield{Name: FieldDescribeClusterResponseClusterId, Ty: schema.TypeStrCompact},
			&schema.Mfield{Name: FieldDescribeClusterResponseControllerId, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldDescribeClusterResponseBrokers, Ty: schema.NewSchema("BrokersV0",
				&schema.Mfield{Name: FieldDescribeClusterResponseBrokersBrokerId, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldDescribeClusterResponseBrokersHost, Ty: schema.TypeStrCompact},
				&schema.Mfield{Name: FieldDescribeClusterResponseBrokersPort, Ty: schema.TypeInt32},
				&schema.Mfield{Name: FieldDescribeClusterResponseBrokersRack, Ty: schema.TypeStrCompactNullable},
				&schema.SchemaTaggedFields{Name: FieldDescribeClusterResponseBrokersTags},
			)},
			&schema.Mfield{Name: FieldDescribeClusterResponseClusterAuthorizedOperations, Ty: schema.TypeInt32},
			&schema.SchemaTaggedFields{Name: FieldDescribeClusterResponseTags},
		),
	}
}

const (
	// FieldDescribeClusterResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldDescribeClusterResponseThrottleTimeMs = "ThrottleTimeMs"
	// FieldDescribeClusterResponseErrorCode is: The top-level error code, or 0 if there was no error
	FieldDescribeClusterResponseErrorCode = "ErrorCode"
	// FieldDescribeClusterResponseBrokers is: Each broker in the response.
	FieldDescribeClusterResponseBrokers = "Brokers"
	// FieldDescribeClusterResponseBrokersBrokerId is: The broker ID.
	FieldDescribeClusterResponseBrokersBrokerId = "BrokerId"
	// FieldDescribeClusterResponseBrokersHost is: The broker hostname.
	FieldDescribeClusterResponseBrokersHost = "Host"
	// FieldDescribeClusterResponseBrokersPort is: The broker port.
	FieldDescribeClusterResponseBrokersPort = "Port"
	// FieldDescribeClusterResponseBrokersRack is: The rack of the broker, or null if it has not been assigned to a rack.
	FieldDescribeClusterResponseBrokersRack = "Rack"
	// FieldDescribeClusterResponseBrokersTags is: The tagged fields.
	FieldDescribeClusterResponseBrokersTags = "Tags"
	// FieldDescribeClusterResponseTags is: The tagged fields.
	FieldDescribeClusterResponseTags = "Tags"
	// FieldDescribeClusterResponseErrorMessage is: The top-level error message, or null if there was no error.
	FieldDescribeClusterResponseErrorMessage = "ErrorMessage"
	// FieldDescribeClusterResponseClusterId is: The cluster ID that responding broker belongs to.
	FieldDescribeClusterResponseClusterId = "ClusterId"
	// FieldDescribeClusterResponseControllerId is: The ID of the controller broker.
	FieldDescribeClusterResponseControllerId = "ControllerId"
	// FieldDescribeClusterResponseClusterAuthorizedOperations is: 32-bit bitfield to represent authorized operations for this cluster.
	FieldDescribeClusterResponseClusterAuthorizedOperations = "ClusterAuthorizedOperations"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeClusterResponse.json
const originalDescribeClusterResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 60,
  "type": "response",
  "name": "DescribeClusterResponse",
  "validVersions": "0",
  "flexibleVersions": "0+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The top-level error code, or 0 if there was no error" },
    { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+", "default": "null",
      "about": "The top-level error message, or null if there was no error." },
    { "name": "ClusterId", "type": "string", "versions": "0+",
      "about": "The cluster ID that responding broker belongs to." },
    { "name": "ControllerId", "type": "int32", "versions": "0+", "default": "-1", "entityType": "brokerId",
      "about": "The ID of the controller broker." },
    { "name": "Brokers", "type": "[]DescribeClusterBroker", "versions": "0+",
      "about": "Each broker in the response.", "fields": [
      { "name": "BrokerId", "type": "int32", "versions": "0+", "mapKey": true, "entityType": "brokerId",
        "about": "The broker ID." },
      { "name": "Host", "type": "string", "versions": "0+",
        "about": "The broker hostname." },
      { "name": "Port", "type": "int32", "versions": "0+",
        "about": "The broker port." },
      { "name": "Rack", "type": "string", "versions": "0+", "nullableVersions": "0+", "default": "null",
        "about": "The rack of the broker, or null if it has not been assigned to a rack." }
    ]},
    { "name": "ClusterAuthorizedOperations", "type": "int32", "versions": "0+", "default": "-2147483648",
      "about": "32-bit bitfield to represent authorized operations for this cluster." }
  ]
}
`
