package edge

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
	path := filepath.Join(genGuide.RootPath, "asynq.go")
	return util.ParseTemplateFile("asynq/edge/asynq.go.tpl", ae.pjInfo, path)
}
