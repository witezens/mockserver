package resolver

import (
	"fmt"
	"net/url"
)

type MockRule struct {
	Param    string
	Template string
	Source   string
}

type MockResolver struct {
	Rules map[string][]MockRule
}

func (r *MockResolver) ResolveFile(service, resource, method string, body map[string]interface{}, query url.Values) string {
	key := fmt.Sprintf("%s_%s", service, resource)
	rules := r.Rules[key]

	for _, rule := range rules {
		if rule.Source == "query" {
			if values, ok := query[rule.Param]; ok && len(values) > 0 {
				return fmt.Sprintf(rule.Template, resource, values[0], method)
			}
		} else {
			if value, ok := body[rule.Param]; ok {
				return fmt.Sprintf(rule.Template, resource, value, method)
			}
		}
	}

	return fmt.Sprintf("%s.%s.json", resource, method)
}
