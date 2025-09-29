package template

import "embed"

//go:embed chain config connect data edge gin infrastructures main redis service
var TemplateFS embed.FS
