package project

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

type InfraList map[string]interface{}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Api struct {
	Server Server `json:"server"`
}

type App struct {
	SecretKey string `json:"secret_key"`
}

type ConfigJSON struct {
	App    App       `json:"app"`
	Api    Api       `json:"api"`
	Infras InfraList `json:"infras"`
}

func (p *Project) GenerateJSONConfig() error {
	infras, err := p.getInfrasConfig()
	if err != nil {
		return err
	}

	secretKey := genSecretKey()
	app := App{
		SecretKey: secretKey,
	}

	api := Api{
		Server: Server{
			Host: "0.0.0.0",
			Port: "8080",
		},
	}

	data := ConfigJSON{
		App:    app,
		Api:    api,
		Infras: infras,
	}

	f, err := os.Create(p.Info.RootPath + "/config.json")
	if err != nil {
		return err
	}
	defer f.Close()

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = f.Write(jsonData)
	return err
}

func (p *Project) getInfrasConfig() (InfraList, error) {
	infraList := make(InfraList)
	for _, infra := range p.Infras {
		infraJson, err := infra.JsonConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to load infrastructure config: %e", err)
		}
		infraList[infra.Name()] = infraJson
	}

	return infraList, nil
}

func genSecretKey() string {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return ""
	}
	return hex.EncodeToString(key)
}
