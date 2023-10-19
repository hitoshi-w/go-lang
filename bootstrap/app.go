package bootstrap

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Application struct {
	Env *Env
	Gorm *gorm.DB
	Redis *redis.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Gorm = NewGorm(app.Env)
	app.Redis = NewRedis(app.Env)

	return *app
}

