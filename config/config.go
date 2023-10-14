package config

import "github.com/kelseyhightower/envconfig"

type (
	DBConfig struct {
		Charset   string `envconfig:"CHARSET"`
		Password  string `envconfig:"PASSWORD"`
		Collation string `envconfig:"COLLATION"`
		Host      string `envconfig:"HOST"`
		Name      string `envconfig:"NAME"`
		Port      string `envconfig:"PORT"`
		Username  string `envconfig:"USERNAME"`
	}
)

type Configuration struct {
	DB DBConfig
}

var config Configuration

func Initialize() (*Configuration, error) {
	if err := envconfig.Process("db", &config.DB); err != nil {
		return nil, err
	}

	return &config, nil
}
