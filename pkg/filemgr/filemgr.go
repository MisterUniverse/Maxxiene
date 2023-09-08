package filemgr

import (
	"fmt"
	"os"
)

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// CreateDirectory creates a directory at the given path.
func CreateDirectory(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", path, err)
	}
	return nil
}

// WriteEnvFile writes key-value pairs to a .env file.
func WriteEnvFile(path string, values map[string]string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Failed to create .env file: %w", err)
	}
	defer file.Close()

	for key, value := range values {
		if _, err := fmt.Fprintf(file, "%s=%s\n", key, value); err != nil {
			return fmt.Errorf("Failed to write to .env file: %w", err)
		}
	}

	return nil
}

// CreateFile creates an empty file at the given path.
func CreateFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Failed to create file %s: %w", path, err)
	}
	file.Close()
	return nil
}
