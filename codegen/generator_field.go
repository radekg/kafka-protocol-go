package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

type fields []field

func (d fields) fieldsForVersion(v int64, flexible bool) fields {
	resp := fields{}
	for _, f := range d {
		if f.Tag == nil && appliesToVersion(f.Name, f.Versions, v) {
			resp = append(resp, f)
		}
	}
	if len(resp) > 0 && flexible {
		resp = append(resp, field{
			Name:             "Tags",
			Type:             "schemaTags",
			Versions:         fmt.Sprintf("%d+", v),
			NullableVersions: "",
			About:            "The tagged fields.",
			Fields:           fields{},
		})
	}
	return resp
}

func (d fields) tagsForVersion(v int64, flexible bool) fields {
	resp := fields{}
	for _, f := range d {
		if f.Tag != nil && appliesToVersion(f.Name, *f.TaggedVersions, v) {
			resp = append(resp, f)
		}
	}
	return resp
}

type field struct {
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	Versions         string  `json:"versions"`
	NullableVersions string  `json:"nullableVersions"`
	Tag              *int    `json:"tag,omitempty"`
	TaggedVersions   *string `json:"taggedVersions,omitempty"`
	About            string  `json:"About"`
	Fields           fields  `json:"fields"`
}

func (d *field) nullable(v int64) bool {
	if d.NullableVersions == "" {
		return false
	}
	minVersion, err := strconv.ParseInt(strings.TrimSuffix(d.NullableVersions, "+"), 10, 64)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("failed parsing versions with + suffix: %s", d.Name))
	}
	if v >= minVersion {
		return true
	}
	return false
}

func (d *field) generateSchema(params *schemaGeneratorMetadata) {

	params.NameRegistry.builder.add(d.Name)
	defer params.NameRegistry.builder.pop()

	newField := &irField{
		APIName:           d.Name,
		APIVersion:        params.Appender.GetAPIVersion(),
		APIType:           d.Type,
		Fields:            []fieldAppender{},
		Flexible:          params.Appender.GetFlexible(),
		ConstantFieldName: params.NameRegistry.builder.generate(),
		Nullable:          d.nullable(params.Appender.GetAPIVersion()),
		Whitespace: strings.Repeat("	", params.NestLevel),
	}

	schemaTypeSpec := getSchemaType(d.Type, params.Appender.GetFlexible(), newField.Nullable)

	params.NameRegistry.register(newField.ConstantFieldName, &fieldWithDoc{
		name: d.Name,
		doc:  d.About,
	})

	if schemaTypeSpec.RequiresCompactArrayWrap {

		if !schemaTypeSpec.IsSchema {
			newField.SchemaType = schemaTypeSpec.fullName()
		} else {
			// Caution: for printing in comments the nest level is 1 higher than actual code:
			collectedTags := collectTags(d.Fields, params.Appender.GetAPIVersion(), params.Appender.GetFlexible(), params.CommonStructs, params.NameRegistry, params.NestLevel+2)
			fieldsForVersion := d.Fields.fieldsForVersion(params.Appender.GetAPIVersion(), params.Appender.GetFlexible())
			// if we have no fields, maybe we have a common struct
			if len(fieldsForVersion) == 0 {
				fieldsForVersion = params.CommonStructs.forVersion(d.Type, params.Appender.GetAPIVersion(), params.Appender.GetFlexible())
			}
			for _, field := range fieldsForVersion {
				field.generateSchema(&schemaGeneratorMetadata{
					Appender:      newField,
					CommonStructs: params.CommonStructs,
					NestLevel:     params.NestLevel + 1,
					NameRegistry:  params.NameRegistry,
					Tags:          collectedTags,
				})
			}
		}

		parsedTemplate, err := template.New("arraycompact").Parse(fieldCompactArrayTemplate)
		if err != nil {
			panic(err)
		}
		var output bytes.Buffer
		if err := parsedTemplate.Execute(&output, newField); err != nil {
			panic(err)
		}
		newField.Rendered = output.String()
		params.Appender.Append(newField)

	} else if schemaTypeSpec.RequiresArrayWrap {

		if !schemaTypeSpec.IsSchema {
			newField.SchemaType = schemaTypeSpec.fullName()
		} else {
			// Caution: for printing in comments the nest level is 1 higher than actual code:
			collectedTags := collectTags(d.Fields, params.Appender.GetAPIVersion(), params.Appender.GetFlexible(), params.CommonStructs, params.NameRegistry, params.NestLevel+2)
			fieldsForVersion := d.Fields.fieldsForVersion(params.Appender.GetAPIVersion(), params.Appender.GetFlexible())
			// if we have no fields, maybe we have a common struct
			if len(fieldsForVersion) == 0 {
				fieldsForVersion = params.CommonStructs.forVersion(d.Type, params.Appender.GetAPIVersion(), params.Appender.GetFlexible())
			}
			for _, field := range fieldsForVersion {
				field.generateSchema(&schemaGeneratorMetadata{
					Appender:      newField,
					CommonStructs: params.CommonStructs,
					NestLevel:     params.NestLevel + 1,
					NameRegistry:  params.NameRegistry,
					Tags:          collectedTags,
				})
			}
		}

		parsedTemplate, err := template.New("array").Parse(fieldArrayTemplate)
		if err != nil {
			panic(err)
		}
		var output bytes.Buffer
		if err := parsedTemplate.Execute(&output, newField); err != nil {
			panic(err)
		}
		newField.Rendered = output.String()
		params.Appender.Append(newField)

	} else {

		if d.Type == "schemaTags" {

			newField.Tags = params.Tags

			parsedTemplate, err := template.New("schematags").Parse(fieldSchemaTagsTemplate)
			if err != nil {
				panic(err)
			}
			var output bytes.Buffer
			if err := parsedTemplate.Execute(&output, newField); err != nil {
				panic(err)
			}
			newField.Rendered = output.String()
			params.Appender.Append(newField)

		} else {

			if !schemaTypeSpec.IsSchema {
				newField.SchemaType = schemaTypeSpec.fullName()
			} else {
				// Caution: for printing in comments the nest level is 1 higher than actual code:
				collectedTags := collectTags(d.Fields, params.Appender.GetAPIVersion(), params.Appender.GetFlexible(), params.CommonStructs, params.NameRegistry, params.NestLevel+2)
				fieldsForVersion := d.Fields.fieldsForVersion(params.Appender.GetAPIVersion(), params.Appender.GetFlexible())
				// if we have no fields, maybe we have a common struct
				if len(fieldsForVersion) == 0 {
					fieldsForVersion = params.CommonStructs.forVersion(d.Type, params.Appender.GetAPIVersion(), params.Appender.GetFlexible())
				}
				for _, field := range fieldsForVersion {
					field.generateSchema(&schemaGeneratorMetadata{
						Appender:      newField,
						CommonStructs: params.CommonStructs,
						NestLevel:     params.NestLevel + 1,
						NameRegistry:  params.NameRegistry,
						Tags:          collectedTags,
					})
				}
			}

			parsedTemplate, err := template.New("field").Parse(fieldSimpleTypeTemplate)
			if err != nil {
				panic(err)
			}
			var output bytes.Buffer
			if err := parsedTemplate.Execute(&output, newField); err != nil {
				panic(err)
			}
			newField.Rendered = output.String()
			params.Appender.Append(newField)

		}

	}

}

func (d *field) generateTags(params *schemaGeneratorMetadata) {

	params.NameRegistry.builder.add(d.Name)
	defer params.NameRegistry.builder.pop()

	newField := &irTag{
		APIName:           d.Name,
		APIVersion:        params.Appender.GetAPIVersion(),
		Fields:            []fieldAppender{},
		Flexible:          params.Appender.GetFlexible(),
		ConstantFieldName: params.NameRegistry.builder.generate(),
		Nullable:          d.nullable(params.Appender.GetAPIVersion()),
		Whitespace: strings.Repeat("	", params.NestLevel),
	}

	schemaTypeSpec := getSchemaType(d.Type, params.Appender.GetFlexible(), newField.Nullable)

	params.NameRegistry.register(newField.ConstantFieldName, &fieldWithDoc{
		name: d.Name,
		doc:  d.About,
	})

	if schemaTypeSpec.RequiresCompactArrayWrap {

		if !schemaTypeSpec.IsSchema {
			newField.SchemaType = schemaTypeSpec.fullName()
		} else {
			fieldsForVersion := d.Fields.fieldsForVersion(params.Appender.GetAPIVersion(), params.Appender.GetFlexible())
			// if we have no fields, maybe we have a common struct
			if len(fieldsForVersion) == 0 {
				fieldsForVersion = params.CommonStructs.forVersion(d.Type, params.Appender.GetAPIVersion(), params.Appender.GetFlexible())
			}
			for _, field := range fieldsForVersion {
				field.generateSchema(&schemaGeneratorMetadata{
					Appender:      newField,
					CommonStructs: params.CommonStructs,
					NestLevel:     params.NestLevel + 1,
					NameRegistry:  params.NameRegistry,
				})
			}
		}

		parsedTemplate, err := template.New("arraycompact").Parse(fieldCompactArrayTemplate)
		if err != nil {
			panic(err)
		}
		var output bytes.Buffer
		if err := parsedTemplate.Execute(&output, newField); err != nil {
			panic(err)
		}
		newField.Rendered = output.String()
		params.Appender.Append(newField)

	} else if schemaTypeSpec.RequiresArrayWrap {

		if !schemaTypeSpec.IsSchema {
			newField.SchemaType = schemaTypeSpec.fullName()
		} else {
			fieldsForVersion := d.Fields.fieldsForVersion(params.Appender.GetAPIVersion(), params.Appender.GetFlexible())
			// if we have no fields, maybe we have a common struct
			if len(fieldsForVersion) == 0 {
				fieldsForVersion = params.CommonStructs.forVersion(d.Type, params.Appender.GetAPIVersion(), params.Appender.GetFlexible())
			}
			for _, field := range fieldsForVersion {
				field.generateSchema(&schemaGeneratorMetadata{
					Appender:      newField,
					CommonStructs: params.CommonStructs,
					NestLevel:     params.NestLevel + 1,
					NameRegistry:  params.NameRegistry,
				})
			}
		}

		parsedTemplate, err := template.New("array").Parse(fieldArrayTemplate)
		if err != nil {
			panic(err)
		}
		var output bytes.Buffer
		if err := parsedTemplate.Execute(&output, newField); err != nil {
			panic(err)
		}
		newField.Rendered = output.String()
		params.Appender.Append(newField)

	} else {

		if d.Type == "schemaTags" {
			parsedTemplate, err := template.New("schematags").Parse(fieldSchemaTagsTemplate)
			if err != nil {
				panic(err)
			}
			var output bytes.Buffer
			if err := parsedTemplate.Execute(&output, newField); err != nil {
				panic(err)
			}
			newField.Rendered = output.String()
			params.Appender.Append(newField)

		} else {

			if !schemaTypeSpec.IsSchema {
				newField.SchemaType = schemaTypeSpec.fullName()
			} else {
				fieldsForVersion := d.Fields.fieldsForVersion(params.Appender.GetAPIVersion(), params.Appender.GetFlexible())
				// if we have no fields, maybe we have a common struct
				if len(fieldsForVersion) == 0 {
					fieldsForVersion = params.CommonStructs.forVersion(d.Type, params.Appender.GetAPIVersion(), params.Appender.GetFlexible())
				}
				for _, field := range fieldsForVersion {
					field.generateSchema(&schemaGeneratorMetadata{
						Appender:      newField,
						CommonStructs: params.CommonStructs,
						NestLevel:     params.NestLevel + 1,
						NameRegistry:  params.NameRegistry,
					})
				}
			}

			parsedTemplate, err := template.New("field").Parse(fieldSimpleTypeTemplate)
			if err != nil {
				panic(err)
			}
			var output bytes.Buffer
			if err := parsedTemplate.Execute(&output, newField); err != nil {
				panic(err)
			}
			newField.Rendered = output.String()
			params.Appender.Append(newField)

		}

	}

}
