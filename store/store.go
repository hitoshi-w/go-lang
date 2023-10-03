package store

import (
	// "errors"

	"github.com/hitoshi-w/go-lang/entity"
)

type TaskStore struct {
	LastID entity.TaskID
	Tasks map[entity.TaskID]*entity.Task
}

// var (
// 	Tasks = &TaskStore{Tasks: map[int]*entity.Task}

// 	ErrNotFound = errors.New("not found")
// )

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t

	return t.ID, nil
}

func (ts *TaskStore) All() entity.Tasks {
	tasks := make([]*entity.Task, len(ts.Tasks))

	for i, t := range ts.Tasks {
		tasks[i-1] = t
	}

	return tasks
}