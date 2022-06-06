package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init32DescribeConfigsResponse() []schema.Schema {

	return []schema.Schema{
		// Message: DescribeConfigsResponse, API Key: 32, Version: 0
		schema.NewSchema("DescribeConfigsResponse:v0",
			&schema.Mfield{Name: FieldDescribeConfigsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDescribeConfigsResponseResults, Ty: schema.NewSchema("[]DescribeConfigsResult:v0",
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsErrorMessage, Ty: schema.TypeStrNullable},
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsResourceName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDescribeConfigsResponseResultsConfigs, Ty: schema.NewSchema("[]DescribeConfigsResourceResult:v0",
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsValue, Ty: schema.TypeStrNullable},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsReadOnly, Ty: schema.TypeBool},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsIsSensitive, Ty: schema.TypeBool},
				)},
			)},
		),

		// Message: DescribeConfigsResponse, API Key: 32, Version: 1
		schema.NewSchema("DescribeConfigsResponse:v1",
			&schema.Mfield{Name: FieldDescribeConfigsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDescribeConfigsResponseResults, Ty: schema.NewSchema("[]DescribeConfigsResult:v1",
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsErrorMessage, Ty: schema.TypeStrNullable},
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsResourceName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDescribeConfigsResponseResultsConfigs, Ty: schema.NewSchema("[]DescribeConfigsResourceResult:v1",
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsValue, Ty: schema.TypeStrNullable},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsReadOnly, Ty: schema.TypeBool},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsConfigSource, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsIsSensitive, Ty: schema.TypeBool},
					&schema.Array{Name: FieldDescribeConfigsResponseResultsConfigsSynonyms, Ty: schema.NewSchema("[]DescribeConfigsSynonym:v1",
						&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsSynonymsName, Ty: schema.TypeStr},
						&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsSynonymsValue, Ty: schema.TypeStrNullable},
						&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsSynonymsSource, Ty: schema.TypeInt8},
					)},
				)},
			)},
		),

		// Message: DescribeConfigsResponse, API Key: 32, Version: 2
		schema.NewSchema("DescribeConfigsResponse:v2",
			&schema.Mfield{Name: FieldDescribeConfigsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDescribeConfigsResponseResults, Ty: schema.NewSchema("[]DescribeConfigsResult:v2",
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsErrorMessage, Ty: schema.TypeStrNullable},
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsResourceName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDescribeConfigsResponseResultsConfigs, Ty: schema.NewSchema("[]DescribeConfigsResourceResult:v2",
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsValue, Ty: schema.TypeStrNullable},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsReadOnly, Ty: schema.TypeBool},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsConfigSource, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsIsSensitive, Ty: schema.TypeBool},
					&schema.Array{Name: FieldDescribeConfigsResponseResultsConfigsSynonyms, Ty: schema.NewSchema("[]DescribeConfigsSynonym:v2",
						&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsSynonymsName, Ty: schema.TypeStr},
						&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsSynonymsValue, Ty: schema.TypeStrNullable},
						&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsSynonymsSource, Ty: schema.TypeInt8},
					)},
				)},
			)},
		),

		// Message: DescribeConfigsResponse, API Key: 32, Version: 3
		schema.NewSchema("DescribeConfigsResponse:v3",
			&schema.Mfield{Name: FieldDescribeConfigsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldDescribeConfigsResponseResults, Ty: schema.NewSchema("[]DescribeConfigsResult:v3",
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsErrorMessage, Ty: schema.TypeStrNullable},
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsResourceName, Ty: schema.TypeStr},
				&schema.Array{Name: FieldDescribeConfigsResponseResultsConfigs, Ty: schema.NewSchema("[]DescribeConfigsResourceResult:v3",
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsName, Ty: schema.TypeStr},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsValue, Ty: schema.TypeStrNullable},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsReadOnly, Ty: schema.TypeBool},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsConfigSource, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsIsSensitive, Ty: schema.TypeBool},
					&schema.Array{Name: FieldDescribeConfigsResponseResultsConfigsSynonyms, Ty: schema.NewSchema("[]DescribeConfigsSynonym:v3",
						&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsSynonymsName, Ty: schema.TypeStr},
						&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsSynonymsValue, Ty: schema.TypeStrNullable},
						&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsSynonymsSource, Ty: schema.TypeInt8},
					)},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsConfigType, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsDocumentation, Ty: schema.TypeStrNullable},
				)},
			)},
		),

		// Message: DescribeConfigsResponse, API Key: 32, Version: 4
		schema.NewSchema("DescribeConfigsResponse:v4",
			&schema.Mfield{Name: FieldDescribeConfigsResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldDescribeConfigsResponseResults, Ty: schema.NewSchema("[]DescribeConfigsResult:v4",
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsErrorCode, Ty: schema.TypeInt16},
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsErrorMessage, Ty: schema.TypeStrCompactNullable},
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsResourceType, Ty: schema.TypeInt8},
				&schema.Mfield{Name: FieldDescribeConfigsResponseResultsResourceName, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldDescribeConfigsResponseResultsConfigs, Ty: schema.NewSchema("[]DescribeConfigsResourceResult:v4",
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsName, Ty: schema.TypeStrCompact},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsValue, Ty: schema.TypeStrCompactNullable},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsReadOnly, Ty: schema.TypeBool},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsConfigSource, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsIsSensitive, Ty: schema.TypeBool},
					&schema.ArrayCompact{Name: FieldDescribeConfigsResponseResultsConfigsSynonyms, Ty: schema.NewSchema("[]DescribeConfigsSynonym:v4",
						&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsSynonymsName, Ty: schema.TypeStrCompact},
						&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsSynonymsValue, Ty: schema.TypeStrCompactNullable},
						&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsSynonymsSource, Ty: schema.TypeInt8},
						&schema.SchemaTaggedFields{Name: FieldDescribeConfigsResponseResultsConfigsSynonymsTags},
					)},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsConfigType, Ty: schema.TypeInt8},
					&schema.Mfield{Name: FieldDescribeConfigsResponseResultsConfigsDocumentation, Ty: schema.TypeStrCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldDescribeConfigsResponseResultsConfigsTags},
				)},
				&schema.SchemaTaggedFields{Name: FieldDescribeConfigsResponseResultsTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldDescribeConfigsResponseTags},
		),
	}

}

const (

	// FieldDescribeConfigsResponseResults is: The results for each resource.
	FieldDescribeConfigsResponseResults = "Results"

	// FieldDescribeConfigsResponseResultsConfigs is: Each listed configuration.
	FieldDescribeConfigsResponseResultsConfigs = "Configs"

	// FieldDescribeConfigsResponseResultsConfigsConfigSource is: The configuration source.
	FieldDescribeConfigsResponseResultsConfigsConfigSource = "ConfigSource"

	// FieldDescribeConfigsResponseResultsConfigsConfigType is: The configuration data type. Type can be one of the following values - BOOLEAN, STRING, INT, SHORT, LONG, DOUBLE, LIST, CLASS, PASSWORD
	FieldDescribeConfigsResponseResultsConfigsConfigType = "ConfigType"

	// FieldDescribeConfigsResponseResultsConfigsDocumentation is: The configuration documentation.
	FieldDescribeConfigsResponseResultsConfigsDocumentation = "Documentation"

	// FieldDescribeConfigsResponseResultsConfigsIsSensitive is: True if this configuration is sensitive.
	FieldDescribeConfigsResponseResultsConfigsIsSensitive = "IsSensitive"

	// FieldDescribeConfigsResponseResultsConfigsName is: The configuration name.
	FieldDescribeConfigsResponseResultsConfigsName = "Name"

	// FieldDescribeConfigsResponseResultsConfigsReadOnly is: True if the configuration is read-only.
	FieldDescribeConfigsResponseResultsConfigsReadOnly = "ReadOnly"

	// FieldDescribeConfigsResponseResultsConfigsSynonyms is: The synonyms for this configuration key.
	FieldDescribeConfigsResponseResultsConfigsSynonyms = "Synonyms"

	// FieldDescribeConfigsResponseResultsConfigsSynonymsName is: The synonym name.
	FieldDescribeConfigsResponseResultsConfigsSynonymsName = "Name"

	// FieldDescribeConfigsResponseResultsConfigsSynonymsSource is: The synonym source.
	FieldDescribeConfigsResponseResultsConfigsSynonymsSource = "Source"

	// FieldDescribeConfigsResponseResultsConfigsSynonymsTags is: The tagged fields.
	FieldDescribeConfigsResponseResultsConfigsSynonymsTags = "Tags"

	// FieldDescribeConfigsResponseResultsConfigsSynonymsValue is: The synonym value.
	FieldDescribeConfigsResponseResultsConfigsSynonymsValue = "Value"

	// FieldDescribeConfigsResponseResultsConfigsTags is: The tagged fields.
	FieldDescribeConfigsResponseResultsConfigsTags = "Tags"

	// FieldDescribeConfigsResponseResultsConfigsValue is: The configuration value.
	FieldDescribeConfigsResponseResultsConfigsValue = "Value"

	// FieldDescribeConfigsResponseResultsErrorCode is: The error code, or 0 if we were able to successfully describe the configurations.
	FieldDescribeConfigsResponseResultsErrorCode = "ErrorCode"

	// FieldDescribeConfigsResponseResultsErrorMessage is: The error message, or null if we were able to successfully describe the configurations.
	FieldDescribeConfigsResponseResultsErrorMessage = "ErrorMessage"

	// FieldDescribeConfigsResponseResultsResourceName is: The resource name.
	FieldDescribeConfigsResponseResultsResourceName = "ResourceName"

	// FieldDescribeConfigsResponseResultsResourceType is: The resource type.
	FieldDescribeConfigsResponseResultsResourceType = "ResourceType"

	// FieldDescribeConfigsResponseResultsTags is: The tagged fields.
	FieldDescribeConfigsResponseResultsTags = "Tags"

	// FieldDescribeConfigsResponseTags is: The tagged fields.
	FieldDescribeConfigsResponseTags = "Tags"

	// FieldDescribeConfigsResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldDescribeConfigsResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/DescribeConfigsResponse.json
const originalDescribeConfigsResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
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
  "apiKey": 32,
  "type": "response",
  "name": "DescribeConfigsResponse",
  // Version 1 adds ConfigSource and the synonyms.
  // Starting in version 2, on quota violation, brokers send out responses before throttling.
  // Version 4 enables flexible versions.
  "validVersions": "0-4",
  "flexibleVersions": "4+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "0+",
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "Results", "type": "[]DescribeConfigsResult", "versions": "0+",
      "about": "The results for each resource.", "fields": [
      { "name": "ErrorCode", "type": "int16", "versions": "0+",
        "about": "The error code, or 0 if we were able to successfully describe the configurations." },
      { "name": "ErrorMessage", "type": "string", "versions": "0+", "nullableVersions": "0+",
        "about": "The error message, or null if we were able to successfully describe the configurations." },
      { "name": "ResourceType", "type": "int8", "versions": "0+",
        "about": "The resource type." },
      { "name": "ResourceName", "type": "string", "versions": "0+",
        "about": "The resource name." },
      { "name": "Configs", "type": "[]DescribeConfigsResourceResult", "versions": "0+",
        "about": "Each listed configuration.", "fields": [
        { "name": "Name", "type": "string", "versions": "0+",
          "about": "The configuration name." },
        { "name": "Value", "type": "string", "versions": "0+", "nullableVersions": "0+",
          "about": "The configuration value." },
        { "name": "ReadOnly", "type": "bool", "versions": "0+",
          "about": "True if the configuration is read-only." },
        { "name": "IsDefault", "type": "bool", "versions": "0",
          "about": "True if the configuration is not set." },
        // Note: the v0 default for this field that should be exposed to callers is
        // context-dependent. For example, if the resource is a broker, this should default to 4.
        // -1 is just a placeholder value.
        { "name": "ConfigSource", "type": "int8", "versions": "1+", "default": "-1", "ignorable": true,
          "about": "The configuration source." },
        { "name": "IsSensitive", "type": "bool", "versions": "0+",
          "about": "True if this configuration is sensitive." },
        { "name": "Synonyms", "type": "[]DescribeConfigsSynonym", "versions": "1+", "ignorable": true,
          "about": "The synonyms for this configuration key.", "fields": [
          { "name": "Name", "type": "string", "versions": "1+",
            "about": "The synonym name." },
          { "name": "Value", "type": "string", "versions": "1+", "nullableVersions": "0+",
            "about": "The synonym value." },
          { "name": "Source", "type": "int8", "versions": "1+",
            "about": "The synonym source." }
        ]},
        { "name": "ConfigType", "type": "int8", "versions": "3+", "default": "0", "ignorable": true,
          "about": "The configuration data type. Type can be one of the following values - BOOLEAN, STRING, INT, SHORT, LONG, DOUBLE, LIST, CLASS, PASSWORD" },
        { "name": "Documentation", "type": "string", "versions": "3+", "nullableVersions": "0+", "ignorable": true,
          "about": "The configuration documentation." }
      ]}
    ]}
  ]
}
`
