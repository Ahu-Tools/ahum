package service

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Ahu-Tools/ahum/pkg/project"
	"github.com/Ahu-Tools/ahum/pkg/util"
)

type ServiceData struct {
	Name        string
	PackageName string
}

type templateData struct {
	PackageName string
	Service     ServiceData
}

type Service struct {
	project     *project.Project
	serviceData ServiceData
}

func NewServiceData(name, pkgName string) *ServiceData {
	return &ServiceData{
		name,
		pkgName,
	}

}

func NewService(project *project.Project, svcData ServiceData) *Service {
	return &Service{
		project,
		svcData,
	}
}

func (s Service) Generate(statusChan chan string) error {
	statusChan <- "Generating service directories structure..."
	err := s.generateBasicDirs()
	if err != nil {
		return err
	}

	statusChan <- "Generate service files..."
	err = s.generateFiles()
	if err != nil {
		return err
	}

	err = s.project.GoSweep(statusChan)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) generateFiles() error {
	t := map[string][]string{
		"chain":   {"chain", "handler", "entity"},
		"service": {"service", "handler", "entity"},
		"data":    {"repo", "entity"},
	}

	for pkg, files := range t {
		for _, file := range files {
			err := s.parseTemplate(pkg, file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s Service) parseTemplate(pkg string, template string) error {
	tmplData := s.tmplData()
	r := filepath.Join(s.project.GenGuide.RootPath, fmt.Sprintf("/%s/", pkg), s.serviceData.PackageName)

	path := filepath.Join(r, fmt.Sprintf("/%s.go", template))

	return util.ParseTemplateFile(
		fmt.Sprintf("template/%s/%s.go.tpl", pkg, template), tmplData, path)
}

func (s Service) generateBasicDirs() error {
	r, p := s.project.GenGuide.RootPath, s.project.GenGuide.DirPerms

	for _, pkg := range []string{"chain", "service", "data"} {
		err := os.MkdirAll(fmt.Sprintf("%s/%s/%s", r, pkg, s.serviceData.PackageName), p)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s Service) tmplData() templateData {
	return templateData{
		PackageName: s.project.Info.PackageName,
		Service:     s.serviceData,
	}
}
