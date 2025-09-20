package project

import (
	"os"
	"path/filepath"

	"github.com/Ahu-Tools/AhuM/pkg/config"
	gen "github.com/Ahu-Tools/AhuM/pkg/generation"
)

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
