package gin

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/util"
	"github.com/iancoleman/strcase"
)

type GinServer struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type GinConfig struct {
	Server GinServer `json:"server"`
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
	return []string{}, nil
}

func (g *Gin) JsonConfig() (any, error) {
	return g.ginConfig, nil
}

func (g *Gin) Load() (string, error) {
	return "", nil
}

func (g *Gin) Generate(status chan string, genGuide project.GenerationGuide) error {
	err := util.ParseTemplateFile("template/gin/gin.go.tpl", g.pj, genGuide.RootPath+"/gin.go")
	if err != nil {
		return err
	}

	err = g.AddVersion("v1", genGuide)
	if err != nil {
		return err
	}

	err = g.AddEntity("v1", "hello", genGuide)
	if err != nil {
		return err
	}

	err = g.AddHandler("v1", "hello", "world", genGuide)
	if err != nil {
		return err
	}
	return nil
}

func (g *Gin) AddHandler(versionName, entityName, handlerName string, genGuide project.GenerationGuide) error {
	handlerName = strcase.ToCamel(handlerName)
	entityName = strcase.ToLowerCamel(entityName)
	versionName = strcase.ToLowerCamel(versionName)

	eHandlerPath := filepath.Join(genGuide.RootPath, versionName, entityName, "/handler.go")
	handlerData, err := os.ReadFile(eHandlerPath)
	if err != nil {
		return err
	}

	insertions := map[string]string{
		"handlers": fmt.Sprintf(
			`func (h Handler) %s(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "Hello World!",
				})
			}`,
			handlerName),
	}

	newFile, err := util.ModifyCodeByMarkers(handlerData, insertions)
	if err != nil {
		return err
	}

	err = os.WriteFile(eHandlerPath, newFile, genGuide.FilePerms)
	if err != nil {
		return err
	}

	return g.RegisterHandler(versionName, entityName, handlerName, genGuide)
}

func (g *Gin) RegisterHandler(versionName, entityName, handlerName string, genGuide project.GenerationGuide) error {
	eRoutePath := filepath.Join(genGuide.RootPath, versionName, entityName, "/route.go")

	routeData, err := os.ReadFile(eRoutePath)
	if err != nil {
		return err
	}

	insertions := map[string]string{
		"routes": fmt.Sprintf(`r.GET("/%s", h.%s)`, strcase.ToKebab(handlerName), handlerName),
	}
	newFile, err := util.ModifyCodeByMarkers(routeData, insertions)
	if err != nil {
		return err
	}

	return os.WriteFile(eRoutePath, newFile, genGuide.FilePerms)
}

func (g *Gin) AddEntity(versionName, entityName string, genGuide project.GenerationGuide) error {
	entityName = strcase.ToLowerCamel(entityName)
	versionName = strcase.ToLowerCamel(versionName)

	ePath := filepath.Join(genGuide.RootPath, versionName, entityName)
	err := os.Mkdir(ePath, genGuide.DirPerms)
	if err != nil {
		return err
	}

	eRoutePath := filepath.Join(ePath, "route.go")
	payload := map[string]string{
		"EntityName": entityName,
	}
	err = util.ParseTemplateFile("template/gin/entity.route.go.tpl", payload, eRoutePath)
	if err != nil {
		return err
	}

	eHandlerPath := filepath.Join(ePath, "handler.go")
	err = util.ParseTemplateFile("template/gin/entity.handler.go.tpl", payload, eHandlerPath)
	if err != nil {
		return err
	}

	return g.RegisterEntity(versionName, entityName, genGuide)
}

func (g *Gin) RegisterEntity(versionName, entityName string, genGuide project.GenerationGuide) error {
	vRegPath := filepath.Join(genGuide.RootPath, versionName, "/registrar.go")

	regData, err := os.ReadFile(vRegPath)
	if err != nil {
		return err
	}

	insertions := map[string]string{
		"imports":  fmt.Sprintf("\"%s/edge/gin/%s/%s\"", g.pj.PackageName, versionName, entityName),
		"entities": fmt.Sprintf("%s.RegisterRoutes(r.Group(\"%s\"))", entityName, strcase.ToKebab(entityName)),
	}

	newRegData, err := util.ModifyCodeByMarkers(regData, insertions)
	if err != nil {
		return err
	}

	return os.WriteFile(vRegPath, newRegData, genGuide.FilePerms)
}

func (g *Gin) AddVersion(versionName string, genGuide project.GenerationGuide) error {
	versionName = strcase.ToLowerCamel(versionName)

	vPath := filepath.Join(genGuide.RootPath, versionName)
	err := os.Mkdir(vPath, genGuide.DirPerms)
	if err != nil {
		return err
	}

	vRegPath := filepath.Join(vPath, "registrar.go")

	payload := map[string]string{
		"VersionName": versionName,
	}

	err = util.ParseTemplateFile("template/gin/version.registrar.go.tpl", payload, vRegPath)
	if err != nil {
		return err
	}

	return g.RegisterVersion(versionName, genGuide)
}

func (g *Gin) RegisterVersion(versionName string, genGuide project.GenerationGuide) error {
	ginPath := filepath.Join(genGuide.RootPath, "gin.go")
	regData, err := os.ReadFile(ginPath)
	if err != nil {
		return err
	}

	insertions := map[string]string{
		"imports":  fmt.Sprintf("\"%s/edge/gin/%s\"", g.pj.PackageName, versionName),
		"versions": fmt.Sprintf("%s.RegisterVersion(r.Group(\"%s\"))", versionName, versionName),
	}

	newRegData, err := util.ModifyCodeByMarkers(regData, insertions)
	if err != nil {
		return err
	}

	return os.WriteFile(ginPath, newRegData, genGuide.FilePerms)
}
