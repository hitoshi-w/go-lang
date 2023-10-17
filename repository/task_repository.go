package repository

import (
	"github.com/hitoshi-w/go-lang/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	FindById(model.TaskID) (model.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (tr *taskRepository) FindById(id model.TaskID) (model.Task, error) {
	task := model.Task{}
  if err := tr.db.First(&task, id).Error; err != nil {
		return model.Task{}, err
	}
	return task, nil
}