package connect

import "github.com/Ahu-Tools/AhuM/pkg/project"

func LoadConnectFromProject(p project.Project) *Connect {
	edges := p.GetEdgesByName()
	c := edges[Name].(*Connect)
	return c
}
