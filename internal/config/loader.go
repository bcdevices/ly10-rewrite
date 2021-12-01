package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// Config file structure with a map of urls.
type Config struct {
	URLs map[string]string `json:"urls"`
}

// Load loads a config file.
func Load(fileName string) (*Config, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to open %s", fileName))
	}

	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to parse $s", fileName))
	}
	return &config, nil
}

// AlternativeURL returns the alternative url for a location.
func (c *Config) AlternativeURL(originalURL string) string {
	for k, v := range c.URLs {
		if originalURL == k {
			return v
		}
	}

	return originalURL
}
