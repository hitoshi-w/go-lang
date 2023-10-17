package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env *Env
	Gorm *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Gorm = NewGorm(app.Env)

	return *app
}

