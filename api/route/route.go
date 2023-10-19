package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/hitoshi-w/go-lang/bootstrap"
	"gorm.io/gorm"
)

type Router struct {
	Gin *gin.Engine
}

func Setup(env *bootstrap.Env, redis *redis.Client, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewTaskRouter(db, publicRouter)
	NewSignupRouter(env, db, publicRouter)
}
