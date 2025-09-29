package postgres

import (
	"path/filepath"
	"strings"

	"github.com/Ahu-Tools/ahum/pkg/config"
	"github.com/Ahu-Tools/ahum/pkg/project"
	"github.com/Ahu-Tools/ahum/pkg/util"
)

const Name = "postgres"

type Postgres struct {
	projectInfo project.ProjectInfo
	jsonConfig  PostgresConfig
}

func NewPostgres(projectInfo project.ProjectInfo, postgresJSONConfig PostgresConfig) *Postgres {
	return &Postgres{
		projectInfo: projectInfo,
		jsonConfig:  postgresJSONConfig,
	}
}

func (c Postgres) Pkgs() ([]string, error) {
	tmplName := "config_imports.go.tpl"
	tmplPath := "infrastructures/postgres/" + tmplName

	data := map[string]string{
		"PackageName": c.projectInfo.PackageName,
		"Name":        c.Name(),
	}

	importsStr, err := util.ParseTemplateString(tmplPath, data)
	if err != nil {
		return nil, err
	}

	imports := strings.Split(importsStr, "\n")
	for i, l := range imports {
		imports[i] = strings.TrimSpace(l)
	}

	return imports, nil
}

func (c Postgres) Load() (string, error) {
	path := filepath.Join("infrastructures/postgres/loadconfig.go.tpl")

	return util.ParseTemplateString(path, nil)
}

func (pc Postgres) Name() string {
	return Name
}

func (pc Postgres) JsonConfig() any {
	return pc.jsonConfig
}

func init() {
	project.RegisterInfraLoader(Name, Loader)
}

func Loader(pj project.Project, cfgGroup string) (project.Infra, error) {
	genGuide, err := pj.GetConfigGenGuide()
	if err != nil {
		return nil, err
	}

	cfg, err := config.LoadConfigByGroup[PostgresConfig](cfgGroup, &Postgres{}, *genGuide)
	if err != nil {
		return nil, err
	}

	return NewPostgres(pj.Info, *cfg), nil
}
