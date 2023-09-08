package filemgr

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type DirBackup struct {
	Source      string
	Destination string
}

func NewDirBackup(source, destination string) *DirBackup {
	return &DirBackup{
		Source:      source,
		Destination: destination,
	}
}

func (d *DirBackup) Backup() error {
	return ZipDir(d.Source, d.Destination)
}

// ZipDir
func ZipDir(src string, dest string) error {
	zipFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	writer := zip.NewWriter(zipFile)
	defer writer.Close()

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = filepath.Join(filepath.Base(src), path[len(src):])

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		w, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(w, file)
		return err
	})
}
