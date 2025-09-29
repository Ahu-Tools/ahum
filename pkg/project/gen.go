package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	gen "github.com/Ahu-Tools/ahum/pkg/generation"
	"github.com/Ahu-Tools/ahum/pkg/util"
)

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

	err = p.GenerateConfig(statusChan)
	if err != nil {
		return err
	}

	statusChan <- "Generating edge..."
	err = p.GenerateEdge()
	if err != nil {
		return err
	}

	statusChan <- "Generating main.go..."
	err = p.GenerateMain()
	if err != nil {
		return err
	}

	statusChan <- "Adding edge..."
	err = p.GenEdges(statusChan)
	if err != nil {
		return err
	}

	statusChan <- "Generating infrastructures..."
	err = p.GenInfras(statusChan)
	if err != nil {
		return err
	}

	err = p.GoSweep(statusChan)
	if err != nil {
		return err
	}

	return nil
}

func (p Project) GenerateMain() error {
	mainPath := filepath.Join(p.GenGuide.RootPath, "/main.go")
	return util.ParseTemplateFile("main/main.go.tpl", p.Info, mainPath)
}

func (p Project) GenerateConfig(statusChan chan string) error {
	cfg := p.GetConfig()
	cfgGenGuide, err := p.GetConfigGenGuide()
	if err != nil {
		return err
	}

	statusChan <- "Generating config.json..."
	err = cfg.GenerateJSON(statusChan, *cfgGenGuide)
	if err != nil {
		return err
	}

	statusChan <- "Generating config.go..."
	return cfg.Generate(statusChan, *cfgGenGuide)
}

func (p Project) GenerateEdge() error {
	edgePath := filepath.Join(p.GenGuide.RootPath, "/edge/edge.go")
	return util.ParseTemplateFile("edge/edge.go.tpl", p.Info, edgePath)
}

func (p Project) GoInit() error {
	return os.WriteFile(p.GenGuide.RootPath+"/go.mod", []byte("module "+p.Info.PackageName+"\n\ngo "+p.Info.GoVersion+"\n"), p.GenGuide.FilePerms)
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

func createBasicDirs(genGuide gen.Guide, infras []Infra) error {
	err := os.Mkdir(genGuide.RootPath+"/bin", genGuide.DirPerms)
	if err != nil {
		return err
	}

	err = os.Mkdir(genGuide.RootPath+"/chain", genGuide.DirPerms)
	if err != nil {
		return err
	}

	err = os.Mkdir(genGuide.RootPath+"/cmd", genGuide.DirPerms)
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
