package postgres

import (
	"bytes"
	"os"
	"strings"
	"text/template"

	"github.com/Ahu-Tools/AhuM/pkg/project"
)

type PostgresConfig struct {
	projectInfo project.ProjectInfo
}

func NewPostgresConfig(projectInfo project.ProjectInfo) *PostgresConfig {
	return &PostgresConfig{
		projectInfo: projectInfo,
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

	data := map[string]string{
		"PackageName": c.projectInfo.PackageName,
		"Name":        c.Name(),
	}
	err = tmpl.ExecuteTemplate(&importsBytes, tmplName, data)
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

func (pc PostgresConfig) Name() string {
	return "postgres"
}
