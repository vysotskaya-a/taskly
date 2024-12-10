package models

import "time"

type Task struct {
	Id          string    `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`
	ProjectId   string    `json:"project_id" db:"project_id"`
	ExecutorId  string    `json:"executor_id" db:"executor_id"`
	Deadline    time.Time `json:"deadline" db:"deadline"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
