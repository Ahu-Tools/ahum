package redis

import (
	"path/filepath"

	"github.com/Ahu-Tools/AhuM/pkg/config"
	gen "github.com/Ahu-Tools/AhuM/pkg/generation"
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/util"
)

type Redis struct {
	pj     project.ProjectInfo
	config Config
}

func NewRedis(pj project.ProjectInfo, config Config) *Redis {
	return &Redis{
		pj:     pj,
		config: config,
	}
}

const Name = "redis"

// We want to implement project.Infra for Redis
func (r *Redis) Name() string {
	return Name
}

func (r *Redis) Pkgs() ([]string, error) {
	return []string{}, nil
}

func (r *Redis) JsonConfig() any {
	return r.config
}

func (r *Redis) Load() (string, error) {
	return "", nil
}

func (r *Redis) Generate(status chan string, genGuide gen.Guide) error {
	clientPath := filepath.Join(genGuide.RootPath, "client.go")
	return util.ParseTemplateFile("template/redis/client.go.tpl", r.pj, clientPath)
}

func init() {
	project.RegisterInfraLoader(Name, Loader)
}

func Loader(pj project.Project, cfgGroup string) (project.Infra, error) {
	genGuide, err := pj.GetConfigGenGuide()
	if err != nil {
		return nil, err
	}

	cfg, err := config.LoadConfigByGroup[Config](cfgGroup, &Redis{}, *genGuide)
	if err != nil {
		return nil, err
	}

	return NewRedis(pj.Info, *cfg), nil
}
