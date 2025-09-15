package project

import "path/filepath"

func (p *Project) GenerateInfras(statusChan chan string) error {
	for _, infra := range p.Infras {
		infraPath := filepath.Join(p.GenGuide.RootPath, "/infrastructure/", infra.Name())
		infraGuide := NewGenerationGuide(infraPath, p.GenGuide.DirPerms, p.GenGuide.FilePerms)
		err := infra.Generate(statusChan, infraGuide)
		if err != nil {
			return err
		}
	}
	return nil
}
