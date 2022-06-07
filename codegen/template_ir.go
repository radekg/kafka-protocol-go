package main

type fieldAppender interface {
	Append(field fieldAppender)
	GetAPIVersion() int64
	GetFlexible() bool
}

type irConstant struct {
	Doc  string
	Name string
}

type irField struct {
	APIName           string
	APIVersion        int64
	APIType           string
	ConstantFieldName string
	Fields            []fieldAppender
	Flexible          bool
	Nullable          bool
	Rendered          string
	SchemaType        string
	Whitespace        string
	Tags              []*irTag
}

func (d *irField) Append(field fieldAppender) {
	d.Fields = append(d.Fields, field)
}

func (d *irField) GetAPIVersion() int64 {
	return d.APIVersion
}

func (d *irField) GetFlexible() bool {
	return d.Flexible
}

type irTag struct {
	APIName           string
	APIVersion        int64
	APIType           string
	ConstantFieldName string
	Fields            []fieldAppender
	Flexible          bool
	Nullable          bool
	Rendered          string
	SchemaType        string
	Tag               int
	Whitespace        string
}

func (d *irTag) Append(field fieldAppender) {
	d.Fields = append(d.Fields, field)
}

func (d *irTag) GetAPIVersion() int64 {
	return d.APIVersion
}

func (d *irTag) GetFlexible() bool {
	return d.Flexible
}

type irVersion struct {
	APIKey     int
	APIName    string
	APIVersion int64
	Flexible   bool
	Fields     []fieldAppender
}

func (d *irVersion) Append(field fieldAppender) {
	d.Fields = append(d.Fields, field)
}

func (d *irVersion) GetAPIVersion() int64 {
	return d.APIVersion
}

func (d *irVersion) GetFlexible() bool {
	return d.Flexible
}

type ir struct {
	APIKey        int
	APIName       string
	Versions      []*irVersion
	ConstantNames []string
	Constants     map[string]*irConstant
}
