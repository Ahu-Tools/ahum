package postgres

import (
	"bytes"
	"os"
	"strings"
	"text/template"

	"github.com/Ahu-Tools/AhuM/pkg/project"
)

const Name = "postgres"

type Postgres struct {
	projectInfo project.ProjectInfo
	jsonConfig  PostgresConfig
}

func NewPostgres(projectInfo project.ProjectInfo, postgresJSONConfig PostgresConfig) *Postgres {
	return &Postgres{
		projectInfo: projectInfo,
		jsonConfig:  postgresJSONConfig,
	}
}

func (c Postgres) Pkgs() ([]string, error) {
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

func (c Postgres) Load() (string, error) {
	resultBytes, err := os.ReadFile("template/infrastructures/postgres/loadconfig.go.tpl")
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}

func (pc Postgres) Name() string {
	return Name
}

func (pc Postgres) JsonConfig() any {
	return pc.jsonConfig
}
