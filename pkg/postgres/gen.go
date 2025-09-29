package postgres

import (
	"os"
	"path/filepath"
	"text/template"

	gen "github.com/Ahu-Tools/ahum/pkg/generation"
)

func (p Postgres) Generate(statusChan chan string, genGuide gen.Guide) error {
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

func (p Postgres) generateBasicDirs(genGuide gen.Guide) error {
	return os.Mkdir(genGuide.RootPath+"/migrations", genGuide.DirPerms)
}

func (p Postgres) generateBasicFiles(genGuide gen.Guide) error {
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

func generateConfig(genGuide gen.Guide) error {
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

func generateConnection(genGuide gen.Guide) error {
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
