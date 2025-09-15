package project

import "path/filepath"

func (p *Project) GenerateInfras(statusChan chan string, genGuide GenerationGuide) error {
	for _, infra := range p.InfrasConfig {
		infraPath := filepath.Join(genGuide.RootPath, "/infrastructure/", infra.Name())
		infraGuide := NewGenerationGuide(infraPath, genGuide.DirPerms, genGuide.FilePerms)
		err := infra.Generate(statusChan, infraGuide)
		if err != nil {
			return err
		}
	}
	return nil
}
