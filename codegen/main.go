package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type genTypesFlags []string

func (i *genTypesFlags) String() string {
	return strings.Join(*i, ",")
}

func (i *genTypesFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var kafkaSourcesRoot string
var genType string

func main() {

	flag.StringVar(&kafkaSourcesRoot, "kafka-source-root", "", "The path to the root of Kafka source code")
	flag.StringVar(&genType, "gen-type", "", "The type to generate")
	flag.Parse()

	sourcePath := filepath.Join("clients/src/main/resources/common/message", fmt.Sprintf("%s.json", genType))
	path := filepath.Join(kafkaSourcesRoot, sourcePath)

	statResult, err := os.Stat(path)
	if err != nil {
		panic(errors.Wrapf(err, "Could not stat the requested gen-type for Kafka root: %s", kafkaSourcesRoot))
	}
	if statResult.IsDir() {
		panic(errors.Wrapf(err, "The requested gen-type '%s' appears to be a directory", path))
	}

	// Read everything, there shouldn't be that much...
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		panic(errors.Wrapf(err, "Could not read file for gen-type at '%s'", path))
	}

	reader := bufio.NewScanner(bytes.NewReader(bs))
	jsonLines := []string{}
	for reader.Scan() {
		line := reader.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "//") {
			continue
		}
		if strings.TrimSpace(line) == "" {
			continue
		}
		jsonLines = append(jsonLines, line)
	}

	fullMessage := strings.Join(jsonLines, "\n")
	msg := &messageDefinition{}
	if err := json.Unmarshal([]byte(fullMessage), msg); err != nil {
		panic(errors.Wrapf(err, "The schema for gen-type '%s' could not be decoded as JSON", path))
	}

	gen := &schemaGenerator{
		originalInput: string(bs),
		sourcePath:    sourcePath,
	}

	fmt.Println(gen.generate(msg))

}
