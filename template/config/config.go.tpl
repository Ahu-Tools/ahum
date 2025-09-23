package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	{{ range .ConfigGroups}}
		{{ range .GetConfigurables}}
			{{range .Pkgs}}
	"{{.}}"
			{{end}}
		{{end}}
	{{ end }}
	//@ahum: imports
)

// NewConfig loads configuration from environment variables or .env file.
func CheckConfigs() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("No config file found!")
		} else {
			panic(fmt.Errorf("Error happened during loading config file: %e", err))
		}
	}

	secretKey := viper.Get("app.secret_key")
	if _, ok := secretKey.(string); !ok {
		log.Fatal("app.secret_key not set. Please provide it.")
	}

}

func ConfigInfras() error {
	{{ range .ConfigGroups}}
	
	// @ahum:{{.Name}}.group
		{{ range .GetConfigurables}}

	// @ahum:{{.Name}}.load
	{{.Load}}
	// @ahum:end.{{.Name}}.load

		{{end}}
	// @ahum:end.{{.Name}}.group

	{{ end }}

	//@ahum: loads

	return nil
}
