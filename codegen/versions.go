package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func appliesToVersion(name, rawVersionString string, v int64) bool {

	if strings.HasSuffix(rawVersionString, "+") {
		minVersion, err := strconv.ParseInt(strings.TrimSuffix(rawVersionString, "+"), 10, 64)
		if err != nil {
			errors.Wrap(err, fmt.Sprintf("failed parsing versions with + suffix: %s", name))
		}
		if v >= minVersion {
			return true
		}
	}
	if strings.Contains(rawVersionString, "-") {
		parts := strings.Split(rawVersionString, "-")
		minVersion, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			errors.Wrap(err, fmt.Sprintf("failed parsing versions - delimited: %s", name))
		}
		maxVersion, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			errors.Wrap(err, fmt.Sprintf("failed parsing versions - delimited: %s", name))
		}
		if v >= minVersion && v <= maxVersion {
			return true
		}
	}
	return false
}
