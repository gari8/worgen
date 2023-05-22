package gen

import (
	"bytes"
	"fmt"
	"golang.org/x/tools/txtar"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type File txtar.File

type Archive struct {
	Files   []File
	Comment []byte
}

func NewArchive(buf []byte) *Archive {
	ar := txtar.Parse(buf)
	var files []File
	for _, f := range ar.Files {
		files = append(files, File(f))
	}
	return &Archive{Files: files}
}

func Format(a *Archive) []byte {
	var buf bytes.Buffer
	buf.Write(fixNL(a.Comment))
	for _, f := range a.Files {
		fmt.Fprintf(&buf, "-- %s --\n", f.Name)
		buf.Write(fixNL(f.Data))
	}
	return buf.Bytes()
}

func fixNL(data []byte) []byte {
	if len(data) == 0 || data[len(data)-1] == '\n' {
		return data
	}
	d := make([]byte, len(data)+1)
	copy(d, data)
	d[len(data)] = '\n'
	return d
}

func isExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func (f File) CreateFile() error {
	var err error
	if len(bytes.TrimSpace(f.Data)) == 0 {
		return nil
	}

	path := filepath.Join(".", filepath.FromSlash(f.Name))

	if _, err := isExist(path); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return err
	}

	if filepath.Ext(path) == ".tmpl" {
		path = strings.Replace(path, ".tmpl", "", 1)
	}
	w, err := os.Create(path)
	if err != nil {
		return err
	}
	defer w.Close()

	r := bytes.NewReader(f.Data)
	if _, err := io.Copy(w, r); err != nil {
		return err
	}

	fmt.Println("create", path)

	return nil
}
