package config

import (
	"fmt"
	"path/filepath"

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
