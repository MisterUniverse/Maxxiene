package filemgr

import "io/ioutil"

type FileBackup struct {
	Source      string
	Destination string
}

func NewFileBackup(source, destination string) *FileBackup {
	return &FileBackup{
		Source:      source,
		Destination: destination,
	}
}

func (f *FileBackup) Backup() error {
	return CopyFile(f.Source, f.Destination)
}

// CopyFile function
func CopyFile(src, dest string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dest, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
