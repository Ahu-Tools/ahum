package config

import (
	"fmt"
	"path/filepath"
	"strings"

	gen "github.com/Ahu-Tools/AhuM/pkg/generation"
	"github.com/Ahu-Tools/AhuM/pkg/util"
)

func LoadConfig[T any](genGuide gen.Guide, name string) (*T, error) {
	configPath := filepath.Join(genGuide.RootPath, "config.json")
	regP, err := util.LoadJSONPathToStruct[RegisterMap](configPath, "registrar")
	if err != nil {
		return nil, err
	}

	reg := *regP

	jsonPath, ok := reg[name]
	if !ok {
		return nil, fmt.Errorf("there's no configuration registered with name: %s", name)
	}

	config, err := util.LoadJSONPathToStruct[T](configPath, jsonPath)
	return config, err
}

func AddConfigByGroup(group string, cfg Configurable, genGuide gen.Guide) error {
	configPath := filepath.Join(genGuide.RootPath, "config.json")
	err := util.AddElementToJSON(configPath, group, cfg.Name(), cfg.JsonConfig())
	if err != nil {
		return err
	}

	goCfgPath := filepath.Join(genGuide.RootPath, "config.go")
	load, err := cfg.Load()
	if err != nil {
		return err
	}
	load = fmt.Sprintf(`//@ahum:%s.load
	%s
	//@ahum:end.%s.load`, cfg.Name(), load, cfg.Name())

	pkgs, err := cfg.Pkgs()
	if err != nil {
		return err
	}
	pkgs = util.Map(pkgs, func(pkg string) string {
		return fmt.Sprintf(`"%s"`, pkg)
	})
	insertions := map[string]string{
		"imports":                          strings.Join(pkgs, "\n"),
		fmt.Sprintf("end.%s.group", group): load,
	}

	return util.ModifyCodeByMarkersFile(goCfgPath, insertions, genGuide.FilePerms)
}
