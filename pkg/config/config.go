package config

type Config struct {
	Database   *Database
	ServerPort int
}

type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	SSLMode  bool
}

func New() *Config {
	databaseCfg := &Database{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "password",
		Name:     "postgres",
		SSLMode:  false,
	}

	cfg := &Config{
		Database:   databaseCfg,
		ServerPort: 4321,
	}

	// switch strings.ToLower(os.Getenv("ENV")) {
	// case "production":
	// 	loadProductionConfig(cfg)
	// case "staging":
	// 	loadStagingConfig(cfg)
	// case "test":
	// 	loadTestConfig(cfg)
	// default:
	// 	loadDevelopmentConfig(cfg)
	// }
	return cfg
}
