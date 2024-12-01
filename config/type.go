package config

type Config struct {
	DB    DBConfig    `json:"db"`
	Redis RedisConfig `json:"redis"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Database string `json:"database"`
	Schema   string `json:"schema"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}
