package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type messageDefinition struct {
	APIKey           int    `json:"apiKey"`
	Type             string `json:"type"`
	Name             string `json:"name"`
	ValidVersions    string `json:"validVersions"`
	FlexibleVersions string `json:"flexibleVersions"`
	Fields           fields `json:"fields"`
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

	lines := []string{
		"package messages",
		"",
		`import "github.com/radekg/kafka-protocol-go/schema"`,
		"",
		fmt.Sprintf("func init%d%s() []schema.Schema {", d.APIKey, d.Name),
		"",
		fmt.Sprintf("	return []schema.Schema{"),
		"",
	}

	for i := int64(0); i <= max; i = i + 1 {

		lines = append(lines, fmt.Sprintf(`		// Message: %s, API Key: %d, Version: %d`, d.Name, d.APIKey, i))

		line := fmt.Sprintf(`		schema.NewSchema("%sv%d"`, d.Name, i)
		isFlexible := i >= minFlexible
		fieldsForVersion := d.Fields.forVersion(i, isFlexible)
		if len(fieldsForVersion) == 0 {
			line = fmt.Sprintf("%s),", line)
			lines = append(lines, line)
		} else {

			line = fmt.Sprintf("%s, ", line)
			lines = append(lines, line)

			for _, field := range fieldsForVersion {
				lines = append(lines, field.generateSchema(i, 3, isFlexible, nameRegistry))
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
	sortedKeys := nameRegistry.sortedKeys()
	for _, key := range sortedKeys {
		lines = append(lines, fmt.Sprintf(`	// %s is: %s`, key, nameRegistry.fieldNames[key].doc))
		lines = append(lines, fmt.Sprintf(`	%s = "%s"`, key, nameRegistry.fieldNames[key].name))
	}
	lines = append(lines, ")")
	lines = append(lines, "")

	return strings.Join(lines, "\n")

}
