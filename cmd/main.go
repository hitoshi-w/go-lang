package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hitoshi-w/go-lang/api/route"
	"github.com/hitoshi-w/go-lang/bootstrap"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Gorm

	redis := app.Redis

	gin := gin.Default()

	route.Setup(env, redis, db, gin)

	gin.Run(":3000")
}
