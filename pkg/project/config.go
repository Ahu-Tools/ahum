package project

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/Ahu-Tools/AhuM/pkg/postgres"
)

type Config struct {
	PackageName string
	Pkgs        []string
	Infras      []InfraConfig
}

type InfraConfig interface {
	Pkgs() ([]string, error)
	Load() (string, error)
}

func (p *Project) GenerateConfig() error {
	infras, err := p.getInfraConfigs()
	if err != nil {
		return err
	}

	config := Config{
		PackageName: p.PackageName,
		Infras:      infras,
	}

	tmplName := "config.go.tpl"
	tmplPath := "template/config/" + tmplName
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	filePath := filepath.Join(p.RootPath + "/config/config.go")
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.ExecuteTemplate(f, tmplName, config)
}

func (p *Project) getInfraConfigs() ([]InfraConfig, error) {
	infras := make([]InfraConfig, 0)

	for _, db := range p.Dbs {
		var infraConfig InfraConfig

		switch db {
		case POSTGRES:
			infraConfig = postgres.NewPostgresConfig(p.PackageName)
		}

		infras = append(infras, infraConfig)
	}

	return infras, nil
}
