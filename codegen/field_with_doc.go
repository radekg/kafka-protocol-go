package main

// Carries over the field nae and its documentation extracted from input JSON.
// This information is used to generate constants at the end of each file.
type fieldWithDoc struct {
	name string
	doc  string
}
