package filemgr

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"
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
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}
	return nil
}

// WriteEnvFile writes key-value pairs to a .env file.
func WriteEnvFile(path string, values map[string]string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create .env file: %w", err)
	}
	defer file.Close()

	for key, value := range values {
		if _, err := fmt.Fprintf(file, "%s=%s\n", key, value); err != nil {
			return fmt.Errorf("failed to write to .env file: %w", err)
		}
	}

	return nil
}

// CreateFile creates an empty file at the given path.
func CreateFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", path, err)
	}
	file.Close()
	return nil
}

func ShredDirectory(directoryPath string, passes int) error {
	// List all files in the directory
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return err
	}

	// Shred each file in the directory
	for _, file := range files {
		filePath := filepath.Join(directoryPath, file.Name())
		if file.IsDir() {
			// Recursively shred sub-directories
			if err := ShredDirectory(filePath, passes); err != nil {
				return err
			}
		} else {
			// Shred the file
			if err := ShredFile(filePath, passes); err != nil {
				return err
			}
		}
	}

	// Remove the directory itself
	return os.Remove(directoryPath)
}

func ShredFile(filename string, passes int) error {
	// Convert filename to UTF-16 for CreateFile
	p, err := windows.UTF16PtrFromString(filename)
	if err != nil {
		return err
	}

	// Open the file with write permission
	handle, err := windows.CreateFile(
		p,
		windows.GENERIC_WRITE,
		windows.FILE_SHARE_WRITE,
		nil,
		windows.OPEN_EXISTING,
		windows.FILE_ATTRIBUTE_NORMAL,
		0,
	)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(handle)

	// Get file size
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	// Initialize a large data slice filled with random bytes
	data := make([]byte, 65536) // 64 KB
	if _, err := rand.Read(data); err != nil {
		return err
	}

	// Perform multiple overwrite passes
	for pass := 0; pass < passes; pass++ {
		_, err = windows.SetFilePointer(handle, 0, nil, windows.FILE_BEGIN)
		if err != nil {
			return err
		}

		for offset := int64(0); offset < fileSize; offset += int64(len(data)) {
			var bytesWritten uint32
			if err = windows.WriteFile(handle, data, &bytesWritten, nil); err != nil {
				return err
			}
		}
	}
	windows.CloseHandle(handle)
	// Close the file handle and remove the file
	if err = os.Remove(filename); err != nil {
		return err
	}

	return nil
}
