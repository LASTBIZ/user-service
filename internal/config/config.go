package config

import (
	"lastbiz/user-service/pkg/logging"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GRPCPort string `yaml:"grpc_port" env:"GRPC_PORT" env-required:"true"`
	Postgres struct {
		Host     string `yaml:"host" env:"POSTGRES_HOST" env-required:"true"`
		User     string `yaml:"host" env:"POSTGRES_USER" env-required:"true"`
		Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-required:"true"`
		DB       string `yaml:"db" env:"POSTGRES_DATABASE" env-required:"true"`
		Port     string `yaml:"port" env:"POSTGRES_PORT" env-required:"true"`
	} `yaml:"postgresql"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			helpText := "LastMBiz user-service by https://github.com/Suro4ek"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
