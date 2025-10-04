package infra

import (
	"path/filepath"

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
	statusChan <- "Generating asynq..."

	path := filepath.Join(genGuide.RootPath, "config.go")
	err := util.ParseTemplateFile("asynq/infra/config.go.tpl", ae.pjInfo, path)
	if err != nil {
		return err
	}

	path = filepath.Join(genGuide.RootPath, "asynq.go")
	err = util.ParseTemplateFile("asynq/infra/asynq.go.tpl", ae.pjInfo, path)
	if err != nil {
		return err
	}

	return nil
}
