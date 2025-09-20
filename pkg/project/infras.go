package project

import (
	"os"
	"path/filepath"
)

func (p *Project) GenerateInfras(statusChan chan string) error {
	for _, infra := range p.Infras {
		infraGuide, err := p.GetInfraGenGuide(infra)
		if err != nil {
			return err
		}

		err = infra.Generate(statusChan, *infraGuide)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) GetInfraGenGuide(infra Infra) (*GenerationGuide, error) {
	infraPath := filepath.Join(p.GenGuide.RootPath, "/infrastructure/", infra.Name())
	err := os.Mkdir(infraPath, p.GenGuide.DirPerms)
	if !os.IsExist(err) {
		return nil, err
	}

	return NewGenerationGuide(infraPath, p.GenGuide.DirPerms, p.GenGuide.FilePerms), nil
}
