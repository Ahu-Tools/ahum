package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	{{ range .Edges}}
		{{ range .Pkgs}}
	"{{.}}"
		{{end}}
	{{ end }}

	{{ range .Infras}}
		{{ range .Pkgs}}
	"{{.}}"
		{{end}}
	{{ end }}
)

// NewConfig loads configuration from environment variables or .env file.
func CheckConfigs() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables.")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("No config file found!")
		} else {
			panic(fmt.Errorf("Error happened during loading config file: %e", err))
		}
	}

	host := viper.Get("api.server.host")
	if _, ok := host.(string); !ok {
		log.Fatal("NO API HOST SPECIFIED")
	}

	port := viper.Get("api.server.port")
	if _, ok := port.(string); !ok {
		log.Fatal("NO API PORT SPECIFIED")
	}

	secretKey := viper.Get("app.secret_key")
	if _, ok := secretKey.(string); !ok {
		log.Fatal("app.secret_key not set. Please provide it.")
	}

}

func ConfigInfras() error {
	{{ range .Edges}}
	{{.Load}}
	{{ end }}

	{{ range .Infras}}
	{{.Load}}
	{{ end }}

	return nil
}
