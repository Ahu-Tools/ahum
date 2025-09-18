package project

import (
	"os"
	"path/filepath"

	"github.com/Ahu-Tools/AhuM/pkg/util"
)

func (p *Project) AddEdges(statusChan chan string) error {
	for _, edge := range p.Edges {
		err := p.AddEdge(edge, statusChan)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) AddEdge(edge Edge, statusChan chan string) error {
	err := p.addEdgeToStart(edge)
	if err != nil {
		return err
	}

	edgePath := filepath.Join(p.GenGuide.RootPath, "/edge/", edge.Name())
	err = os.Mkdir(edgePath, p.GenGuide.DirPerms)
	if err != nil {
		return err
	}

	edgeGuide := NewGenerationGuide(edgePath, p.GenGuide.DirPerms, p.GenGuide.FilePerms)
	return edge.Generate(statusChan, edgeGuide)
}

func (p *Project) addEdgeToStart(edge Edge) error {
	path := filepath.Join(p.GenGuide.RootPath, "/edge/edge.go")
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	insertions := map[string]string{
		"imports": "\"" + p.Info.PackageName + "/edge/" + edge.Name() + "\"\n",
		"edges":   edge.Name() + ".New(),\n",
	}

	newFile, err := util.ModifyCodeByMarkers(file, insertions)
	if err != nil {
		return err
	}

	return os.WriteFile(path, newFile, p.GenGuide.FilePerms)
}
