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

func (d fields) forVersion(v int64, flexible bool) fields {
	resp := fields{}
	for _, f := range d {
		if f.appliesToVersion(v) {
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

func (d *field) appliesToVersion(v int64) bool {

	// a tag is not a field
	if d.Tag != nil {
		return false
	}

	if strings.HasSuffix(d.Versions, "+") {
		minVersion, err := strconv.ParseInt(strings.TrimSuffix(d.Versions, "+"), 10, 64)
		if err != nil {
			errors.Wrap(err, fmt.Sprintf("failed parsing versions with + suffix: %s", d.Name))
		}
		if v >= minVersion {
			return true
		}
	}
	if strings.Contains(d.Versions, "-") {
		parts := strings.Split(d.Versions, "-")
		minVersion, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			errors.Wrap(err, fmt.Sprintf("failed parsing versions - delimited: %s", d.Name))
		}
		maxVersion, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			errors.Wrap(err, fmt.Sprintf("failed parsing versions - delimited: %s", d.Name))
		}
		if v >= minVersion && v <= maxVersion {
			return true
		}
	}
	return false
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

func (d *field) generateSchema(commons commonStructs, parent fieldAppender, nest int, namRegistry *fieldNameRegistry) {

	namRegistry.builder.add(d.Name)
	defer namRegistry.builder.pop()

	newField := &irField{
		APIName:           d.Name,
		APIVersion:        parent.GetAPIVersion(),
		Fields:            []*irField{},
		Flexible:          parent.GetFlexible(),
		ConstantFieldName: namRegistry.builder.generate(),
		Nullable:          d.nullable(parent.GetAPIVersion()),
		Whitespace: strings.Repeat("	", nest),
	}

	schemaTypeSpec := getSchemaType(d.Type, parent.GetFlexible(), newField.Nullable)

	namRegistry.register(newField.ConstantFieldName, &fieldWithDoc{
		name: d.Name,
		doc:  d.About,
	})

	if schemaTypeSpec.RequiresCompactArrayWrap {

		if !schemaTypeSpec.IsSchema {
			newField.SchemaType = schemaTypeSpec.fullName()
		} else {
			fieldsForVersion := d.Fields.forVersion(parent.GetAPIVersion(), parent.GetFlexible())
			// if we have no fields, maybe we have a common struct
			if len(fieldsForVersion) == 0 {
				fieldsForVersion = commons.forVersion(d.Type, parent.GetAPIVersion(), parent.GetFlexible())
			}
			for _, field := range fieldsForVersion {
				field.generateSchema(commons, newField, nest+1, namRegistry)
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
		parent.Append(newField)

	} else if schemaTypeSpec.RequiresArrayWrap {

		if !schemaTypeSpec.IsSchema {
			newField.SchemaType = schemaTypeSpec.fullName()
		} else {
			fieldsForVersion := d.Fields.forVersion(parent.GetAPIVersion(), parent.GetFlexible())
			// if we have no fields, maybe we have a common struct
			if len(fieldsForVersion) == 0 {
				fieldsForVersion = commons.forVersion(d.Type, parent.GetAPIVersion(), parent.GetFlexible())
			}
			for _, field := range fieldsForVersion {
				field.generateSchema(commons, newField, nest+1, namRegistry)
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
		parent.Append(newField)

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
			parent.Append(newField)

		} else {

			if !schemaTypeSpec.IsSchema {
				newField.SchemaType = schemaTypeSpec.fullName()
			} else {
				fieldsForVersion := d.Fields.forVersion(parent.GetAPIVersion(), parent.GetFlexible())
				// if we have no fields, maybe we have a common struct
				if len(fieldsForVersion) == 0 {
					fieldsForVersion = commons.forVersion(d.Type, parent.GetAPIVersion(), parent.GetFlexible())
				}
				for _, field := range fieldsForVersion {
					field.generateSchema(commons, newField, nest+1, namRegistry)
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
			parent.Append(newField)

		}

	}

}
