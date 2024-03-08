package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host         string
	Port         string
	DatabasePort string
	User         string
	Password     string
	DBName       string
}

func NewConfig() (*Config, error) {
	viper.SetConfigFile("services/article/configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{
		Host:         viper.GetString("database.host"),
		Port:         viper.GetString("server.port"),
		DatabasePort: viper.GetString("database.port"),
		User:         viper.GetString("database.user"),
		Password:     viper.GetString("database.password"),
		DBName:       viper.GetString("database.dbname"),
	}

	return cfg, nil
}
