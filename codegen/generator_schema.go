package main

import (
	"fmt"
	"strings"
)

type schemaGeneratorMetadata struct {
	CommonStructs commonStructs
	Appender      fieldAppender
	NestLevel     int
	NameRegistry  *fieldNameRegistry
	Tags          []*irTag
}

type schemaGenerator struct {
	originalInput string
	sourcePath    string
}

func (d *schemaGenerator) generate(message *messageDefinition) string {
	lines := []string{message.generateSchema()}
	lines = append(lines, fmt.Sprintf("// Generated from Apache Kafka source code file: %s", d.sourcePath))
	lines = append(lines, fmt.Sprintf("const original%sInput = `%s`", message.Name, strings.ReplaceAll(d.originalInput, "`", "'")))
	lines = append(lines, "")
	return strings.Join(lines, "\n")
}
