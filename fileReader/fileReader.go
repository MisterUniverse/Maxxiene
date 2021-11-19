package fileReader

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	path string
}

func (c *Config) Path() string {
	if c == nil {
		return "usr/home"
	}
	return c.path
}

type errWriter struct {
	io.Writer
	err error
}

func ReadConfig() (buf []byte, err error) {
	home := os.Getenv("HOME")
	path := filepath.Join(home, ".settings.xml")
	if path != "" {
		config, err := ReadFile(path)
		return config, err
	}
	return nil, nil
}

func ReadFile(path string) (buf []byte, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err = ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func ConvertFileBufferToStr(buf []byte) s string {
	str := string(buf[:len(buf)])
	return str
}

func (e *errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	var n int
	n, e.err = e.Writer.Write(buf)
	return n, nil
}

func WriteAll(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	if err != nil {
		return fmt.Errorf("write failed: %v", err)
	}
	return nil
}

func WriteConfig(w io.Writer, conf *Config) error {
	buf, err := json.Marshal(conf)
	if err != nil {
		return fmt.Errorf("could not marshal config: %v", err)
	}
	if err := WriteAll(w, buf); err != nil {
		return fmt.Errorf("could not write config: %v", err)
	}
	return nil
}
