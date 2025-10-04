package edge

import (
	"github.com/Ahu-Tools/ahum/pkg/project"
	asynq "github.com/Ahu-Tools/ahum/pkg/tui/asynq/edge"
	"github.com/Ahu-Tools/ahum/pkg/tui/basic"
	"github.com/Ahu-Tools/ahum/pkg/tui/connect"
	"github.com/Ahu-Tools/ahum/pkg/tui/gin"
	"github.com/charmbracelet/huh"
)

func GetEdges(pj project.ProjectInfo) []huh.Option[basic.RouterModel] {
	return []huh.Option[basic.RouterModel]{
		huh.NewOption[basic.RouterModel]("Gin", gin.NewForm(pj)),
		huh.NewOption[basic.RouterModel]("Connect", connect.NewForm(pj)),
		huh.NewOption[basic.RouterModel]("Asynq", asynq.NewForm(pj)),
	}
}
