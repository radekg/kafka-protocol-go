package schema

// bound field
type boundField struct {
	def    Field
	index  int
	schema *schema
}

func (f *boundField) GetDef() Field {
	return f.def
}
