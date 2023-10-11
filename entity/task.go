package entity

import "time"

// defined type を使って独自の型を定義することで、誤った型の代入を防ぐ
type TaskID int
type TaskStatus string

const (
	TaskStatusToDo TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone TaskStatus = "done"
)

type Task struct {
	ID TaskID `json:"id"`
	Title string `json:"title"`
	Status TaskStatus `json:"status"`
	Created time.Time `json:"created"`
}

type Tasks []*Task