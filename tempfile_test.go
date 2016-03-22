package ioutilx

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestTempFile(t *testing.T) {
	root, err := ioutil.TempDir("", "")
	if err != nil {
		t.Error(err)
	}
	err = os.RemoveAll(root)
	if err != nil {
		t.Error(err)
	}
	_, err = TempFile(root, "")
	if err != nil {
		t.Error(err)
	}
	os.RemoveAll(root)
}

func TestTempDir(t *testing.T) {
	root, err := ioutil.TempDir("", "")
	if err != nil {
		t.Error(err)
	}
	err = os.RemoveAll(root)
	if err != nil {
		t.Error(err)
	}
	_, err = TempDir(root, "")
	if err != nil {
		t.Error(err)
	}
	os.RemoveAll(root)
}
