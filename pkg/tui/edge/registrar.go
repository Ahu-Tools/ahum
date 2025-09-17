package edge

import (
	"github.com/Ahu-Tools/AhuM/pkg/tui/gin"
	"github.com/charmbracelet/huh"
)

func GetEdges() []huh.Option[Form] {
	return []huh.Option[Form]{
		huh.NewOption[Form]("gin", gin.NewForm()),
	}
}
