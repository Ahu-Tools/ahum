package gin

// We want to implement project.Edge for Gin
func (g *Gin) Name() string {
	return Name
}

func (g *Gin) Pkgs() ([]string, error) {
	return []string{}, nil
}

func (g *Gin) JsonConfig() any {
	return g.ginConfig
}

func (g *Gin) Load() (string, error) {
	return "", nil
}
