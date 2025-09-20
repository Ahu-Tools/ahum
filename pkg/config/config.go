package config

type Config struct {
	PackageName  string
	ConfigGroups []ConfigurableGroup
}

func NewConfig(packageName string, cfgGroups []ConfigurableGroup) *Config {
	return &Config{
		PackageName:  packageName,
		ConfigGroups: cfgGroups,
	}
}

type ConfigurableGroup interface {
	Name() string
	GetConfigurables() []Configurable
}

type Configurable interface {
	Pkgs() ([]string, error)
	Load() (string, error)
	Name() string
	JsonConfig() any
}
