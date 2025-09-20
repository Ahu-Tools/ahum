package config

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"

	gen "github.com/Ahu-Tools/AhuM/pkg/generation"
)

type ConfigMap map[string]interface{}

type App struct {
	SecretKey string `json:"secret_key"`
}

func (c Config) GenerateJSON(statusChan chan string, genGuide gen.Guide) error {
	configMap := make(ConfigMap)

	secretKey := genSecretKey()
	app := App{
		SecretKey: secretKey,
	}
	configMap["app"] = app

	for _, cfGroup := range c.ConfigGroups {
		groupConfig := make(ConfigMap)
		for _, cfg := range cfGroup.GetConfigurables() {
			groupConfig[cfg.Name()] = cfg.JsonConfig()
		}
		configMap[cfGroup.Name()] = groupConfig
	}

	configPath := filepath.Join(genGuide.RootPath, "config.json")
	f, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer f.Close()

	jsonData, err := json.MarshalIndent(configMap, "", "  ")
	if err != nil {
		return err
	}

	_, err = f.Write(jsonData)
	return err
}

func genSecretKey() string {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return ""
	}
	return hex.EncodeToString(key)
}
