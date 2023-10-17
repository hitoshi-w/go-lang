package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	Gin *gin.Engine
}

func Setup(db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewTaskRouter(db, publicRouter)
}
