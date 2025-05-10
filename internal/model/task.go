package model

import "time"

type Task struct {
	Id          int
	Description string
	CreatedAt   time.Time
	CompletedAt *time.Time
}
