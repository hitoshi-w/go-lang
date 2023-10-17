package bootstrap

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	DBCharset   string `envconfig:"DB_CHARSET"`
	DBPassword  string `envconfig:"DB_PASSWORD"`
	DBCollation string `envconfig:"DB_COLLATION"`
	DBHost      string `envconfig:"DB_HOST"`
	DBName      string `envconfig:"DB_NAME"`
	DBPort      string `envconfig:"DB_PORT"`
	DBUsername  string `envconfig:"DB_USERNAME"`
	AppEnv      string `envconfig:"APP_ENV"`
}

func NewEnv() *Env {
	env := Env{}

	err := envconfig.Process("", &env)

	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
