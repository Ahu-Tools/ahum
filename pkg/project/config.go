package project

import (
	"path/filepath"

	"github.com/Ahu-Tools/AhuM/pkg/util"
)

type Config struct {
	PackageName string
	Pkgs        []string
	Infras      []Infra
}

type Edge interface {
	Generate(chan string, GenerationGuide) error
	Pkgs() ([]string, error)
	Load() (string, error)
	Name() string
	JsonConfig() (any, error)
	StartCode() string
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

	tmplPath := "template/config/config.go.tpl"
	filePath := filepath.Join(p.Info.RootPath + "/config/config.go")

	return util.ParseTemplateFile(tmplPath, config, filePath)
}
