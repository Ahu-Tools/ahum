package project

import "path/filepath"

func (p *Project) AddEdges(statusChan chan string) error {
	for _, edge := range p.Edges {
		p.AddEdge(edge, statusChan)
	}
	return nil
}

func (p *Project) AddEdge(edge Edge, statusChan chan string) error {
	edgePath := filepath.Join(p.GenGuide.RootPath, "/edge/", edge.Name())
	edgeGuide := NewGenerationGuide(edgePath, p.GenGuide.DirPerms, p.GenGuide.FilePerms)
	err := edge.Generate(statusChan, edgeGuide)
	if err != nil {
		return err
	}
	return nil
}
