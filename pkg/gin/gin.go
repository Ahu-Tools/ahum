package gin

import "github.com/Ahu-Tools/AhuM/pkg/project"

type GinServer struct {
	Host string
	Port string
}

type GinConfig struct {
	Server GinServer
}

type Gin struct {
	pj        *project.ProjectInfo
	ginConfig GinConfig
}

func NewGinServer(host, port string) *GinServer {
	return &GinServer{
		Host: host,
		Port: port,
	}
}

func NewGinConfig(srv GinServer) *GinConfig {
	return &GinConfig{
		srv,
	}
}

func NewGin(pj *project.ProjectInfo, ginConfig GinConfig) *Gin {
	return &Gin{
		pj,
		ginConfig,
	}
}

//We want to implement project.Edge for Gin

func (g *Gin) Name() string {
	return "gin"
}

func (g *Gin) Pkgs() ([]string, error) {
	return []string{"github.com/gin-gonic/gin"}, nil
}

func (g *Gin) JsonConfig() (any, error) {
	return g.ginConfig, nil
}

func (g *Gin) Load() (string, error) {
	return "", nil
}

func (g *Gin) Generate(status chan string, genGuide project.GenerationGuide) error {
	return nil
}
