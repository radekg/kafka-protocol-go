package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"text/template"
)

type messageDefinition struct {
	APIKey           int           `json:"apiKey"`
	Type             string        `json:"type"`
	Name             string        `json:"name"`
	ValidVersions    string        `json:"validVersions"`
	FlexibleVersions string        `json:"flexibleVersions"`
	Fields           fields        `json:"fields"`
	CommonStructs    commonStructs `json:"commonStructs"`
}

func (d *messageDefinition) maxVersion() (int64, error) {
	parts := strings.Split(d.ValidVersions, "-")
	if len(parts) == 1 {
		return strconv.ParseInt(parts[0], 10, 64)
	}
	return strconv.ParseInt(parts[1], 10, 64)
}

func (d *messageDefinition) minFlexibleVersion() (int64, error) {
	if d.FlexibleVersions == "none" {
		return math.MaxInt64, nil
	}
	if d.FlexibleVersions != "" {
		return strconv.ParseInt(strings.ReplaceAll(d.FlexibleVersions, "+", ""), 10, 64)
	}
	return -1, nil
}

func (d *messageDefinition) generateSchema() string {

	max, err := d.maxVersion()
	if err != nil {
		panic(err)
	}
	minFlexible, err := d.minFlexibleVersion()
	if err != nil {
		panic(err)
	}

	nameRegistry := &fieldNameRegistry{
		builder:    &fieldNameBuilder{prefix: []string{fmt.Sprintf("Field%s", d.Name)}},
		fieldNames: map[string]*fieldWithDoc{},
	}

	generatorIr := &ir{
		APIKey:    d.APIKey,
		APIName:   d.Name,
		Versions:  []*irVersion{},
		Constants: map[string]*irConstant{},
	}

	for i := int64(0); i <= max; i = i + 1 {
		isFlexible := i >= minFlexible

		generatorIrVersion := &irVersion{
			APIKey:     d.APIKey,
			APIName:    d.Name,
			APIVersion: i,
			Flexible:   isFlexible,
		}

		collectedTags := collectTags(d.Fields, i, isFlexible, d.CommonStructs, nameRegistry, 4)

		fieldsForVersion := d.Fields.fieldsForVersion(i, isFlexible)
		// if we have no fields, maybe we have a common struct
		if len(fieldsForVersion) == 0 {
			fieldsForVersion = d.CommonStructs.forVersion(d.Type, generatorIrVersion.GetAPIVersion(), generatorIrVersion.GetFlexible())
		}
		for _, field := range fieldsForVersion {
			field.generateSchema(&schemaGeneratorMetadata{
				Appender:      generatorIrVersion,
				CommonStructs: d.CommonStructs,
				NestLevel:     3,
				NameRegistry:  nameRegistry,
				Tags:          collectedTags,
			})
		}
		generatorIr.Versions = append(generatorIr.Versions, generatorIrVersion)
	}

	generatorIr.ConstantNames = nameRegistry.sortedKeys()
	for _, key := range generatorIr.ConstantNames {
		generatorIr.Constants[key] = &irConstant{
			Doc:  nameRegistry.fieldNames[key].doc,
			Name: nameRegistry.fieldNames[key].name,
		}
	}

	parsedTemplate, err := template.New("output").Parse(outputTemplate)
	if err != nil {
		panic(err)
	}
	var output bytes.Buffer
	if err := parsedTemplate.Execute(&output, generatorIr); err != nil {
		panic(err)
	}

	return output.String()

}
