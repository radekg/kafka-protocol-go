package main

import (
	"sort"
	"strings"
)

type fieldNameBuilder struct {
	prefix []string
}

func (d *fieldNameBuilder) add(v string) {
	d.prefix = append(d.prefix, v)
}

func (d *fieldNameBuilder) pop() {
	d.prefix = d.prefix[0 : len(d.prefix)-1]
}

func (d *fieldNameBuilder) generate() string {
	return strings.Join(d.prefix, "")
}

type fieldNameRegistry struct {
	builder    *fieldNameBuilder
	fieldNames map[string]*fieldWithDoc
}

func (d *fieldNameRegistry) register(k string, v *fieldWithDoc) {
	d.fieldNames[k] = v
}

func (d *fieldNameRegistry) sortedKeys() []string {
	keys := []string{}
	for key := range d.fieldNames {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
