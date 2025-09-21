package project

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	gen "github.com/Ahu-Tools/AhuM/pkg/generation"
	"github.com/Ahu-Tools/AhuM/pkg/util"
)

type ProjectInfo struct {
	PackageName string
	GoVersion   string
	RootPath    string
}

type Project struct {
	Info     ProjectInfo
	Infras   []Infra
	Edges    []Edge
	GenGuide gen.Guide
}

func NewProjectInfo(packageName, goVersion, rootPath string) *ProjectInfo {
	return &ProjectInfo{
		packageName,
		goVersion,
		rootPath,
	}
}

func NewProject(info ProjectInfo, infras []Infra, edges []Edge) Project {
	return Project{
		Info:     info,
		Infras:   infras,
		GenGuide: *gen.DefaultGuide(info.RootPath),
		Edges:    edges,
	}
}

func LoadProjectInfo(path string) (ProjectInfo, error) {
	rootPath := filepath.Clean(path)
	if filepath.IsAbs(path) {
		var err error
		rootPath, err = filepath.Localize(path)
		if err != nil {
			return ProjectInfo{}, err
		}
	}

	goModPath := filepath.Join(rootPath, "go.mod")
	content, err := os.ReadFile(goModPath)
	if err != nil {
		return ProjectInfo{}, fmt.Errorf("could not find go.mod in %s: %w", rootPath, err)
	}

	lines := strings.Split(string(content), "\n")
	var packageName, goVersion string
	for _, line := range lines {
		if strings.HasPrefix(line, "module") {
			packageName = strings.TrimSpace(strings.TrimPrefix(line, "module"))
		}
		if strings.HasPrefix(line, "go") {
			goVersion = strings.TrimSpace(strings.TrimPrefix(line, "go"))
		}
	}

	if packageName == "" {
		return ProjectInfo{}, errors.New("could not find module name in go.mod")
	}
	if goVersion == "" {
		return ProjectInfo{}, errors.New("could not find go version in go.mod")
	}

	return ProjectInfo{
		PackageName: packageName,
		GoVersion:   goVersion,
		RootPath:    rootPath,
	}, nil
}

func LoadProject(path string) (*Project, error) {
	info, err := LoadProjectInfo(path)
	if err != nil {
		return nil, err
	}

	// For now, we don't need to load infrastructure and edge configurations
	// to add a new service. So we can initialize it as empty.
	project := NewProject(info, []Infra{}, []Edge{})
	project.GenGuide = *gen.DefaultGuide(info.RootPath)
	err = project.LoadEdges()
	if err != nil {
		return nil, err
	}

	err = project.LoadInfras()
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (p *Project) LoadEdges() error {
	p.Edges = make([]Edge, 0)
	for _, loader := range edgeLoaders {
		edge, err := loader(*p, EdgesGroup)
		if err, ok := err.(util.JsonError); ok {
			if err.Code == util.NOT_FOUND_JERR {
				continue
			}
		}
		if err != nil {
			return err
		}
		p.Edges = append(p.Edges, edge)
	}
	return nil
}

func (p *Project) LoadInfras() error {
	p.Infras = make([]Infra, 0)
	for _, loader := range infraLoaders {
		infra, err := loader(*p, InfrasGroup)
		if err, ok := err.(util.JsonError); ok {
			if err.Code == util.NOT_FOUND_JERR {
				continue
			}
		}
		if err != nil {
			return err
		}
		p.Infras = append(p.Infras, infra)
	}
	return nil
}

func (p *Project) GetEdgesByName() map[string]Edge {
	edgesByName := make(map[string]Edge)
	for _, edge := range p.Edges {
		edgesByName[edge.Name()] = edge
	}
	return edgesByName
}
