package project

import (
	"os"
	"path/filepath"
	"text/template"
)

type Config struct {
	PackageName string
	Pkgs        []string
	Infras      []Infra
}

type Infra interface {
	Generate(chan string, GenerationGuide) error
	Pkgs() ([]string, error)
	Load() (string, error)
	Name() string
	JsonConfig() (any, error)
}

func (p *Project) GenerateConfig() error {
	config := Config{
		PackageName: p.Info.PackageName,
		Infras:      p.Infras,
	}

	tmplName := "config.go.tpl"
	tmplPath := "template/config/" + tmplName
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	filePath := filepath.Join(p.Info.RootPath + "/config/config.go")
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.ExecuteTemplate(f, tmplName, config)
}
