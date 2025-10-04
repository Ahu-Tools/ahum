package infra

import (
	"fmt"
	"path/filepath"

	"github.com/Ahu-Tools/ahum/pkg/util"
)

const Name = "asynq"

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type Config struct {
	Redis   Redis
	Modules map[string]map[string]any
}

func DefaultConfig() Config {
	return Config{
		Redis: Redis{
			Host:     "localhost",
			Port:     6379,
			Username: "admin",
			Password: "1234",
			DB:       0,
		},
	}
}

// We want to implement project.Edge for Connect
func (ae *Asynq) Name() string {
	return Name
}

func (ae *Asynq) Pkgs() ([]string, error) {
	return []string{
		fmt.Sprintf("%s/infrastructure/%s", ae.pjInfo.PackageName, Name),
	}, nil
}

func (ae *Asynq) JsonConfig() any {
	cfg := make(map[string]any)
	cfg["redis"] = ae.Config.Redis
	for modName, modCfg := range ae.Config.Modules {
		cfg[modName] = modCfg
	}
	return cfg
}

func (ae *Asynq) Load() (string, error) {
	path := filepath.Join("asynq/infra/loadconfig.go.tpl")

	return util.ParseTemplateString(path, nil)
}
