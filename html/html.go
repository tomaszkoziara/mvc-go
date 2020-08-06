package html

import (
	"fmt"
	"io"
	"text/template"
)

// ITemplate is template interface.
type ITemplate interface {
	Load(path string) error
	Apply(w io.Writer, name string, data interface{}) error
}

// Template is the wrapper around templates.
type Template struct {
	template *template.Template
}

// Load loads all templates in a path.
func (t *Template) Load(path string) error {
	tmpl, err := template.ParseGlob(path)
	if err != nil {
		return fmt.Errorf("parsing view files: %w", err)
	}

	t.template = tmpl
	return nil
}

// Apply executes a template and writes it in the output.
func (t *Template) Apply(w io.Writer, name string, data interface{}) error {
	return t.template.ExecuteTemplate(w, name, data)
}

// NewTemplate returns a new template.
func NewTemplate() *Template {
	return &Template{}
}
