package main

import (
	"fmt"
	"strings"
)

type commonStructs []commonStruct

type commonStruct struct {
	Name     string `json:"name"`
	Versions string `json:"versions"`
	Fields   fields `json:"fields"`
}

func (d commonStructs) forVersion(t string, v int64, flexible bool) fields {

	if strings.HasPrefix(t, "[]") {
		t = strings.TrimPrefix(t, "[]")
	}

	for _, cs := range d {
		if cs.Name == t {
			if !appliesToVersion(cs.Name, cs.Versions, v) {
				return fields{}
			}
			resp := fields{}
			for _, f := range cs.Fields {
				if f.Tag == nil && appliesToVersion(f.Name, f.Versions, v) {
					resp = append(resp, f)
				}
			}
			if len(resp) > 0 && flexible {
				resp = append(resp, field{
					Name:             "Tags",
					Type:             "schemaTags",
					Versions:         fmt.Sprintf("%d+", v),
					NullableVersions: "",
					About:            "The tagged fields.",
					Fields:           fields{},
				})
			}
			return resp
		}
	}

	return fields{}

}
