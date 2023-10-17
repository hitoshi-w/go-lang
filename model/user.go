package model

import "time"

type UserID int64

// ユーザーモデル
type User struct {
	ID        UserID
	Name      string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}