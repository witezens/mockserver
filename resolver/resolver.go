package resolver

import (
	"fmt"
	"mock-server/mockcache"
	"mock-server/utils"
	"net/url"
)

const defaultVersion = "api/v1"

type MockRule struct {
	Param     string
	Template  string
	Source    string // body or query
	Versioned bool
}

type MockResolver struct {
	Rules map[string][]MockRule
}

// fileExists check if file exists in cache (Raw or Parsed)
func fileExists(filename string) bool {
	_, inRaw := mockcache.GlobalCache.Raw[filename]
	_, inParsed := mockcache.GlobalCache.Parsed[filename]
	return inRaw || inParsed
}

func (r *MockResolver) ResolveFile(service, resource, method string, body map[string]interface{}, query url.Values) string {
	key := fmt.Sprintf("%s_%s", service, resource)
	rules := r.Rules[key]

	for _, rule := range rules {
		version := defaultVersion
		if !rule.Versioned {
			version = ""
		}

		if rule.Source == "query" {
			if values, ok := query[rule.Param]; ok && len(values) > 0 {
				candidate := utils.BuildMockPathWithParam(service, resource, rule.Param, values[0], method, version)
				if fileExists(candidate) {
					utils.Logger.Infof("📄 Matched rule (query param): %s", candidate)
					return candidate
				}
			}
		} else { // source = body
			if value, ok := body[rule.Param]; ok {
				candidate := utils.BuildMockPathWithParam(service, resource, rule.Param, value, method, version)
				if fileExists(candidate) {
					utils.Logger.Infof("📄 Matched rule (body param): %s", candidate)
					return candidate
				}
			}
		}
	}

	// Fallback versioning decision
	version := defaultVersion
	for _, rule := range rules {
		if !rule.Versioned {
			version = ""
			break
		}
	}

	defaultFile := utils.BuildMockPath(service, resource, method, version)
	if fileExists(defaultFile) {
		utils.Logger.Infof("📄 Fallback to default mock file: %s", defaultFile)
		return defaultFile
	}

	utils.Logger.Warnf("❌ No mock file matched for %s/%s [%s]", service, resource, method)
	return ""
}
