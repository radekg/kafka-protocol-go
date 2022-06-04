package main

import (
	"fmt"
	"strings"
)

type typeDef struct {
	Name             string
	SupportsFlexible bool
}

type typeDescriptor struct {
	Name                     string
	IsSchema                 bool
	RequiresArrayWrap        bool
	RequiresCompactArrayWrap bool
	SupportsFlexible         bool
}

func (d *typeDescriptor) fullName() string {
	return fmt.Sprintf("schema.%s", d.Name)
}

func (d *typeDescriptor) string() string {
	if !d.IsSchema {
		return fmt.Sprintf("{name: %s, array: %v, compact: %v}", d.fullName(), d.RequiresArrayWrap, d.RequiresCompactArrayWrap)
	}
	return fmt.Sprintf("*schema.Schema{..., array: %v, compact: %v}", d.RequiresArrayWrap, d.RequiresCompactArrayWrap)
}

var schemaTypeMap = map[string]typeDef{
	"bool":       {Name: "TypeBool"},
	"bytes":      {Name: "TypeBytes", SupportsFlexible: true},
	"int8":       {Name: "TypeInt8"},
	"int16":      {Name: "TypeInt16"},
	"int32":      {Name: "TypeInt32"},
	"int64":      {Name: "TypeInt64"},
	"float64":    {Name: "TypeFloat64"},
	"records":    {Name: "TypeBytes", SupportsFlexible: true},
	"schemaTags": {Name: "SchemaTags"},
	"string":     {Name: "TypeStr", SupportsFlexible: true},
	"uint16":     {Name: "TypeUint16"},
	"uint32":     {Name: "TypeUint32"},
	"uuid":       {Name: "TypeUuid"},
}

func getSchemaType(t string, flexible, nullable bool) *typeDescriptor {
	if strings.HasPrefix(t, "[]") {
		t = strings.TrimPrefix(t, "[]")
		if mt, ok := schemaTypeMap[t]; ok {
			if flexible {
				return &typeDescriptor{
					Name:             fmt.Sprintf("%sCompactArray", mt.Name),
					SupportsFlexible: mt.SupportsFlexible,
				}
			}
			return &typeDescriptor{
				Name: fmt.Sprintf("%sArray", mt.Name),
				// IsArray:          true,
				SupportsFlexible: mt.SupportsFlexible,
			}
		}
		if string(t[0]) != strings.ToLower(string(t[0])) {
			if flexible {
				return &typeDescriptor{
					Name:                     t,
					IsSchema:                 true,
					RequiresArrayWrap:        true,
					RequiresCompactArrayWrap: true,
					SupportsFlexible:         true,
				}
			}
			return &typeDescriptor{
				Name:              t,
				IsSchema:          true,
				RequiresArrayWrap: true,
				SupportsFlexible:  true,
			}
		}
	}
	if mt, ok := schemaTypeMap[t]; ok {
		if flexible && nullable {
			return &typeDescriptor{
				Name:             fmt.Sprintf("%sCompactNullable", mt.Name),
				SupportsFlexible: mt.SupportsFlexible,
			}
		}
		if !flexible && nullable {
			return &typeDescriptor{
				Name:             fmt.Sprintf("%sNullable", mt.Name),
				SupportsFlexible: mt.SupportsFlexible,
			}
		}
		if flexible && !nullable && mt.SupportsFlexible {
			return &typeDescriptor{
				Name:             fmt.Sprintf("%sCompact", mt.Name),
				SupportsFlexible: mt.SupportsFlexible,
			}
		}
		return &typeDescriptor{
			Name:             mt.Name,
			SupportsFlexible: mt.SupportsFlexible,
		}
	}
	if string(t[0]) != strings.ToLower(string(t[0])) {
		return &typeDescriptor{
			Name:              t,
			IsSchema:          true,
			RequiresArrayWrap: false,
			SupportsFlexible:  false,
		}
	}
	return &typeDescriptor{
		Name:             fmt.Sprintf("!!! unknown type: %s", t),
		SupportsFlexible: false,
	}
}
