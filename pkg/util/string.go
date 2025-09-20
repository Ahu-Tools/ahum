package util

import (
	"strings"

	"github.com/iancoleman/strcase"
)

func ToPkgName(s string) string {
	return strings.ReplaceAll(strcase.ToSnake(s), "_", "")
}
