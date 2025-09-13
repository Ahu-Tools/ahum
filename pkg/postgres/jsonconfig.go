package postgres

type PostgresJSONConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	SSLMode  string `json:"sslmode"`
}

func NewPostgresJSONConfig(user, password, dbName, host, port, sslMode string) *PostgresJSONConfig {
	return &PostgresJSONConfig{
		User:     user,
		Password: password,
		DbName:   dbName,
		Host:     host,
		Port:     port,
		SSLMode:  sslMode,
	}
}

func DefaultPostgresJSONConfig() *PostgresJSONConfig {
	return &PostgresJSONConfig{
		User:     "postgres",
		Password: "postgres",
		DbName:   "my_app",
		Host:     "127.0.0.1",
		Port:     "5432",
		SSLMode:  "disable",
	}
}

func (pc PostgresJSONConfig) Name() string {
	return "postgres"
}

func (pc PostgresJSONConfig) Config() (any, error) {
	return pc, nil
}
