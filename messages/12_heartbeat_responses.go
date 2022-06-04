package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init12HeartbeatResponse() []schema.Schema {

	return []schema.Schema{

		// Message: HeartbeatResponse, API Key: 12, Version: 0
		schema.NewSchema("HeartbeatResponsev0",
			&schema.Mfield{Name: FieldHeartbeatResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: HeartbeatResponse, API Key: 12, Version: 1
		schema.NewSchema("HeartbeatResponsev1",
			&schema.Mfield{Name: FieldHeartbeatResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldHeartbeatResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: HeartbeatResponse, API Key: 12, Version: 2
		schema.NewSchema("HeartbeatResponsev2",
			&schema.Mfield{Name: FieldHeartbeatResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldHeartbeatResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: HeartbeatResponse, API Key: 12, Version: 3
		schema.NewSchema("HeartbeatResponsev3",
			&schema.Mfield{Name: FieldHeartbeatResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldHeartbeatResponseErrorCode, Ty: schema.TypeInt16},
		),

		// Message: HeartbeatResponse, API Key: 12, Version: 4
		schema.NewSchema("HeartbeatResponsev4",
			&schema.Mfield{Name: FieldHeartbeatResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldHeartbeatResponseErrorCode, Ty: schema.TypeInt16},
			&schema.SchemaTaggedFields{Name: FieldHeartbeatResponseTags},
		),
	}
}

const (
	// FieldHeartbeatResponseErrorCode is: The error code, or 0 if there was no error.
	FieldHeartbeatResponseErrorCode = "ErrorCode"
	// FieldHeartbeatResponseTags is: The tagged fields.
	FieldHeartbeatResponseTags = "Tags"
	// FieldHeartbeatResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldHeartbeatResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/HeartbeatResponse.json
const originalHeartbeatResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 12,
  "type": "response",
  "name": "HeartbeatResponse",
  // Version 1 adds throttle time.
  //
  // Starting in version 2, on quota violation, brokers send out responses before throttling.
  //
  // Starting from version 3, heartbeatRequest supports a new field called groupInstanceId to indicate member identity across restarts.
  //
  // Version 4 is the first flexible version.
  "validVersions": "0-4",
  "flexibleVersions": "4+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "1+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "0+",
      "about": "The error code, or 0 if there was no error." }
  ]
}
`
