package project

import (
	"fmt"
	"os"
	"os/exec"
)

const DefaultDirPerms = 0775
const DefaultFilePerms = 0664

func (p Project) Generate(statusChan chan string) error {
	defer close(statusChan)

	statusChan <- "Generating project directories structure..."
	err := os.MkdirAll(p.Info.RootPath, DefaultDirPerms)
	if err != nil {
		return err
	}

	err = createBasicDirs(p.Info.RootPath, p.InfrasJson)
	if err != nil {
		return err
	}

	statusChan <- "Initialising go.mod file..."
	err = goInit(p.Info.RootPath, p.Info.PackageName, p.Info.GoVersion)
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

	statusChan <- "Running go mod tidy..."
	err = goModTidy(p.Info.RootPath)
	if err != nil {
		return err
	}

	statusChan <- "Running go mod download..."
	err = goModDownload(p.Info.RootPath)
	if err != nil {
		return err
	}

	statusChan <- "Running go fmt..."
	err = goFmt(p.Info.RootPath)
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

func createBasicDirs(rootPath string, infras []JSONInfra) error {
	err := os.Mkdir(rootPath+"/bin", DefaultDirPerms)
	if err != nil {
		return err
	}

	err = os.Mkdir(rootPath+"/chain", DefaultDirPerms)
	if err != nil {
		return err
	}

	err = os.MkdirAll(rootPath+"/cmd/api", DefaultDirPerms)
	if err != nil {
		return err
	}

	err = os.Mkdir(rootPath+"/config", DefaultDirPerms)
	if err != nil {
		return err
	}

	err = os.Mkdir(rootPath+"/data", DefaultDirPerms)
	if err != nil {
		return err
	}

	err = os.Mkdir(rootPath+"/docs", DefaultDirPerms)
	if err != nil {
		return err
	}

	err = os.Mkdir(rootPath+"/edge", DefaultDirPerms)
	if err != nil {
		return err
	}

	err = os.Mkdir(rootPath+"/infrastructure", DefaultDirPerms)
	if err != nil {
		return err
	}

	for _, v := range infras {
		err = os.Mkdir(rootPath+"/infrastructure/"+v.Name(), DefaultDirPerms)
		if err != nil {
			return err
		}
	}

	err = os.Mkdir(rootPath+"/service", DefaultDirPerms)
	if err != nil {
		return err
	}

	return nil
}
