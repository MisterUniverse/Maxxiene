package filemgr

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
	"time"

	toml "github.com/pelletier/go-toml"
	"golang.org/x/sys/windows"
)

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func IsFileThere(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func BackupAllFromMap(paths map[string]string, dst string) {
	for _, path := range paths {
		if err := backup(path, dst); err != nil {
			fmt.Printf("Failed to backup %s: %s\n", path, err)
		}
	}
}

func BackupSingle(path, dst string) {
	if err := backup(path, dst); err != nil {
		fmt.Printf("Failed to backup %s: %s\n", path, err)
	}
}

func backup(path, dst string) error {
	timestamp := time.Now().Format("20060102-150405")
	backupFileName := filepath.Join(dst, fmt.Sprintf("%s-%s", filepath.Base(path), timestamp))

	var backupable Backupable

	if IsDir(path) {
		backupable = NewDirBackup(path, fmt.Sprintf("%s.zip", backupFileName))
	} else {
		backupable = NewFileBackup(path, backupFileName+".mxbkup")
	}

	if err := backupable.Backup(); err != nil {
		return err
	}
	fmt.Println("Backup successful:", backupFileName)
	return nil
}

// CreateDirectory creates a directory at the given path.
func CreateDirectory(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}
	return nil
}

// CreateTOMLFile creates a TOML file using a given map and file name
func CreateTOMLFile(settings map[string]map[string]string, filename string) error {
	// Create or open the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)

	// Encode the map into the file
	err = encoder.Encode(settings)
	if err != nil {
		return err
	}

	return nil
}

// WriteEnvFile writes key-value pairs to a .env file.
func CreateEnvFile(path string, values map[string]string) error {
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
	if IsFileThere(path) {
		fmt.Println("File already exist: ", path)
		return nil
	}

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
