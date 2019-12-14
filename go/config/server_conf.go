package config

func LoadServerConfig() *ServerConfig {
	return &ServerConfig{
		Addr: "localhost",
		Port: "8080",
	}
}

type ServerConfig struct {
	Addr string
	Port string
}
