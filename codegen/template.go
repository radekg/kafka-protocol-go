package main

const fieldArrayTemplate = `{{- $fieldsLen := len .Fields }}
{{- if eq $fieldsLen 0 }}
{{ .Whitespace }}&schema.Array{Name: {{ .ConstantFieldName }}, Ty: {{ .SchemaType }}},
{{- else }}
{{ .Whitespace }}&schema.Array{Name: {{ .ConstantFieldName }}, Ty: schema.NewSchema("{{ .APIType }}:v{{ .APIVersion }}",{{ range $index, $field := .Fields }}{{ $field.Rendered }}{{ end }}
{{ .Whitespace }})},{{ end }}`

const fieldCompactArrayTemplate = `{{- $fieldsLen := len .Fields }}
{{- if eq $fieldsLen 0 }}
{{ .Whitespace }}&schema.ArrayCompact{Name: {{ .ConstantFieldName }}, Ty: {{ .SchemaType }}},
{{- else }}
{{ .Whitespace }}&schema.ArrayCompact{Name: {{ .ConstantFieldName }}, Ty: schema.NewSchema("{{ .APIType }}:v{{ .APIVersion }}",{{ range $index, $field := .Fields }}{{ $field.Rendered }}{{ end }}
{{ .Whitespace }})},{{ end }}`

const fieldSimpleTypeTemplate = `{{- $fieldsLen := len .Fields }}
{{- if eq $fieldsLen 0 }}
{{ .Whitespace }}&schema.Mfield{Name: {{ .ConstantFieldName }}, Ty: {{ .SchemaType }}},
{{- else }}
{{ .Whitespace }}&schema.Mfield{Name: {{ .ConstantFieldName }}, Ty: schema.NewSchema("{{ .APIType }}:v{{ .APIVersion }}",{{ range $index, $field := .Fields }}{{ $field.Rendered }}{{ end }}
{{ .Whitespace }})},{{ end }}`

const fieldSchemaTagsTemplate = `
{{ .Whitespace }}&schema.SchemaTaggedFields{Name: {{ .ConstantFieldName }}},
{{- $tagsLen := len .Tags }}
{{- if gt $tagsLen 0 }}
{{ .Whitespace }}/** Applicable tags:
{{ range $index, $tag := .Tags }}{{ .Whitespace }}{{ $tag.Rendered }}{{ end }}
{{ .Whitespace }}**/
{{- end }}
`

const fieldTagTemplate = `{{- $root := . }}
{{- range $index, $field := .Fields }}
{{ $root.Whitespace }}{{ $root.Tag }}: {{ $root.APIName }} (type: {{ $root.APIType }}) = {{ $field.Rendered }}
{{- end }}
`

const outputTemplate = `package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init{{.APIKey}}{{.APIName}}() []schema.Schema {

	return []schema.Schema{

		{{- range $index, $version := .Versions}}
		// Message: {{ $version.APIName }}, API Key: {{ $version.APIKey }}, Version: {{ $version.APIVersion }}
		{{- $fieldsLen := len $version.Fields }}
		{{- if eq $fieldsLen 0 }}
		schema.NewSchema("{{ $version.APIName }}:v{{ $version.APIVersion }}"),
		{{ else }}
		schema.NewSchema("{{ $version.APIName }}:v{{ $version.APIVersion }}",{{ range $index, $field := .Fields }}{{ $field.Rendered }}{{ end }}
		),
		{{ end }}
		{{ end }}
	}

}

const (
	{{- $constants := .Constants }}
	{{ range $index, $key := .ConstantNames }}
	{{- $huh := (index $constants $key) }}
	// {{ $key }} is: {{ $huh.Doc }}
	{{ $key }} = "{{ $huh.Name }}"
	{{end}}
)

`
