package project

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

type InfraList map[string]interface{}
type EdgeList map[string]interface{}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type App struct {
	SecretKey string `json:"secret_key"`
}

type ConfigJSON struct {
	App    App       `json:"app"`
	Infras InfraList `json:"infras"`
	Edges  EdgeList  `json:"edges"`
}

func (p *Project) GenerateJSONConfig() error {
	infras, err := p.getInfrasConfig()
	if err != nil {
		return err
	}

	edges, err := p.getEdgesConfig()
	if err != nil {
		return err
	}

	secretKey := genSecretKey()
	app := App{
		SecretKey: secretKey,
	}

	data := ConfigJSON{
		App:    app,
		Infras: infras,
		Edges:  edges,
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

func (p *Project) getEdgesConfig() (EdgeList, error) {
	edges := make(EdgeList)
	for _, edge := range p.Edges {
		edgeJson, err := edge.JsonConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to load edge config: %e", err)
		}
		edges[edge.Name()] = edgeJson
	}

	return edges, nil
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
