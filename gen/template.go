package gen

import (
	"bytes"
	"embed"
	"github.com/gari8/worgen/config"
	"io/fs"
	"path"
	"strings"
	"text/template"
)

var (
	//go:embed _template/*
	templateDir embed.FS
)

type Template struct {
	*config.Config
	*embed.FS
}

func NewTemplate(c *config.Config) *Template {
	return &Template{Config: c, FS: &templateDir}
}

func (t Template) ReadTemplates(rootPath string) ([]byte, error) {
	var archive Archive
	tmpl := template.New("").Delims("@@", "@@")
	err := fs.WalkDir(t.FS, ".", func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		content, err := t.FS.ReadFile(filePath)
		if err != nil {
			return err
		}

		var buf, bufFileName bytes.Buffer

		templ := template.Must(tmpl.Clone()).New(filePath)
		if _, err := templ.Parse(string(content)); err != nil {
			return err
		}
		if err := templ.Execute(&buf, t.Config); err != nil {
			return err
		}

		tFileName := template.Must(tmpl.Clone()).New("filename")
		if _, err := tFileName.Parse(filePath); err != nil {
			return err
		}
		if err := tFileName.Execute(&bufFileName, t.Config); err != nil {
			return err
		}

		archive.Files = append(archive.Files, File{
			Name: path.Join(rootPath, strings.Replace(bufFileName.String(), "_template", t.Config.AppName, 1)),
			Data: buf.Bytes(),
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return Format(&archive), nil
}
