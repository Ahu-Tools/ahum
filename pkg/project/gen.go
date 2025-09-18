package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Ahu-Tools/AhuM/pkg/util"
)

const DefaultDirPerms = 0775
const DefaultFilePerms = 0664

type GenerationGuide struct {
	RootPath  string
	DirPerms  os.FileMode
	FilePerms os.FileMode
}

func NewGenerationGuide(rootPath string, DirPerms, FilePerms os.FileMode) GenerationGuide {
	return GenerationGuide{
		RootPath:  rootPath,
		DirPerms:  DirPerms,
		FilePerms: FilePerms,
	}
}

func DefaultGenerationGuide(rootPath string) GenerationGuide {
	return NewGenerationGuide(rootPath, DefaultDirPerms, DefaultFilePerms)
}

func (p Project) Generate(statusChan chan string) error {
	defer close(statusChan)

	statusChan <- "Generating project directories structure..."
	err := os.MkdirAll(p.GenGuide.RootPath, p.GenGuide.DirPerms)
	if err != nil {
		return err
	}

	err = createBasicDirs(p.GenGuide, p.Infras)
	if err != nil {
		return err
	}

	statusChan <- "Initialising go.mod file..."
	err = p.GoInit()
	if err != nil {
		return err
	}

	statusChan <- "Generating config.json..."
	err = p.GenerateJSONConfig()
	if err != nil {
		return err
	}

	statusChan <- "Generating config.go..."
	err = p.GenerateConfig()
	if err != nil {
		return err
	}

	statusChan <- "Generating edge..."
	err = p.GenerateEdge()
	if err != nil {
		return err
	}

	statusChan <- "Adding edge..."
	err = p.AddEdges(statusChan)
	if err != nil {
		return err
	}

	statusChan <- "Generating infrastructures..."
	err = p.GenerateInfras(statusChan)
	if err != nil {
		return err
	}

	err = p.GoSweep(statusChan)
	if err != nil {
		return err
	}

	return nil
}

func (p Project) GenerateEdge() error {
	edgePath := filepath.Join(p.GenGuide.RootPath, "/edge/edge.go")
	return util.ParseTemplateFile("template/edge/edge.go.tpl", p.Info, edgePath)
}

func (p Project) GoInit() error {
	return os.WriteFile(p.GenGuide.RootPath+"/go.mod", []byte("module "+p.Info.PackageName+"\n\ngo "+p.Info.GoVersion+"\n"), DefaultFilePerms)
}

func (p Project) GoSweep(statusChan chan string) error {
	statusChan <- "Running go mod tidy..."
	err := p.GoModTidy()
	if err != nil {
		return err
	}

	statusChan <- "Running go mod download..."
	err = p.GoModDownload()
	if err != nil {
		return err
	}

	statusChan <- "Running go fmt..."
	err = p.GoFmt()
	if err != nil {
		return err
	}

	return nil
}

func (p Project) GoModTidy() error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = p.GenGuide.RootPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go mod tidy failed!\n\nDetails:\n%s", string(output))
	}
	return nil
}

func (p Project) GoModDownload() error {
	cmd := exec.Command("go", "mod", "download")
	cmd.Dir = p.GenGuide.RootPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go mod download failed!\n\nDetails:\n%s", string(output))
	}
	return nil
}

func (p Project) GoFmt() error {
	cmd := exec.Command("go", "fmt", "./...")
	cmd.Dir = p.GenGuide.RootPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go mod fmt failed!\n\nDetails:\n%s", string(output))
	}
	return nil
}

func createBasicDirs(genGuide GenerationGuide, infras []Infra) error {
	err := os.Mkdir(genGuide.RootPath+"/bin", genGuide.DirPerms)
	if err != nil {
		return err
	}

	err = os.Mkdir(genGuide.RootPath+"/chain", genGuide.DirPerms)
	if err != nil {
		return err
	}

	err = os.MkdirAll(genGuide.RootPath+"/cmd/api", genGuide.DirPerms)
	if err != nil {
		return err
	}

	err = os.Mkdir(genGuide.RootPath+"/config", genGuide.DirPerms)
	if err != nil {
		return err
	}

	err = os.Mkdir(genGuide.RootPath+"/data", genGuide.DirPerms)
	if err != nil {
		return err
	}

	err = os.Mkdir(genGuide.RootPath+"/docs", genGuide.DirPerms)
	if err != nil {
		return err
	}

	err = os.Mkdir(genGuide.RootPath+"/edge", genGuide.DirPerms)
	if err != nil {
		return err
	}

	err = os.Mkdir(genGuide.RootPath+"/infrastructure", genGuide.DirPerms)
	if err != nil {
		return err
	}

	for _, v := range infras {
		err = os.Mkdir(genGuide.RootPath+"/infrastructure/"+v.Name(), genGuide.DirPerms)
		if err != nil {
			return err
		}
	}

	err = os.Mkdir(genGuide.RootPath+"/service", genGuide.DirPerms)
	if err != nil {
		return err
	}

	return nil
}
