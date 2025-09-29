package connect

import (
	"github.com/Ahu-Tools/ahum/pkg/config"
	"github.com/Ahu-Tools/ahum/pkg/project"
)

const Name = "connect"

// We want to implement project.Edge for Connect
func (g *Connect) Name() string {
	return Name
}

func (g *Connect) Pkgs() ([]string, error) {
	return []string{}, nil
}

func (g *Connect) JsonConfig() any {
	return g.ConnectConfig
}

func (g *Connect) Load() (string, error) {
	return "", nil
}

func init() {
	project.RegisterEdgeLoader(Name, Loader)
}

func Loader(pj project.Project, cfgGroup string) (project.Edge, error) {
	genGuide, err := pj.GetConfigGenGuide()
	if err != nil {
		return nil, err
	}

	cfg, err := config.LoadConfigByGroup[ConnectConfig](cfgGroup, &Connect{}, *genGuide)
	if err != nil {
		return nil, err
	}

	return NewConnect(&pj.Info, *cfg), nil
}
