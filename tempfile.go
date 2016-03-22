package ioutilx

import (
	"io/ioutil"
	"os"
)

// TempFile is a wrapper of ioutil.TempFile with 2 differences:
// 1. It returns a string instead of *os.file
// 2. It will create the directory if dir doesn't exist
func TempFile(dir, prefix string) (string, error) {
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return "", err
	}
	file, err := ioutil.TempFile(dir, prefix)
	if err != nil {
		return "", err
	}
	return file.Name(), file.Close()
}

// TempDir is a wrapper of ioutil.TempDir with 1 difference:
// 1. It will create the directory if dir doesn't exist
func TempDir(dir, prefix string) (string, error) {
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return "", err
	}
	return ioutil.TempDir(dir, prefix)
}
