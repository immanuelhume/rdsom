package internal

import (
	"embed"
	"go/format"
	"io"
	"path"
	"text/template"
)

//go:embed templates
var templates embed.FS

type Templater struct {
	name  string        // name of the .gotmpl file
	out   io.ReadWriter // where to write the result to
	noFmt bool          // true if the Go formatter should not be run
}

func NewTemplater(wr io.ReadWriter, name string) *Templater {
	return &Templater{
		name: name,
		out:  wr,
	}
}

func (t *Templater) SkipFmt() *Templater {
	t.noFmt = true
	return t
}

func (t *Templater) Do(data interface{}) error {
	tmpl, err := template.New(t.name).Funcs(funcs).ParseFS(templates, path.Join("templates", t.name))
	if err != nil {
		return err
	}
	if err := tmpl.Execute(t.out, data); err != nil {
		return err
	}
	if t.noFmt {
		return nil
	}
	src, err := io.ReadAll(t.out)
	if err != nil {
		return err
	}
	src, err = format.Source(src)
	if err != nil {
		return err
	}
	_, err = t.out.Write(src)
	if err != nil {
		return err
	}
	return nil
}
