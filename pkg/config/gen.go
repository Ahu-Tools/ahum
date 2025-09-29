package config

import (
	"path/filepath"

	gen "github.com/Ahu-Tools/ahum/pkg/generation"
	"github.com/Ahu-Tools/ahum/pkg/util"
)

func (c Config) Generate(statusChan chan string, genGuide gen.Guide) error {
	tmplPath := "config/config.go.tpl"
	filePath := filepath.Join(genGuide.RootPath, "config.go")

	return util.ParseTemplateFile(tmplPath, c, filePath)
}
