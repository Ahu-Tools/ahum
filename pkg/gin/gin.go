package gin

import (
	"fmt"
	"os"
	"path/filepath"

	gen "github.com/Ahu-Tools/ahum/pkg/generation"
	"github.com/Ahu-Tools/ahum/pkg/project"
	"github.com/Ahu-Tools/ahum/pkg/util"
	"github.com/iancoleman/strcase"
)

const Name = "gin"

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

func LoadGinFromProject(pj project.Project) *Gin {
	edges := pj.GetEdgesByName()
	return edges[Name].(*Gin)
}

func (g *Gin) Generate(status chan string, genGuide gen.Guide) error {
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

func (g *Gin) AddHandler(versionName, entityName, handlerName string, genGuide gen.Guide) error {
	handlerName = strcase.ToCamel(handlerName)
	entityName = util.ToPkgName(entityName)
	versionName = util.ToPkgName(versionName)
	eHandlerPath := filepath.Join(genGuide.RootPath, versionName, entityName, "/handler.go")

	payload := map[string]any{
		"HandlerName": handlerName,
	}
	handleFunc, _ := util.ParseTemplateString("template/gin/entity.handle_func.go.tpl", payload)
	insertions := map[string]string{
		"handlers": handleFunc,
	}

	err := util.ModifyCodeByMarkersFile(eHandlerPath, insertions, genGuide.FilePerms)
	if err != nil {
		return err
	}

	return g.registerHandler(versionName, entityName, handlerName, genGuide)
}

func (g *Gin) registerHandler(versionName, entityName, handlerName string, genGuide gen.Guide) error {
	eRoutePath := filepath.Join(genGuide.RootPath, versionName, entityName, "/route.go")

	insertions := map[string]string{
		"routes": fmt.Sprintf(`r.GET("/%s", h.%s)`, strcase.ToKebab(handlerName), handlerName),
	}
	return util.ModifyCodeByMarkersFile(eRoutePath, insertions, genGuide.FilePerms)
}

func (g *Gin) AddEntity(versionName, entityName string, genGuide gen.Guide) error {
	entityName = util.ToPkgName(entityName)
	versionName = util.ToPkgName(versionName)

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

	return g.registerEntity(versionName, entityName, genGuide)
}

func (g *Gin) registerEntity(versionName, entityName string, genGuide gen.Guide) error {
	vRegPath := filepath.Join(genGuide.RootPath, versionName, "/registrar.go")

	insertions := map[string]string{
		"imports":  fmt.Sprintf("\"%s/edge/gin/%s/%s\"", g.pj.PackageName, versionName, entityName),
		"entities": fmt.Sprintf("%s.RegisterRoutes(r.Group(\"%s\"))", entityName, strcase.ToKebab(entityName)),
	}

	return util.ModifyCodeByMarkersFile(vRegPath, insertions, genGuide.FilePerms)
}

func (g *Gin) AddVersion(versionName string, genGuide gen.Guide) error {
	versionName = util.ToPkgName(versionName)

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

	return g.registerVersion(versionName, genGuide)
}

func (g *Gin) registerVersion(versionName string, genGuide gen.Guide) error {
	ginPath := filepath.Join(genGuide.RootPath, "gin.go")

	insertions := map[string]string{
		"imports":  fmt.Sprintf("\"%s/edge/gin/%s\"", g.pj.PackageName, versionName),
		"versions": fmt.Sprintf("%s.RegisterVersion(r.Group(\"%s\"))", versionName, versionName),
	}

	return util.ModifyCodeByMarkersFile(ginPath, insertions, genGuide.FilePerms)
}
