package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hitoshi-w/go-lang/api/route"
	"github.com/hitoshi-w/go-lang/bootstrap"
)

func main() {
	app := bootstrap.App()

	db := app.Gorm

	gin := gin.Default()

	route.Setup(db, gin)

	gin.Run()
}
