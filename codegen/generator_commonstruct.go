package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type commonStructs []commonStruct

type commonStruct struct {
	Name     string `json:"name"`
	Versions string `json:"versions"`
	Fields   fields `json:"fields"`
}

func (d *commonStruct) appliesToVersion(v int64) bool {
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

func (d commonStructs) forVersion(t string, v int64, flexible bool) fields {

	if strings.HasPrefix(t, "[]") {
		t = strings.TrimPrefix(t, "[]")
	}

	for _, cs := range d {
		if cs.Name == t {
			if !cs.appliesToVersion(v) {
				return fields{}
			}
			resp := fields{}
			for _, f := range cs.Fields {
				if f.appliesToVersion(v) {
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
