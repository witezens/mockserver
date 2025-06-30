package utils

import "fmt"

func BuildMockPath(service, resource, method string, version string) string {
	if version == "" {
		return fmt.Sprintf("%s/%s.%s.json", service, resource, method)
	}
	return fmt.Sprintf("%s/%s/%s.%s.json", service, version, resource, method)
}

func BuildMockPathWithParam(service, resource, param string, value interface{}, method string, version string) string {
	if version == "" {
		return fmt.Sprintf("%s/%s__%s_%v.%s.json", service, resource, param, value, method)
	}
	return fmt.Sprintf("%s/%s/%s__%s_%v.%s.json", service, version, resource, param, value, method)
}
