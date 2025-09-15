package project

import (
	"fmt"
	"os"
	"os/exec"
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
	err = goInit(p.GenGuide.RootPath, p.Info.PackageName, p.Info.GoVersion)
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

	statusChan <- "Generating infrastructures..."
	err = p.GenerateInfras(statusChan)
	if err != nil {
		return err
	}

	statusChan <- "Running go mod tidy..."
	err = goModTidy(p.GenGuide.RootPath)
	if err != nil {
		return err
	}

	statusChan <- "Running go mod download..."
	err = goModDownload(p.GenGuide.RootPath)
	if err != nil {
		return err
	}

	statusChan <- "Running go fmt..."
	err = goFmt(p.GenGuide.RootPath)
	if err != nil {
		return err
	}
	return nil
}

func goInit(rootPath, packageName, goVersion string) error {
	return os.WriteFile(rootPath+"/go.mod", []byte("module "+packageName+"\n\ngo "+goVersion+"\n"), DefaultFilePerms)
}

func goModTidy(rootPath string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = rootPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go mod tidy failed!\n\nDetails:\n%s", string(output))
	}
	return nil
}

func goModDownload(rootPath string) error {
	cmd := exec.Command("go", "mod", "download")
	cmd.Dir = rootPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go mod download failed!\n\nDetails:\n%s", string(output))
	}
	return nil
}

func goFmt(rootPath string) error {
	cmd := exec.Command("go", "fmt", "./...")
	cmd.Dir = rootPath
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
