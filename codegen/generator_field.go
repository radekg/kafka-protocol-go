package main

import (
	"fmt"
	"strconv"
	"strings"

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
	Name             string `json:"name"`
	Type             string `json:"type"`
	Versions         string `json:"versions"`
	NullableVersions string `json:"nullableVersions"`
	About            string `json:"About"`
	Fields           fields `json:"fields"`
}

func (d *field) appliesToVersion(v int64) bool {
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

func (d *field) generateSchema(version int64, nest int, flexible bool, namRegistry *fieldNameRegistry) string {

	namRegistry.builder.add(d.Name)
	defer namRegistry.builder.pop()

	linePrefix := strings.Repeat("	", nest)

	isNullable := d.nullable(version)
	typeSpec := getSchemaType(d.Type, flexible, isNullable)
	lines := []string{}

	fullVarName := namRegistry.builder.generate()
	namRegistry.register(fullVarName, &fieldWithDoc{
		name: d.Name,
		doc:  d.About,
	})

	if typeSpec.RequiresCompactArrayWrap {

		//namRegistry.register(fullVarName, d.Name)
		line := fmt.Sprintf(`%s&schema.ArrayCompact{Name: %s, Ty: `, linePrefix, fullVarName)

		if !typeSpec.IsSchema {
			lines = append(lines, fmt.Sprintf("%s%s},", line, typeSpec.fullName()))
		} else {
			lines = append(lines, fmt.Sprintf(`%sschema.NewSchema("%sV%d",`, line, d.Name, version))
			// Get subfields:
			fieldsForVersion := d.Fields.forVersion(version, flexible)
			for _, field := range fieldsForVersion {
				lines = append(lines, field.generateSchema(version, nest+1, flexible, namRegistry))
			}
			lines = append(lines, fmt.Sprintf("%s)},", linePrefix))
		}

	} else if typeSpec.RequiresArrayWrap {

		line := fmt.Sprintf(`%s&schema.Array{Name: %s, Ty: `, linePrefix, fullVarName)
		if !typeSpec.IsSchema {
			lines = append(lines, fmt.Sprintf("%s%s},", line, typeSpec.fullName()))
		} else {
			lines = append(lines, fmt.Sprintf(`%sschema.NewSchema("%sV%d",`, line, d.Name, version))
			// Get subfields:
			fieldsForVersion := d.Fields.forVersion(version, flexible)
			for _, field := range fieldsForVersion {
				lines = append(lines, field.generateSchema(version, nest+1, flexible, namRegistry))
			}
			lines = append(lines, fmt.Sprintf("%s)},", linePrefix))
		}

	} else {

		if d.Type == "schemaTags" {
			lines = append(lines, fmt.Sprintf(`%s&schema.SchemaTaggedFields{Name: %s},`, linePrefix, fullVarName))
		} else {
			line := fmt.Sprintf(`%s&schema.Mfield{Name: %s, Ty: `, linePrefix, fullVarName)
			if !typeSpec.IsSchema {
				lines = append(lines, fmt.Sprintf("%s%s},", line, typeSpec.fullName()))
			} else {
				lines = append(lines, fmt.Sprintf(`%sschema.NewSchema("%sV%d",`, line, d.Name, version))
				// Get subfields:
				fieldsForVersion := d.Fields.forVersion(version, flexible)
				for _, field := range fieldsForVersion {
					lines = append(lines, field.generateSchema(version, nest+1, flexible, namRegistry))
				}
				lines = append(lines, fmt.Sprintf("%s)},", linePrefix))
			}
		}

	}

	return strings.Join(lines, "\n")
}
