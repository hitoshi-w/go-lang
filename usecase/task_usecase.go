package usecase

import (
	"github.com/hitoshi-w/go-lang/model"
	"github.com/hitoshi-w/go-lang/repository"
)

type TaskUsecase interface {
	FindById(id model.TaskID) (model.Task, error)
}

type taskUsecase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUsecase(tr repository.TaskRepository) TaskUsecase {
	return &taskUsecase{
		taskRepository: tr,
	}
}

func (tu *taskUsecase) FindById(id model.TaskID) (model.Task, error) {
	return tu.taskRepository.FindById(id)
}
