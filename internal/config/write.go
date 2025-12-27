package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func (c *Config) write() error {
	buffer, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("failed to serialize config to json: %w", err)
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, buffer, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config to file %s: %w", filePath, err)
	}

	return nil
}
