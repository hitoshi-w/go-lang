package bootstrap

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	DBCharset              string `envconfig:"DB_CHARSET"`
	DBPassword             string `envconfig:"DB_PASSWORD"`
	DBCollation            string `envconfig:"DB_COLLATION"`
	DBHost                 string `envconfig:"DB_HOST"`
	DBName                 string `envconfig:"DB_NAME"`
	DBPort                 string `envconfig:"DB_PORT"`
	DBUsername             string `envconfig:"DB_USERNAME"`
	AppEnv                 string `envconfig:"APP_ENV"`
	Port                   int    `envconfig:"PORT"`
	RedisHost              string `envconfig:"REDIS_HOST"`
	RedisPassword          string `envconfig:"REDIS_PASSWORD"`
	RedisPort              string `envconfig:"REDIS_PORT"`
	AccessTokenSecret      string `envconfig:"ACCESS_TOKEN_SECRET"`
	AccessTokenExpiryHour  int    `envconfig:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenSecret     string `envconfig:"REFRESH_TOKEN_SECRET"`
	RefreshTokenExpiryHour int    `envconfig:"REFRESH_TOKEN_EXPIRY_HOUR"`
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
