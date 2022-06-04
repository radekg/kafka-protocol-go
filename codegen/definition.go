package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
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

func (d *typeDescriptor) FullName() string {
	return fmt.Sprintf("schema.%s", d.Name)
}

func (d *typeDescriptor) String() string {
	if !d.IsSchema {
		return fmt.Sprintf("{name: %s, array: %v, compact: %v}", d.FullName(), d.RequiresArrayWrap, d.RequiresCompactArrayWrap)
	}
	return fmt.Sprintf("*schema.Schema{..., array: %v, compact: %v}", d.RequiresArrayWrap, d.RequiresCompactArrayWrap)
}

var typeMap = map[string]typeDef{
	"bool":       typeDef{Name: "TypeBool"},
	"bytes":      typeDef{Name: "TypeBytes", SupportsFlexible: true},
	"int8":       typeDef{Name: "TypeInt8"},
	"int16":      typeDef{Name: "TypeInt16"},
	"int32":      typeDef{Name: "TypeInt32"},
	"int64":      typeDef{Name: "TypeInt64"},
	"float64":    typeDef{Name: "TypeFloat64"},
	"records":    typeDef{Name: "TypeBytes", SupportsFlexible: true},
	"schemaTags": typeDef{Name: "SchemaTags"},
	"string":     typeDef{Name: "TypeStr", SupportsFlexible: true},
	"uint16":     typeDef{Name: "TypeUint16"},
	"uint32":     typeDef{Name: "TypeUint32"},
	"uuid":       typeDef{Name: "TypeUuid"},
}

func GetType(t string, flexible, nullable bool) *typeDescriptor {
	if strings.HasPrefix(t, "[]") {
		t = strings.TrimPrefix(t, "[]")
		if mt, ok := typeMap[t]; ok {
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
	if mt, ok := typeMap[t]; ok {
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
	return &typeDescriptor{
		Name:             fmt.Sprintf("!!! unknown type: %s", t),
		SupportsFlexible: false,
	}
}

type Fields []Field

func (d Fields) GetForVersion(v int64, flexible bool) Fields {
	resp := Fields{}
	for _, f := range d {
		if f.AppliesToVersion(v) {
			resp = append(resp, f)
		}
	}
	if len(resp) > 0 && flexible {
		resp = append(resp, Field{
			Name:             "Tags",
			Type:             "schemaTags",
			Versions:         fmt.Sprintf("%d+", v),
			NullableVersions: "",
			About:            "The tagged fields.",
			Fields:           Fields{},
		})
	}
	return resp
}

type Field struct {
	Name             string `json:"name"`
	Type             string `json:"type"`
	Versions         string `json:"versions"`
	NullableVersions string `json:"nullableVersions"`
	About            string `json:"About"`
	Fields           Fields `json:"fields"`
}

func (d *Field) AppliesToVersion(v int64) bool {
	if strings.HasSuffix(d.Versions, "+") {
		minVersion, err := strconv.ParseInt(strings.TrimSuffix(d.Versions, "+"), 10, 64)
		if err != nil {
			errors.Wrap(err, fmt.Sprintf("failed parsing versions with + suffix: %s", d.Name))
		}
		if v >= minVersion {
			return true
		}
	}
	if strings.Contains(d.Versions, "-") {
		parts := strings.Split(d.Versions, "-")
		minVersion, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			errors.Wrap(err, fmt.Sprintf("failed parsing versions - delimited: %s", d.Name))
		}
		maxVersion, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			errors.Wrap(err, fmt.Sprintf("failed parsing versions - delimited: %s", d.Name))
		}
		if v >= minVersion && v <= maxVersion {
			return true
		}
	}
	return false
}

func (d *Field) Nullable(v int64) bool {
	if d.NullableVersions == "" {
		return false
	}
	minVersion, err := strconv.ParseInt(strings.TrimSuffix(d.NullableVersions, "+"), 10, 64)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("failed parsing versions with + suffix: %s", d.Name))
	}
	if v >= minVersion {
		return true
	}
	return false
}

type MessageDefinition struct {
	ApiKey           int    `json:"apiKey"`
	Type             string `json:"type"`
	Name             string `json:"name"`
	ValidVersions    string `json:"validVersions"`
	FlexibleVersions string `json:"flexibleVersions"`
	Fields           Fields `json:"fields"`
}

func (d *MessageDefinition) MaxVersion() (int64, error) {
	parts := strings.Split(d.ValidVersions, "-")
	if len(parts) == 1 {
		return strconv.ParseInt(parts[0], 10, 64)
	}
	return strconv.ParseInt(parts[1], 10, 64)
}

func (d *MessageDefinition) MinFlexibleVersion() (int64, error) {
	if d.FlexibleVersions != "" {
		if d.FlexibleVersions == "none" {
			return math.MaxInt64, nil
		}
		return strconv.ParseInt(strings.ReplaceAll(d.FlexibleVersions, "+", ""), 10, 64)
	}
	return -1, nil
}
