package infra

import (
	"github.com/Ahu-Tools/AhuM/pkg/tui/postgres"
)

func GetInfras() map[string]Form {
	return map[string]Form{
		"PostgreSQL": postgres.NewPostgresForm(),
	}
}
