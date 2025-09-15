package util

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func ParseTemplateFile(tmplPath string, data any, saveTo string) error {
	rubbish := strings.Split(tmplPath, "/")
	tmplName := rubbish[len(rubbish)-1]

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	path := filepath.Clean(saveTo)
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(f, tmplName, data)
}
