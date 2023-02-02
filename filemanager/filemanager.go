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

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func copy(srcFile, dstFile string) error {
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
			if err := copy(sourcePath, destPath); err != nil {
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
