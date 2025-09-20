package project

import (
	"os"
	"path/filepath"

	"github.com/Ahu-Tools/AhuM/pkg/config"
	gen "github.com/Ahu-Tools/AhuM/pkg/generation"
	"github.com/Ahu-Tools/AhuM/pkg/util"
)

type Infra interface {
	config.Configurable
	Generate(status chan string, genGuide gen.Guide) error
}

type InfraConfig struct {
	cfgs []config.Configurable
}

func NewInfraConfig(cfgs []Infra) InfraConfig {
	return InfraConfig{
		cfgs: util.Map(cfgs, func(infra Infra) config.Configurable { return infra }),
	}
}

func (InfraConfig) Name() string {
	return "infras"
}

func (e InfraConfig) GetConfigurables() []config.Configurable {
	return e.cfgs
}

func (p *Project) GenerateInfras(statusChan chan string) error {
	for _, infra := range p.Infras {
		infraGuide, err := p.GetInfraGenGuide(infra)
		if err != nil {
			return err
		}

		err = infra.Generate(statusChan, *infraGuide)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) GetInfraGenGuide(infra Infra) (*gen.Guide, error) {
	infraPath := filepath.Join(p.GenGuide.RootPath, "/infrastructure/", infra.Name())
	err := os.Mkdir(infraPath, p.GenGuide.DirPerms)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}

	return gen.NewGuide(infraPath, p.GenGuide.DirPerms, p.GenGuide.FilePerms), nil
}
