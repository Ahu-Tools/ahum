package edge

const Name = "asynq"

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type Config struct {
	Concurrency int            `json:"concurrency"`
	Queues      map[string]int `json:"queues"`
	Redis       Redis          `json:"redis"`
}

func DefaultConfig() Config {
	return Config{
		Concurrency: 10,
		Queues:      map[string]int{},
		Redis: Redis{
			Host:     "localhost",
			Port:     6379,
			Username: "admin",
			Password: "1234",
			DB:       0,
		},
	}
}

// We want to implement project.Edge for Connect
func (ae *Asynq) Name() string {
	return Name
}

func (ae *Asynq) Pkgs() ([]string, error) {
	return []string{}, nil
}

func (ae *Asynq) JsonConfig() any {
	return ae.Config
}

func (ae *Asynq) Load() (string, error) {
	return "", nil
}
