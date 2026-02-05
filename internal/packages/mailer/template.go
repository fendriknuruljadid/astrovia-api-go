package mailer

import (
	"bytes"
	"html/template"
	"path/filepath"
)

func RenderTemplate(name string, data any) (string, error) {
	path := filepath.Join("mailer", "templates", name)

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
