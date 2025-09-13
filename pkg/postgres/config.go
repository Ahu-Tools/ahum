package postgres

type PostgresConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	SSLMode  string `json:"sslmode"`
}

func NewPostgresConfig(user, password, dbName, host, port, sslMode string) *PostgresConfig {
	return &PostgresConfig{
		User:     user,
		Password: password,
		DbName:   dbName,
		Host:     host,
		Port:     port,
		SSLMode:  sslMode,
	}
}

func DefaultPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		User:     "postgres",
		Password: "postgres",
		DbName:   "my_app",
		Host:     "127.0.0.1",
		Port:     "5432",
		SSLMode:  "disable",
	}
}

func (pc PostgresConfig) Name() string {
	return "postgres"
}

func (pc PostgresConfig) Config() (any, error) {
	return pc, nil
}
