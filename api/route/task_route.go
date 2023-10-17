package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hitoshi-w/go-lang/api/controller"
	"github.com/hitoshi-w/go-lang/repository"
	"github.com/hitoshi-w/go-lang/usecase"
	"gorm.io/gorm"
)

func NewTaskRouter(db *gorm.DB, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr),
	}
	group.GET("/task/:id", tc.Show)
}
