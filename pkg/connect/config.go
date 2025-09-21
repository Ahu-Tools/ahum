package connect

const Name = "connect"

// We want to implement project.Edge for Connect
func (g *Connect) Name() string {
	return Name
}

func (g *Connect) Pkgs() ([]string, error) {
	return []string{}, nil
}

func (g *Connect) JsonConfig() any {
	return g.ConnectConfig
}

func (g *Connect) Load() (string, error) {
	return "", nil
}
