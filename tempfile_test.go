package ioutilx

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestTempFile(t *testing.T) {
	root, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	err = os.RemoveAll(root)
	if err != nil {
		t.Fatal(err)
	}
	_, err = TempFile(root, "")
	if err != nil {
		t.Fatal(err)
	}
	os.RemoveAll(root)
}

func TestTempFileWithContent(t *testing.T) {
	content := "hello world\n你好世界"
	file, err := TempFileWithContent("", "", content)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(file)
	// defer os.Remove(file)
	nbcontent, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	ncontent := string(nbcontent)
	if content != ncontent {
		t.Errorf("Expecting content:\n%s\nGet instead:\n%s", content, ncontent)
	}
}

func TestTempDir(t *testing.T) {
	root, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	err = os.RemoveAll(root)
	if err != nil {
		t.Fatal(err)
	}
	_, err = TempDir(root, "")
	if err != nil {
		t.Fatal(err)
	}
	os.RemoveAll(root)
}
