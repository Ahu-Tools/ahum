package postgres

import (
	"os"

	gen "github.com/Ahu-Tools/ahum/pkg/generation"
	"github.com/Ahu-Tools/ahum/pkg/util"
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
	tmplNamePath := "infrastructures/postgres/" + tmplName

	return util.ParseTemplateFile(tmplNamePath, nil, genGuide.RootPath+"/config.go")
}

func generateConnection(genGuide gen.Guide) error {
	//Generate connection.go using connection.go.tpl template
	tmplName := "connection.go.tpl"
	tmplNamePath := "infrastructures/postgres/" + tmplName

	return util.ParseTemplateFile(tmplNamePath, nil, genGuide.RootPath+"/connection.go")
}
