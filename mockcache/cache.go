package mockcache

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Cache struct {
	Parsed map[string]map[string]interface{} // for files that directly modifies dynamically
	Raw    map[string][]byte                 // for plain responses (JSON, XML, etc. without changes)
	mu     sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		Parsed: make(map[string]map[string]interface{}),
		Raw:    make(map[string][]byte),
	}
}

func (c *Cache) Load(basePath string) error {
	return filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		relPath, _ := filepath.Rel(basePath, path)
		ext := strings.ToLower(filepath.Ext(relPath))

		if ext != ".json" && ext != ".xml" && ext != ".txt" {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}

		c.mu.Lock()
		defer c.mu.Unlock()

		if ext == ".json" && isDynamicJSON(path) {
			var parsed map[string]interface{}
			if err := json.Unmarshal(content, &parsed); err != nil {
				return fmt.Errorf("failed to parse file %s: %w", path, err)
			}
			c.Parsed[relPath] = parsed
			log.Printf("[cache] Parsed mock loaded: %s", relPath)
		} else {
			c.Raw[relPath] = content
			log.Printf("[cache] Raw mock loaded: %s", relPath)
		}

		return nil
	})
}

func isDynamicJSON(path string) bool {
	lower := strings.ToLower(filepath.Base(path))

	return strings.Contains(lower, "__dynamic__") ||
		strings.Contains(lower, ".dynamic.") ||
		strings.Contains(lower, "__template__")
}

var GlobalCache = NewCache()
