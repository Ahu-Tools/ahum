package config

import (
	"path/filepath"

	gen "github.com/Ahu-Tools/AhuM/pkg/generation"
	"github.com/Ahu-Tools/AhuM/pkg/util"
)

func (c Config) Generate(statusChan chan string, genGuide gen.Guide) error {
	tmplPath := "template/config/config.go.tpl"
	c.RootPath = filepath.Clean(genGuide.RootPath)
	filePath := filepath.Join(c.RootPath, "/config.go")

	return util.ParseTemplateFile(tmplPath, c, filePath)
}
