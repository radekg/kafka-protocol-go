package messages

import "github.com/radekg/kafka-protocol-go/schema"

var headers = []schema.Schema{}

func initHeaderRegistry() {
	headers = initHeaders()
}

func init() {
	initHeaderRegistry()
}
