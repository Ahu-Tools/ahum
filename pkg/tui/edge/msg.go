package edge

import "github.com/Ahu-Tools/AhuM/pkg/project"

type EdgesMsg struct {
	Ok    bool
	Edges []project.Edge
}

type EdgeMsg struct {
	Edge project.Edge
}
