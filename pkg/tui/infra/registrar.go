package infra

import (
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
	"github.com/Ahu-Tools/AhuM/pkg/tui/postgres"
	"github.com/charmbracelet/huh"
)

func GetInfras(p project.ProjectInfo) []huh.Option[basic.RouterModel] {
	return []huh.Option[basic.RouterModel]{
		huh.NewOption[basic.RouterModel]("PostgreSQL", postgres.NewPostgresForm(p)),
	}
}
