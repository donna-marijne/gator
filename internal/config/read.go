package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error) {
	var c Config

	filePath, err := getConfigFilePath()
	if err != nil {
		return c, err
	}

	buffer, err := os.ReadFile(filePath)
	if err != nil {
		return c, fmt.Errorf("failed to read the file at %s: %w", filePath, err)
	}

	err = json.Unmarshal(buffer, &c)
	if err != nil {
		return c, fmt.Errorf("failed to read json from the file at %s: %w", filePath, err)
	}

	return c, nil
}
