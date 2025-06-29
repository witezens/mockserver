package utils

import "net/url"

func ToURLValues(m map[string]string) url.Values {
	values := url.Values{}
	for k, v := range m {
		values.Add(k, v)
	}
	return values
}
