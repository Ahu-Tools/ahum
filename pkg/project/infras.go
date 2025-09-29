package project

import (
	"os"
	"path/filepath"

	"github.com/Ahu-Tools/ahum/pkg/config"
	gen "github.com/Ahu-Tools/ahum/pkg/generation"
	"github.com/Ahu-Tools/ahum/pkg/util"
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

func (p *Project) GenInfras(statusChan chan string) error {
	for _, infra := range p.Infras {
		err := p.GenInfra(infra, statusChan)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) GenInfra(infra Infra, statusChan chan string) error {
	infraGuide, err := p.GetInfraGenGuide(infra)
	if err != nil {
		return err
	}

	return infra.Generate(statusChan, *infraGuide)
}

func (p *Project) AddInfra(infra Infra, statusChan chan string) error {
	cfgGen, err := p.GetConfigGenGuide()
	if err != nil {
		return nil
	}

	err = config.AddConfigByGroup(InfrasGroup, infra, *cfgGen)
	if err != nil {
		return err
	}

	err = p.GenInfra(infra, statusChan)
	if err != nil {
		return err
	}

	err = p.GoSweep(statusChan)
	if err != nil {
		return err
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
