package cache

import (
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// ExcludeList holds URL prefixes to skip during checking.
type ExcludeList struct {
	Domains []string `yaml:"domains"`
}

// IsExcluded returns true if the URL starts with any excluded prefix.
func (e *ExcludeList) IsExcluded(url string) bool {
	for _, d := range e.Domains {
		if strings.HasPrefix(url, d) {
			return true
		}
	}
	return false
}

// LoadExcludeList reads an exclude.yaml file.
func LoadExcludeList(path string) (*ExcludeList, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var excl ExcludeList
	if err := yaml.Unmarshal(data, &excl); err != nil {
		return nil, err
	}
	return &excl, nil
}

// HealthEntry stores metadata about a single entry.
type HealthEntry struct {
	URL        string    `yaml:"url"`
	Name       string    `yaml:"name"`
	Status     string    `yaml:"status"` // healthy, inactive, stale, archived, dead
	Stars      int       `yaml:"stars,omitempty"`
	Forks      int       `yaml:"forks,omitempty"`
	LastPush   time.Time `yaml:"last_push,omitempty"`
	HasLicense  bool      `yaml:"has_license,omitempty"`
	HasReadme   bool      `yaml:"has_readme,omitempty"`
	CheckedAt   time.Time `yaml:"checked_at"`
	Category    string    `yaml:"category,omitempty"`
	Description string    `yaml:"description,omitempty"`
}

// HealthCache is the full YAML cache file.
type HealthCache struct {
	Entries []HealthEntry `yaml:"entries"`
}

// LoadHealthCache reads a health_cache.yaml file. Returns empty cache if file doesn't exist.
func LoadHealthCache(path string) (*HealthCache, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &HealthCache{}, nil
		}
		return nil, err
	}
	var hc HealthCache
	if err := yaml.Unmarshal(data, &hc); err != nil {
		return nil, err
	}
	return &hc, nil
}

// SaveHealthCache writes the cache to a YAML file.
func SaveHealthCache(path string, hc *HealthCache) error {
	data, err := yaml.Marshal(hc)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// Merge updates the cache with new entries, replacing existing ones by URL.
func (hc *HealthCache) Merge(entries []HealthEntry) {
	index := make(map[string]int)
	for i, e := range hc.Entries {
		index[e.URL] = i
	}
	for _, e := range entries {
		if i, exists := index[e.URL]; exists {
			hc.Entries[i] = e
		} else {
			index[e.URL] = len(hc.Entries)
			hc.Entries = append(hc.Entries, e)
		}
	}
}
