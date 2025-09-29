package connect

import "github.com/Ahu-Tools/ahum/pkg/project"

func LoadConnectFromProject(p project.Project) *Connect {
	edges := p.GetEdgesByName()
	c := edges[Name].(*Connect)
	return c
}
