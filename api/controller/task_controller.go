package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hitoshi-w/go-lang/model"
	"github.com/hitoshi-w/go-lang/usecase"
)

type TaskController struct {
	TaskUsecase usecase.TaskUsecase
}

func (tc *TaskController) Show(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	taskID := model.TaskID(id)

	task, err := tc.TaskUsecase.FindById(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, task)
}
