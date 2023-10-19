package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hitoshi-w/go-lang/api/controller"
	"github.com/hitoshi-w/go-lang/bootstrap"
	"github.com/hitoshi-w/go-lang/repository"
	"github.com/hitoshi-w/go-lang/usecase"
	"gorm.io/gorm"
)

func NewSignupRouter(env *bootstrap.Env, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	sc := &controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur),
		Env: env,
	}
	group.POST("/signup", sc.Signup)
}
