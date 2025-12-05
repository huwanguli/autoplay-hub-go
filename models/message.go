package models

type Message struct {
	DeviceURL string `json:"device_id"`
	TaskID    int64  `json:"task_id"`
	Content   string `json:"content"`
}
