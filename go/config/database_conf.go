package config

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		Driver:   "mysql",
		Host:     "localhost",
		Port:     "3307",
		User:     "root",
		Password: "password",
		Database: "note_app",
	}
}

type DBConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Database string
}
