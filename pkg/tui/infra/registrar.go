package infra

import (
	"github.com/Ahu-Tools/AhuM/pkg/tui/postgres"
	"github.com/charmbracelet/huh"
)

func GetInfras() []huh.Option[Form] {
	return []huh.Option[Form]{
		huh.NewOption[Form]("PostgreSQL", postgres.NewPostgresForm()),
	}
}
