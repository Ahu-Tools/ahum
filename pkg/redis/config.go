package redis

type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewConfig(host string, port int, username, password string) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func DefaultConfig() *Config {
	return &Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Username: "default",
		Password: "1234",
	}
}
