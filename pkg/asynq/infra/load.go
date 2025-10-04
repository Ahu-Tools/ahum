package infra

import (
	"github.com/Ahu-Tools/ahum/pkg/config"
	"github.com/Ahu-Tools/ahum/pkg/project"
)

func init() {
	project.RegisterInfraLoader(Name, loader)
}

func loader(pj project.Project, cfgGroup string) (project.Infra, error) {
	cfgGuide, err := pj.GetConfigGenGuide()
	if err != nil {
		return nil, err
	}

	cfg, err := config.LoadConfigByGroup[Config](cfgGroup, &Asynq{}, *cfgGuide)
	if err != nil {
		return nil, err
	}

	return NewAsynq(*cfg, pj.Info), nil
}
