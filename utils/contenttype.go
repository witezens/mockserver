package utils

import (
	"strings"
)

func GetContentType(filename string) string {
	switch {
	case strings.HasSuffix(filename, ".json"):
		return "json"
	case strings.HasSuffix(filename, ".xml"):
		return "xml"
	case strings.HasSuffix(filename, ".txt"):
		return "plain"
	default:
		return "octet-stream"
	}
}
