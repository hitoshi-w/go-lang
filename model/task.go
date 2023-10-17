package model

import "time"

// defined type を使って独自の型を定義することで、誤った型の代入を防ぐ
type TaskID int64
type TaskStatus string

const (
	TaskStatusToDo TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone TaskStatus = "done"
)

type Task struct {
	ID TaskID
	Title string
	Status TaskStatus
	CreatedAt time.Time
	UpdatedAtAt time.Time
}
