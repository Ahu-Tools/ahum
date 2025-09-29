package project

import (
	"os"
	"path/filepath"

	"github.com/Ahu-Tools/ahum/pkg/config"
	gen "github.com/Ahu-Tools/ahum/pkg/generation"
)

type EdgeLoader func(pj Project, cfgGroup string) (Edge, error)
type InfraLoader func(pj Project, cfgGroup string) (Infra, error)

var edgeLoaders = make(map[string]EdgeLoader)
var infraLoaders = make(map[string]InfraLoader)

func (p *Project) GetConfig() *config.Config {
	edgesCfg := NewEdgeConfig(p.Edges)
	infrasCfg := NewInfraConfig(p.Infras)
	cfgGroups := []config.ConfigurableGroup{
		edgesCfg,
		infrasCfg,
	}
	return config.NewConfig(p.Info.PackageName, cfgGroups)
}

func (p Project) GetConfigGenGuide() (*gen.Guide, error) {
	path := filepath.Join(p.GenGuide.RootPath, "config")
	err := os.Mkdir(path, p.GenGuide.DirPerms)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}
	return gen.NewGuide(path, p.GenGuide.DirPerms, p.GenGuide.FilePerms), nil
}

func RegisterEdgeLoader(name string, el EdgeLoader) {
	edgeLoaders[name] = el
}

func RegisterInfraLoader(name string, il InfraLoader) {
	infraLoaders[name] = il
}
