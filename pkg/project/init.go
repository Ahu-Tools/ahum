package project

import "os"

const DefaultDirPerms = 755

func (p Project) Generate() error {
	err := os.MkdirAll(p.RootPath, DefaultDirPerms)
	if err != nil {
		return err
	}

	err = createBasicDirs(p.RootPath)
	if err != nil {
		return err
	}

	return p.GenerateConfig()
}

func createBasicDirs(rootPath string) error {
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

	err = os.MkdirAll(rootPath+"/infrastructure/db", DefaultDirPerms)
	if err != nil {
		return err
	}

	err = os.Mkdir(rootPath+"/service", DefaultDirPerms)
	if err != nil {
		return err
	}

	return nil
}
