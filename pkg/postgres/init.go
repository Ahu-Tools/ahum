package postgres

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/Ahu-Tools/AhuM/pkg/project"
)

func (p Postgres) Generate(statusChan chan string, genGuide project.GenerationGuide) error {
	statusChan <- "Generating postgresql directories structure..."
	err := p.generateBasicDirs(genGuide)
	if err != nil {
		return err
	}

	statusChan <- "Generate postgresql basic files..."
	err = p.generateBasicFiles(genGuide)
	if err != nil {
		return err
	}

	return nil
}

func (p Postgres) generateBasicDirs(genGuide project.GenerationGuide) error {
	return os.MkdirAll(genGuide.RootPath+"/migrations", genGuide.DirPerms)
}

func (p Postgres) generateBasicFiles(genGuide project.GenerationGuide) error {
	err := generateConfig(genGuide)
	if err != nil {
		return err
	}

	err = generateConnection(genGuide)
	if err != nil {
		return err
	}

	return nil
}

func generateConfig(genGuide project.GenerationGuide) error {
	//Generate config.go using config.go.tpl template
	tmplName := "config.go.tpl"
	tmplNamePath := "template/infrastructures/postgres/" + tmplName

	tmpl, err := template.ParseFiles(tmplNamePath)
	if err != nil {
		return err
	}

	path := filepath.Join(genGuide.RootPath + "/config.go")

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(f, tmplName, nil)
}

func generateConnection(genGuide project.GenerationGuide) error {
	//Generate connection.go using connection.go.tpl template
	tmplName := "connection.go.tpl"
	tmplNamePath := "template/infrastructures/postgres/" + tmplName

	tmpl, err := template.ParseFiles(tmplNamePath)
	if err != nil {
		return err
	}

	path := filepath.Join(genGuide.RootPath + "/connection.go")

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(f, tmplName, nil)
}
