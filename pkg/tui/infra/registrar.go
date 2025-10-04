package infra

import (
	"github.com/Ahu-Tools/ahum/pkg/project"
	"github.com/Ahu-Tools/ahum/pkg/tui/asynq/infra"
	"github.com/Ahu-Tools/ahum/pkg/tui/basic"
	"github.com/Ahu-Tools/ahum/pkg/tui/postgres"
	"github.com/Ahu-Tools/ahum/pkg/tui/redis"
	"github.com/charmbracelet/huh"
)

func GetInfras(p project.ProjectInfo) []huh.Option[basic.RouterModel] {
	return []huh.Option[basic.RouterModel]{
		huh.NewOption[basic.RouterModel]("PostgreSQL", postgres.NewPostgresForm(p)),
		huh.NewOption[basic.RouterModel]("Redis", redis.NewForm(p)),
		huh.NewOption[basic.RouterModel]("Asynq", infra.NewForm(p)),
	}
}
