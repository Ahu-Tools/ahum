package edge

import (
	"os"
	"path/filepath"
	"strings"

	gen "github.com/Ahu-Tools/ahum/pkg/generation"
	"github.com/Ahu-Tools/ahum/pkg/project"
	"github.com/Ahu-Tools/ahum/pkg/util"
	"github.com/iancoleman/strcase"
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

	err = ae.AddTaskHandler("world", "hello", "v1", genGuide)
	if err != nil {
		return err
	}

	return nil
}

func (ae *Asynq) AddTaskHandler(taskName, modName, verName string, genGuide gen.Guide) error {
	taskName = strcase.ToCamel(taskName)
	verName = strings.ToLower(verName)
	modName = strings.ToLower(modName)

	path := filepath.Join(genGuide.RootPath, modName, verName, modName+".go")
	data := map[string]any{
		"TaskName":      taskName,
		"TaskNameKebab": strcase.ToKebab(taskName),
	}
	payload, _ := util.ParseTemplateString("asynq/edge/payload_struct.go.tpl", data)
	handler, _ := util.ParseTemplateString("asynq/edge/handle_func.go.tpl", data)
	insertions := map[string]string{
		"payloads": payload,
		"handlers": handler,
	}

	err := util.ModifyCodeByMarkersFile(path, insertions, genGuide.FilePerms)
	if err != nil {
		return err
	}

	path = filepath.Join(genGuide.RootPath, modName, verName, "registrar.go")
	taskType, _ := util.ParseTemplateString("asynq/edge/task_type.go.tpl", data)
	taskRegister, _ := util.ParseTemplateString("asynq/edge/task_register.go.tpl", data)
	insertions = map[string]string{
		"types":     taskType,
		"registers": taskRegister,
	}
	return util.ModifyCodeByMarkersFile(path, insertions, genGuide.FilePerms)
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
