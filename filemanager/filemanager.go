package filemanager

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type WorkingDir struct {
	Path string
}

type WriteTypeOptions interface{}

func AppendFile(file *os.File, w WriteTypeOptions) {
	switch v := w.(type) {
	case string:
		len, err := file.WriteString(v + "\n")
		if err != nil {
			fmt.Println("failed writing file: ", err)
		}
		fmt.Printf("Length: %d bytes\n", len)

	}

	fmt.Printf("File Name: %s\n", file.Name())
}

func Copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer out.Close()

	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}

	defer in.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

func CopyDirectory(scrDir, dest string) error {
	entries, err := ioutil.ReadDir(scrDir)
	if err != nil {
		fmt.Println(err)
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CreateIfNotExists(destPath, 0755); err != nil {
				return err
			}
			if err := CopyDirectory(sourcePath, destPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := CopySymLink(sourcePath, destPath); err != nil {
				return err
			}
		default:
			if err := Copy(sourcePath, destPath); err != nil {
				return err
			}
		}

		isSymlink := entry.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(destPath, entry.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

func CopySymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(link, dest)
}

func CreateIfNotExists(dir string, perm os.FileMode) error {
	if FileExists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

func CreateWorkingDirectory(fileName string) (string, error) {
	userCache, err := os.UserCacheDir()
	if err != nil {
		return "Error: Can't find user dir", err
	}

	workDir := filepath.Join(userCache, fileName)

	fail := CreateIfNotExists(workDir, 755)
	if fail != nil {
		return "ERROR: ", fail
	}

	return workDir, nil
}

func Delete(path string) (string, error) {
	err := os.RemoveAll(path)
	if err != nil {
		return "Can't remove dir/file\n", err
	}
	return "Item deleted.\n", nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// TODO: List dir
func List(path string) ([]string, error) {
	return nil, nil
}

// Reads the whole file into memory and returns byte array
func ReadFile(path string) []byte {

	if FileExists(path) {
		f, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		}
		return f
	}

	return nil
}
