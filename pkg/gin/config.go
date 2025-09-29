package gin

import (
	"github.com/Ahu-Tools/ahum/pkg/config"
	"github.com/Ahu-Tools/ahum/pkg/project"
)

// We want to implement project.Edge for Gin
func (g *Gin) Name() string {
	return Name
}

func (g *Gin) Pkgs() ([]string, error) {
	return []string{}, nil
}

func (g *Gin) JsonConfig() any {
	return g.ginConfig
}

func (g *Gin) Load() (string, error) {
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

	ginConfig, err := config.LoadConfigByGroup[GinConfig](cfgGroup, &Gin{}, *genGuide)
	if err != nil {
		return nil, err
	}

	return NewGin(&pj.Info, *ginConfig), nil
}
