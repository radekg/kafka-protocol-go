package main

import (
	"bytes"
	"strings"
	"text/template"
)

func collectTags(f fields, version int64, flexible bool,
	cs commonStructs,
	nameRegistry *fieldNameRegistry, nestLevel int) []*irTag {
	tagsForVersion := f.tagsForVersion(version, flexible)
	respTags := []*irTag{}
	for _, tag := range tagsForVersion {
		newTag := &irTag{
			APIName:           tag.Name,
			APIVersion:        version,
			Fields:            []fieldAppender{},
			Flexible:          flexible,
			ConstantFieldName: nameRegistry.builder.generate(),
			Nullable:          false,
			Tag:               *tag.Tag,
			Whitespace: strings.Repeat("	", nestLevel),
		}
		tag.generateSchema(&schemaGeneratorMetadata{
			Appender:      newTag,
			CommonStructs: cs,
			NestLevel:     nestLevel,
			NameRegistry:  nameRegistry,
		})

		parsedTemplate, err := template.New("tag").Parse(fieldTagTemplate)
		if err != nil {
			panic(err)
		}
		var output bytes.Buffer
		if err := parsedTemplate.Execute(&output, newTag); err != nil {
			panic(err)
		}
		newTag.Rendered = output.String()

		respTags = append(respTags, newTag)
	}
	return respTags
}
