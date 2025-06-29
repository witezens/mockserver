package utils

import "fmt"

// BuildKey builds the standardized key for handler registry
func BuildKey(service, resource string) string {
	return fmt.Sprintf("%s/%s", service, resource)
}
