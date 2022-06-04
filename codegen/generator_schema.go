package main

import (
	"fmt"
	"strings"
)

type fieldWithDoc struct {
	name string
	doc  string
}

type fieldNameBuilder struct {
	prefix []string
}

func (d *fieldNameBuilder) add(v string) {
	d.prefix = append(d.prefix, v)
}

func (d *fieldNameBuilder) pop() {
	d.prefix = d.prefix[0 : len(d.prefix)-1]
}

func (d *fieldNameBuilder) generate() string {
	return strings.Join(d.prefix, "")
}

type fieldNameRegistry struct {
	builder    *fieldNameBuilder
	fieldNames map[string]*fieldWithDoc
}

func (d *fieldNameRegistry) register(k string, v *fieldWithDoc) {
	d.fieldNames[k] = v
}

func (d *Field) Schema(version int64, nest int, flexible bool, namRegistry *fieldNameRegistry) string {

	namRegistry.builder.add(d.Name)
	defer namRegistry.builder.pop()

	linePrefix := strings.Repeat("	", nest)

	isNullable := d.Nullable(version)
	typeSpec := GetType(d.Type, flexible, isNullable)
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
			lines = append(lines, fmt.Sprintf("%s%s},", line, typeSpec.FullName()))
		} else {
			lines = append(lines, fmt.Sprintf(`%sschema.NewSchema("%sV%d",`, line, d.Name, version))
			// Get subfields:
			fieldsForVersion := d.Fields.GetForVersion(version, flexible)
			for _, field := range fieldsForVersion {
				lines = append(lines, field.Schema(version, nest+1, flexible, namRegistry))
			}
			lines = append(lines, fmt.Sprintf("%s)},", linePrefix))
		}

	} else if typeSpec.RequiresArrayWrap {

		line := fmt.Sprintf(`%s&schema.Array{Name: %s, Ty: `, linePrefix, fullVarName)
		if !typeSpec.IsSchema {
			lines = append(lines, fmt.Sprintf("%s%s},", line, typeSpec.FullName()))
		} else {
			lines = append(lines, fmt.Sprintf(`%sschema.NewSchema("%sV%d",`, line, d.Name, version))
			// Get subfields:
			fieldsForVersion := d.Fields.GetForVersion(version, flexible)
			for _, field := range fieldsForVersion {
				lines = append(lines, field.Schema(version, nest+1, flexible, namRegistry))
			}
			lines = append(lines, fmt.Sprintf("%s)},", linePrefix))
		}

	} else {

		if d.Type == "schemaTags" {
			lines = append(lines, fmt.Sprintf(`%s&schema.SchemaTaggedFields{Name: %s},`, linePrefix, fullVarName))
		} else {
			line := fmt.Sprintf(`%s&schema.Mfield{Name: %s, Ty: `, linePrefix, fullVarName)
			if !typeSpec.IsSchema {
				line = fmt.Sprintf("%s%s},", line, typeSpec.FullName())
			}
			lines = append(lines, line)
		}

	}

	return strings.Join(lines, "\n")
}

func (d *MessageDefinition) Schema() string {

	max, err := d.MaxVersion()
	if err != nil {
		panic(err)
	}
	minFlexible, err := d.MinFlexibleVersion()
	if err != nil {
		panic(err)
	}

	nameRegistry := &fieldNameRegistry{
		builder:    &fieldNameBuilder{prefix: []string{fmt.Sprintf("Field%s", d.Name)}},
		fieldNames: map[string]*fieldWithDoc{},
	}

	lines := []string{
		"package messages",
		"",
		`import "github.com/radekg/kafka-protocol-go/schema"`,
		"",
		fmt.Sprintf("func init%d%s() []schema.Schema {", d.ApiKey, d.Name),
		"",
		fmt.Sprintf("	return []schema.Schema{"),
		"",
	}

	for i := int64(0); i <= max; i = i + 1 {

		lines = append(lines, fmt.Sprintf(`		// Message: %s, API Key: %d, Version: %d`, d.Name, d.ApiKey, i))

		line := fmt.Sprintf(`		schema.NewSchema("%sv%d"`, d.Name, i)
		isFlexible := i >= minFlexible
		fieldsForVersion := d.Fields.GetForVersion(i, isFlexible)
		if len(fieldsForVersion) == 0 {
			line = fmt.Sprintf("%s),", line)
			lines = append(lines, line)
		} else {

			line = fmt.Sprintf("%s, ", line)
			lines = append(lines, line)

			for _, field := range fieldsForVersion {
				lines = append(lines, field.Schema(i, 3, isFlexible, nameRegistry))
			}

			lines = append(lines, "		),") // end top level schema
			lines = append(lines, "")
		}

	}

	// end wrapping function:
	lines = append(lines, "	}")
	lines = append(lines, "}")
	lines = append(lines, "")
	lines = append(lines, "const (")
	for name, value := range nameRegistry.fieldNames {
		lines = append(lines, fmt.Sprintf(`	// %s is: %s`, name, value.doc))
		lines = append(lines, fmt.Sprintf(`	%s = "%s"`, name, value.name))
	}
	lines = append(lines, ")")
	lines = append(lines, "")

	return strings.Join(lines, "\n")

}

type schemaGenerator struct {
	originalInput string
	sourcePath    string
}

func (d *schemaGenerator) generate(message *MessageDefinition) string {
	lines := []string{message.Schema()}
	lines = append(lines, fmt.Sprintf("// Generated from Apache Kafka source code file: %s", d.sourcePath))
	lines = append(lines, fmt.Sprintf("const original%sInput = `%s`", message.Name, strings.ReplaceAll(d.originalInput, "`", "'")))
	lines = append(lines, "")
	return strings.Join(lines, "\n")
}
