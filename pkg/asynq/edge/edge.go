package edge

import (
	"os"
	"path/filepath"
	"strings"

	gen "github.com/Ahu-Tools/ahum/pkg/generation"
	"github.com/Ahu-Tools/ahum/pkg/project"
	"github.com/Ahu-Tools/ahum/pkg/util"
)

type Asynq struct {
	Config Config
	pjInfo project.ProjectInfo
}

func NewAsynq(cfg Config, pjInfo project.ProjectInfo) *Asynq {
	return &Asynq{
		Config: cfg,
		pjInfo: pjInfo,
	}
}

func (ae *Asynq) Generate(statusChan chan string, genGuide gen.Guide) error {
	path := filepath.Join(genGuide.RootPath, "asynq.go")
	err := util.ParseTemplateFile("asynq/edge/asynq.go.tpl", ae.pjInfo, path)
	if err != nil {
		return err
	}

	err = ae.AddModule("v1", "hello", genGuide)
	if err != nil {
		return err
	}

	return nil
}

func (ae *Asynq) AddModule(verName, modName string, genGuide gen.Guide) error {
	verName = strings.ToLower(verName)
	modName = strings.ToLower(modName)

	dirPath := filepath.Join(genGuide.RootPath, modName, verName)
	err := os.MkdirAll(dirPath, genGuide.DirPerms)
	if err != nil {
		return err
	}

	path := filepath.Join(dirPath, "registrar.go")
	data := map[string]any{
		"PackageName": ae.pjInfo.PackageName,
		"VersionName": verName,
		"ModuleName":  modName,
	}

	err = util.ParseTemplateFile("asynq/edge/module_version.registrar.go.tpl", data, path)
	if err != nil {
		return err
	}

	path = filepath.Join(dirPath, modName+".go")
	return util.ParseTemplateFile("asynq/edge/module.go.tpl", data, path)
}
