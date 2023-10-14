package infrastructure

import "github.com/gin-gonic/gin"

type Router struct {
	Gin *gin.Engine
}

func InitializeRouter() *Router {
	r := &Router{Gin: gin.Default()}
	return r
}

func (r *Router) Run() error {
	return r.Gin.Run()
}
