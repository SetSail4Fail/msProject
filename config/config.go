package config

import (
	"fmt"
	"path"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		DB DB `yaml:"database"`
	}

	DB struct {
		Host     string `env-required:"true" yaml:"host"		env-default:"localhost"`
		Port     int    `env-required:"true" yaml:"port" 		env-default:"6444"`
		Name     string `env-required:"true" yaml:"name" 		env:"xenous"`
		Password string `env-required:"true" yaml:"password"	env:"xenous"`
		DBname   string `env-required:"true" yaml:"dbname" 		env:"mydatabase"`
	}
)

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{DB: DB{}}

	err := cleanenv.ReadConfig(path.Join("./", configPath), cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("error updating env: %w", err)
	}

	return cfg, nil
}
