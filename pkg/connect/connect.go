package connect

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	gen "github.com/Ahu-Tools/AhuM/pkg/generation"
	"github.com/Ahu-Tools/AhuM/pkg/project"
	"github.com/Ahu-Tools/AhuM/pkg/util"
	"github.com/iancoleman/strcase"
)

type ConnectServer struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type ConnectConfig struct {
	Server ConnectServer `json:"server"`
}

type Connect struct {
	pj            *project.ProjectInfo
	ConnectConfig ConnectConfig
}

func NewConnectServer(host, port string) *ConnectServer {
	return &ConnectServer{
		Host: host,
		Port: port,
	}
}

func NewConnectConfig(srv ConnectServer) *ConnectConfig {
	return &ConnectConfig{
		srv,
	}
}

func NewConnect(pj *project.ProjectInfo, ConnectConfig ConnectConfig) *Connect {
	return &Connect{
		pj,
		ConnectConfig,
	}
}

func (g *Connect) Generate(status chan string, genGuide gen.Guide) error {
	err := util.ParseTemplateFile("template/connect/connect.go.tpl", g.pj, genGuide.RootPath+"/connect.go")
	if err != nil {
		return err
	}

	err = util.ParseTemplateFile("template/connect/buf.yaml.tpl", g.pj, genGuide.RootPath+"/buf.yaml")
	if err != nil {
		return err
	}

	err = util.ParseTemplateFile("template/connect/buf.gen.yaml.tpl", g.pj, genGuide.RootPath+"/buf.gen.yaml")
	if err != nil {
		return err
	}

	err = g.AddService("hello", genGuide)
	if err != nil {
		return err
	}

	err = g.AddVersion("v1", "hello", genGuide)
	if err != nil {
		return err
	}

	err = g.AddMethod("world", "hello", "v1", genGuide)
	if err != nil {
		return err
	}
	return nil
}

func (g *Connect) AddMethod(methodName, serviceName, versionName string, genGuide gen.Guide) error {
	svcName := util.ToPkgName(serviceName)
	vName := util.ToPkgName(versionName)
	methName := strcase.ToCamel(methodName)

	payload := map[string]any{
		"MethodName": methName,
	}
	method, _ := util.ParseTemplateString("template/connect/method.proto.tpl", payload)
	messages, _ := util.ParseTemplateString("template/connect/messages.proto.tpl", payload)

	protoPath := filepath.Join(genGuide.RootPath, svcName, vName, strcase.ToSnakeWithIgnore(serviceName, ".")+".proto")
	insertion := map[string]string{
		"methods":  method,
		"messages": messages,
	}

	err := util.ModifyFileByMarkersFile(protoPath, insertion, genGuide.FilePerms)
	if err != nil {
		return err
	}

	err = g.BufGenerate(genGuide)
	if err != nil {
		return err
	}

	return g.implementMethod(serviceName, versionName, methName, genGuide)
}

func (g *Connect) implementMethod(serviceName, versionName, methodName string, genGuide gen.Guide) error {
	payload := map[string]any{
		"ServiceName": serviceName,
		"VersionName": versionName,
		"PackageName": g.pj.PackageName,
		"MethodName":  methodName,
		"Lowerer":     util.ToPkgName,
	}
	method, _ := util.ParseTemplateString("template/connect/method.go.tpl", payload)
	edgePath := filepath.Join(genGuide.RootPath, serviceName, versionName, "edge.go")
	insertions := map[string]string{
		"methods": method,
	}

	return util.ModifyCodeByMarkersFile(edgePath, insertions, genGuide.FilePerms)
}

func (g *Connect) AddVersion(versionName, serviceName string, genGuide gen.Guide) error {
	svcName := util.ToPkgName(serviceName)
	vName := util.ToPkgName(versionName)
	vPath := filepath.Join(genGuide.RootPath, svcName, vName)

	err := os.Mkdir(vPath, genGuide.DirPerms)
	if err != nil {
		return err
	}

	payload := map[string]any{
		"ServiceName": serviceName,
		"VersionName": versionName,
		"PackageName": g.pj.PackageName,
		"Snaker":      func(s string) string { return strcase.ToSnakeWithIgnore(s, ".") },
		"Lowerer":     util.ToPkgName,
	}

	protoPath := filepath.Join(vPath, strcase.ToSnakeWithIgnore(serviceName, ".")+".proto")
	err = util.ParseTemplateFile("template/connect/edge.proto.tpl", payload, protoPath)
	if err != nil {
		return err
	}

	err = g.BufGenerate(genGuide)
	if err != nil {
		return err
	}

	edgePath := filepath.Join(vPath, "edge.go")
	err = util.ParseTemplateFile("template/connect/edge.go.tpl", payload, edgePath)
	if err != nil {
		return err
	}

	vRegPath := filepath.Join(vPath, "registrar.go")
	err = util.ParseTemplateFile("template/connect/version.registrar.go.tpl", payload, vRegPath)
	if err != nil {
		return nil
	}

	return g.registerVersion(serviceName, versionName, genGuide)
}

func (g *Connect) BufGenerate(genGuide gen.Guide) error {
	cmd := exec.Command("buf", "generate")
	cmd.Dir = genGuide.RootPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("'buf generate' execution failed!\nDetails:\n%s", string(output))
	}
	return nil
}

func (g *Connect) registerVersion(serviceName, versionName string, genGuide gen.Guide) error {
	svcName := util.ToPkgName(serviceName)
	vName := util.ToPkgName(versionName)
	regPath := filepath.Join(genGuide.RootPath, svcName, "registrar.go")
	insertions := map[string]string{
		"imports":  fmt.Sprintf(`%s "%s/edge/connect/%s/%s"`, vName, g.pj.PackageName, svcName, vName),
		"versions": fmt.Sprintf("%s.RegisterVersion(mux)", vName),
	}

	return util.ModifyCodeByMarkersFile(regPath, insertions, genGuide.FilePerms)
}

func (g *Connect) AddService(serviceName string, genGuide gen.Guide) error {
	svcName := util.ToPkgName(serviceName)
	svcPath := filepath.Join(genGuide.RootPath, svcName)

	err := os.Mkdir(svcPath, genGuide.DirPerms)
	if err != nil {
		return err
	}

	svcRegPath := filepath.Join(svcPath, "registrar.go")
	payload := map[string]any{
		"ServiceName": serviceName,
		"Lowerer":     util.ToPkgName,
	}

	err = util.ParseTemplateFile("template/connect/service.registrar.go.tpl", payload, svcRegPath)
	if err != nil {
		return err
	}

	return g.registerService(serviceName, genGuide)
}

func (g *Connect) registerService(serviceName string, genGuide gen.Guide) error {
	svcName := util.ToPkgName(serviceName)
	regPath := filepath.Join(genGuide.RootPath, "connect.go")
	insertions := map[string]string{
		"imports":  fmt.Sprintf(`"%s/edge/%s/%s"`, g.pj.PackageName, g.Name(), svcName),
		"services": fmt.Sprintf("%s.RegisterService(mux)", svcName),
	}

	return util.ModifyCodeByMarkersFile(regPath, insertions, genGuide.FilePerms)
}
