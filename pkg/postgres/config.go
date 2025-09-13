package postgres

import (
	"bytes"
	"os"
	"strings"
	"text/template"
)

type PostgresConfig struct {
	PackageName string
}

func NewPostgresConfig(packageName string) *PostgresConfig {
	return &PostgresConfig{
		PackageName: packageName,
	}
}

func (c PostgresConfig) Pkgs() ([]string, error) {
	tmplName := "config_imports.go.tpl"
	tmplPath := "template/infrastructures/postgres/" + tmplName
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return nil, err
	}

	var importsBytes bytes.Buffer
	err = tmpl.ExecuteTemplate(&importsBytes, tmplName, c)
	if err != nil {
		return nil, err
	}

	imports := strings.Split(importsBytes.String(), "\n")
	for i, l := range imports {
		imports[i] = strings.TrimSpace(l)
	}

	return imports, nil
}

func (c PostgresConfig) Load() (string, error) {
	resultBytes, err := os.ReadFile("template/infrastructures/postgres/loadconfig.go.tpl")
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
