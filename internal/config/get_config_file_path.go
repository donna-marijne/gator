package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func getConfigFilePath() (string, error) {
	homeDirPath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home dir: %w", err)
	}

	return filepath.Join(homeDirPath, ".gatorconfig.json"), nil
}
