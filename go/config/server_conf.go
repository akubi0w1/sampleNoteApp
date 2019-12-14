package config

type ServerConfig struct {
	Addr string
	Port string
}

func LoadServerConfig() *ServerConfig {
	return &ServerConfig{
		Addr: "localhost",
		Port: "8080",
	}
}
