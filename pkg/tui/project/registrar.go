package project

import (
	"github.com/Ahu-Tools/AhuM/pkg/postgres"
	"github.com/Ahu-Tools/AhuM/pkg/tui/basic"
)

func GetInfras() map[string]basic.RouterModel {
	return map[string]basic.RouterModel{
		"PostgreSQL": postgres.NewPostgresForm(),
	}
}
