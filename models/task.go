package models

import "time"

type Task struct {
	ID         int64     `json:"task_id" db:"id"`
	TaskName   string    `json:"task_name" db:"task_name"`
	ScriptID   int64     `json:"script_id,string" binding:"required" db:"script_id"`
	Status     int       `json:"status" db:"status"`
	Log        string    `json:"log" db:"log_content"`
	UserID     int64     `json:"user_id" db:"user_id"`
	ExecutedAt time.Time `json:"executed_at" db:"executed_at"`
}

type ApiTask struct {
	TaskID     int64  `json:"task_id" db:"id"`
	TaskName   string `json:"task_name" db:"task_name"`
	ScriptName string `json:"script_name,string"`
	Status     int    `json:"status" db:"status"`
}

type ApiTaskDetail struct {
	Task
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
