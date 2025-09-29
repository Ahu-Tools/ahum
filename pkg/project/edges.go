package project

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Ahu-Tools/ahum/pkg/config"
	gen "github.com/Ahu-Tools/ahum/pkg/generation"
	"github.com/Ahu-Tools/ahum/pkg/util"
)

const EdgesGroup = "edges"
const InfrasGroup = "infras"

type Edge interface {
	config.Configurable
	Generate(status chan string, genGuide gen.Guide) error
}

type EdgeConfig struct {
	cfgs []config.Configurable
}

func NewEdgeConfig(cfgs []Edge) EdgeConfig {
	return EdgeConfig{
		cfgs: util.Map(cfgs, func(edge Edge) config.Configurable { return edge }),
	}
}

func (EdgeConfig) Name() string {
	return EdgesGroup
}

func (e EdgeConfig) GetConfigurables() []config.Configurable {
	return e.cfgs
}

func (p *Project) GenEdges(statusChan chan string) error {
	for _, edge := range p.Edges {
		err := p.GenEdge(edge, statusChan)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) GenEdge(edge Edge, statusChan chan string) error {
	err := p.addEdgeToStart(edge)
	if err != nil {
		return err
	}

	edgeGuide, err := p.GetEdgeGenGuide(edge)
	if err != nil {
		return err
	}

	return edge.Generate(statusChan, *edgeGuide)
}

func (p *Project) GetEdgeGenGuide(edge Edge) (*gen.Guide, error) {
	edgePath := filepath.Join(p.GenGuide.RootPath, "/edge/", edge.Name())
	err := os.Mkdir(edgePath, p.GenGuide.DirPerms)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}
	return gen.NewGuide(edgePath, p.GenGuide.DirPerms, p.GenGuide.FilePerms), nil
}

func (p *Project) AddEdge(edge Edge, statusChan chan string) error {
	cfgGen, err := p.GetConfigGenGuide()
	if err != nil {
		return nil
	}

	err = config.AddConfigByGroup(EdgesGroup, edge, *cfgGen)
	if err != nil {
		return err
	}

	err = p.GenEdge(edge, statusChan)
	if err != nil {
		return err
	}

	err = p.GoSweep(statusChan)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) addEdgeToStart(edge Edge) error {
	path := filepath.Join(p.GenGuide.RootPath, "/edge/edge.go")

	insertions := map[string]string{
		"imports":  fmt.Sprintf(`"%s/edge/%s"`, p.Info.PackageName, edge.Name()),
		EdgesGroup: edge.Name() + ".New(),",
	}

	return util.ModifyCodeByMarkersFile(path, insertions, p.GenGuide.FilePerms)
}
