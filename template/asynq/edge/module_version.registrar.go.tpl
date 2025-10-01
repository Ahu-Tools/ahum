package {{.VersionName}}

import (
	"{{.PackageName}}/edge/asynq"
	//@ahum: imports
)

const version = "{{.VersionName}}"

// A list of task types.
const (
	ModuleName        = "{{.ModuleName}}"
	//@ahum: types
)

func GetPattern(handler string) string {
	return ModuleName + ":" + handler
}

func init() {
	// @ahum: registers
}
