package edge

import (
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
)

type Form interface {
	basic.RouterModel
	InitProjectInfo(project.ProjectInfo)
}
