package main

type fieldAppender interface {
	Append(field *irField)
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
	ConstantFieldName string
	Fields            []*irField
	Flexible          bool
	Nullable          bool
	Rendered          string
	SchemaType        string
	Whitespace        string
}

func (d *irField) Append(field *irField) {
	d.Fields = append(d.Fields, field)
}

func (d *irField) GetAPIVersion() int64 {
	return d.APIVersion
}

func (d *irField) GetFlexible() bool {
	return d.Flexible
}

type irVersion struct {
	APIKey     int
	APIName    string
	APIVersion int64
	Flexible   bool
	Fields     []*irField
}

func (d *irVersion) Append(field *irField) {
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